# Connecting to other machines via SSH

In this lab you will learn how to switch to other machines.

## Check the other machine

You can find the IP address of the machine you will connect to in the README file you received in the beginning of the training.

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
# connect as root => you have to fill in the password (which I will tell you ;) )
ssh root@$OTHER_MACHINE_IP

# => you are now in the home folder of the other machine as user root

# create a file
touch <YOUR_NAME>.txt

# verify the file exists => note that also the other trainees are creating files
ls

# exit the VM
exit
```

> Note: connecting via password is considered insecure. So let's create a safer way of authentication.

## Create a new key pair

```bash
# create your key pair (you can leave all inputs blank)
ssh-keygen

# verify that the key pair has been created properly
ls -alh ~/.ssh
```

## Copy the public key to the destination machine

```bash
# copy your public key
ssh-copy-id -i ~/.ssh/id_rsa.pub root@$OTHER_MACHINE_IP

# connect to the other machine => note that you are not asked for the password anymore
ssh root@$OTHER_MACHINE_IP

# you can verify the allowed public keys via the following
cat ~/.ssh/authorized_keys

# exit the other machine again
exit
```

> Note: now it is best practice to disable password authentication on the other machine for security reasons. This can be done by changing the sshd config of the other machine (`PermitRootLogin` and `PasswordAuthentication` in `/etc/ssh/sshd_config.d/`). But this is out of scope for this training.

## SSH Config File

You can create an SSH config file which comes in handy. Create the file on the source machine in the location `~/.ssh/config`

Add the following content to the file

```bash
cat <<EOF > /root/.ssh/config
Host other-machine
    HostName $OTHER_MACHINE_IP
    User root
EOF
```

```bash
# verify the SSH config file
cat /root/.ssh/config

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
ssh other-machine "hostname && ls -alh"
```

## Copying directories

```bash
# create a directory with a file to copy
mkdir -p myfolder
echo "something" > myfolder/file.txt

# scp needs -r to copy a whole directory
scp -r myfolder/ other-machine:

# rsync is faster for repeated copies (only transfers changes)
rsync -av myfolder/ other-machine:myfolder/
```
