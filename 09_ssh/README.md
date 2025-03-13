# Connecting to other machines via SSH

In this lab you will learn how to switch to other machines.

## Check the other machine

You can find the IP address of the machine you will connect in the README file you received in the beginning of the training.

```bash
# store the IP into an env variable
export OTHER_MACHINE_IP=<IP.OF.OTHER.MACHINE>
echo $OTHER_MACHINE_IP

# verify the machine is reachable
ping $OTHER_MACHINE_IP

# verify that sshd is running on port 22
nmap -p22 $OTHER_MACHINE_IP
```

## Connecting via password

```bash
# connect as root => you have to fill in the password, which is
ssh root@$OTHER_MACHINE_IP

# => you are now in the home folder of the other machine as user root

# create a file
touch <YOUR_NAME>.txt

# verify the file exists => note that also the other trainees are creating files
ls

# exit the VM
exit
```

> Note connecting via password is considered insecure. So let's create some safer way of authentication.

## Create a new key pair

```bash
# create your key pair (you can leave all inputs blank)
ssh-keygen

# verify that the key pair has been created properly
ls -alh ~/.ssh
```

## Copy the public key to the destination machine

```bash
# copy the public key
ssh-copy-id -i ~/.ssh/id_rsa.pub root@$OTHER_MACHINE_IP

# connect to the other machine => note that you are not asked for the password anymore
ssh root@$OTHER_MACHINE_IP

# you can verify the allowed public keys via the following
cat ~/.ssh/authorized_keys

# exit the other machine again
exit
```

## Connect from the source machine to the destination machine

```bash
# connect to the destination machine (note the machine name: root@training-lf-ssh)
ssh root@training-lf-ssh

# switch back to the source machine (note the machine name: root@training-lf)
exit
```

> Note now it is best practice to disable password authentication on the other machine, for security reasons. This can be done via changing the sshd config of the other machine (`PermitRootLogin` and `PasswordAuthentication` in `/etc/ssh/sshd_config.d/`). But this is out of scope for this training.

## SSH Config File

You can create a ssh config file which comes in handy. Create the file on the source machine in the location `~/.ssh/config`

Add the following content to the file

```config
Host other-machine
    HostName <IP.OF.OTHER.MACHINE>
    User root
```

```bash
# connect to the other machine
ssh other-machine

# switch back to the source machine
exit
```

## Copying files to other machines

```bash
# create a file
echo "something" > <YOUR_NAME>2.txt

# copy the file to the destination machine into its home folder (in our case /root/)
scp <YOUR_NAME>2.txt other-machine:
```

## Execute a command on another machine

```bash
# you can trigger commands via ssh on the destination machine like this
ssh other-machine "hostname && ls -alh && cat ~/<YOUR_NAME>2.txt"
```
