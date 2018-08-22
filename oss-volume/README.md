## Aliyun OSS Volume Sample

The following sample is for running a simple application using Aliyun OSS Volume

Serverless Kubernetes's OSS Volume is based on flexvolume mechanism.

## Test It Out

1. Create a OSS volume by Aliyun web console

2. Deploy application

```
kubectl create -f ./oss-volume.yaml
```

Check on the OSS directory in pod using this command: 

```
kubectl exec -it ubuntu-pod ls /oss
```

Delete pod

```
kubectl delete -f oss-volume.yaml
```

