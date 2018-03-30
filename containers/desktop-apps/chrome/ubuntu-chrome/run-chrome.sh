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
#docker ps -aq -f status=exited | xargs -r docker rm
state=$(docker inspect --format "{{.State.Running}}" chrome 2>/dev/null)
if [[ "$state" == "false" ]]; then
    docker rm chrome
fi

docker run -it \
	--net host \
	--cpuset-cpus 0,1 \
	--memory 3gb \
	-v /tmp/.x11-unix:/tmp/.x11-unix \
	-e DISPLAY=unix$DISPLAY \
	-v google-chrome:/data \
    -v downloads:/home/chrome/Downloads \
	--security-opt seccomp=chrome.json \
	--device /dev/snd \
	--device /dev/dri \
	-v ${XDG_RUNTIME_DIR}/pulse/native:${XDG_RUNTIME_DIR}/pulse/native \
	-e PULSE_SERVER=unix:${XDG_RUNTIME_DIR}/pulse/native \
	-v /dev/shm:/dev/shm \
	--name chrome \
	alejandrox1/chrome
