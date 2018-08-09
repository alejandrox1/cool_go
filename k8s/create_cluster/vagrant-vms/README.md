# Creating a Bionic Box

This tutorial follows the 
[creating a vagrant box from an existing one](https://scotch.io/tutorials/how-to-create-a-vagrant-base-box-from-an-existing-one)

See what [vagrant boxes are available](https://app.vagrantup.com/boxes/search).

We are going to go from Ubuntu 16.04 to Ubuntu 18.04.

## Get a base box
```
 $ vagrant init ubuntu/xenial64
A `Vagrantfile` has been placed in this directory. You are now
ready to `vagrant up` your first virtual environment! Please read
the comments in the Vagrantfile as well as documentation on
`vagrantup.com` for more information on using Vagrant.
```

Now, boot into the machine.
```
 $ vagrant up
```

Once that is done, ssh into the box.
```
 $ vagrant ssh
```

## Upgrade the system
Once inside the box, make sure the system is as up-to-date as possible.
```
 $ sudo apt update
 $ sudo apt upgrade
 $ sudo apt dist-upgrade
```

Remove the packages that are no longer required.
```
 $ sudo apt autoremove
```

We will upgrade our system following the ubuntu way. For this we will need to
install one more package.
```
 $ sudo apt install update-manager-core
 $ sudo do-release-upgrade
```

Now, to do th upgrade (at the time of this writing we have to use the `-d` flag):
```
 $ sudo do-release-upgrade -d
```

And just follow the instructions.

When the whole process end, you will need to restart the system. As we are
working insde a VM the ssh conection will be droped.
When this happens run `vagrant ssh` again and be patient, it may take a little
while.


One little change we will make: change the hostname.
```
$ sudo hostname ubuntu-bionic
```

Again, just exit and `vagrant ssh` back.

## Prep the box for posterity
Before we package this box, we will zeroout the drive. This is done in order to
fix framgmentation issues with the underlying disk and will permit an efficient
compression in a couple more steps.
```
sudo dd if=/dev/zero of=/EMPTY bs=1M
sudo rm -f /EMPTY
```

Clear the bash history.
```
 $ cat /dev/null > ~/.bash_history && history -c && exit
```

Repackage the machine we just created into a new Vagrant Base Box.
```
 $ vagrant package --output ubuntubionic64.box
```

This `.box` file is a tarred zip package containing the following:
* Vagrantfile
* `*.vmdk` - virtual harddisk drive
* `*.ovf` - virtual hardware definitions
* metadata.json - provider information

To make this box available see 
[creating a new vagrant box](https://www.vagrantup.com/docs/vagrant-cloud/boxes/create.html)

If you want to start completely from scratch 
[download an ubuntu server](https://www.ubuntu.com/download/server) and follow
this great tutorial: 
[building a vagrantbox from start to finish](https://www.engineyard.com/blog/building-a-vagrant-box)
