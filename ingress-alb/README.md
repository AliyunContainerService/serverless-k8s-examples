## Deploying Ingress Demo

In this example we deploy a simple web application and then configure Application Load Balancer for that application using the Ingress resource.

## Prerequisites

You need to create the Serverless K8s cluster in the available region and AZs for ALB
https://help.aliyun.com/document_detail/258300.html

## Test It Out

Install and config the ALB Ingress Controller
https://www.alibabacloud.com/help/zh/ack/serverless-kubernetes/user-guide/manage-the-alb-ingress-controller


Deploy the Cafe Application.

Create the coffee and the tea deployments, services, and ingress:

```
kubectl apply -f ./ingress-cafe-demo.yaml
```

Check on the status of the pod using this command: 

```
kubectl get pod
kubectl get svc
kubectl get ingress
```

You will see the service and ingress resources as following

```
$ kubectl get pod
NAME                     READY     STATUS    RESTARTS   AGE
coffee-f5cd54465-f82f5   1/1       Running   0          16m
coffee-f5cd54465-gm8qm   1/1       Running   0          16m
tea-6bcb468bfc-76fsd     1/1       Running   0          16m
tea-6bcb468bfc-jkjzw     1/1       Running   0          16m
tea-6bcb468bfc-mnw7h     1/1       Running   0          16m

$ kubectl get svc --all-namespaces
NAMESPACE   NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
default     coffee-svc   ClusterIP   None         <none>        80/TCP    98m
default     kubernetes   ClusterIP   10.0.0.1     <none>        443/TCP   3d8h
default     tea-svc      ClusterIP   None         <none>        80/TCP    98m
...

$ kubectl get ingress
NAME           CLASS   HOSTS   ADDRESS                                                  PORTS   AGE
cafe-ingress   alb     *       alb-xxxx.cn-zhangjiakou.alb.aliyuncs.com   80      99m
```



To get coffee:
```

INGRESS_ADDRESS=$(kubectl get ingress cafe-ingress -o jsonpath="{.status.loadBalancer.ingress[0].hostname}")

$ curl $INGRESS_ADDRESS/coffee
Server address: 192.168.54.127:80
Server name: coffee-7f99b48c49-mw9h8
Date: 13/Jul/2023:15:07:14 +0000
URI: /coffee
Request ID: e635ce56f1188568ceadddc28b1caace
```

If your rather prefer tea:
```
$ curl $INGRESS_ADDRESS/tea
Server address: 192.168.54.127:80
Server name: coffee-7f99b48c49-mw9h8
Date: 13/Jul/2023:15:07:53 +0000
URI: /tea
Request ID: 9b739572ac6d4f0d6e4a4e9268a6ff78
```

Delete deployments, services, and ingress

```
kubectl delete -f ./ingress-cafe-demo.yaml
```

