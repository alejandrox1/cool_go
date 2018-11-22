# Bootstrap a Kubernetes Cluster
Based off:
* [How To Create a Kubernetes 1.10 Cluster Using Kubeadm on CentOS 7](https://blog.markhorhost.com/how-to-create-a-kubernetes-1-10-cluster-using-kubeadm-on-centos-7/)
* [How To Create a Kubernetes 1.11 Cluster Using Kubeadm on Ubuntu 18.04](https://blog.markhorhost.com/how-to-create-a-kubernetes-1-11-cluster-using-kubeadm-on-ubuntu-18-04/)


## Setting up the workspace directory and Ansible inventory file

`hosts` will serve as our inventory file. It specifies server information such
as IP addresses, remote users, and groups servers so these can be targeted as a
single unit when executing ansible commands.

Our inventory file consist of two ansible groups: masters and workers.
For example, if we look at the masters group, we will see a server entry that
will list the node's IP address and tells ansible that we should run commands
as the root user.


## Installing Kubernetes dependencies

* Docker - container runtime.
* kubeadm - CLI tool to install and configure various componenets of the
  cluster.
* kubelet - system service that runs on all nodes and handles node-level
  operations.
* kubectl - CLI tool used for interacting with the cluster though its API
  server.

### Creating ssh keys
```
$ ssh-keygen -b 4096 -f gcloud -N '' -C kubernetes
```

## Check cluster

```
kubectl get nodes
```

If you get something like `The connection to the server localhost:8080 was
refused - did you specify the right host or port?` on GCP, then try setting
`$HOME` to the correct place so that `kubectl` knows to use
`/home/kubernetes/.kube/config` 
(the kubernetes user is hardcoded for this example).

If you want to have access to your cluster from another host, try:
```
$ sudo kubectl proxy --address='0.0.0.0' --port=80 --accept-hosts='^*$'
```

and in your kubeconfig file change the server field to `http://<master IP>:80`.

## Test MetalLB
Check the controller and speaker are working:
```
$ kubectl get pods -n metallb-system
```

```
$ kubectl run nginx --image=nginx --port=80
$ kubectl expose deployment nginx --type=LoadBalancer --name=nginx-service
```
