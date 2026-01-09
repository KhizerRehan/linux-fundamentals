# Linux Fundamentals

## Setup the training environment

1. Open [Github Codespaces](https://github.com/codespaces) and create your new `cloudnativetrainings/linux-fundamentals` codespace.
1. Copy the files `README.MD`, `ssh-config` and `ssh-private-key` into the directory named `.secrets` into your codespace. You can drag and top those two files into the browser.
1. Run the following commands:

```bash


# fix the permissions of your private key
chmod 0400 ./.secrets/ssh-private-key

# connect to your VM
ssh -F ./.secrets/ssh-config linux-fundamentals-vm
```

## Teardown the training environment

1. Delete your `cloudnativetrainings/linux-fundamentals` codespace via [Github Codespaces](https://github.com/codespaces).
