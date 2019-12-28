## Deploying Ingress Demo

In this example we deploy a simple web application and then configure load balancing for that application using the Nginx-Ingress resource.

## Deploy nginx-ingress-controller

Firstly we need to deploy nginx-ingress-controller in Serverless Kubernetes cluster.

```
$ kubectl create -f ./nginx-ingress-controller.yaml
...

$ kubectl -n kube-system get deploy|grep nginx
nginx-ingress-controller   2         2         2            2           40s

$ kubectl -n kube-system get svc|grep nginx
nginx-ingress-lb   LoadBalancer   ****   ****   80:32665/TCP,443:30338/TCP   1m
```


## Test It Out

Deploy the Cafe Application.

Create the coffee and the tea deployments, services, and ingress:

```
kubectl create -f ./ingress-cafe-demo.yaml
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
NAMESPACE     NAME               TYPE           CLUSTER-IP      EXTERNAL-IP      PORT(S)                      AGE
default       coffee-svc         ClusterIP      None            <none>           80/TCP                       7m7s
default       kubernetes         ClusterIP      172.19.0.1      <none>           443/TCP                      80m
default       tea-svc            ClusterIP      None            <none>           80/TCP                       7m7s
kube-system   nginx-ingress-lb   LoadBalancer   172.19.4.82     139.196.2.186    80:30848/TCP,443:32167/TCP   9m4s
...

$ kubectl get ingress
NAME           HOSTS              ADDRESS         PORTS     AGE
cafe-ingress   cafe.example.com   139.196.2.186   80        16m
```

To get coffee:
```
$ curl -H "Host:cafe.example.com" 139.196.2.186/coffee
Server address: 192.168.133.107:80
Server name: coffee-f5cd54465-f82f5
Date: 28/Jun/2018:11:49:30 +0000
URI: /coffee
Request ID: 2b81fbb5ba3e22a0ae7eb5f1806d4ce2
```

If your rather prefer tea:
```
$ curl -H "Host:cafe.example.com" 139.196.2.186/tea
Server address: 192.168.133.109:80
Server name: tea-6bcb468bfc-mnw7h
Date: 28/Jun/2018:11:50:17 +0000
URI: /tea
Request ID: 517c32f5dccc6ab88e4593f7c0ef00d5
```

Delete deployments, services, and ingress

```
kubectl delete -f ./ingress-cafe-demo.yaml
kubectl delete -f ./nginx-ingress-controller.yaml
```

