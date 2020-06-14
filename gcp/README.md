## script to provision GKE cluster

When using GKE, the kubectl context must be called gke1 for other scripts to work. the ```gke1.sh``` script helps with the task. 

This script also provision [ext-dns](https://github.com/kubernetes-sigs/external-dns) which automatically create DNS record based on host setting in [Istio Gateway](https://istio.io/latest/docs/reference/config/networking/gateway/) object. This component is optional. A google service account credential JSON file is required for this to work.


```
# retrieve credential for the new cluster and store locally

./gke1.sh 

# do all the above create a new cluster called gke1 the do all the above

./gke1.sh -f



```