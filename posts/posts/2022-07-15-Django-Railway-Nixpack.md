{
  "type": "posts",
  "title": "Deploying Django Project with Railway Nixpacks",
  "description": "Configuring and Seting up nixpacks for deploying django project on Railway app. Interacting with the nixpacks CLI for building and creating deployable django applications.",
  "date": "2022-07-16 16:15:00",
  "status": "published",
  "slug": "django-deploy-railway-nixpacks",
  "tags": [
    "django",
    "python",
    "railway"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/django-deploy-railway-nixpacks.png",
  "series": [
    "Django-Deployment",
    "Django-Series"
  ]
}

## Introduction

We have seen how to deploy a Django application on railway app in the [previous article](https://www.meetgor.com/django-deploy-railway/) of the [series](https://www.meetgor.com/series/django-deployment/). We deployed the django project using the Heroku Buildpacks under the hood. The railway app provides a couple of options on how to build your web application. Currently, there are three options, one of which is the [heroku buildpacks](https://devcenter.heroku.com/articles/heroku-20-stack) which is the default one, second, we have the nixpacks which we will see today, and the third is the [Paketo buildpack](https://paketo.io/). 

## What is a Buildpack?

A buildpack is a set of programs that turns your source code into a container image. So it is basically a tool for converting your application into a deployment-ready state with the help of containerization technology. Buildpacks allow us to extract away the steps for deploying an application. The layer of abstraction for converting a source code into a deployable code/container is played by the build packs.

### Benefits of buildpacks

Buildpacks as we discussed act as a layer of abstraction from converting source code into the deployable containers, so it helps in avoiding manually writing dependencies and installing them. But on top of that, they can even detect low-level changes to the source code, i.e. if a dependency is changed, it would know which dependencies are altered and which components can be fetched from the cache. Caching is a great property of buildpacks which enhances the performance and deployment time. Since we can specify the commands for the build process, the buildpacks are customizable and hence provide a solid foundation for professional applications.

For more references on Buildpacks, you can follow up with the great article on [dockerfile vs buildpacks](https://technology.doximity.com/articles/buildpacks-vs-dockerfiles).

## What are Nixpacks?

[Nixpacks](https://nixpacks.com/docs) are quite similar to buildpacks, but they have their own set of technologies used in managing and installing builds for the application. Nixpacks as the name suggests uses [Nix Packages](https://search.nixos.org/packages) for creating and installing dependencies and [Docker](https://www.docker.com/) for building images and running containers.

Nixpacks are quite cool as most of the applications require little or almost no configuration. You don't need a lot of knowledge of docker or nix technologies, everything is abstracted for you. Nixpacks uses nix packages for installing the runtime environment and the dependencies for the applications. It also caches support for detecting the core modules or packages in the runtime, so it gets a huge boost in deployment performance. Minimal deployment time, simultaneously makes it developer friendly and improves the quality of the application.

- Abstracted technologies like Nix and Docker.
- Caching of dependencies and steps to build.
- Customizable at most of the steps.
- Extensible and Developer friendly.

## Installing Nixpacks

Nixpacks are the [nix packges](https://search.nixos.org/packages) which are used with the source code to create a buildpack of their own. The nix packages take in the source code of your application and convert it into a simple OCI image e.g. Docker image that can be run on various environments. It is similar to buildpacks but it's not the same instead it is better and uses a different ecosystem.

Railway has a cool [CLI](https://nixpacks.com/docs/cli) tool for creating nixpacks on your local system. You can install the nixpack cli from the official documentation site. There are a couple of ways to install them on your system. 
You can install the nixpack with the curl command in your terminal which is one of the simplest ways to install it.

```
curl -fsSL https://raw.githubusercontent.com/railwayapp/nixpacks/master/install.sh | bash
```

![Railway Nixpacks Install](https://res.cloudinary.com/techstructive-blog/image/upload/v1657725394/blog-media/railway-nixpacks-install.png)

Once it has been installed we can check the status if nixpacks was successfully installed or not.

```
nixpacks --version

OR 

nixpacks
```

So that's how we install nixpacks CLI into the system. Now, we can move into the configuration required for creating a nixpack from the Django application.

## Creating a Django Project

So, I assume here, you have your django project ready. You can pick any django project and configure it as mentioned below.

### Creating requirements file

It is a good practice to include the `requirements.txt` file in any python based applications. It becomes really easy to set up and give the project a spin. It's often the case, that we have to use a virtual environment for creating a pip file, otherwise, we might conflict and mix up the globally installed packages with project-specific dependencies.

```
pip freeze > requirements.txt
```

This will create a `requirements.txt` file in the current folder, which will contain the list of all the dependencies with the version mentioned in. We also have other package management system like [pipenv](https://www.meetgor.com/pipenv-intro/) or [poetry](https://python-poetry.org/).


## Creating a Nixpack for the project

Finally, we can now start creating nixpacks from the source code of the django application. So, make sure the Django project is running on your local system. Set up a virtualenv and database as per your local environment. We will be using the build command for creating the nixpack of our django project.

```
nixpacks build .
```

<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1657881139/blog-media/nixpacks-demo-base.mp4" type="video/mp4">
</video>

This is the simplest of commands you can run to create a nixpack. You need to be in the source file where all your files are located. In Django, we call it the `BASE_DIR`, the same folder where your `manage.py` resides. So, be on that path and run the above command. It would pick up a few things from the source code itself. There are also some default values picked up after analyzing the source code, for example in the case if it detects the django project, it attaches the `python manage.py migrate && gunicorn <project_name>.wsgi` command as the default build command.

Similarly, there are other default options like the version of the runtime, installation steps, etc. We can see that the port is not accessible from the docker container that we ran it is because of two reasons either we have not exposed the port or the port is not bound with gunicorn. We will see that configuration later, it's quite easy to bind the local port to the docker container.

### Steps for creating Nixpacks

Now, we need to understand the process of the creation of nixpacks. There were several stages that you can see in the above video. 

1. Planning 
2. Building

#### Planning Phase

The first step was quite important, to understand the runtime environment. This is critical because it can decide how our application installs, runs, or even performs in the actual container environment. This is just the phase where we define the sets of commands or nix packages to be used for the building of the application.

You would get a clear look in the nixpack cli, it gives a list of the planned instructions before the build process starts in the build command. You can get the plan before building the application by using the `plan` command in the CLI.

```
nixpacks plan .
```
![Nixpacks Plan Command](https://res.cloudinary.com/techstructive-blog/image/upload/v1657960771/blog-media/nixpacks-plan-command.gif)

![Nixpacks Plan Command Output](https://res.cloudinary.com/techstructive-blog/image/upload/v1657960816/blog-media/nixpacks-plan-cmd.png)

So, this gives a `nix` output of all the configuration needed to build and install the application and it's dependencies. When you would have built the application i.e. the base build command, it also shows in the CLI, some useful information before moving to the build step.

![Nixpacks Plan Steps in Build Command](https://res.cloudinary.com/techstructive-blog/image/upload/v1657960880/blog-media/nixpacks-build-plan.png)

So the list of the following procedure is listed in the planning phase of building the application.

- Packages/base runtime environment
- Installation commands
- Build commands
- Commands to Start Application 

In the case of django, the following correspond to the planning attributes.

- Python version as the application runtime.
- Installing dependencies via pip/pipenv/poetry
- Building the Django app (collecting static files, setting database, etc)
- Running the Django app (gunicorn/Nginx web server to run the app)

So, hopefully, this gives you a better understanding of what is going on in the nixpack CLI. We are trying to automate the process of deployment and building the application for repeated deployments i.e. in a continuous integration/delivery system.

#### Building Phase

This is the phase, where actual installation, setup, and configuration takes place at the application level. In the build phase, we have several layers where things happen like installing the core packages for the application, installing dependencies, setting/generating up the necessary steps for the proper running of the application, and finally steps to run the application. This process creates an OSI(a standard used in containers/virtualization technology) image of the application. We can run the image and thus create a container of the application. You will need docker or any other containerization tool for the building of images and running containers.

![Django Build Command](https://res.cloudinary.com/techstructive-blog/image/upload/v1657961801/blog-media/nixpacks-build-command.gif)

For the Django application, we have several steps of the application to be followed. We can use python as the base runtime environment, install the python packages, perform other commands like setting up admin accounts, static/media files, database setup, and finally run the steps to run the application at the container level.

- Installing python as a [Nix Package](https://search.nixos.org/packages?channel=22.05&show=python38&from=0&size=50&sort=relevance&type=packages&query=python)
- Installing all python packages provided in the `requirements.txt`/`Pipfile`/`pyproject.toml` files.
- Call commands from the environment to set up the applications like `createsuperuser`, `collectstatic`, `makemigrations`, pull data, management commands, etc. There are a lot of things that can be done here depending on the application.
- Finally, the step to run the Django app, usually gunicorn/Nginx server is used for running the django application on a port with the host.

So, this is what the build phase does, this is the whole and soul of the nixpacks. We literally do every installation setup and configuration of the application. Though the planning phase is equally important, a single missing data can ruin the build phase.


### Understanding the build phase for Django

We ran the build command for creating the nixpack build, this started with planning the application configuration and then building up the application. This build phase was also divided into further processes like installing, running commands, copying actual source code to an image, and all the docker-related stuff that is required to create an image for a django application.

![Django Application build nixpacks](https://res.cloudinary.com/techstructive-blog/image/upload/v1657961691/blog-media/nixpacks-build-process.png)

So, we have used the Debian nixpack, which sets as the base runtime for the application. Railway provides a [package](https://github.com/railwayapp/nixpacks/pkgs/container/nixpacks) of the Debian image as the base runtime for our application. This is where we will run all the build processes on. This Debian image will be used for installing all types of dependencies and layers of language-specific runtime installation in the form of [nix packages](https://search.nixos.org/packages). 

Now, we have the base image of debian, the nixpack build command actually starts executing the `Dockerfile`, this is an auto-generated step after the planning phase. So, with the instructions in `Dockerfile`, steps are executed one after other just as a normal Docker image build. This will generate the image for the application and after a while, because this process takes a while on the first iteration locally, after the build process has been completed, it will give a container name for you to run. This can be used to test the application locally, in production, the container is instantly created after the image has been generated.

```
RUN:

docker run -it <container_id_or_name>

```

![Nixpacks build command run container](https://res.cloudinary.com/techstructive-blog/image/upload/v1657966274/blog-media/nixpacks-build-run.png)

This is the command for running your application, this marks the end of the build process and also the build command provided with the CLI.

### Build Command Parameters

The build command in the nixpacks CLI provides a few parameters or arguments to customize how to output the result and build the application, you can definitely provide the configuration in the application source code itself, but it is nicer to have it locally before deploying the application.

#### Give a name to the Nixpack Image/Container

The first parameter which might be helpful is to provide a name to the application at the build time. This becomes useful for running the container, this helps in avoiding long container names and giving a context of the nixpack.

```
nixpacks build . --name <project_name>

OR 

nixpacks build . -n <project_name>

```

![Nixpacks Build Command Name Image](https://res.cloudinary.com/techstructive-blog/image/upload/v1657962253/blog-media/nixpacks-build-name.png)

![Nixpacks Build Command name run](https://res.cloudinary.com/techstructive-blog/image/upload/v1657962328/blog-media/nixpacks-build-name-image-run.png)

This gives a name to the image which has been built. Thereby providing a better context for the user to run the image and create a container out of it.

#### Output the Built Image to a folder

This is the command that can output the built application into a  provided folder. This parameter will not run the docker step thereafter; hence, no image is created if you provide an output folder. Though the folder will contain the `Dockerfile` and `environment.nix` files for creating the image and running the container. **Make sure the output folder is NOT in the application folder itself, it will result in errors.** The output command will not create an image but the process will be definitely executed in order to generate the `Dockerfile` and `environment.nix` files.

```
nixpacks build . --out ../blog_image

OR

nixpacks build . -o ../blog_image
```

![Nixpacks Build Command Output folder](https://res.cloudinary.com/techstructive-blog/image/upload/v1657962407/blog-media/nixpacks-build-output.png)


**Dockerfile**

![Nixpacks Build command ouptut](https://res.cloudinary.com/techstructive-blog/image/upload/v1657962479/blog-media/nixpacks-build-output-folder.png)


**environment.nix File**

![Nixpacks environment.nix file](https://res.cloudinary.com/techstructive-blog/image/upload/v1657969127/blog-media/nixpacks-build-environment-nix-file.png)

So, this will output the built application into the provided path. The output folder should necessarily be out of the application folder as it makes no sense to output in the same folder as the application since the nixpacks CLI will consider the folder as the application folder.

### Provide a Install/Build/Start Command

We can provide the commands at the install phase/build/start phase of the application to the build command in order to build the app with non-default or custom commands. This will add up to the docker steps that will involve making the build for the application.

```
nixpacks build . --build-cmd 'python manage.py collectstatic'

OR

nixpacks build . -b 'python manage.py collectstatic'
```

![Nixpacks Build Command Providing install/build/start commands](https://res.cloudinary.com/techstructive-blog/image/upload/v1657962514/blog-media/nixpacks-build-build-command.png)

These kinds of parameters can be passed similarly for `install-cmd` and `start-cmd` as `-i` and `-s` respectively. We can further chain up the commands and customize the build process as per the application's requirements.

### Providing environment variables to image

The environment variable can be passed to the build command for parsing to the application. This can be used for parsing additional or optional environment variables to the application image.

```
nixpacks build . --env 'NAME=VALUE'

nixpack build . --env 'DATABASE_URL=postgres://postgres:postgres@localhost:5432/techstructive_blog'
```

![Nixpacks Build Comand parsing environment variables](https://res.cloudinary.com/techstructive-blog/image/upload/v1657963255/blog-media/nixpacks-build-env-variable-db-url.png)

![Nixpacks Build Command env variable migrate](https://res.cloudinary.com/techstructive-blog/image/upload/v1657963302/blog-media/nixpacks-build-db-migrate-env.png)

Here, we provide the environment variable `DATABASE_URL` to the build command and this is parsed to the application image. Thereby when the image is run as a container, it is parsed as a normal environment variable and thereby is available for utilization from the application.

For further references on the build command arguments, you can follow the [documentation of nixpack](https://nixpacks.com/docs/cli) by railway app.

### Creating a Procfile

This is important for telling any buildpack in this case nixpacks to understand the process to start for this web application. For django it is simply to add the web process like to mention the `wsgi` app with the project name. We can use the gunicorn as the web server in production.

```

# Procfile

web: gunicorn <django_project_name>.wsgi

```

This is the Procfile, this is a file type without the extension. So, this is a typical Django application Procfile, you can also use the other variants of Procfile for applying migration for every web process start-up.

```

# Procfile

web: python manage.py migrate && gunicorn <django_project_name>.wsgi

```

As we saw in the base build command, the local server was not able to listen to the gunicorn server in the container, so we need to bind the gunicorn server to the local port.

```

# Procfile

web: python manage.py migrate && gunicorn <django_project_name>.wsgi -b :8000

```

So, we use the `-b` option in the gunicorn command to bind the port in the container to the port in the local machine. Now, if we build the application and forward the port to the `8000` port in the local machine, we will see our application running 

<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1657964597/blog-media/nixpacks-local-bind.mp4" type="video/mp4">
</video>

```
docker run -p 8000:8000 -it <container_id> 

```

### Specifying the Python version

This is created for specifying the python version for building the Django application or any other python app.

```
# .python-version

3.10
```

Save the `.python-version` file with just the python version like `3.6`, `3.10`, etc. and this will be picked by the build command while creating the build image.

## Deploying the Django Application

After we looked at the nix picks specifications, we can now deploy our Django application with nixpacks on Railway. So, you can follow up with the [Railway Deployment]() Article for setting up your Django app for deployment at the railway. This usually involves a few steps like creating Procfile(not necessary but recommended), requirements.txt(necessary to pull dependencies), and the python version which is chosen as `3.8` as default. The further steps are to create a GitHub repository to link with the Railway app and create a PostgreSQL database service on the railway platform.

### Create configuration files

As we have seen we will require a `requirements.txt` file, `Pipfile` or a `pyproject.toml` file for listing out and installing dependencies for our django application. This can be done with various commands like:

```
# For requirements.txt and virtualenv
pip freeze > requirements.txt

# Autogenerated Pipfile for pipenv
# Autogenerated pyproject.toml for poetry

```

So, this file should be present on the base directory of the django application in order for the nixpack to pick up and install the python packages. Also, for customization of the start command in the build process, you can create a `Procfile` as discussed earlier in order to run commands to start the Django web server.

The python version can be specified with the `.python-version` file with just the version name as `3.9`, `3.10`, etc. OR we can add an environment variable `NIXPACKS_PYTHON_VERSION` to the python version we want.

### Create and Linkup a GitHub repository for existing Django projects

We can create a GitHub repository and link up the project to the Railway platform and thereby creating an automated build for every push. 

The below video will explain how to set up the GitHub repository for the Railway app.

<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1652970864/blog-media/django-deployment/railway_project_init.webm" type="video/mp4">
</video>

<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1657965284/blog-media/railway-platform-github.mp4" type="video/mp4">
</video>

### Setup environment variables 

We can use `python-environ` to set up environment variables in a Django application, we will require these environment variables for attributes like `SECRET_KEY`, `DATABASE_URL`, email-stuff, etc. These are quite handy to avoid leaking sensitive information to the open source project on GitHub.

You can install the `python-environ` package with pip or any other package management tool as follows:

```
pip install python-environ
```

After installing the package, we can set up the environment variable in the settings file.

``` python
# <project_name>/settings.py

import os
from dotenv import load_dotenv

BASE_DIR = Path(__file__).resolve().parent.parent

load_dotenv(os.path.join(BASE_DIR, ".env"))
```

After loading the environment variables, we can access them with `os.env("ENV_NAME", default="")`, this will load the environment variable with the name or we can provide a default value.

### Attach a PostgreSQL database service

You can add a PostgreSQL database service to your Django Railway app by attaching a service. This will add a new service along with the django application, so these two act as different entities within a railway project. 

<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1652963718/blog-media/django-deployment/postgres_spinup_railway_d2xkpt.mp4" type="video/mp4">
</video>

You can then access the `DATABASE_URL` from the connect settings and copy the database URL and set it as an environment variable in the django railway project service. This will link up the Django app and the PostgreSQL database. While setting it up locally, you can use the `.env` file and add the environment variable there.

```
# environment variable
DATABASE_URL=postgres://username:password@hostname:port/db_name

# local database postgres
DATABASE_URL=postgres://postgres:postgres@localhost:5432/db_name
```

<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1657964943/blog-media/railway-postgres-spinup.mp4" type="video/mp4">
</video>


This will set up the database environment variable and you can access these from the settings.py file with the `env.db` function as follows:

```python
env.db("DATABASE_URL", default=<local_database_url>)
```

So, we can finally use the database from the Railway app in our Django application once the environment variable is correctly used.

### Choose the Buildpack

We can choose a buildpack for our Django application in the Railway platform, we have options like 

1. Heroku Buildpack
2. Railway Nixpacks
3. Paketo Buildpack

As of the writing of the article, on 16th July 21, the Railway has made `Nixpacks` the default buildpack for an application :) It was the `Heroku` Buildpack as a default earlier. So, that is a cool thing, you can toggle these settings for choosing the buildpacks from the project settings.

Railway Dashboard Choose BuildPack

<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1657964699/blog-media/nixpacks-railway-dashboard.mp4" type="video/mp4">
</video>

### Deploy to Railway with Nixpacks

Now, we have seen how to set up the nixpacks, we had the Postgres database setup, and we can finally deploy our application to the railway platform with nixpacks.


<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1657965560/blog-media/railway-nixpacks-deploy.mp4" type="video/mp4">
</video>

So, we simply can configure the source code or include the `environment.nix` file into the source code for desired behavior. The nixpack selection can be done based on the source code or the presence of `environment.nix` and that's why we can rely on expected behaviors from the deployment builds. 

## Summary

So, nixpacks is a great way to deploy an application, for me it's an automated docker deployment, it basically creates docker images of the application and runs it with the appropriate environment. There is a lot of language support on nixpacks currently on Railway, you can check them out on the official website. Every programming language has specific requirements for managing or installing dependencies and packages, the nixpacks manage them automatically for us.

## Conclusion

So, from this post of the [Django Deployment Series](https://www.meetgor.com/series/django-deployment/), we were able to understand how to deploy a Django application on the Railway app with Nixpacks which are a very intuitive way to deploy apps. We covered what are nixpacks, the process of building an application with nixpacks, and deploying a existing, new Django project on the railway with nixpacks. We also explored the various commands provided in the nixpacks CLI to build. plan a Django application.

Hopefully, you were able to understand the concept of nixpacks and how they can automate the process of containerization and deployment. Thank you for reading, if you have any queries or feedback, you can leave them down in the comments or freely drop them on social media. Happy Coding :)
