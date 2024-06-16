#!/bin/bash

deployments=("core-v2" "db-v2" )

cpu_percent=50
min_replicas=3
max_replicas=10

for deployment in "${deployments[@]}"; do
  if kubectl get hpa "$deployment" &> /dev/null; then
    echo "HPA for deployment $deployment already exists."
  else
    echo "Creating HPA for deployment $deployment."
    kubectl autoscale deployment "$deployment" --cpu-percent="$cpu_percent" --min="$min_replicas" --max="$max_replicas"
  fi
done
