#
# utilities functions shared by various scripts
#

#
# determine the the ingress host:port of the services
#
# use load balancer external IP when running in GKE
# use service nodeport when running with minikube
#
function get_ingress_addr 
{
    CTX=$(kubectl config current-context)

    ISTIO=$(kubectl get namespaces | grep istio-system | wc -l )
    if [ "$ISTIO" = "0" ]; then
        return
    fi

    case "$CTX" in

        "gke1")
            DEMOBANK_HOST=$(kubectl get svc istio-ingressgateway -n istio-system -o jsonpath="{.status.loadBalancer.ingress[*].ip}")
            DEMOBANK_PORT=31400
            ;;

        "minikube")
            DEMOBANK_HOST=$(minikube ip)
            DEMOBANK_PORT=$(kubectl get svc istio-ingressgateway -n istio-system -o jsonpath="{.spec.ports[?(@.port=="31400")].nodePort}")
            ;;

        "")
            echo No current kubectl context?...
            ;;

    esac
}