In directory hw06 

Run the following to set up database
```shell
kubectl apply -f ./deployments/kube/postgresql-secret.yaml 

helm repo add bitnami https://charts.bitnami.com/bitnami
helm install postgresql-hw06 bitnami/postgresql -f ./deployments/postgresql-values.yaml
```

Run the following to apply migrations 
```shell
kubectl apply -f ./deployments/kube/migrations
```

After a second check if migrations where applied successfully. COMPLETIONS must be "1/1"
```shell
kubectl get job hw06-migrate-up-job
```

Run to deploy the service
```shell
kubectl apply -f ./deployments/kube/app
```

To test the solution in Postman 
1. import test/user-api-postman-collection.json 
1. change base_url variable to minikube address (minikube ip)
1. run the collection

Cleanup
```shell
kubectl delete -f ./deployments/kube/app
kubectl delete -f ./deployments/kube/migrations
helm delete postgresql-hw06
kubectl delete pvc -l app.kubernetes.io/instance=postgresql-hw06
```