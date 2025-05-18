---
type: til
title: "Python Pipreqs: Generate requirements file from the imported packages"
description: "Exploring the pipreqs package that allows to list all the dependencies or packages which are imported in a python project"
date: 2022-09-14 21:57:30
status: published
slug: python-pipreqs
tags: ['python',]
---

## Introduction

[Pipreqs](https://pypi.org/project/pipreqs/) is a python package that allows us to list all the pacakges which are imported in a python project. This is a great package for reducing the amount of redundant packages for a project. 

## Install pipreqs

You can install pipreqs with one of the many ways with pip, pipx, or any other pacakge management tool. I personally use pipreqs with `pipx` as it remains isolated from the rest of my project dependencies.

### Using simple pip install

We can install with pip by creating a virtual environment or in a existing virtual environment.

```
pip install virtualenv venv
source venv/bin/activate

pip install pipreqs
```

### Using pipx

We can install pipreqs with pipx. [Pipx](https://pypi.org/project/pipx/) is also a python package but used as a tool to install any cli specific tool with the isolated environment.

```
pipx install pipreqs
pipx run pipreqs
```

## Using pipreqs

We need to specify the encoding, which is used for reading the files while capturing the imports from the project.

```
pipx run pipreqs --encoding=utf-8 .
```

Additionaly, we can specify the `path` or filename where it will be used to save the imported packages. The `--savepath` option takes in the path to the file where you want to generate the list of the packages to be installed.

```
pipx run pipreqs --encoding=utf-8 --savepath reqs.txt . 
```

Though this doesn't guarentee all the requirements for a file, it is really helpful for explicitly used packages in the python project.
