## Knative on ASK (Developer Preview)

The following sample is for deploying a Knative application on Serverless Kubernetes Cluster

## Test It Out

#### Create Serverless Kubernetes cluster

#### Enable Knative

Please join the Knative DingTalk Group, and send your request to admin with your cluster id
Or, open the ticket for that

#### Ensure Knative is enabled

Execute the following commands and checking if the Knative is enabled

```
$ kubectl get crd | grep services.serving.knative.dev
services.serving.knative.dev                         2020-07-14T03:14:30Z
```

Get the "ingress-gateway" endpoint for Knative

```
$ kubectl -n knative-serving get svc
NAME              TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)        AGE
ingress-gateway   LoadBalancer   172.19.10.249   xx.xx.xx.xx     80:31305/TCP   26m
```

#### Deploy Knative application


The simple sample application for Knative serving

```
$ cat coffee.yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: coffee
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/target: "10"
    spec:
      containers:
        - image: registry.cn-hangzhou.aliyuncs.com/knative-sample/helloworld-go:160e4dc8
          env:
            - name: TARGET
              value: "coffee"
```


Deploy application 


```
kubectl apply -f ./coffee.yaml
```

Check on the status of the pod and Knative Service using following commands:

```
$ kubectl get pod

NAME                                       READY   STATUS    RESTARTS   AGE
coffee-x9fhg-deployment-85fd649c7d-xvw5j   2/2     Running   0          43s

$ kubectl get ksvc

NAME     URL                                 LATESTCREATED   LATESTREADY    READY   REASON
coffee   http://coffee.default.example.com   coffee-x9fhg    coffee-x9fhg   True
```

Access the sample application

```
$ LB_ENDPOINT=$(kubectl get service ingress-gateway -n knative-serving -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Curl URL with host name
$ curl -H "Host: coffee.default.example.com" http://${LB_ENDPOINT}
Hello coffee!
```

Test autoscaling for the sample application

**Prerequisite** Install [hey](https://github.com/rakyll/hey) for performance testing


Run the load test with following configurations

* -z 30s # Duration in 30 seconds
* -c 90 # with 90 workers


```
$ hey -z 30s -c 90 --host "coffee.default.example.com" "http://${LB_ENDPOINT}/?sleep=100"

Summary:
  Total:	30.1015 secs
  Slowest:	0.5165 secs
  Fastest:	0.0086 secs
  Average:	0.1250 secs
  Requests/sec:	720.0644

  Total data:	303450 bytes
  Size/request:	14 bytes

Response time histogram:
  0.009 [1]	|
  0.059 [2213]	|■■■■■■■■
  0.110 [10912]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.161 [3902]	|■■■■■■■■■■■■■■
  0.212 [1341]	|■■■■■
  0.263 [1232]	|■■■■■
  0.313 [1206]	|■■■■
  0.364 [584]	|■■
  0.415 [214]	|■
  0.466 [57]	|
  0.516 [13]	|


Latency distribution:
  10% in 0.0588 secs
  25% in 0.0765 secs
  50% in 0.0973 secs
  75% in 0.1435 secs
  90% in 0.2583 secs
  95% in 0.2996 secs
  99% in 0.3666 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0002 secs, 0.0086 secs, 0.5165 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0000 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.1246 secs, 0.0086 secs, 0.5164 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0025 secs

Status code distribution:
  [200]	21675 responses
```

And you can watch the replica changes with following commands

```
$ kubectl get pod -w
NAME                                               READY   STATUS    RESTARTS   AGE
coffee-x9fhg-deployment-reserve-679896b76c-bj469   2/2     Running   0          2m32s
coffee-x9fhg-deployment-85fd649c7d-p66h5           0/2     Pending   0          0s
coffee-x9fhg-deployment-reserve-679896b76c-bj469   2/2     Running   0          2m45s
coffee-x9fhg-deployment-85fd649c7d-mx2vd           0/2     Pending   0          0s
coffee-x9fhg-deployment-85fd649c7d-2kbz8           0/2     Pending   0          0s
coffee-x9fhg-deployment-85fd649c7d-2c9l4           0/2     Pending   0          0s
coffee-x9fhg-deployment-85fd649c7d-x2x72           0/2     Pending   0          0s
coffee-x9fhg-deployment-85fd649c7d-k7n2c           0/2     Pending   0          0s
coffee-x9fhg-deployment-85fd649c7d-btrf6           0/2     Pending   0          0s
coffee-x9fhg-deployment-85fd649c7d-2884f           0/2     Pending   0          0s
coffee-x9fhg-deployment-85fd649c7d-p66h5           0/2     Pending   0          11s
```



#### Delete the application

```
$ kubectl delete -f coffee.yaml
```

