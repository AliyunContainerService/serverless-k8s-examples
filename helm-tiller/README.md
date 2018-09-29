## Using Helm in serverless kubernetes

## Install Tiller in your serverless cluster

Deploy tiller server

```
# kubectl apply -f tiller.yaml
```

Check status of the tiller pod:

```
# kubectl get pod
NAME                             READY     STATUS    RESTARTS   AGE
tiller-deploy-5f5c88db48-555nf   1/1       Running   0          24m
```

## Use Helm
```
# helm ls

# helm install alpine
NAME:   your-pike
LAST DEPLOYED: Sat Sep 29 07:24:59 2018
NAMESPACE: viking-c14f93dcb603f496c887ae05798046e4f
STATUS: DEPLOYED

RESOURCES:
==> v1/Pod
NAME              READY  STATUS   RESTARTS  AGE
your-pike-alpine  0/1    Pending  0         1s

# helm ls
NAME     	REVISION	UPDATED                 	STATUS  	CHART       	NAMESPACE
your-pike	1       	Sat Sep 29 07:24:59 2018	DEPLOYED	alpine-0.1.0	viking-c14f93dcb603f496c887ae05798046e4f

# kubectl get pod
NAME                             READY     STATUS    RESTARTS   AGE
tiller-deploy-5f5c88db48-555nf   1/1       Running   0          30m
your-pike-alpine                 1/1       Running   0          1m

# kubectl get configmap
NAME           DATA      AGE
your-pike.v1   1         2m

# helm del --purge your-pike
release "your-pike" deleted
```

## Delete Tiller Server

```
kubectl delete -f tiller.yaml
```

### Limitation
Please mind that serverless kubernetes's helm only support create resources in one namespace. Namespace and CRD are not allowed to create.

