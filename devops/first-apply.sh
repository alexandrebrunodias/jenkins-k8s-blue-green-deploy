#!/bin/bash

kubectl apply -f db/service.yml
kubectl apply -f db/statefulset.yml

kubectl apply -f app/service.yml
kubectl apply -f app/deployment.yml


# kubectl delete -f db/service.yml
# kubectl delete -f db/statefulset.yml

# kubectl delete -f app/service.yml
# kubectl delete -f app/deployment.yml