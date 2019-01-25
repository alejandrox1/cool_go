# Monitoring Kubernetes with the Prometheus operator

## Creating a cluster
We will use an EKS cluster for this tutorial.
We will employ [`weaveworks/eksctl`](https://github.com/weaveworks/eksctl) to
create the cluster.

```
$ eksctl create cluster --auto-kubeconfig -f cluster.yaml
```

You will need to have `cloudFormation::CreateStack` and `eks::ClusterCreate`
permissions for this to work.
We mention this policies in particular since they have to be manually created.
If you did come accross policy errors, then you can start over by running:

```
$ eksctl delete cluster --region=<region> --name=<name>
```

### Interacting with the cluster
```
$ go get -u github.com/kubernetes-sigs/aws-iam-authenticator
```

### Add users
```
mapUsers: |
 - userarn: arn:aws:iam::555555555555:user/admin
   username: admin
   groups:
   - system:masters
```
