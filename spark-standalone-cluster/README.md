## Spark standalone cluster running over serverless kubernetes

The following document describes the deployment of a Spark Standalone Cluster .


## Test It Out

Deploy application

```
kubectl apply -f spark-one.yml
```

Check status of the deployments/pods/services:

```
kubectl get deployment
kubectl get pod | grep myspark
kubectl get service myspark-master-webui
```

Access the sample application

```
LB_ENDPOINT=$(kubectl get service myspark-master-webui  -o jsonpath="{.status.loadBalancer.ingress[*].ip}")

# Open browser with URL in MacOSX
curl http://${LB_ENDPOINT}:8080
open http://${LB_ENDPOINT}:8080
```

Run an spark job over the spark cluster just created

```
<your_spark_home>/bin/spark-submit  run-example --master spark://${LB_ENDPOINT}:7077 --executor-memory 512M  org.apache.spark.examples.SparkPi

```

Submit a job to spark cluster and check results via http://${LB_ENDPOINT}:8080 and verify application UI http://${LB_ENDPOINT}:4040

```
~/Documents/spark/spark/spark-2.2.0-bin-hadoop2.7/bin/spark-submit \
  --class org.apache.spark.examples.SparkPi \
  --master spark://${LB_ENDPOINT}:7077 \
  --deploy-mode cluster \
  --supervise \
  --executor-memory 512M \
  --conf "spark.cores.max=2" \
  --conf "spark.executor.cores=1" \
  --conf "spark.testing.reservedMemory=128000000" \
  --executor-cores 1 \
  --driver-memory 256M \
/opt/spark/examples/jars/spark-examples_2.11-2.2.0.jar \
  100000

```


Delete application


```
kubectl delete -f spark-one.yml
```
