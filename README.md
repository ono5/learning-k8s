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
