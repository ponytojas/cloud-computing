#/bin/bash
kubectl apply -n default -f ./volumes/\*.yaml
kubectl apply -n default -f ./services/\*.yaml
kubectl apply -n default -f ./servicesAccount/\*.yaml
kubectl apply -n default -f ./deployments/\*.yaml
kubectl apply -n default -f ./ingress/\*.yaml