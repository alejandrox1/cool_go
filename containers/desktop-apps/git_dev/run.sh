#!/bin/bash
#-
#-                          :-) Testing sd2e-cli (-:
#-
#- Testing nvironment for the development of github.com/sd2e/sd2e-cli fork on
#- github.com/alejandrox1/sd2e-cli.
#-
##
## Usage:
##
##      ./run.sh [-v|-h] [-b|--branch]
##
##      -v|--version print to stdout any information related to this script.
##
##      -h|--help print to stdout any help information included in the header
##                of the script.
##
set -e
set -o pipefail
export blue="\e[1;34m"                                                         
export green="\e[32m"                                                           
export red="\e[1;31m"
export reset="\e[0m"

# Parameters
CONTAINER="sd2e-cli-dev"
REPO="https://github.com/alejandrox1/sd2e-cli"

# Input parameters

# Parse command line arguments.
while [[ "$#" > 0 ]]; do
    arg="$1"

    case "${arg}" in
        -v|--version)
            echo "$(grep "^#-" ${BASH_SOURCE[0]} | cut -c 4-)"
            exit 0
            ;;
        -h|--help)
            echo "$(grep "^##" ${BASH_SOURCE[0]} | cut -c 4-)"
            exit 0
            ;;
        *)
            >&2 echo "Unknown command-line option: '${arg}'."
            exit 1
            ;;
    esac
    shift
done



echo -e "${red}SD2E-CLI development starting...${reset}"
docker build --no-cache --force-rm \
    --build-arg UID=$UID \
    --build-arg USER=$USER \
    --build-arg REPO=$REPO \
    -t $CONTAINER . && \
    docker run \
    -v ~/.gitconfig:/home/$USER/.gitconfig \
    -v ~/.git-credentials:/home/$USER/.git-credentials \
    -w /home/$USER/$(basename $REPO) \
    --rm -it $CONTAINER
