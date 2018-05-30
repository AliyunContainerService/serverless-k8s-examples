## Service Descovery Sample

The following sample is for running a simple nginx application to demonstrate service descovery usage

## Test It Out

1. Ensure PrivateZone product enabled (https://dns.console.aliyun.com/#/privateZone/list)
2. Create Serverless cluster, and enable "PrivateZone Service Descovery" in console.
3. Deploy application

```
kubectl create -f ./ngix.yaml
```

Check on the status of the pod using this command: 

```
kubectl get pod
kubectl logs $nginx-pod -c outbound
```

In same vpc, we can access the services by several domain names:
```
ping nginx-service-headless.$NAMESPACE.svc.cluster.local
ping nginx-service-headless

ping nginx-service-intranet.$NAMESPACE.svc.cluster.local
ping nginx-service-intranet
```

Delete deployment and services

```
kubectl delete -f nginx.yaml
```

