# Services 

## Create a `ClusterIP` type service
```
$ kubectl get svc
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP   94m
kubia        ClusterIP   10.101.154.207   <none>        80/TCP    5m11s
```

```
$ kubectl exec -it kubia-4pksp bash
...

root@kubia-4pksp:/# curl 10.101.154.207
10.44.0.0:54052 - kubia-4pksproot@kubia-4pksp:/# curl 10.101.154.207/health
okroot@kubia-4pksp:/#
```

## Ingress resource

```
$ openssl genrsa -out tls.key 4096
$ req -new x509 -key tls.key -out tls.cert -days 360 0subj /CN=alejandrox1.com

$ kubectl create secret tls tls-cert --cert=tls.cert --key.tls.key
```


## References
* [NGINX Ingress Controller: Bare-metal considerations](https://kubernetes.github.io/ingress-nginx/deploy/baremetal/)
    * [NGINX Ingress Controller](https://kubernetes.github.io/ingress-nginx/deploy/)
