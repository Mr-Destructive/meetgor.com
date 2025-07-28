{
  "title": "Django Basics: Setup and Installation",
  "status": "published",
  "slug": "django-basics-setup",
  "date": "2025-04-05T12:33:50Z"
}

<h2>Introduction</h2>
<p>The crucial aspect of starting to learn any framework is the ease to set it up and Django by far is the easiest of the options out there. There is just a few lines of code to install django if you already have python installed in your system. In this article, we see how to setup a django project along with a virtual environment.</p>
<p>If you already have python and pip installed, you can move on to the <a href="#setting-up-virtual-environment-in-python">virtual environment setup</a>.</p>
<h2>Installing Python and PIP</h2>
<p>Django is a python based framework so that makes sense to have Python installed along with its package manager to use Django.</p>
<p>To install Python, you can visit the official <a href="https://www.python.org/downloads/">Python</a> website to download any relevant version for your system (recommended 3.7 and above).</p>
<p>Mostly the Python installation comes with the option to install <code>pip</code>(python's package manager) but if you missed that, that's fine, you can install the <a href="https://bootstrap.pypa.io/get-pip.py">get-pip.py</a> file into your system and run the below code:</p>
<pre><code>python get-pip.py   
</code></pre>
<p>Make sure the include the relative path to the file if you are not in the same folder as the file.</p>
<p>So, that should be python setup in your local machine. To check that python was installed correctly, type in <code>python --version</code> and <code>pip --version</code> to check if they return any version number. IF they do, Congratulations !! You installed Python successfully and if not, don't worry there might be some simple issues that can be googled out and resolved easily.</p>
<hr>
<p>Let's move on to the actual setting of the Django project set up.</p>
<h2>Setting up Virtual Environment in python</h2>
<p>Virtual Environment is a software which isolates the installation of dependencies and libraries for a specific project, making it a clean and safe environment for deployment as well as maintenance.</p>
<p>In Python, we have a virtual environment package known as <code>virtualenv</code> that does this thing. It is for installing the Python related packages into a isolated folder. So, we can install the <code>virtualenv</code> package in python by following the following steps:</p>
<h3>Installing Virtualenv</h3>
<p>Firstly, install the virtual environment package, it's not mandatory but it keeps things simple and easy for your project in correspondence to the entire OS. So in python, we have a module to create the virtual environment pretty easily,</p>
<pre><code>pip install virtualenv
</code></pre>
<p>You can use <code>pip3</code> or <code>pip -m</code>, or however you install normal python modules. This just installs the python virtual environment, we need to create a virtual environment in the current folder.</p>
<h3>Creating a virtual environment</h3>
<p>We need to create the environment so as to give the Python interpreter an indication to consider the current folder as an isolated Python environment. We need to create a virtual environment in the current folder, so for that navigate to the folder where you want to create the project and enter the following command:</p>
<pre><code>virtualenv venv
</code></pre>
<p>Here, <code>venv</code> can be anything like <code>env</code> just for your understanding and simplicity it's a standard name kept for the same. After this, you will see a folder of the same name i.e. <code>venv</code> or any other name you have used. This is the folder where python will keep every installation private to the local folder itself.</p>
<h3>Activating Virtual environment</h3>
<p>Now, we need to activate the virtual environment, this means that any thing installed in the prompt with the virtualenv activated will be isolated from the entire system and will be installed on in the virtual environment. To activate the environment, we can use the command :</p>
<h4>for Linux/macOS :</h4>
<pre><code>source venv/bin/activate
</code></pre>
<h4>for Windows:</h4>
<pre><code>venv\Scriptsctivate
</code></pre>
<p>After this, your command prompt will have a <code>(venv)</code> attached in the beginning. This indicates you are in a virtual environment, things you do here, may it be module installation or any configuration related to python will stay in the local folder itself.</p>
<h2>Installing Django</h2>
<p>After the virtual environment is set up and activated, you can install Django and get started with it. Django is a python module or package, which can be easily installed using its package manager <code>pip</code>.</p>
<p>Install Django using pip:</p>
<pre><code>pip install django
</code></pre>
<h2>Create a Django Project</h2>
<p>After the installation is completed, you can start a Django project in the current folder from the django package we installed. There are several commands available in the django module which you can execute in the command line that we'll discuss later.
For now, we will use the command <code>startproject</code> this is one of the <a href="https://github.com/django/django/tree/main/django/core/management/commands">management commands</a> in Django. The <a href="https://docs.djangoproject.com/en/3.2/ref/django-admin/">django-admin</a> is a command line utility for doing the administrative tasks related to Django.</p>
<pre><code>django-admin startproject myproject
</code></pre>
<p>Here <code>myproject</code> can be your project name. After this, you will see one new folder and one file pop up.</p>
<p>Namely, the <code>&lt;project-name&gt;</code> folder and <code>manage.py</code> file. We don't have to touch the <code>manage.py</code> file but we use it in most of the commands to use the Django functionalities, it is quite similar to the <code>django-admin</code> command.</p>
<p>You can now run your basic server using the command :</p>
<pre><code>python manage.py runserver
</code></pre>
<p>OR</p>
<p>You can use <code>djagno-admin</code> command, but you need to set certain environment variables and modify the settings.py file as per the project-name. You can use the <code>django-admin</code> as the steps given in the django <a href="https://docs.djangoproject.com/en/3.2/ref/django-admin/#cmdoption-settings">documentation</a>.</p>
<p>The output of the command <code>python manage.py runserver</code> should be visible in the browser at <code>https://127.0.0.1:8000</code> as below :</p>
<p><img src="https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/screenshotr_2021-11-20T15-40-50.png" alt="Django-Base-Project"></p>
<p>That's it the base django project is installed in your system. To stop the server simply press <code>Ctrl+C</code>.</p>
<p>Follow the below GIF for a clear understanding of those instructions:</p>
<p><img src="https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/djp2.gif" alt="Django-basics-part2-setup"></p>
<hr>
<h2>Quick-Setup-Script</h2>
<p>You can avoid manually typing the commands once you get the idea of the process in setting up a django project by executing a simple shell script (for Linux/macOS) or a batch script (for Windows). The script looks something like:</p>
<p>For Linux/macOS:</p>
<pre><code class="language-shell">#!/usr/bin/env bash

mkdir $1
cd $1
pip install virtualenv
virtualenv venv
source venv/bin/activate

pip install django
django-admin startproject $1 .
clear
</code></pre>
<p>Save as commands.sh file</p>
<p>For Windows:</p>
<pre><code class="language-batch">mkdir %1 
cd %1
pip install virtualenv
virtualenv env
call env\Scriptsctivate

pip install django
django-admin startproject %1 .
cls

</code></pre>
<p>save as commands.bat file</p>
<p>For further instructions you can checkout the <a href="https://github.com/Mr-Destructive/django-quick-setup-script">GitHub repository</a> or a detailed <a href="https://mr-destructive.github.io/techstructive-blog/django/web-development/python/2021/08/15/Django-Quick-Setup.html">article</a> about it.</p>
<h2>Conclusion</h2>
<p>From this section, we were able to setup the Django project in our local system. In the next part, we will cover the <code>folder structure</code> of the Django project. We won't directly go into the code part because that is very easy once you understand the flow of the framework and its internal working. So, thanks for reading and Happy Coding :)</p>
