---
templateKey: til 
title: "Django App from Scratch Using Docker with Debian Base Image"
description: "Creating a django basic application with configuration for static files, templates and user authentication using docker and debian base image."
date: 2022-05-24 22:30:00
status: published-til
slug: django-app-from-scratch
tags: ['docker', 'web-development', 'django',]
---

## Pull a Fresh Debian Image

Create a docker container from a Debian image, the following command can be used to pull a debain 11-slim image and create a container of it, also enter into the container in a interactive environment `-it` mode. 

```
docker run -v $(pwd):/var/www --rm -it -p 8001:80 debian:11-slim
```

## Create a Django App from Shell script

Now, since we are inside a Debian container, we can enter a few commands, you can refer to the Mark Gibney's [GitHub repository](https://github.com/no-gravity/web_app_from_scratch) for the script.

```
apt update

apt install wget

wget https://raw.githubusercontent.com/no-gravity/web_app_from_scratch/main/django.sh
```

Also, if you want to do a few adjustments, you can install an editor, get used to vim or use nano.

```
apt install vim 

OR

apt install nano
```

```
chmod +x django.sh
bash django.sh
```

I also have a few adjustment of the original script, that accepts a project name and creates a django project based on the positional parameter given to it. You can get it from the [quick-setup-script repository](https://github.com/Mr-Destructive/quick-setup-scripts/blob/main/django_docker.sh) or directly the [script](https://raw.githubusercontent.com/Mr-Destructive/quick-setup-scripts/main/django_docker.sh).

To use the above file, you need to execute the command as :

```
chmod +x django_docker.sh
bash django_docker.sh <project_name>
```

This will generate the project in the `/var/www/` folder with the name of the project. The script will prompt you with a few things for setting up at some iterations like basic application setup , `static file` configuration, `basic tempalte` setup and the `user authentication` setup.

## Copy the contents from the docker container

You can copy the contents of the folder into your local machine by entering the [cp](https://docs.docker.com/engine/reference/commandline/cp/) command in docker. 

```
docker cp <container_id>:/var/www/<project_name> /path/in_local_machine/
```

This will copy the project in the docker container into the local machine for you to extend and tweak it as per your needs.

That's a basic Django Project Setup by using Docker with a Debian Image.
