```
gcloud init
```

## Adding keys manually

```
USER=centos
SSHDIR=/home/centos/.ssh
adduser $USER
mkdir -p $SSHDIR
echo -e "${USER}\n${USER}" | (passwd --stdin ${USER})
echo "${USER}   ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
```

Now add your public key to `/home/centos/.ssh/authorized_keys`.
Then

```
chmod -R 600 ${SSHDIR}/*
chown -R ${USER}:${USER} ${SSHDIR}
```
