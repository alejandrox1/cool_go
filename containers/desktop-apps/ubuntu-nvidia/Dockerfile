FROM ubuntu:18.04
MAINTAINER Jorge Alarcon Ochoa

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
    apt-get install -y gcc-5

# GM107M [GeForce GTX 960M]
RUN apt-get install -y nvidia-driver-390 nvidia-cuda-toolkit 

CMD ["/bin/bash"]
