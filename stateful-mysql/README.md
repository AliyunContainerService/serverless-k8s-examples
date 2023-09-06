## Deploying a stateful MySQL cluster

In this example we deploy a highly available MySQL topology on ACK Serverless.

## Prerequisites

* [A ACK Serverless cluster with more that one virtual nodes](https://www.alibabacloud.com/help/en/ack/serverless-kubernetes/user-guide/create-an-ask-cluster-2?spm=a2c63.p38356.0.0.1bd47252kJUeed)
* [Spread Elastic Container Instance-based pods across zones and configure affinities](https://www.alibabacloud.com/help/en/ack/serverless-kubernetes/user-guide/spread-eci-based-pods-across-zones-and-configure-affinities-1?spm=a2c63.p38356.0.0.78166818nVVT8p)
* Addon `csi-provisioner` has been installed, and its version is >= v1.22.9.

## Test It Out
Clone this repository and enter this working directroy.

1. Create a namespace for the StatefulSet.
```bash
kubectl create namespace mysql
```

2. Create a Secret for the StatefulSet.
```bash
kubectl apply -n mysql -f secret.yaml
```

3. Create a StorageClass for the StatefulSet.
```bash
kubectl apply -f storageclass.yaml
```

4. Deploy the StatefulSet.
```bash
kubectl apply -n mysql -f stateful-mysql.yaml
```

5. Get the StatefulSet and verify whether the deployment is successful.
```bash
kubectl get statefulset -n mysql
```

6. Get pods.
```bash
kubectl -n mysql get po -o wide
```

Check the colume NODE. It's expected that pods are in different nodes/zones.
```bash
NAME     READY   STATUS    RESTARTS   AGE     IP             NODE                            NOMINATED NODE   READINESS GATES
dbc1-0   1/1     Running   0          2m28s   10.1.214.213   virtual-kubelet-cn-hangzhou-j   <none>           <none>
dbc1-1   1/1     Running   0          106s    10.3.146.113   virtual-kubelet-cn-hangzhou-h   <none>           <none>
dbc1-2   1/1     Running   0          64s     10.4.232.78    virtual-kubelet-cn-hangzhou-g   <none>           <none>
```
