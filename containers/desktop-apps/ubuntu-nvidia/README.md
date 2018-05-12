# Nvidia on Ubuntu 18.04

This [Dockerfile](Dockerfile) shows the intallation procedure for nvidia drivers and a
functioning cuda toolkit on Ubuntu 18.04.

The bash script [run-unv.sh](run-unv.sh) shows how to build a docker image and
how to run the corresponding container by attaching the appropiate devices to
the container.

This container will have a working version of the drivers along with the CUDA
compilers and any accompanying tools for debugging or profiling you code.

Furthermore, we attach a volume to the current working directory in order to
share source code and exevutables betwene the container and the machine.

For other purposes, we also have an installation script to use the nvidia
container runtime which enables sharing nvidia drivers between host and
containers. This script canbe found on
[alejandrox1/dev_env/nvidia](https://github.com/alejandrox1/dev_env/tree/master/nvidia).
