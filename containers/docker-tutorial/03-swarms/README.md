# Docker Swarm
Run app as a bonafide swarm on a cluster of docker machines.

`docker swarm init`  enables swarm mode and will make the current machine the
swarm manager.
In order to have workers, run `docker swarm join` in other machines.

## Create a Cluster
1. Let's start off by creating a couple VMs using `dcker-machine` with the help
   of the `virtualbox` driver:
   ```
   docker-machine create --driver virtualbox myvm1
   docker-machine create --driver virutalbox myvm2

   vboxmanage list vms

   docker-machine ls
   ```

## Initialise the Swarm
1. Make `myvm1` swarm manager:
   ```
   docker-machine ssh myvm1 "docker swarm init --advertise-addr <myvm1 ip>"
   ```
   NOTE: always run `docker swarm init` and `docker swarm join` with port
   `2377` or no port at all (let docker chose the default).
        1. `2377` is the swarm management port.
        2. `2376` is the docker daemon port.


   This will give some output:
   ```
   Swarm initialized: current node (bn94melmls8681siphq3wc0di) is now a manager.

   To add a worker to this swarm, run the following command:

       docker swarm join --token
       SWMTKN-1-1v02ltcey66eel1ww0ld4tvqygub1t9ouhdcoofa0e3y6zgmmo-2xqhndofe2r4cnnblwpsukpkz
       192.168.99.100:2377

    To add a manager to this swarm, run 'docker swarm join-token manager' and 
    follow the instructions.
    ```

2. Let `myvm2` join the swarm as a worker by running the pre-configured `docker
   swarm join` command given during the init process.
   ```
   docker-machine ssh myvm2 "docker swarm join --token
   SWMTKN-1-0y6rcsf2nh04ktiq5ipm8dx9gs2qadpqvnuweozmy0u3cgtf94-1l8k3ssd0y9zee131pahst2hl
   192.168.99.102:2377"
   ```

3. Duble check:
   ```
   docker-machine ssh myvm1 "docker node ls"
   ```

## Deploy

### Configure a Shell for the Manager
Instead of using `docker-machine ssh` to work on a node, we can configure a
shell to talk directly to the remote docker daemon.
Using this method, we can deploy with local yaml files - no need to move them
around.

1. Set the shell environment by evaluating the last command given from running
   `docker-machine env myvm1`
   ```
   eval $(docker-machine env myvm1)
   ```

   NOTE: to move files around:
   ```
   docker-machine scp <file> <machine>:~
   ```

2. Deploy,
   ```
   docker stack deploy -c docker-compose.yaml getstartedlab
   docker service ls
   ```

   For an image in a private registry,
   ```
   docker login registry.example.com

   docker stack deploy --with-registry-auth -c docker-compose.yml getstartedlab
   ```

## Cleanup
Tear everything down...
```
docker-machine stack rm getstartedlab

docker-machine ssh myvm2 "docker swarm leave"

docker-machine ssh myvm1 "docker swarm leave --force"

eval $(docker-machine env -u)
```
