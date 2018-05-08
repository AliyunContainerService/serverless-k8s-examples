## Aliyun NAS Sample

The following sample is for running a simple application using Aliyun NAS

## Test It Out

1. Create a NAS volume by Aliyun web console
Note that the vpc of the nas volume should be same as your serverless cluster.

2. Deploy application

```
kubectl create -f ./nas-volume.yaml
```

Check on the status of the pod using this command: 

```
kubectl get pod ubuntu
kubectl logs ubuntu
```

Delete pod

```
kubectl delete -f nas-volume.yaml
```

