deployment:
	@kubectl apply -f k8s/deployment.yaml && watch -n1 kubectl get pods

delete-deployment:
	@kubectl delete -f k8s/deployment.yaml

stress-test:
	kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 200s -c 70 "http://goserver-service:8080/health"

port-forward:
	kubectl port-forward deployment/goserver 8080:8080
