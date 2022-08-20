# App

## Docker Image Push Docker Hub

### 1. Create Docker Hub Repository
Create the bleow repository
`hono1029/go-redis`

### 2. Cerate Docker Image
```
$ docker build -t hono1029/go-redis .
$ docker push hono1029/go-redis
```

## Helm Chart

### Ceate Helm Chart
```
$ helm create apicharts
```

### How to deploy with Helm Chart?

[helm install](https://helm.sh/docs/helm/helm_install/#helm-install)

```
# helm install [NAME] [CHART] [flags]
helm install myapi apicharts/

export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=apicharts,app.kubernetes.io/instance=myapi" -o jsonpath="{.items[0].metadata.name}")

export CONTAINER_PORT=$(kubectl get pod --namespace default $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")

echo "Visit http://127.0.0.1:8080 to use your application"

kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT
```

### How to undeploy with helm Chart?

```
$ helm uninstall myapi
release "myapi" uninstalled
```

### How to package Helm Chart?

```
$ helm package apicharts/
Successfully packaged chart and saved it to: /Users/hoge/Desktop/k8s/app/apicharts-0.1.0.tgz

$ helm install myapi apicharts-0.1.0.tgz
$ helm uninstall myapi
release "myapi" uninstalled
```
