# Pods

## Create a pod resource
```
$ kubectl create -f kubia-manual.yaml

$ kubectl get pods
NAME                    READY   STATUS    RESTARTS   AGE
kubia-manual            1/1     Running   0          3s
```

## Forward a local network port to a port in a pod
```
$ kubectl port-forward kubia-manual 8888:8080
Forwarding from 127.0.0.1:8888 -> 8080
Forwarding from [::1]:8888 -> 8080
Handling connection for 8888
Handling connection for 8888
```

The last two lines come from testing out our app:
```
$ curl localhost:8888
127.0.0.1:57658 - kubia-manual
$ curl localhost:8888
127.0.0.1:57660 - kubia-manual
```

If you want to see the logs for the `kubia` container in the `kubia-manual`
pod:
```
$ kubectl logs kubia-manual -c kubia
2018/11/25 01:06:59 127.0.0.1:57658 - kubia-manual
2018/11/25 01:07:01 127.0.0.1:57660 - kubia-manual
```

# Creating a namespace
```
$ kubectl apply -f custom-ns.yaml
namespace/custom-namespace created

$ kubectl get ns
NAME               STATUS   AGE
custom-namespace   Active   7s
default            Active   41h
kube-public        Active   41h
kube-system        Active   41h
metallb-system     Active   20h
```
