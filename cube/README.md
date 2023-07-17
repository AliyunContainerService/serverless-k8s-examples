## Deploying a sample application for Cube

The following sample is for deploying a Cube application.


## Test It Out

1. Create Serverless Kubernetes cluster

2. Modify the ```cube.yaml``` file with the image URL in your region. 
  
    * Hangzhou region: registry-vpc.cn-hangzhou.aliyuncs.com/acs/ack-cube
    * Beijing region: registry-vpc.cn-beijing.aliyuncs.com/acs/ack-cube
    * Shenzhen region: registry-vpc.cn-shenzhen.aliyuncs.com/acs/ack-cub

3. Deploy application


```
kubectl apply -f cube.yaml
```



Check on the status of the pod using this command: 

```
kubectl get pod -l app=cube
kubectl get service -l app=cube
```

You will see the service result as following

```
$ kubectl get pod -l app=cube

NAME                    READY   STATUS    RESTARTS   AGE
cube-5695d76876-pnbjv   1/1     Running   0          10m

$ kubectl get service -l app=cube

NAME           TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)        AGE
cube-service   LoadBalancer   172.16.99.92   xxx.xxx.xx.xx   80:30170/TCP   10m
```

Access the Cube application

```
LB_ENDPOINT=$(kubectl get service cube-service  -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Open browser with URL in MacOSX
open http://$LB_ENDPOINT
```

4. Delete deployments and services

```
kubectl delete -f ./cube.yaml
```
