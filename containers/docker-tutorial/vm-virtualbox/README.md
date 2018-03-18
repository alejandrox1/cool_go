# Provisioning VMs

Does not require VTX/AMD-V or special Linux kernel features.

## Get Info
* List avail OSs
  ```
  vboxmanage list ostypes
  ```

* List all VMs
  ```
  vboxmanage list vms
  ```

## Manage VMs
* Create a VM `vboxmanage createvm --name "ubuntu server" --ostype "Ubuntu_64" --register`
  will output something like this:
  ```
  Virtual machine 'ubuntu server' is created and registered.
  UUID: 55882de3-e2eb-4d36-8234-fd4aec9273d8
  Settings file: '/home/user/VirtualBox VMs/ubuntu server/ubuntu server.vbox'
  ```

* Get rid of a VM
  ```
  vboxmanage unregistervm "ubuntu server" --delete
  ```

## Get Going
* Start running a VM
  ```
  vboxmanage startvm "ubuntu server" --type headless
  ```

* Power off
  ```
  vboxmanage controlvm "ubuntu server" poweroff --type headless
  ```

* Pause and resume
  ```
  vboxmanage controlvm "ubuntu server" pause --type headless

  vboxmanage controlvm "ubuntu server" resume --type headless
  ```

## Working with your VM
The best way to get into a guest machine is via port fowarding.
By this point we should already have an interface using NAT.

```
vboxmanage mofidyvm "ubuntu server" --nic1 nat
vboxmanage modifyvm "ubuntu server" --natpf1 "ssh,tcp,,2222,,22"
vboxmanage modifyvm $VM_NAME --natdnshostresolver1 on

# Check added rules.
vboxmanage showvminfo "ubuntu server" | grep 'Rule'
```

Then simply,
```
sh -p 2222 username@127.0.0.1
```

# References
* [sshing into a vm](https://leemendelowitz.github.io/blog/ubuntu-server-virtualbox.html)

* [vm
  cheatsheet](https://www.perkin.org.uk/posts/create-virtualbox-vm-from-the-command-line.html)
