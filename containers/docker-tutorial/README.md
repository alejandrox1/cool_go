# Hello, Docker!

1. Setup your Docker environment.
2. Build an image and run it as a container.
3. Scale app to multiple containers.
4. Distribute the app across a cluster.
5. Stack services by adding a backend database.
6. Deploy app to production.


# Fun Facts
* Accessing the name of the host when inside a container retrieves the
  container ID.


# CMDs
* `docker info`

* `docker ps`


## Clening up

## image
* `build`       Build an image from a Dockerfile
* `history`     Show the history of an image
* `import`      Import the contents from a tarball to create a filesystem image
* `inspect`     Display detailed information on one or more images
* `load`        Load an image from a tar archive or STDIN
* `ls`          List images
* `prune`       Remove unused images
* `pull`        Pull an image or a repository from a registry
* `push`        Push an image or a repository to a registry
* `rm`          Remove one or more images
* `save`        Save one or more images to a tar archive (streamed to STDOUT by default)
* `tag`         Create a tag `TARGET_IMAGE` that refers to `SOURCE_IMAGE`


## Homemade
* Session inside container
 ```
 docker exec -i -t <conatiner name> bash
 ```

* Delete all containers
 ```
 docker ps -a -q | xargs docker rm
 ```

* Delete all "exited" containers
 ```
 docker ps -aq -f status=exited | xargs docker rm
 ```

* Delete all images
 ```
 docker image ls -a -q | xargs docker rmi
 ```

* Delete all untagged images
 ```
 docker rmi $(docker images | grep "^<none>" | awk '{print $3}')
 ```


# Extras Sides
* [Docker configurations for Linux](https://docs.docker.com/install/linux/linux-postinstall/)

* [Docker Trusted Registry overview](https://docs.docker.com/datacenter/dtr/2.2/guides/)
