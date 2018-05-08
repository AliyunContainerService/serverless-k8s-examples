## Basic Test Sample

The following document demonstrates the basic pod functionality like secret/configmap/volumes/initContainer/CommandArgs/Envs

## Test It Out

Deploy application

```
kubectl create -f ./basic-pod.yaml
```

Check on the status of the pod using this command: 

```
kubectl get pod basic-pod
kubectl logs basic-pod
```

Delete pod

```
kubectl delete -f ./basic-pod.yaml
```

