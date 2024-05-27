#!/bin/bash

# Function to check if Metrics Server is running
function check_metrics_server {
  kubectl get deployment metrics-server -n kube-system &> /dev/null
  if [ $? -eq 0 ]; then
    return 0  # Metrics Server is running
  else
    return 1  # Metrics Server is not running
  fi
}

# Function to deploy Metrics Server
function deploy_metrics_server {
  echo "Deploying Metrics Server..."
  kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
  if [ $? -eq 0 ]; then
    echo "Metrics Server deployed successfully."
  else
    echo "Failed to deploy Metrics Server."
    exit 1
  fi
}

# Main script execution
echo "Checking if Metrics Server is running..."
if check_metrics_server; then
  echo "Metrics Server is already running."
else
  echo "Metrics Server is not running."
  deploy_metrics_server
fi
