# Docker Services

Let's scale up our application and enable load-balancing.

In a distributed application, different pieces of the app are called
**services**.

To define, run, and sale services we will use `docker-compose`.



A single cntainer running in a service is called a task.


# Run the App
1. Run `docker swarm init`

Which will output something like this,
```
Swarm initialized: current node (7bw3rwmoa7yozlea0dzqfsteu) is now a manager.

To add a worker to this swarm, run the following command:

docker swarm join --token SWMTKN-1-4n1oeb9cl0bwl3c4g6n61c78npdmg6bp6f7fiu8h4iljk5lo8g-29afylsqqc59bx09dzu83newk 192.168.254.12:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```

2. Give app a name by running
```
docker stack deploy -c docker-compose.yaml getstartedlab
```

3. Check it out:
```
docker service ls

# List the tasks in your service,
docker service ps getstartedlab_web
```

# Scale the App
To scale up or down, simply re-run
```
dockr stack deploy -c docker-compose.yaml getstartedlab
```
Docker will perform an in-place update.


# Cleaning Up
```
docker stack rm getstartedlab

docker swarm leave --force
```
