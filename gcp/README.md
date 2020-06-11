## Use ext-dns to update Cloud DNS automatically whenever a new service is exposed
```
kubectl create secret generic gcloud-config --from-file=gcloud-config.json -n default
kubectl apply -f ext-dns.yaml -n default
```