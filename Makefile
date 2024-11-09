deployment:
	@kubectl apply -f k8s/deployment.yaml && watch -n1 kubectl get pods

delete-deployment:
	@kubectl delete -f k8s/deployment.yaml
