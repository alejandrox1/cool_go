```
$ kubectl get po
NAME      READY   STATUS              RESTARTS   AGE
fortune   0/2     ContainerCreating   0          5s

kubernetes at instance-1 in ~
$ kubectl port-forward fortune 8080:80
Forwarding from 127.0.0.1:8080 -> 80
Forwarding from [::1]:8080 -> 80
Handling connection for 8080
```

In another terminal:
```
$ curl localhost:8080
Your manuscript is both good and original, but the part that is good is not
original and the part that is original is not good.
		-- Samuel Johnson
```
