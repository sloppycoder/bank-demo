#!/bin/bash

# set -e

if [ -z "${GCP_PROJECT_ID}" ]; then
    GCP_PROJECT_ID=$(gcloud config get-value project)
    echo using default GCP project $GCP_PROJECT_ID
fi

CLUSTER_NAME=gke1

function create_cluster {

gcloud beta container clusters create "${CLUSTER_NAME}" \
--project "${GCP_PROJECT_ID}" --zone "us-east1-b" \
--no-enable-basic-auth \
--release-channel "regular" \
--machine-type "n1-standard-2" --image-type "COS" \
--disk-type "pd-standard" --disk-size "50" \
--metadata disable-legacy-endpoints=true \
--num-nodes "3" \
--default-max-pods-per-node "110" \
--scopes \
"https://www.googleapis.com/auth/devstorage.read_only",\
"https://www.googleapis.com/auth/logging.write",\
"https://www.googleapis.com/auth/monitoring",\
"https://www.googleapis.com/auth/servicecontrol",\
"https://www.googleapis.com/auth/service.management.readonly",\
"https://www.googleapis.com/auth/trace.append" \
--enable-stackdriver-kubernetes \
--enable-ip-alias \
--network "projects/${GCP_PROJECT_ID}/global/networks/default" \
--subnetwork "projects/${GCP_PROJECT_ID}/regions/us-east1/subnetworks/default" \
--no-enable-master-authorized-networks \
--addons HorizontalPodAutoscaling,HttpLoadBalancing,Istio,ApplicationManager \
--istio-config auth=MTLS_PERMISSIVE \
--enable-autoupgrade \
--enable-autorepair \
--max-surge-upgrade 1 \
--max-unavailable-upgrade 0 \
--labels project=bank-demo

sleep 5

}

function deploy_tools {
    if [ -f "gcloud-config.json" ]; then
        kubectl create secret generic gcloud-config --from-file=gcloud-config.json -n default
        kubectl apply -f ext-dns.yaml -n default
    else 
        echo ext-dns not deployed
        echo did you get a service account key file?
    fi
}


if [ "$1" = "--create" ]; then
    echo creating cluster $CLUSTER_NAME 
    create_cluster
    deploy_tools
fi

# delete existing gke1 context before retrieving a new one
kubectl config delete-context ${CLUSTER_NAME}
gcloud container clusters get-credentials ${CLUSTER_NAME}
OUTPUT=$(kubectl config get-contexts -o name | grep ${CLUSTER_NAME} | sort | head -1)

if [ -z "$OUTPUT" ]; then
    echo Unable to get kubectl contexts, something is wrong...
    echo to create a new cluster, use --create option
    exit 1
fi

kubectl config rename-context ${OUTPUT} ${CLUSTER_NAME}
kubectl config get-contexts
