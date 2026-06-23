# Package Management

In this lab you will learn how to make use of a package manager to add and remove the installed packages on a Linux machine.

## Install Packages from the official Ubuntu Repositories

```bash

# try to run `tree` (this command will fail because `tree` is not installed yet)
tree

# install tree (this command will fail because we have to update the package metainfo first)
apt install tree

# update the package metainfo (ignore the output concerning upgradable packages, we will cover this later)
apt update

# install `tree`
apt install tree

# make use of `tree`
tree

# list all installed packages
apt list --installed

# let's filter for the lines with the word tree in it (we will cover pipes and grep later in detail)
apt list --installed | grep tree

# print out the location of the executable `tree`
which tree

# remove the package `tree` again
apt remove tree

# now the tool `which` will deliver an empty output because we uninstalled `tree`
which tree

# upgrade all installed packages
# if you get a dialog with the header `Daemons using outdated libraries` some Linux services have to be restarted, so the changes are taken into account. Please select all of them.
apt upgrade

# you can chain commands like this
apt update && apt upgrade -y
```

## Install Packages from other Repositories

Some packages are not contained in the official Ubuntu repositories. To install them you have to add the repository and its GPG key. We will try this via installing Docker.

```bash
# add some additional packages (note you can span commands over several lines via the character `\`)
apt install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# add the official docker gpg key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg

# verify gpg key file was created
ls /etc/apt/keyrings/

# set up the repository
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

# verify that the repository got added
cat /etc/apt/sources.list.d/docker.list

# update apt package metadata with new packages from the recently added repository
apt update

# install docker (note that apt also takes care of the packages docker-ce depends on)
apt install docker-ce

# verify the docker installation (you should get info about the versions of docker)
docker --version
```
