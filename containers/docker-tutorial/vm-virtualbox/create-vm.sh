#!/bin/bash

set -e

VMNAME="$1"
ISO_PATH=~/Downloads/iosimages/ubuntu-16.04.4-desktop-amd64.iso

if [[ -z "$VMNAME" ]]; then
    exit 0
fi

set -x 
vboxmanage createvm --name "$VMNAME" --ostype Ubuntu_64 --register
vboxmanage storagectl "$VMNAME" --name "IDE Controller" --add ide
vboxmanage storageattach "$VMNAME" --storagectl "IDE Controller" \
    --port 0 --device 0 --type dvddrive --medium "$ISO_PATH"
vboxmanage modifyvm "$VMNAME" --memory 1024 --vram 128
vboxmanage modifyvm "$VMNAME" --nic1 nat
vboxmanage modifyvm "$VMNAME" --natpf1 "ssh,tcp,,2222,,22"
vboxmanage modifyvm "$VMNAME" --natdnshostresolver1 on

vboxmanage startvm "$VMNAME" 

vboxmanage list runningvms
