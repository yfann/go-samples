
## prometheus adapter

helm install prometheus-adapter -f prometheus-adapter.yaml stable/prometheus-adapter --namespace monitor
helm delete prometheus-adapter --namespace monitor

kubectl exec -it prometheus-adapter-86b4d79bd6-6bz5t bash -n monitor
kubectl logs prometheus-adapter-86b4d79bd6-6bz5t -n monitor
kubectl logs prometheus-server-57c4bc5b9d-npt28 -c prometheus-server -n monitor


kubectl get --raw /apis/custom.metrics.k8s.io/v1beta1
kubectl get --raw /apis/custom.metrics.k8s.io/v1beta1/namespaces/monitor/pods/*/container_cpu_cfs_periods_total

+ issues
unable to update list of all metrics: unable to fetch metrics for query "container_cpu_cfs_periods_total": Get http://prometheus-server.monitor.svc.cluster.local:9090/api/v1/series?match%5B%5D=container_cpu_cfs_periods_total&start=1599753597.435: dial tcp 10.106.179.242:9090: connect: connection refused

## hsc-sample-application

helm install  hsc-sample-hpa -f ./hsc-sample-application/values.yaml ./hsc-sample-application --namespace monitor
helm delete hsc-sample-hpa --namespace monitor

kubectl get hpa -n monitor
kubectl describe hpa hpa-test -n monitor