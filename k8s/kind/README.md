# The app
1. `cobra init kind`

2. Use glog

3. `cobra add build`

4. `pkg/build/images.go` which will copy the files necessary to build a
   `kind-node`, it will compile the entrypoint, and build the image using the
    Docker cli.
```
go build
$ ./kind build image --source images/node/
```

5. `cobra add create`

`pkg/cluster` implements method to create and manage the node cluster.

```
$ kind create
```

6. `cobra add delete`
```
$ kind delete
```

# Bazel and vedor
```
$ dep init
```

```
$ bazel run //:gazelle
```

```
$ bazel build //:kind
INFO: Analysed target //:kind (0 packages loaded).
INFO: Found 1 target...
Target //:kind up-to-date:
  bazel-bin/linux_amd64_stripped/kind
INFO: Elapsed time: 0.129s, Critical Path: 0.00s
INFO: 0 processes.
INFO: Build completed successfully, 1 total action
```

To clean up
```
$ bazel clean
```
