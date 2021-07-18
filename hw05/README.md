To test the solution apply yamls 
```shell
kubectl apply -f deployments/
```
and then run the following
```shell
curl http://$(minikube ip)/health -H "HOST: arch.homework"
curl http://$(minikube ip)/otusapp/nbelyakov/health -H "HOST: arch.homework"

curl http://$(minikube ip)/otusapp/nbelyakov/a -H "HOST: arch.homework"
curl http://$(minikube ip)/a -H "HOST: arch.homework"
```