---
templateKey: til
title: "Adding SSH Keys for Multiple Accounts in Git"
description: "Setting up SSH config for using multiple accounts for Git repositories."
status: published-til
slug: git-ssh-multiple-accounts
tags: ["git", "github"]
date: 2024-03-22 22:30:00
---

Let's  say you have multiple github accounts. One for your personal projects, one for your company that you work at, and one other remote repository account (let's say gitlab).

You are juggling with multiple accounts, you should not waste much time and pick a SSH from those remote repository and pull it in your local machine, that makes the process just smooth and saves a ton of time.

### Create a SSH Key

To create a SSH key, in linux you can use `ssh-keygen` command.

```bash
ssh-keygen -t ed25519 -C "alice@example.com"
```

The above command will prompt you for two things

1. The location where you want to store the key
2. The passphrase for accessing the key


### Add SSH Key to Github

Locate to the `ssh` folder and copy the generated `.pub` file to your `github` account.

For example, if you have created the key at `~/.ssh/your_name` then copy the contents of the file `~/.ssh/your_name.pub` to your clipbaord.

Navigate to your `github` account and in the settings, `SSH and GPG keys` tab, click on `Add SSH key` and copy the contents of your clipboard to the `Key` field.


### Configuring the SSH keys for multiple accounts

```config
Host your_company
  HostName github.com
  User git
  IdentityFile ~/.ssh/your_company

Host your_name
  HostName github.com
  User git
  IdentityFile ~/.ssh/your_name

Host some_name
  HostName gitlab.com
  User git
  IdentityFile ~/.ssh/some_name
```

You can change the `Host` config tag values in the `~/.ssh/conFig`

The next time you clone/create a repository on those remote git providers, you need to specify the ssh key for that account.

For example, if you have a repository `github.com/StartUp_company/some_wired_project` then you can specify the remote as `git@your_company.com:StartUp_company/some_wired_project`. Here, the `git@your_company` is the `Host` value tag from the `~/.ssh/config`. If that repository is from your `your_company` organisation/user scope, you need to add the `git@your_company` tag, if that's your project, simply add `git@your_name` before the repository url i.e. `your_name/repo_name` which would set the origin as `git@your_name:your_name/repo_name`, here the 1st `your_name` is the tag from the `Host` config and the 2nd `your_name` is the github username.

So, in summary if you wanted to use multiple accounts in the same machine, you can understand in the following example:

```bash
ssh -T git@your_name

git clone https://github.com/your_name/repo_name
```

However, you will need to authenticate with the ssh keys in this way everytime you push/pull a repository. So for that, you can set the origin with the `git@your_name` tag as the host for automatically authenticating the ssh keys on every push/pull or other activities.

Thanks for reading, Happy Coding :)
