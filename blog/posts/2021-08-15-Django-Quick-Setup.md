---
templateKey: blog-post
title: "Django Quick Setup Script"
subtitle: "A Quick simple script to set up the Django project."
date: 2021-08-15 18:50:46 +0530
status: published
tags: ['django', 'web-development', 'python', 'bash',]
slug: django-setup-script
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643286698/blogmedia/dyfjizyq2zmdo8fk4ars.png
---

## Introduction

Once you learn something very neatly and understand it very clearly, it feels like you are wasting time doing things over and over again. Same things in setting up a Django project, you feel like you could save those typing minutes by making a script. 

In this article, we will make a script a batch script for windows and a BASH script for Linux/macOS. We will make a virtual environment using python and then install the libraries and dependencies like in this case we will install Django. You can also tinker with these scripts and install other dependencies if you want like Django rest framework, crispy forms, markdown, Redis, etc. We will also make a Django project using the positional parameter passed before running the script from the command line. 

### Python development environment
This article assumes you have a python environment setup. If you don't you must install Python from the  [official website](https://www.python.org/downloads/)  as per your operating system. Also, you should have pip installed and configured correctly. You can install pip by following the  [official documentation](https://pip.pypa.io/en/stable/)  for the specific operating systems.

## Steps in Django Project Setup 

So, If you are already familiar with the Django project setup, you can directly use the scripts provided in the next few sections. You can also visit  [this GitHub repository](https://github.com/Mr-Destructive/django-quick-setup-script)  if you have any issues and errors.  

If you are new to django, let me first explain the process of django project setup. 
- ### Initialize a VirtualEnvironment (Recommended but not necessary)

Virtual Environment in Python is a great way of localizing the dependencies and frameworks only in the particular folder, it allows the developer to separate things out and keep them distinct, Additionally, when deploying or sharing the repository, the other developers can install the dependencies in the requirement.txt file in their local environment flawlessly. 

So, it is always recommended to use python virtualenv when working with python frameworks or libraries. We can set it up by simple pip install and then giving it a name.

```
pip install virtualenv
``` 

This will install the package/tool using pip. 

After that has been properly installed, we can now give it an appropriate name 
```
virtualenv mytest
```

The `virtualenv` is the command and `mytest` can be any name, generally `env` or `venv` is preferred but it should be understandable to the user.  You will now be able to see the folder inside of your current directory named as `mytest` or the name you've given to it. 

**Windows**

Now if you are on windows, you can activate the virtual environment by using the command :
```
mytest\Scripts\activate
``` 
here mytest is the name of your virtual env it can be anything as per your choice. This will now activate the virtualenv which will be shown by `(mytest)` before the command prompt. 

**Linux/macOS**

For Linux or macOS, you can use the command: 
```
source mytest/Scripts/activate.sh
```
In the above command, `mytest` can be anything that you have used while creating the virtualenv. This should activate the vrtualenv and will be indicated by `(mytest)` before the prompt in the terminal. 


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1629023409389/kEe5AVAsr.png)
From the above image, we can see that we created and activated an virtualenv in python in a folder. 
So, this is the walkthrough for setting up the virtualenv for a Django project, now we will move to install Django in this environment.

- ### Installing Django using pip

This is quite straightforward. You can use `pip install django` or `pip3 install django` or the normal way you install a library from pip. 

- ### Creating a Django project

To create a django project, we use the django-admin command like:
```
django-admin startproject mywebsite
```
This will create a folder called `mywebsite` or your project name anything you like. Inside the `mywebsite` folder, you will have 2 things: `manage.py` file, and `mywebsite` folder . Yes there will be another `mywebsite` folder inside your project which will have the settings, URLs and other global(project-level) configuration files. The `manage.py` file is the most important file here. You should never touch/edit this file. We use this file to carry out all the operations from running the server to managing the database. 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1629032998253/QQ5QXf4v5.png)

### Setup script  for Windows

The below is a batch file for Windows Operating System. Save the file in a `.bat` extension.
<iframe
  src="https://carbon.now.sh/embed?bg=rgba%28171%2C+184%2C+195%2C+1%29&t=cobalt&wt=none&l=application%2Fx-sh&ds=true&dsyoff=20px&dsblur=68px&wc=true&wa=true&pv=56px&ph=56px&ln=false&fl=1&fm=Hack&fs=14px&lh=133%25&si=false&es=2x&wm=false&code=mkdir%2520%25251%2520%250Acd%2520%25251%250Apip%2520install%2520virtualenv%250Avirtualenv%2520env%250Acall%2520env%255CScripts%255Cactivate%250A%250Apip%2520install%2520django%250Adjango-admin%2520startproject%2520%25251%2520.%250Acls%250A"
  style="width: 803px; height: 366px; border:0; transform: scale(1); overflow:hidden;"
  sandbox="allow-scripts allow-same-origin">
</iframe>

Make sure the file is saved in a `.bat` file and be in the folder where you would like to create the Django project. After being in the appropriate location, enter the following command:

```
commands.bat myproj
```
Here I assume that you have named the script file as `commands.bat`, you can name it anything you like, but I like to keep this a convention. After this you don't need to do anything, everything will be handled by the script. 
You can run the server using 
```
python manage.py runserver
```
This will have the base django project set up on your system. The below is the live demonstration of the script, I have deliberately removed the `cls` command to show the process. It won't break the script if you add this to it. 

![djqss.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1629024613612/Zsqa54_KD.gif)

### Setup script  for Linux/macOS

Copy the code from the below script and save it in a file preferably called `commands.sh`, you can name it anything you want but keep the `.sh` extension after it to identify it as a shell-script.

After that file is saved locally, you can run the script by passing the positional parameter as the name of the Django project. The command will be like:

```
bash commands.sh myproj
```
 
<iframe
  src="https://carbon.now.sh/embed?bg=rgba%28171%2C+184%2C+195%2C+1%29&t=blackboard&wt=none&l=application%2Fx-sh&ds=true&dsyoff=20px&dsblur=68px&wc=true&wa=true&pv=56px&ph=56px&ln=false&fl=1&fm=Hack&fs=14px&lh=133%25&si=false&es=2x&wm=false&code=%2523%21%252Fusr%252Fbin%252Fenv%2520bash%250A%250Amkdir%2520%25241%250Acd%2520%25241%250Apip%2520install%2520virtualenv%250Avirtualenv%2520env%250Asource%2520env%255Cbin%255Cactivate%250A%250Apip%2520install%2520django%250Adjango-admin%2520startproject%2520%25241%2520.%250Aclear"
  style="width: 803px; height: 384px; border:0; transform: scale(1); overflow:hidden;"
  sandbox="allow-scripts allow-same-origin">
</iframe>

From the output of the script, you will have created a Django project inside a virtual environment. We can manually activate the virtual environment. You can experiment it within your system as it can be a lot more customizable. This is just bare bone script to start a Django project but you can add your own things into it.

## Conclusion
Thus, from this little article, you can get a bit lazier in initializing a bare-bone Django project. We were able to understand the structure of the Django project and how to set up a virtual environment powered by Python. 

After understanding those concepts we then moved on to making a script namely a batch file and a shell script to automate the initialization of the Django project. I hope it might have helped in some or another way, Thanks for reading till here. Happy Coding :)

