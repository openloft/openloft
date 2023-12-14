#!/usr/bin/env sh

kubectl get secret vc-isolated-sample -n vc-isolated-sample --template={{.data.config}} | base64 -D > kubeconfig.yaml
export KUBECONFIG=$PWD/kubeconfig.yaml
helm repo add harbor https://helm.goharbor.io
helm install apps harbor/harbor