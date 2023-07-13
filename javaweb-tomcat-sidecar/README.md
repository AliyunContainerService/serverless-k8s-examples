## Java Web Application with Tomcat and Sidecar Container

The following document describes the deployment of a Java Web application using Tomcat. Instead of packaging `war` file inside the Tomcat image or mount the `war` as a volume, we use a sidecar container as `war` file provider.

The orginal example is from Kubernetes examples project

https://github.com/kubernetes/examples/blob/master/staging/javaweb-tomcat-sidecar/

The details for Pod with Sidecar

https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/

## Test It Out

Deploy application

```
kubectl apply -f javaweb.yaml
```

Check status of the deployments/pods/services:

```
kubectl get deployment tomcat-app
kubectl get pod -l app=tomcat-app
kubectl get service tomcat-app-svc
```

Access the sample application

```
LB_ENDPOINT=$(kubectl get service tomcat-app-svc  -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Open browser with URL in MacOSX
open http://$LB_ENDPOINT:8080/sample/
```

Delete application


```
kubectl delete -f javaweb.yaml
```
