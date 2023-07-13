## Deploying Selenium Grid for Web Testing

The following sample is for deploying a Selenium Grid with Chrome/Firefox nodes. More details for Selenium Docker images could be found in https://github.com/SeleniumHQ/docker-selenium

## Test It Out

1. Ensure PrivateZone service is enabled (https://dns.console.aliyun.com/#/privateZone/list)
2. Create Serverless Kubernetes cluster, and check the option for "Using PrivateZone for service descovery" 
3. Deploy application

```
kubectl apply -f selenium-hub.yaml
kubectl apply -f selenium-hub-svc.yaml
kubectl apply -f selenium-node-chrome.yaml
kubectl apply -f selenium-node-firefox.yaml
```

Check on the status of the pod using this command: 

```
kubectl get all
```

You will see the service result as following

```
$ kubectl get all
NAME                           DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/selenium-hub            1         1         1            1           8m
deploy/selenium-node-chrome    1         1         1            1           7m
deploy/selenium-node-firefox   1         1         1            1           7m

NAME                                  DESIRED   CURRENT   READY     AGE
rs/selenium-hub-5dc97596bf            1         1         1         8m
rs/selenium-node-chrome-784b6fb878    1         1         1         7m
rs/selenium-node-firefox-6579d6cfdf   1         1         1         7m

NAME                           DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/selenium-hub            1         1         1            1           8m
deploy/selenium-node-chrome    1         1         1            1           7m
deploy/selenium-node-firefox   1         1         1            1           7m

NAME                                  DESIRED   CURRENT   READY     AGE
rs/selenium-hub-5dc97596bf            1         1         1         8m
rs/selenium-node-chrome-784b6fb878    1         1         1         7m
rs/selenium-node-firefox-6579d6cfdf   1         1         1         7m

NAME                                        READY     STATUS    RESTARTS   AGE
po/selenium-hub-5dc97596bf-6vwkr            1/1       Running   0          8m
po/selenium-node-chrome-784b6fb878-9rl5g    1/1       Running   0          7m
po/selenium-node-firefox-6579d6cfdf-8fplm   1/1       Running   0          7m

NAME                  TYPE           CLUSTER-IP   EXTERNAL-IP      PORT(S)          AGE
svc/selenium-hub      ClusterIP      <none>       <none>           4444/TCP         7m
svc/selenium-hub-lb   LoadBalancer   <none>       139.224.218.84   4444:30420/TCP   7m

```

Access the Selenium Hub

```
LB_ENDPOINT=$(kubectl get service selenium-hub-lb  -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Open browser with URL in MacOSX
open http://$LB_ENDPOINT:4444/grid/console
```

Delete deployments and services

```
kubectl delete -f selenium-node-firefox.yaml
kubectl delete -f selenium-node-chrome.yaml
kubectl delete -f selenium-hub-svc.yaml
kubectl delete -f selenium-hub.yaml
```

