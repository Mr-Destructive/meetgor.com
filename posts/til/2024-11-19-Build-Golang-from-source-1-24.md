---
type: til
title: "Building Golang from Source v1.23 and Above"
description: "Exploring one of the way to install and build golang from source for version 1.23 and above."
status: published
slug: golang-build-from-source-1-24-above
tags: ["go",]
date: 2024-11-19 23:15:00
---

## Introduction

Are you excited to try out the latest golang version, or test out your changes (cooking some serious stuff?), or install some random golang version? Then let’s explore one of the easiest ways to install golang on your system (Linux).

## Building Go from source 1.23+

The process for installing and building golang from source is quite simple.

* Clone the repo
    
* Build the binary
    
* Set Environment Variables
    
* Export the binary to the system path
    

For detailed other steps, you can follow [this guide](https://go.dev/doc/install/source).

### Clone the repo

Just clone the repo from GitHub or Google Git repo.

```bash
git clone https://github.com/golang/go

OR

git clone https://go.googlesource.com/go
```

This will install the golang source code required to build the golang binary and the ecosystem (gofmt + standard library + test suite).

Then let’s navigate to the cloned repo and we can build the golang from the source.

### Build it

We need to run the bash script in the folder to build the binary. They `all.bash` can be run to build the binary which will be stored in the `go/bin` folder `go/bin/go` and `go/bin/gofmt` files. These two binaries will be generated and are required in the Golang ecosystem.

```bash
cd src

./all.bash
```

Once we have the binaries in the specified folder, we can move ahead to make the environment understand where the actual binary is located.

### Environment Variables

The `GOROOT` and `GOPATH` variables are required for the golang ecosystem to work as expected. The `GOROOT` is set as the path to the actual golang source repository, the cloned repository from which we built the binary. This `GOPATH` is the path where Golang stores the external repositories or modules for use anywhere in the system.

```bash
export GOROOT=path_to_clone_repo

export GOPATH=$(HOME)/go

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

The `PATH` environment needs to be updated with the `GOROOT` and `GOPATH` to make the binaries in those paths visible and accessible to the system.

> NOTE: If you are installing the golang from source when you have already a version of golang installed on your system, then you need to make sure you do not mess up the `GOROOT` and `GOPATH`.
> 
> Those could be juse set with the current shell session, as you do not want this golang version permantely on the system, if you do requrie the newly installed golang version as your default, then you can set this environment variables in your shell config.

Finally, we can now set the binary as something different because we do not want it to clash with the default golang version.

### Run the binary

The binary can be stored in the `/usr/local/bin/` to make any binary available in the system from anywhere. This is not necessary but handy if you are going to use it commonly but don’t need it as the default golang version.

```bash
# The binary is stored in the 
# path_to_cloned_repo/bin
# Whatever you want to name the binary

cp bin/go /usr/local/bin/go-dev

OR

cp bin/go /usr/local/bin/go1.24
```

Once this is done. we can check the installed golang version

```bash
go1.24 version
```

With this, you can use it as go1.24 or go-dev as the binary name.

So, that is how we install and build from source any golang version above 1.23.

## Conclusion

For context, I wanted to check out the latest changes in 1.24, so I cloned the repo and after some trial and error of some commands to build the golang version, I was able to do it correctly. So just decided to share it here, hope you found it helpful.

Thank you for reading.

Happy Coding :)
