# Building a container

[docker build docs](https://docs.docker.com/edge/engine/reference/commandline/build/#git-repositories)
```
docker build -t friendlyhello .
```

# Run the app
```
docker run -p 4000:80 friendlyhello
```

# Share the thing!
```
# docker tag image username/repository:tag
docker tag friendlyhello alejandrox1/playing-around:v0.1-rc

docker push alejandrox1/playing-around:v0.1-rc
```

The in some remote station,
```
docker run -d -p 4000:80 alejandrox1/playing-around:v0.1-rc
```
