## Pulling Private Image Sample

The following document demonstrates how to pull private docker image stored in Aliyun Container Registry.

Aliyun Serverless Kubernetes already fully integrated with Aliyun Container Registry, that let you pull private image without handle secret manually.

## Test It Out

1. Push your private docker image into Aliyun Container Registry
Choose image url in following format:
 vpc access:    registry-vpc.region-id.aliyuncs.com/namespace/image:tag
 public access: registry.region-id.aliyuncs.com/namespace/image:tag

2. Use your private image yaml file
Note that you only can use private image in your account.

3. Deploy application

```
kubectl create -f ./private-image-pulling.yaml
```

Check on the status of the pod using this command: 

```
kubectl get pod test-pod
kubectl logs test-pod
```

Delete pod

```
kubectl delete -f ./private-image-pulling.yaml
```

