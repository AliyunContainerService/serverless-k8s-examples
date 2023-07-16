## Deploying WordPress and MySQL

The following sample is for deploying a WordPress site and a MySQL database to demonstrate service descovery with DNS.

The wordpress application can be accessed with [EIP endpoint](https://help.aliyun.com/document_detail/451273.html).

Before the test, you need configure the Security Group of ASK cluster to allow incoming traffic on 80 port. 

## Test It Out

1. Ensure PrivateZone service is enabled (https://dns.console.aliyun.com/#/privateZone/list)
2. Create Serverless Kubernetes cluster, and check the option for "Using PrivateZone for service descovery" 
3. Deploy application

```
kubectl apply -k ./
```

Verify that the Secret exists by running the following command:

```
kubectl get secrets
```

Check on the status of the pod using this command: 

```
kubectl get pod
kubectl get pod -l app=wordpress
```

Access the WordPress application

```
ENDPOINT=$(kubectl get pod -l app=wordpress -l tier=frontend -o jsonpath="{.items[0].metadata.annotations.k8s\.aliyun\.com/allocated-eipAddress}")

echo $ENDPOINT

# Open browser with URL in MacOSX
open http://$ENDPOINT
```

5. Delete deployments and services

```
kubectl delete -k ./
```
