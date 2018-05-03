## Spark standalone cluster running over serverless kubernetes

The following document describes the deployment of a Spark Standalone Cluster .


## Test It Out

Deploy application

```
kubectl apply -f spark-standalone.yml
```

Check status of the deployments/pods/services:

```
kubectl get deployment
kubectl get pod | grep myspark
kubectl get service myspark-master-webui
```

Access the sample application

```
LB_ENDPOINT=$(kubectl get service myspark-master-webui  -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Open browser with URL in MacOSX
curl http://${LB_ENDPOINT}:8080
open http://${LB_ENDPOINT}:8080
```

Delete application


```
kubectl delete -f spark-standalone.yml
```
