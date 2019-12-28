## Serverless Kubernetes Examples

[Serverless Kubernetes](https://www.aliyun.com/product/kubernetes) is part of Container Service for Alibaba Cloud. It enable you to run Kubernetes application without effort for managing servers or clusters. Serverless Kubernetes lets you focus on building your applications instead of managing the infrastructure.

This directory contains a number of examples of how to run real applications with Serverless Kubernetes of Alibaba Cloud


## Quick Start

Create the Serverless Kubernetes and copy the cluster config file to  ```~/.kube/config```

![cluster](./cluster.png)



![config](./config.png)


## Create NAT Gateway (Suggested)

Create NAT Gateway if you want to pull image from internet (e.g. Docker Hub), or your applications want to access internet.

![snat](./SNAT.png)

## Test It Out

Deploy nginx application


```
# Pull image from internet public URI with NAT GW
kubectl run nginx --image nginx:1.13 --replicas=3

# or pull image from Aliyun Container Registry (ACR) through VPC internal URI
kubectl run nginx --image registry-vpc.cn-shanghai.aliyuncs.com/denverdino/nginx:1.13.12 --replicas=3

```

Expose nginx with Elastic Load Balancer(ELB) service 

```
kubectl expose deployment nginx --port=80 --target-port=80 --name=nginx-svc --type=LoadBalancer
```


Get the application status

```
kubectl get deployment nginx
kubectl get pod -l run=nginx
kubectl get service nginx-svc
```

Access nginx application

```
LB_ENDPOINT=$(kubectl get service nginx-svc -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Open browser with URL in MacOSX
open http://${LB_ENDPOINT}
```


Delete the nginx application

```
kubectl delete deployment nginx
```

Delete the nginx service

```
kubectl delete service nginx-svc
```
