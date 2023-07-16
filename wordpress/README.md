## Deploying WordPress and MySQL

The following sample is for deploying a WordPress site and a MySQL database to demonstrate service descovery with DNS

This sample is simplified version from the [official kubernetes sample](https://kubernetes.io/docs/tutorials/stateful-application/mysql-wordpress-persistent-volume/) without the PVC.

The full version could be found in [../wordpress-pvc]
The alternative version with EIP could be found in [../wordpress-eip]


## Test It Out

1. Ensure PrivateZone service is enabled (https://dns.console.aliyun.com/#/privateZone/list)
2. Create Serverless Kubernetes cluster, and check the option for "Using PrivateZone for service descovery" 
3. Modify the ```kustomization.yaml``` file with your password

4. Deploy application


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
kubectl get service -l app=wordpress
```

You will see the service result as following

```
$ kubectl get service -l app=wordpress
NAME              TYPE           CLUSTER-IP   EXTERNAL-IP      PORT(S)        AGE
wordpress         LoadBalancer   <none>       xxx.xxx.xxx.xxx  80:30987/TCP   2m
wordpress-mysql   ClusterIP      <none>       <none>           3306/TCP       2m
```

Access the WordPress application

```
LB_ENDPOINT=$(kubectl get service wordpress  -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Open browser with URL in MacOSX
open http://$LB_ENDPOINT
```

5. Delete deployments and services

```
kubectl delete -k ./
```
