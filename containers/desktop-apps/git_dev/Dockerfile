#    SD2E-CLI
#
# Install sd2e-cli from source. This image specification file allows you to
# specify what branch to build off of.
#
FROM ubuntu:18.04
MAINTAINER alejandrox1

ARG UID=1000
ARG USER=docker
ARG REPO=https://github.com/alejandrox1/sd2e-cli

RUN apt-get update -y && apt-get install -y git build-essential \
    vim bash-completion curl
    # Commenting these out in order to have the system updated - may need to
    # install something during an interactive run
    #rm -rf /var/lib/apt/lists/* && \
    #rm -rf /src/.*deb

RUN useradd -u $UID -ms /bin/bash $USER
USER $USER

WORKDIR /home/$USER

# Set a cool prompt
RUN curl -L https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bash_prompt \
    -o ~/.bash_prompt && \
    echo "source ~/.bash_prompt" >> ~/.bashrc

# Get your repo
RUN git clone --recurse-submodules $REPO

ENV PATH /home/$USER/sd2e-cli/bin:$PATH

CMD ["/bin/bash"]
