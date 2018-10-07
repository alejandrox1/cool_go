# [Create a Kubernetes 1.10 cluster using kubeadm on CentOS 7](https://blog.markhorhost.com/how-to-create-a-kubernetes-1-10-cluster-using-kubeadm-on-centos-7/)


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


## Seting up the master node

```
kubectl get nodes
```
