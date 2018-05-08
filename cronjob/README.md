## CronJob Sample

The following sample is for running automated tasks with cron jobs. The details could be found in  https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/

NOTE: The cronjob pod will be cleaned up automatically after completed. So you need to save  your job result to OSS or RDS, etc.

## Test It Out

Deploy application

```
kubectl create -f ./cronjob.yaml
```

Check on the status of the cron job using this command: 

```
kubectl get cronjob hello
kubectl get jobs --watch
```

Delete cronjob

```
kubectl delete -f cronjob.yaml
```

