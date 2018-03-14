FROM ubuntu:16.04
MAINTAINER Jorge Alarcon <alarcj137@gmail.com>


RUN apt-get update -y && apt-get install -y \ 
    apt-transport-https \ 
    ca-certificates \
	curl \
	libcanberra-gtk* \
	libgl1-mesa-dri \                                                           
	libgl1-mesa-glx \
	--no-install-recommends && \ 
    curl -sSL https://dl.google.com/linux/linux_signing_key.pub | apt-key add - && \
	echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google.list && \
	apt-get update -y && apt-get install -y google-chrome-stable --no-install-recommends && \
	apt-get purge --auto-remove -y curl && \
	rm -rf /var/lib/apt/lists/* && \
	rm -rf /src/*.deb

RUN apt-get update -y && apt-get install -y ca-certificates \
    alsa-base \
	pulseaudio --no-install-recommends &&\
    rm -rf /var/lib/apt/lists/* && \
	rm -rf /src/*.deb

# Add chrome user                                                               
RUN groupadd -r chrome && useradd -r -g chrome -G audio,video,voice chrome && \
    mkdir -p /home/chrome/Downloads /home/chrome/.config/google-chrome /data && \
	chown -R chrome:chrome /home/chrome && chown -R chrome:chrome /data

USER chrome

ENTRYPOINT [ "google-chrome" ]
CMD [ "--user-data-dir=/data" ]
