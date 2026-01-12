# Bash Scripts

In this lab you will learn how to write your own bash scripts.

## Create your own bash script

Create a file called `my-bash-script.sh` via vi with the following content

> #!/bin/bash
> date >> my-bash-script-file.txt

You can do this also without vi like this:

```bash
cat <<EOF > my-bash-script.sh
#!/bin/bash
date >> my-bash-script-file.txt
EOF
```

```bash
# verify the content of your bash script
cat my-bash-script.sh

# try to run the bash script (which will fail due to it is not executable yet)
./my-bash-script.sh

# make the bash script executable
ls -alh my-bash-script.sh
chmod 700 my-bash-script.sh
ls -alh my-bash-script.sh

# run the bash script
./my-bash-script.sh

# verify the bash script worked out
cat my-bash-script-file.txt

# try to run the script without ./
# this will not work - this is a safety mechanism of Linux for avoiding unintended bash script executions
my-bash-script.sh
```

## Create a python script

Create a python script.

```bash
cat <<EOF > my-python-script.py
#!/usr/bin/env python3
print('Hello from Python')
EOF
```

> Note that the Hash-Bang line does not define the binary to run the code explicitely. This will cause that the first binary called python3 in the path will run the script.

```bash
# verify the content of your python script
cat my-python-script.py

# make the python script executable
chmod 700 my-python-script.py

# run the python script
./my-python-script.py
```
