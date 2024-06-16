#!/bin/bash

function part1 {

  kubectl apply -n default -f ./deploy_v1/volumes/\*.yaml
  kubectl apply -n default -f ./deploy_v1/services/\*.yaml
  kubectl apply -n default -f ./deploy_v1/servicesAccount/\*.yaml
  kubectl apply -n default -f ./deploy_v1/deployments/\*.yaml
  kubectl apply -n default -f ./deploy_v1/ingress/\*.yaml
  sh ./deploy_v1/hpa/autoscale.sh
  
  kubectl apply -f ./specific-routing/destination-rules.yaml
  kubectl apply -n default -f ./deploy_v1/virtual-service.yaml

}

function part2 {
  kubectl apply -n default -f ./deploy_v2/deployments/core.yaml
  kubectl apply -n default -f ./deploy_v2/deployments/db.yaml
  sh ./deploy_v2/hpa/autoscale_1.sh

}

function part3 {
  kubectl apply -n default -f ./a-b/virtual-service-ab.yaml
}

function part4 {
  kubectl apply -n default -f ./deploy_v2/deployments/\*.yaml
  sh ./deploy_v2/hpa/autoscale_2.sh

}

function part5 {
  kubectl delete deployments -l version=v2
  kubectl delete virtualservice -n default  mdaw-virtual-service
  kubectl delete virtualservice -n default  db-virtual-service
  kubectl apply -n default -f ./deploy_v1/ingress/gateway.yaml
}

function display_menu {
  echo "Select an option:"
  echo "1) Deploy version 1 of all services"
  echo "2) Deploy version 2 of 2 services"
  echo "3) Routing A/B"
  echo "4) Deploy all services version 2"
  echo "5) Rollback"
  echo "t) Test"
  echo "x) Exit"
}

clear
while true; do
  display_menu
  read -p "Enter your choice: " choice
  case $choice in
    1)
      part1
      clear
      echo "Deployed version 1 of all services."
      echo "HPA will be ready in a few seconds."
      sleep 15
      clear
      ;;
    2)
      part2
      clear
      echo "Deployed version 2 of 2 services."
      echo "HPA will be ready in a few seconds."
      sleep 15
      clear
      ;;
    3)
      part3
      clear
      echo "Enabled A/B routing."
      sleep 5
      clear
      ;;
    4)
      part4
      clear
      echo "Deployed version 2 of all services."
      sleep 5
      clear
      ;;
    5)
      part5
      clear
      echo "Rollback to version 1."
      sleep 5
      clear
      ;;
    t)
      sh ./test.sh
      clear
      ;;
    x)
      echo "Exiting..."
      exit 0
      ;;
    *)
      echo "Invalid choice. Please select a valid option."
      ;;
  esac
done
