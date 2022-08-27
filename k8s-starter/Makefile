.PHONY: up clean watch

up:
	kubectl apply -f store-pod.yaml
	kubectl apply -f store-node-port-service.yaml
	kubectl apply -f products-api-pod.yaml

clean:
	kubectl delete -f store-pod.yaml
	kubectl delete -f store-node-port-service.yaml
	kubectl delete -f products-api-pod.yaml

watch:
	watch -t "kubectl get svc,deployment,pods,ingress"
