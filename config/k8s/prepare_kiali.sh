#!/bin/bash

cd ~/.istio
kubectl apply -f samples/addons
kubectl rollout status deployment/kiali -n istio-system