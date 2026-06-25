# Hello Linux

In this lab you will learn some basic commands to communicate with Linux.

## Getting some Info

```bash
# getting info about yourself
whoami

# getting user ID and group memberships
id

# getting system and kernel information
uname -a

# getting the hostname of the machine
hostname

# getting the current directory
pwd

# show how long the system has been running
uptime

# show disk space usage
df -h

# show memory usage
free -h

# locate where a command binary lives
which mkdir

# print out some text
echo "hello linux"

# print out current timestamp
date

# clearing the screen
clear

# clearing the screen (the fancy way)
<CTRL>+<L>
```

## Working with Directories

```bash
# create a directory
mkdir my-dir

# list the content of the current directory
ls -alh

# create a subdirectory (the command will fail because the directory `my-parent-dir` does not exist yet)
mkdir my-parent-dir/my-sub-dir

# getting help about the mkdir command, try to find the right argument on your own
mkdir --help

# man pages provide detailed documentation
man mkdir

# now create the directory with the `--parents` argument
mkdir -p my-parent-dir/my-sub-dir
ls -alh

# visualize the directory structure recursively
ls -R my-parent-dir

# jump into the sub directory
cd my-parent-dir/my-sub-dir
pwd

# jump into the parent directory
cd ..

# jump into the home directory
cd ~

# jump into the previous directory
cd -

# remove an empty directory
rmdir my-dir
```

## Working with Environment Variables

```bash
# print out all environment variables
env

# add an environment variable
export TRAINING_DIR="/root/"
env

# print out the environment variable (you can avoid typing some characters by pressing Tab after entering `echo $TR`)
echo $TRAINING_DIR

# remove an environment variable
unset TRAINING_DIR
echo $TRAINING_DIR

# inspect the PATH variable to see where the system looks for executables
echo $PATH
```
