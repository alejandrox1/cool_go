# Docker Swarm
Run app as a bonafide swarm on a cluster of docker machines.

1. Let's start off by creating a couple VMs using `dcker-machine` with the help
   of the `virtualbox` driver:
   ```
   docker-machine create --driver virtualbox myvm1
   docker-machine create --driver virutalbox myvm2

   vboxmanage list vms
   ```
