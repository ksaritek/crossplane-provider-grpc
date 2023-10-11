# Crossplane sample that talks through Grpc Server

That project is an example for crossplane provider. Provider talks to grpc server to do CRUD operation

## prepare local k8s cluster

### kind

```bash
kind create cluster --name k8s-crossplane
kind get clusters
kubectl cluster-info --context kind-k8s-crossplane
## export kubeconfig
export KUBECONFIG="$(kind get kubeconfig --name='k8s-crossplane')"
```

### crossplane install

```bash
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update
helm install crossplane \
--namespace crossplane-system \
--create-namespace crossplane-stable/crossplane
```

!!! if you run controller as an image instead of `make run`, use `host.docker.internal` to reach grpc-server running locally from local k8s cluster https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host

## clone [provider userprovider](https://github.com/crossplane/provider-userprovider) and prepare provider

```bash
# fetch the [upbound/build] submodule by running the following
make submodules

# update template with your provider name
export provider_name=UserProvider
make provider.prepare provider=${provider_name}

# add User kind under playground group
export group=playground
export type=User
make provider.addtype provider=${provider_name} group=${group} kind=${type}

# register new APIs and Controllers
## 1. apis/userprovider.go
## 2. internal/controller/userprovider.go

# generate the code
make generate
```
