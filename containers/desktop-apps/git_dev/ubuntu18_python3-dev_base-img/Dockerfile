FROM alejandrox1/ubuntu18-dev
MAINTAINER alejandrox1

RUN curl -L https://repo.continuum.io/miniconda/Miniconda3-latest-Linux-x86_64.sh -o miniconda.sh && \
    bash miniconda.sh -b -p /opt/conda && \
    rm miniconda.sh

ENV PATH /opt/conda/bin:$PATH 


CMD ["/bin/bash"]
