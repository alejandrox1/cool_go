FROM ubuntu:18.04

RUN apt-get update -y \
    && apt-get install -y lsb-release curl gnupg2 \
    git bash-completion vim tree \
    software-properties-common build-essential python-dev python-pip \
    && apt-add-repository ppa:ansible/ansible \
    && apt-get update -y && apt-get install -y ansible \
    && pip install apache-libcloud==0.20.1

RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
    echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    apt-get update -y && apt-get install google-cloud-sdk -y

ADD https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bashrc /root/.bashrc
ADD https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bash_prompt /root/.bash_prompt

WORKDIR /root/compute-video-demo-ansible

CMD ["/bin/bash"]
