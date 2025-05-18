{
    "type": "til",
    "title": "Dockerize a Django project",
    "date": "2022-03-02 19:12:51",
    "status": "published",
    "tags": ["django", "docker"],
    "slug": "dockerize-django-prj"
}

## Dockerize a Django project

We can run our Django projects in a Docker Container by creating a Docker image for our project. It is really easy and intuitive to create a Dockerfile for any given application as it really is a matter of writing commands in a file and basically running it is a isolated environment. To create a Docker image, we first need a Dockerfile. A Dockerfile is simply a Blueprint to create a image in Docker. In this file we specify the instructions/commands/environment variables to create a image for our app to run. 

### Create a Docker File

To create a dockerfile, simply create a file named as `Dockerfile` without any extension(not a .txt file). We need to then add the following code into the file. 

```dockerfile
FROM python:3.9-buster

ENV PYTHONUNBUFFERED=1

WORKDIR /code

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

EXPOSE 8000

CMD ["python","manage.py","runserver","0.0.0.0:8000"]
```

- Let's see what are we doing here, first we are pulling a base image of Python in this case it is [3.9-buster](https://github.com/docker-library/python/blob/a4b368154b7e3c33c76385f1be7a998fcf3123eb/3.9/buster/Dockerfile), but it can be any of the [Official Python images](https://hub.docker.com/_/python) from [dockerhub](https://hub.docker.com). It is simply a environment for our Django project to run. You can even use linux environments like ubuntu, debian or alpine. 

- Next we add a `PYTHONUNBUFFERED` as a environment variable and initialize it to `1`, which basically allows to parse python output straight into the terminal without actually buffering first. 

- We set the `WORKDIR` as the `code` directory, this sets the provided directory as the base to run any command or instructions in the Dockerfile. 

- We then `COPY` the `requirements.txt` file in the current(`.`) directory which is `code` as mentioned earlier. We simply want the `requirements.txt` file in the current directory so we can proceed with the next step which is to install the dependencies for the Django project

- After copying the `requirements.txt` file, we simply install the packages mentioned int the file with `pip` we use the `RUN` command to execute any shell commands in the environment. 

- Next step is to `COPY` all the files from the current folder(local machine) into the docker environment image. So, we have the source code files in the Container. 

- Expose the port `8000`(default) or any other port you would want to run Django in the image. We use the expose command in complement to the `-p` argument when we want to create the container. I have explained the [port-forwarding](https://mr-destructive.github.io/techstructive-blog/docker-port-forward) concept in the previous TIL. 

- Finally we can run the server. We need to parse every command as a comma separated string in the list to the `CMD` which the container actually executes the command when the image is build. So, all the commands previously were to build an image for the django project, the `CMD` will actually run the server in a container after building the image of the django project.  

### Build the image

Using the Dockerfile, we can create the image for the Django project by a simple command.

```
docker build . -t django-app
```

Assuming the Dockerfile is in the same folder from which you are running this command.

![Build the django-app image](https://res.cloudinary.com/techstructive-blog/image/upload/v1646230907/blog-media/jj04subyvkuvfb5obytu.png)

![Build Complete Success](https://res.cloudinary.com/techstructive-blog/image/upload/v1646230988/blog-media/ugjoakqgyhiwelqkyaat.png)

### Run the Image and Create a Container

```
docker run -p 8000:8000 django-app
```

![Create Container](https://res.cloudinary.com/techstructive-blog/image/upload/v1646231023/blog-media/yneuz46burorz4b5vzp4.png)

You can use any port like `3000` or `4000` on your local system by changing the first port number as `3000:8000` and so on. So, the first port number is the local system's port number and the port number followed by a `:` is the port number for the environment inside the container of the django project. The container port is the same which we have `exposed` while creating the image. You can now go to [127.0.0.1:8000](127.0.0.1:8000) or your chosen port in your local environment and see the django app running. 

If you are using a Docker as a Virtual Machine, you need to find the IP-address of that Virtual-Machine(`docker-machine ip`) and use that instead of `localhost` or `127.0.0.1`. 


### Cleaning the Container for iterations

We need to stop the container to run it again after making a few changes in the app if any. Simply pressing the `CTRL + D` option wont stop the container here. We need to stop the container by executing: 

```
docker ps

docker stop <container-id>
```

This will stop the container. Also if you want to remove the images you built, you can run the command:

```
docker image ls

docker rmi -f <image-id>
```

  So, we now have a image of our Django project, this image can be used by anyone inside a docker environment and thus creating much more easier to test/work on a given project irrespective of the system is being used.

