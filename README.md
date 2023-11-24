# OpenLoft

[![License](https://img.shields.io/github/license/openloft/openloft?logo=github)](https://opensource.org/license/mit/) [![Makefile CI](https://github.com/openloft/openloft/actions/workflows/makefile.yml/badge.svg)](https://github.com/openloft/openloft/actions/workflows/makefile.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/openloft/openloft)](https://goreportcard.com/report/github.com/openloft/openloft) <a href="https://github.com/openloft/openloft/graphs/contributors" alt="Contributors"><img src="https://img.shields.io/github/contributors/openloft/openloft" /></a> <img alt="GitHub last commit (branch)" src="https://img.shields.io/github/last-commit/openloft/openloft/main">

----

## Quick Start (~ 1 minute)

### 1. Prerequisites

* Setup a Kubernetes cluster with [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation).

  ```bash
  git clone https://github.com/openloft/kind.git
  cd kind
  # Create a cluster with 1 control-plane node and 2 worker nodes,
  # install nginx ingress controller and
  # deploy fake device plugin to provision 8 NVIDIA GPUs on each worker node.
  just a
  ```

* Deploy [vCluster Operator](https://github.com/openloft/vcluster-operator).

  ```bash
  kubectl apply -f https://raw.githubusercontent.com/openloft/vcluster-operator/main/deploy/manifests.yaml
  ```

### 2. Deploy OpenLoft

```bash
kubectl apply -f https://raw.githubusercontent.com/openloft/openloft/main/deploy/manifests.yaml
```

### 3. Create a vCluster Template

There are two types of vCluster templates in sample, `isolated` and `default`. The `isolated` template will create a vCluster with a [resource quota](https://kubernetes.io/docs/concepts/policy/resource-quotas/). The `default` template will create a vCluster with no resource quota.

The below command will create a `isolated` vCluster template.

```bash
kubectl apply -f https://raw.githubusercontent.com/openloft/openloft/main/config/samples/isolated/storage_v1_virtualclustertemplate.yaml
```

### 4. Create a vCluster Instance

```bash
kubectl apply -f https://raw.githubusercontent.com/openloft/openloft/main/config/samples/isolated/storage_v1_virtualclusterinstance.yaml
```

By running the above command, a vCluster called `isolated-sample-openloft-cn` would be created in the namespace `vc-isolated-sample-openloft-cn`. And an ingress would be created with the name `ingress-isolated-sample-openloft-cn`.

### 5. Retrieving the kube config from the vCluster secret

The secret is prefixed with `vc-` and ends with the vCluster name, so a vCluster instance called `isolated-sample.openloft.cn` would create a secret called `vc-isolated-sample-openloft-cn` in the namespace `vc-isolated-sample-openloft-cn`. You can retrieve the kube config after the vCluster has started via:

```
kubectl get secret vc-isolated-sample-openloft-cn -n vc-isolated-sample-openloft-cn --template={{.data.config}} | base64 -D > kubeconfig.yaml
```

The secret will hold a kube config in this format:

```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0t...
        server: https://isolated-sample.openloft.cn
  name: my-vcluster
contexts:
- context:
    cluster: my-vcluster
    user: my-vcluster
  name: my-vcluster
current-context: my-vcluster
kind: Config
preferences: {}
users:
- name: my-vcluster
  user:
    client-certificate-data: LS0tLS...
    client-key-data: LS0tLS...
```

To be able to access the vCluster, update the `hosts` file with `127.0.0.1` pointing to the vCluster ingress:

e.g. `/etc/hosts` on macOS/Linux:

```
127.0.0.1 isolated-sample.openloft.cn
```

### 6. Accessing the vCluster

You can access the vCluster via:

```bash
kubectl --kubeconfig kubeconfig.yaml get ns
```

## Contributing

* [Open an issue](https://github.com/openloft/openloft/issues/new)
* [Submit a pull request](https://github.com/openloft/openloft/compare)