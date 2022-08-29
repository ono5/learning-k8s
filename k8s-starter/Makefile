.PHONY: up clean watch dry-run install uninstall status upgrade namespace

up:
	kubectl apply -f store-pod.yaml
	kubectl apply -f store-node-port-service.yaml
	kubectl apply -f products-api-pod.yaml

clean:
	kubectl delete -f store-pod.yaml
	kubectl delete -f store-node-port-service.yaml
	kubectl delete -f products-api-pod.yaml

watch:
	# watch -t "kubectl get svc,deployment,pods,ingress --namespace=myshop"
	watch -t "kubectl get svc,deployment,pods,Ingress --namespace=${n}"

dry-run:
	helm install myshop --dry-run ./myshop/

install:
	# ネームスペースを作成し、その上にデプロイすることも可能
	# helm install --create-namespace myshop -n myshop ./myshop/
	helm install myshop ./myshop/

uninstall:
	# helm uninstall myshop --namespace myshop
	helm uninstall myshop

status:
	helm status myshop

upgrade:
	helm upgrade myshop ./myshop

namespace:
	kubectl get namespace
