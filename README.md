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
