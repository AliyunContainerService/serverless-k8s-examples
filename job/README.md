## Job Sample

The following sample describes the deployment of a job. The details could be found in  https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/

NOTE: The job pod will be cleaned up automatically after completed. So you need to save  your job result to OSS or RDS, etc.

## Test It Out

Deploy application

```
kubectl create -f job.yaml
```

Check on the status of the job using this command: 

```
kubectl describe jobs/pi
kubectl get pods --selector=job-name=pi --show-all
```

Delete job

```
kubectl delete -f ./job.yaml
```
