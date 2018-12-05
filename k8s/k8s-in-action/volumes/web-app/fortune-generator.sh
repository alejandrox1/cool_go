#!/usr/bin/env bash

trap "exit" SIGINT
HTMLDOCS=/var/htdocs
mkdir ${HTMLDOCS}

while :
do
    echo $(date) Writing to ${HTMLDOCS}/index.html
    /usr/games/fortune > ${HTMLDOCS}/index.html
    sleep 10
done
