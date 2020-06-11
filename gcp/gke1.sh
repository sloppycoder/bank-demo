#!/bin/bash

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
     
    secret=$(kubectl get secret gcloud-config -o name)
    if [ "$secret" = "" ]; then
        kubectl create secret generic gcloud-config --from-file=gcp/gcloud-config.json -n default
    else 
        echo secret already exists
    fi
    
    kubectl apply -f ext-dns.yaml -n default

}


if [ "$1" = "-f" ]; then
    echo creating cluster $CLUSTER_NAME 
    create_cluster
    deploy_tools
fi

gcloud container clusters get-credentials ${CLUSTER_NAME}

OUTPUT=$(kubectl config get-contexts -o name | grep ${CLUSTER_NAME} | head -1)

if [ -z "$OUTPUT" ]; then
    echo Unable to get kubectl contexts, something is wrong...
    exit 1
fi

kubectl config delete-context ${CLUSTER_NAME}
kubectl config rename-context ${OUTPUT} ${CLUSTER_NAME}
kubectl config get-contexts
