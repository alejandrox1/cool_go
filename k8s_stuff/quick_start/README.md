# Architecture

## [Cluster Networking](https://kubernetes.io/docs/concepts/cluster-administration/networking/#summary)

# Kuernetes Objects
A kubernetes object is a single unit that exists in a cluster.
When creating an object, we are telling the cluster about the desired state we
want to have.

Some of the most important kubernetes objects we willbe working with are the
following:
 
* [Pods](https://kubernetes.io/docs/concepts/workloads/pods/pod/)
 A pod is a group of one or more containers, with shared storage, network, and
 specification for how to run.
 All containers in a pod are reachable via a single IP address.

* [Services](https://kubernetes.io/docs/concepts/services-networking/service/)
 A service is an abstraction defininf a logical set of pods and a policy to
 access them.
 Pods have a life cycle and we need a way to manage accesibility to them
 thought the enterity fo their life cycle.
 By giving pods a certain label, we use a service to route traffic to all pods
 associated with a given label.

* [ReplicaSets](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/)
 ReplicaSets are mostly used through deployments. Their purpose is to label the
 pods to control their replication.

* [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
 A deployment describes the desired state. 
 The purpose of dplyments is to manage pods and replicasets.


