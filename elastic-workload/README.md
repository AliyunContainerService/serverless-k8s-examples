## Elastic Workload Sample(Deploy application in ECI On-demand and Spot mixed mode)

The following sample is for running a simple application to demonstrate elastic workload.
In this sample, we can make 2 replicas of a deployment to be on-demand instances, and the others are eci-spot instances.

Based on Elastic Workload, we can reduce application cost significantly.

## Elastic Workload Introduction
https://yq.aliyun.com/articles/759290

## Deploy ElasticWorkload Controller

Install ElasticWorkload Controller: https://cs.console.aliyun.com/#/k8s/catalog/detail/incubator_ack-kubernetes-elastic-workload


## Create Elastic Workload Application
```
# kubectl apply -f nginx-deploy-elasticworkload.yaml
deployment.extensions/nginx-deploy created
elasticworkload.autoscaling.alibabacloud.com/elasticworkload-nginx created

# kubectl get pod
NAME                                          READY   STATUS    RESTARTS   AGE
nginx-deploy-55d8dcf755-945t4                 1/1     Running   0          52s
nginx-deploy-55d8dcf755-gtjwl                 1/1     Running   0          52s
nginx-deploy-unit-eci-spot-5ffcb58d75-4g2kh   1/1     Running   0          52s
nginx-deploy-unit-eci-spot-5ffcb58d75-g6ztd   1/1     Running   0          52s
nginx-deploy-unit-eci-spot-5ffcb58d75-k5jx2   1/1     Running   0          52s
nginx-deploy-unit-eci-spot-5ffcb58d75-xtrkh   1/1     Running   0          52s
```

## Scale Elastic Workload Application
```
# kubectl get pod
NAME                                          READY   STATUS    RESTARTS   AGE
nginx-deploy-55d8dcf755-945t4                 1/1     Running   0          12m
nginx-deploy-55d8dcf755-gtjwl                 1/1     Running   0          12m
nginx-deploy-unit-eci-spot-5ffcb58d75-4g2kh   1/1     Running   0          12m
nginx-deploy-unit-eci-spot-5ffcb58d75-9m65d   1/1     Running   0          5m34s
nginx-deploy-unit-eci-spot-5ffcb58d75-fdl6g   1/1     Running   0          5m34s
nginx-deploy-unit-eci-spot-5ffcb58d75-fq5cr   1/1     Running   0          5m34s
nginx-deploy-unit-eci-spot-5ffcb58d75-g6ztd   1/1     Running   0          12m
nginx-deploy-unit-eci-spot-5ffcb58d75-k5jx2   1/1     Running   0          12m
nginx-deploy-unit-eci-spot-5ffcb58d75-qlv52   1/1     Running   0          5m34s
nginx-deploy-unit-eci-spot-5ffcb58d75-xtrkh   1/1     Running   0          12m
```

