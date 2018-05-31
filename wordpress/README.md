## Deploying WordPress and MySQL

The following sample is for deploying a WordPress site and a MySQL database to demonstrate service descovery with DNS

## Test It Out

1. Ensure PrivateZone service is enabled (https://dns.console.aliyun.com/#/privateZone/list)
2. Create Serverless Kubernetes cluster, and check the option for "Using PrivateZone for service descovery" 
3. Deploy application

```
kubectl create -f ./wordpress-mysql.yaml
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
open http://${LB_ENDPOINT}
```

Delete deployments and services

```
kubectl delete -f nginx.yaml
```

