## PV/PVC Sample

The following sample is for running a simple application to demonstrate pv/pvc usage.

Firstly, we should install alicloud-disk-controller in Serverless Kubernetes
Please use your own AK in configmap of alicloud-disk-controller.yaml

```
kubectl apply -f alicloud-disk-controller.yaml

kubectl get storageclass
NAME                       PROVISIONER     AGE
alicloud-disk-available    alicloud/disk   159m
alicloud-disk-efficiency   alicloud/disk   159m
alicloud-disk-essd         alicloud/disk   159m
alicloud-disk-ssd          alicloud/disk   159m

kubectl -n kube-system get deploy
NAME                       DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
alicloud-disk-controller   1         1         1            1           160m
```

## Test It Out

1. Deploy disk pvc application

```
kubectl create -f ./pod-with-pvc-disk.yaml
```

After pod running, check the pv/pvc and disk'smounted path.

```
kubectl get pvc
NAME       STATUS   VOLUME                   CAPACITY   ACCESS MODES   STORAGECLASS         AGE
pvc-essd   Bound    d-2zec3qewe49hxcrmk4ko   20Gi       RWO            alicloud-disk-essd   33m

kubectl get pv
NAME                     CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM              STORAGECLASS         REASON   AGE
d-2zec3qewe49hxcrmk4ko   20Gi       RWO            Delete           Bound    default/pvc-essd   alicloud-disk-essd            33m

kubectl exec -it nginx ls /pvc
lost+found
```

Delete resources:
```
kubectl delete -f ././pod-with-pvc-disk.yaml
```

2. Deploy nas pvc application

```
kubectl create -f ./pod-with-pvc-nas.yaml
```

After pod running, check the pv/pvc and disk'smounted path.

```
kubectl get pvc
NAME       STATUS   VOLUME                   CAPACITY   ACCESS MODES   STORAGECLASS         AGE
nfs-pvc    Bound    nfs-pv                   8Gi        RWX                                 19m

kubectl get pv
NAME                     CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM              STORAGECLASS         REASON   AGE
nfs-pv                   8Gi        RWX            Retain           Bound    default/nfs-pvc                                  19m
```

Delete resources:
```
kubectl delete -f ././pod-with-pvc-disk.yaml
```





