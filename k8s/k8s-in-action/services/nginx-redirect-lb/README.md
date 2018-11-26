# Running nginx

```
$ sudo docker build --no-cache -t nginx-lb . 

$ sudo docker run --rm -d -p 80:80 nginx-lb
```
