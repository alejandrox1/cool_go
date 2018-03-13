#!/bin/bash
#
# Source https://github.com/jessfraz/dockerfiles

CSP=chrome-seccomp-prof
xhost +

mkdir -p $CSP
if [ ! -f ${CSP}/chrome.json ];
then
	wget https://raw.githubusercontent.com/jfrazelle/dotfiles/master/etc/docker/seccomp/chrome.json -O ${CSP}/chrome.json
fi

docker ps -aq -f status=exited | xargs docker rm

docker run -it \
	--net host \
	--cpuset-cpus 0 \
	--memory 512mb \
	-v /tmp/.X11-unix:/tmp/.X11-unix \
	-e DISPLAY=unix$DISPLAY \
	-v $HOME/Downloads:/home/chrome/Downloads \
	-v data:$HOME/.config/google-chrome \
	--security-opt seccomp=$CSP/chrome.json \
	--device /dev/snd \
	--device /dev/dri \
	-v /dev/shm:/dev/shm \
	--name chrome \
	alejandrox1/chrome

