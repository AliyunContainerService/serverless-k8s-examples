## Deploying WordPress and MySQL with PVC

The following sample is for deploying a WordPress site and a MySQL database with PVC

More details for the sample is as following
https://kubernetes.io/docs/tutorials/stateful-application/mysql-wordpress-persistent-volume/


## Test It Out

1. Ensure PrivateZone service is enabled (https://dns.console.aliyun.com/#/privateZone/list)
2. Create Serverless Kubernetes cluster, and check the option for "Using PrivateZone for service descovery" 
3. Install and update csi-provisioner addon to Serverless Kubernetes cluster, https://www.alibabacloud.com/help/zh/ack/serverless-kubernetes/user-guide/install-and-update-csi-provisioner 
4. Set default storage class

```
kubectl patch storageclass alicloud-disk-available -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
```

5. Modify the ```kustomization.yaml``` file with your password

6. Deploy application


```
kubectl apply -k ./
```

Verify that the Secret exists by running the following command:

```
kubectl get secrets
```

Verify that a PersistentVolume got dynamically provisioned:

```
kubectl get pvc
```

Check on the status of the pod using this command: 

```
kubectl get pod
kubectl get pod -l app=wordpress
kubectl get service -l app=wordpress
```

You will see the service result as following

```
$ kubectl get service -l app=wordpress
NAME              TYPE           CLUSTER-IP   EXTERNAL-IP      PORT(S)        AGE
wordpress         LoadBalancer   <none>       xxx.xxx.xxx.xxx  80:30987/TCP   2m
wordpress-mysql   ClusterIP      <none>       <none>           3306/TCP       2m
```

Access the WordPress application

```
LB_ENDPOINT=$(kubectl get service wordpress  -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Open browser with URL in MacOSX
open http://$LB_ENDPOINT
```

Delete deployments and services

```
kubectl delete -k ./
```

