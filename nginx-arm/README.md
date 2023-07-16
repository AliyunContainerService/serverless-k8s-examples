## Deploying Arm-based Nginx Demo

In this example we deploy a nginx application with multi-arch image on Arm-based instance.


## Prerequisites

And you can found the available region and AZs for Arm instance
https://ecs-buy.aliyun.com/instanceTypes#/instanceTypeByRegion

## Deploy 

1. Enbale Arm64-based deployment

```
kubectl patch configmap eci-profile -n kube-system --type='json' -p='[{"op": "replace", "path": "/data/enableLinuxArm64Node", "value":"true"}]'
```

And you can verify the enableLinuxArm64Node value is patched correctly with following command.

```
kubectl get configmap eci-profile -n kube-system -o yaml
```

2. Deploy test application

```
kubectl apply -f ./nginx.yaml
```

3. Verify test application

```
$ kubectl get pod
NAME                        READY   STATUS    RESTARTS   AGE
nginx-arm-c4c5b795f-gjtjl   1/1     Running   0          32m
```

4. Test It Out

```
$ kubectl exec nginx-arm-c4c5b795f-gjtjl -ti -- uname -m
aarch64
```

## Delete deployments

```
kubectl delete -f ./nginx.yaml
```