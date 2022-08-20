# Kubernetes


## Reference
https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/

## Enter into Pod

```
# POD -> pod name that you want to enter in
kubectl exec -it POD sh
```

## Watch Command

```
$ kubectl get pod -w
```

## Get Pod's IP address

```
$ kubectl get pod -o wide
NAME    READY   STATUS    RESTARTS   AGE   IP          NODE             NOMINATED NODE   READINESS GATES
debug   1/1     Running   0          22s   10.1.0.16   docker-desktop   <none>           <none>
nginx   1/1     Running   0          22s   10.1.0.15   docker-desktop   <none>           <none>
```

## File Transfer

```
# SRC -> Source file that you want to transfer
# DEST -> Destination File you want to get
kuubectl cp SRC DEST

# Host -> Pod
kubectl cp <src> <pod-name>:<dest>

# Pod -> Host
kubectl cp <pod-name>:<src> <dest>
```

## How to investigate pod

```
# TYPE/NAME -> resource type / resource name
kubectl describe [TYPE/NAME]

# n -> get most recent record until assigned number
kubectl logs [TYPE/NAME][--tail=n]
```

## Confirm Revision

```
kubectl rollout history TYPE/NAME
```

## Rollout

```
kubectl rollout undo TYPE/NAME --to-revision=N
```

## Secret

```
# create Secerat from command
# open
# --from-literal=key=value # assign key value
# --from-file=filename # ceate from file
# NAME -> secret
kubectl create secret generic NAME [optoin]
```

Base64 string

```
echo -N 'TEXT' | base64
```

```
$ kubectl create secret generic sample-secret --from-literal=message='Hello World!' --from-file=./secret/keyfile

$ kubectl get secret
NAME            TYPE     DATA   AGE
sample-secret   Opaque   2      28s

$ kubectl get secret/sample-secret -o yaml
apiVersion: v1
data:
  keyfile: WU9VUi1TRUNSRVQtS0VZCg==
  message: SGVsbG8gV29ybGQh
kind: Secret
metadata:
  creationTimestamp: "2022-08-12T01:42:53Z"
  name: sample-secret
  namespace: default
  resourceVersion: "54737"
  uid: 93e895ff-c260-48db-b43b-90a9efc84b87
type: Opaque
```

# Run pod for debug

```
kubectl run debug --image=centos:7 -it --rm --restart=Never -- sh
```

# Output yaml
```
$ kubectl get service/myapi-apicharts  -o yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    meta.helm.sh/release-name: myapi
    meta.helm.sh/release-namespace: default
  creationTimestamp: "2022-08-20T09:29:48Z"
  labels:
    app.kubernetes.io/instance: myapi
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: apicharts
    app.kubernetes.io/version: 1.16.0
    helm.sh/chart: apicharts-0.1.0
  name: myapi-apicharts
  namespace: default
  resourceVersion: "302292"
  uid: 1d10db3a-0b6b-484e-b6cf-d92929347df6
spec:
  clusterIP: 10.102.168.60
  clusterIPs:
  - 10.102.168.60
  internalTrafficPolicy: Cluster
  ipFamilies:
  - IPv4
  ipFamilyPolicy: SingleStack
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: http
  selector:
    app.kubernetes.io/instance: myapi
    app.kubernetes.io/name: apicharts
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
```
