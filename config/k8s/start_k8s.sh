#!/bin/bash

check_minikube() {
  if ! command -v minikube &> /dev/null
  then
    echo "Minikube is not installed. Please install Minikube and try again."
    exit 1
  else
    echo "Minikube is installed."
  fi
}

check_istio() {
  if ! command -v istioctl &> /dev/null
  then
    echo "Istio is not installed. Please install Istio and try again."
    exit 1
  else
    echo "Istio is installed."
  fi
}

check_minikube
check_istio
minikube start --memory=14350 --cpus=8
istioctl install --set profile=demo -y
kubectl label namespace default istio-injection=enabled

cd ~/.istio
kubectl apply -f samples/addons
kubectl rollout status deployment/kiali -n istio-system

minikube tunnel