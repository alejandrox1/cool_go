FROM ubuntu:18.04
MAINTAINER alejandrox1

RUN apt-get update -y && apt-get install -y git build-essential \
    vim bash-completion curl sudo
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /src/.*deb

# Set a cool prompt and vimrc.
ONBUILD RUN curl -L https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bash_prompt \
    -o ~/.bash_prompt && \
    echo "source ~/.bash_prompt" >> ~/.bashrc && \
    curl -L https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/vimrc \
    -o ~/.vimrc

CMD ["/bin/bash"]
