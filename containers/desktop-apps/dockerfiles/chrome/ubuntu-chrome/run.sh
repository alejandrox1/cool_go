#!/bin/bash

xhost +

# Get your own session.
if [ ! -d google-chrome ];
then
	cp -r ~/.config/google-chrome .
fi

# seccomp rules and guidelines.
if [ ! -f chrome.json ];                                                 
then                                                                            
	wget https://raw.githubusercontent.com/jfrazelle/dotfiles/master/etc/docker/seccomp/chrome.json -O chrome.json
fi

# rm exited containers.
docker ps -aq -f status=exited | xargs -r docker rm

docker run -it \
	--net host \
	--cpuset-cpus 0 \
	--memory 512mb \
	-v /tmp/.x11-unix:/tmp/.x11-unix \
	-e DISPLAY=unix$DISPLAY \
	-v downloads:/home/chrome/Downloads \
	-v google-chrome:/data \
	--security-opt seccomp=chrome.json \
	--device /dev/snd \
	--device /dev/dri \
	-v ${XDG_RUNTIME_DIR}/pulse/native:${XDG_RUNTIME_DIR}/pulse/native \
	-e PULSE_SERVER=unix:${XDG_RUNTIME_DIR}/pulse/native \
	-v /dev/shm:/dev/shm \
	--name uchrome \
	alejandrox1/uchrome
