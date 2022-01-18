# How to connect ec2 to git repo


## connect to EC2

```bash
ssh -i <.pem> <user>@public-ip4
```

## Make sure you are in `/home/ubuntu`
```bash
pwd

#output should be `/home/ubuntu`
```

## Create SSH Keygen
```bash
ssh-keygen

# just enter until done

# check ssh
cd .ssh

# read id_rsa.pub
cat id_rsa.pub

# copy all text in id_rsa.pub file
```

## Go to github web
- Open your github profile
- Open Setting
- Click SSH and GPG Keys
- Click New SSH Key
- Paste your id_rsa.pub text

Now you can clone and pull your repo from your instance.