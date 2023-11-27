# MPI Operator

```bash
kubectl get secret vc-isolated-sample-openloft-cn -n vc-isolated-sample-openloft-cn --template={{.data.config}} | base64 -D > kubeconfig.yaml
kubectl --kubeconfig=kubeconfig.yaml apply -f https://raw.githubusercontent.com/kubeflow/mpi-operator/v0.4.0/deploy/v2beta1/mpi-operator.yaml
# https://github.com/kubeflow/mpi-operator/blob/master/examples/v2beta1/tensorflow-benchmarks/tensorflow-benchmarks.yaml
# XXX: change image/command and resource limits
kubectl --kubeconfig=kubeconfig.yaml apply -f tensorflow-benchmarks.yaml
```