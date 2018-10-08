## Creating an deleting an object

```
kubectl apply -f kuard-pod.yaml
```

```
kubectl delete -f kuard-pod.yaml
```

or
```
kubecl delete pods/kuard
```

### Creating a pod by hand

```
kubectl run kuard --image=docker.io/alejandrox1/some-image
```

See what's going on behind the scenes
```
kubectl get pods
```

Delete the pod (deployment really)
```
kubectl delete deploymnets/kuard
```


### Debugging

```
kubectl exec -it kuard -- bash
```

Copy files to and from
```
kubectl cp kuard:/path/to/file /local/path
```

For logs,
```
kubectl logs kuard [-f] [--previous]
```
