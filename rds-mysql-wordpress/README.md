## WordPress Application using Aliyun RDS

The following document describes the deployment of a wordpress application using aliyun rds.

## Test It Out

1. Create a RDS mysql instance in aliyun web console
Currently mysql version 5.5 and 5.6 is recommended.

2. Configure security group whitelist for your RDS instance
Get security group id from serverless kubernetes cluster info page, and set it
as whitelist for your RDS instance.

3. Configure database password for your RDS instance

4. Update wordpress yaml file with RDS host and password

5. Deploy wordpress application

```
kubectl apply -f wordpress.yaml
```

Check status of the deployments/pods/services:

```
kubectl get deployment wordpress
kubectl get pod -l app=wordpress
kubectl get service wordpress-svc
```

Access the sample application

```
LB_ENDPOINT=$(kubectl get service wordpress-svc  -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Open browser with URL in MacOSX
open http://${LB_ENDPOINT}:80/sample/
```

Delete application


```
kubectl delete -f wordpress.yaml
```
