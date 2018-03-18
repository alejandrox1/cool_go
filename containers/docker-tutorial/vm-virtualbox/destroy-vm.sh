#!/bin/bash

set -e


VMNAME=""
POWEROFF="false"

while [[ $# -gt 0 ]]; do
    key="$1"

    case $key in 
        -n|--name)
            VMNAME="$2"
            shift
            shift
            ;;
        -p|--poweroff)
            POWEROFF="true"
            shift
            ;;
        *)
            echo "You messed up with the flags buddy..."
            break
    esac
done

if [[ -z "$VMNAME" ]]; then
    exit 1
fi
if [[ "$POWEROFF" == "true" ]]; then
    echo "Powering off machine..."
    vboxmanage controlvm "$VMNAME" poweroff --type headless
    echo \n\n\n
fi

echo "Deleting machine..."
vboxmanage unregistervm "$VMNAME" --delete

vboxmanage list runningvms
vboxmanage list vms
