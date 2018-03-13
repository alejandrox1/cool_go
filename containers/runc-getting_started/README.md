# RunC: quick run through

To start a container, we need the following:
* rootfs - directory with a filesystem bundle
* config.json
* runtime.json

## rootfs
Create a rootfs by exporting a docker container.

1. Create a container,
 ```
 sudo docker create --name cont1 alpine sh
 ```

2. Export to tar file,
 ```
 sudo docker export cont1 > alpine.tar
 ```

3. rm container,
 ```
 sudo docker rm -f cont1
 ```

4. Create the rootfs directory,
 ```
 mkdir rootfs
 tar -xf alpine.tar -C rootfs
 ```

## Create specification files
With a rootfs alive we now need to generate the two specification files.
We can use runC to create a default file based on rootfs
```
runc spec
```


## Run your container
```
sudo runc start
```

and to test it out,
```
apk --version
```
