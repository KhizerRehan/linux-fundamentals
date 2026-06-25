# Executables

In this lab you will learn how to handle executables and we will create our own small application.

## The path for executables in Linux

```bash
# the path in Linux is an environment variable, let's take a look at it
echo $PATH

# one of the entries is the following folder (currently nothing is in it)
ls -alh /usr/local/bin
```

## Creating a small application

```bash
# for our application we will need golang on our box, let's look if it is installed
which go

# since it does not exist we have to install it
apt install -y golang

# check if it got installed properly
go version

# take a look at the sourcecode
cat /training/03_executables/my-executable.go

# build the application
go build /training/03_executables/my-executable.go

# the executable with the name `my-executable` should now exist
ls -alh

# let's start our application, note the `./` before the name of the executable
./my-executable

# stopping the application
<CTRL>+<C>
```

## Adding the application to the path

```bash
# switch to the directory /
cd /

# due to the executable does not exist in the current directory / this will not work
./my-executable

# also the binary is not in the path, so this will not provide any info
which my-executable

# let's move the executable to the folder /usr/local/bin (which is in the PATH environment variable)
mv /root/my-executable /usr/local/bin

# now our application is known by the Linux system
which my-executable

# now run the executable from the directory /root/
my-executable
<CTRL>+<C>
```

## Process Management

```bash
# switch back to the training folder
cd ~

# let's start the application in the background and redirect stdout and stderr
my-executable > my-executable.log 2>&1 &

# verify the application is running
cat my-executable.log

# let's find the process id for our application
ps aux

# since this is too much output we grep the line with our application
ps aux | grep my-executable

# let's kill our application
# get the process id via the previous command, it is in the second column
kill -9 PROCESS_ID

# the "task manager" on our box (you can exit htop by pressing 'q')
htop
```
