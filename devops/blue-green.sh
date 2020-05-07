#!/bin/bash

# Credits: https://www.ianlewis.org/en/bluegreen-deployments-kubernetes
# blue-green.sh <servicename> <version> <green-deployment.yaml>
# Deployment name should be <service>-<version>

python3 changeversion.py $2

DEPLOYMENTNAME=$1-$2
SERVICE=$1
VERSION=$2
DEPLOYMENTFILE=$3

kubectl apply -f $DEPLOYMENTFILE

# Wait until the Deployment is ready by checking the MinimumReplicasAvailable condition.
READY=$(kubectl get deploy $DEPLOYMENTNAME -o json | jq '.status.conditions[] | select(.reason == "MinimumReplicasAvailable") | .status' | tr -d '"')
while [[ "$READY" != "True" ]]; do
    READY=$(kubectl get deploy $DEPLOYMENTNAME -o json | jq '.status.conditions[] | select(.reason == "MinimumReplicasAvailable") | .status' | tr -d '"')
    echo "Waiting container to be ready."
    sleep 5
done

# Update the service selector with the new version
kubectl patch svc $SERVICE -p "{\"spec\":{\"selector\": {\"name\": \"${SERVICE}\", \"version\": \"${VERSION}\"}}}"

echo "Done."