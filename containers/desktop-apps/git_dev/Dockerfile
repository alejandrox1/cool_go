FROM ubuntu:18.04
MAINTAINER alejandrox1

ARG UID=1000
ARG USER=docker
ARG REPO=https://github.com/alejandrox1/sd2e-cli
ARG REPO_LOC=sd2e-cli
ARG UPSTREAM=https://github.com/SD2E/sd2e-cli

RUN apt-get update -y && apt-get install -y git build-essential \
    vim bash-completion curl sudo
    # Commenting these out in order to have the system updated - may need to
    # install something during an interactive run
    #rm -rf /var/lib/apt/lists/* && \
    #rm -rf /src/.*deb

# Passwordless sudo.
#RUN useradd -u $UID -ms /bin/bash -G sudo $USER && echo "$USER ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
# Sudo with password.
RUN useradd -u $UID -ms /bin/bash -G sudo $USER && echo "$USER:docker" | chpasswd
USER $USER

WORKDIR /home/$USER

# Set a cool prompt and vimrc.
RUN curl -L https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bash_prompt \
    -o ~/.bash_prompt && \
    echo "source ~/.bash_prompt" >> ~/.bashrc && \
    curl -L https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/vimrc \
    -o ~/.vimrc

# Get your repo.
RUN git clone --recurse-submodules $REPO $REPO_LOC && \
    cd $REPO_LOC && \
    git remote add upstream $UPSTREAM

# Container specific commands.
ENV PATH /home/$USER/$REPO_LOC/bin:$PATH

CMD ["/bin/bash"]
