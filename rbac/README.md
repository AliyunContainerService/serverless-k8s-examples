## RBAC Sample

The following sample is for running a simple application to demonstrate rbac usage.

The pod contains two containers:

1. in-cluster container:  access apiserver by kubernetes go api. See source code in-cluster.go.
2. ubuntu container:  access apiserver by curl command line.

## Test It Out

1. Change the apiserver address in yaml file

2. Deploy application

```
kubectl create -f ./in-cluster.yaml
```

Check on the pod status and log output of each container: 

```
kubectl get pod
NAME                          READY     STATUS    RESTARTS   AGE
in-cluster-7dc6bcf878-m2bwk   2/2       Running   0          1m

kubectl logs -l app=in-cluster -c in-cluster
There are 1 pods in the cluster: in-cluster-7dc6bcf878-m2bwk
There are 1 pods in the cluster: in-cluster-7dc6bcf878-m2bwk
There are 1 pods in the cluster: in-cluster-7dc6bcf878-m2bwk
...


kubectl logs -l app=in-cluster -c ubuntu
{
  "kind": "PodList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/namespaces/viking-6/pods",
    "resourceVersion": "16331994"
  },
  "items": [
    {
      "metadata": {
        "name": "in-cluster-7dc6bcf878-m2bwk",
        "generateName": "in-cluster-7dc6bcf878-",
...
```

Delete deployment and serviceAccount

```
kubectl delete -f ./in-cluster.yaml
```

