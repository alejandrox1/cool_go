Build and push the image up to DockerHub.
```
$ docker build -t alejandrox1/kubia:v1.1.0 .

$ docker push alejandrox1/kubia:v1.0.0
```

create a **ReplicationController** instead of a **Deployment**.
```
kubernetes@instance-1:~$ kubectl run kubia --image=alejandrox1/kubia:v1.0.0 --port=8080 --generator=run/v1
kubectl run --generator=run/v1 is DEPRECATED and will be removed in a future version. Use kubectl create instead.
replicationcontroller/kubia created
kubernetes@instance-1:~$ kubectl get pods
NAME          READY   STATUS    RESTARTS   AGE
kubia-fpwxj   1/1     Running   0          29s
```

`kubectl describe node instance-1` `kubectl describe pods` you'll notice the
existence of two namespaces: `kube-system` and `default`.

## Creating a service object
Create a **LoadBalancer** service object to make the app available to the
outsie.

```
kubernetes@instance-1:~$ kubectl expose rc kubia --type=LoadBalancer --name kubia-http
service/kubia-http exposed
kubernetes@instance-1:~$ kubectl get svc
NAME            TYPE           CLUSTER-IP       EXTERNAL-IP     PORT(S)          AGE
kubernetes      ClusterIP      10.96.0.1        <none>          443/TCP          25h
kubia-http      LoadBalancer   10.105.114.249   192.168.1.240   8080:30138/TCP   4h13m
```

> Looks up a deployment, service, replica set, replication controller or pod by 
> name and uses the selector for that resource as the selector for a new service 
> on the specified port. 
> A deployment or replica set will be exposed as a service only if its selector 
> is convertible to a selector that service supports, i.e. when the selector 
> contains only the `matchLabels` component.
> Note that if no port is specified via --port and the exposed resource has 
> multiple ports, all will be re-used by the new service. 
> Also if no labels are specified, the new service will re-use the labels from 
> the resource it exposes.
> 
> Service types: `ClusterIP`, `NodePort`, `LoadBalancer`, or `ExternalName`. 
> Default is `ClusterIP`.


## Scaling out
```
kubernetes@instance-1:~$ kubectl get rc
NAME    DESIRED   CURRENT   READY   AGE
kubia   1         1         1       5h2m
kubernetes@instance-1:~$ kubectl scale rc kubia --replicas=3
replicationcontroller/kubia scaled
kubernetes@instance-1:~$ kubectl get rc
NAME    DESIRED   CURRENT   READY   AGE
kubia   3         3         3       5h4m
```
