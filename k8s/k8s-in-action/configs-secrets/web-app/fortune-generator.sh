#!/usr/bin/env bash

trap "exit" SIGINT

echo Configured to generate fortune every $INTERVAL seconds

HTMLDOCS=/var/htdocs
mkdir ${HTMLDOCS}

while :
do
    echo $(date) Writing to ${HTMLDOCS}/index.html
    /usr/games/fortune > ${HTMLDOCS}/index.html
    sleep $INTERVAL
done
