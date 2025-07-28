{
  "title": "Django Quick Setup Script",
  "status": "published",
  "slug": "django-setup-script",
  "date": "2025-04-05T12:33:56Z"
}

<h2>Introduction</h2>
<p>Once you learn something very neatly and understand it very clearly, it feels like you are wasting time doing things over and over again. Same things in setting up a Django project, you feel like you could save those typing minutes by making a script.</p>
<p>In this article, we will make a script a batch script for windows and a BASH script for Linux/macOS. We will make a virtual environment using python and then install the libraries and dependencies like in this case we will install Django. You can also tinker with these scripts and install other dependencies if you want like Django rest framework, crispy forms, markdown, Redis, etc. We will also make a Django project using the positional parameter passed before running the script from the command line.</p>
<h3>Python development environment</h3>
<p>This article assumes you have a python environment setup. If you don't you must install Python from the  <a href="https://www.python.org/downloads/">official website</a>  as per your operating system. Also, you should have pip installed and configured correctly. You can install pip by following the  <a href="https://pip.pypa.io/en/stable/">official documentation</a>  for the specific operating systems.</p>
<h2>Steps in Django Project Setup</h2>
<p>So, If you are already familiar with the Django project setup, you can directly use the scripts provided in the next few sections. You can also visit  <a href="https://github.com/Mr-Destructive/django-quick-setup-script">this GitHub repository</a>  if you have any issues and errors.</p>
<p>If you are new to django, let me first explain the process of django project setup.</p>
<ul>
<li>
<h3>Initialize a VirtualEnvironment (Recommended but not necessary)</h3>
</li>
</ul>
<p>Virtual Environment in Python is a great way of localizing the dependencies and frameworks only in the particular folder, it allows the developer to separate things out and keep them distinct, Additionally, when deploying or sharing the repository, the other developers can install the dependencies in the requirement.txt file in their local environment flawlessly.</p>
<p>So, it is always recommended to use python virtualenv when working with python frameworks or libraries. We can set it up by simple pip install and then giving it a name.</p>
<pre><code>pip install virtualenv
</code></pre>
<p>This will install the package/tool using pip.</p>
<p>After that has been properly installed, we can now give it an appropriate name</p>
<pre><code>virtualenv mytest
</code></pre>
<p>The <code>virtualenv</code> is the command and <code>mytest</code> can be any name, generally <code>env</code> or <code>venv</code> is preferred but it should be understandable to the user.  You will now be able to see the folder inside of your current directory named as <code>mytest</code> or the name you've given to it.</p>
<p><strong>Windows</strong></p>
<p>Now if you are on windows, you can activate the virtual environment by using the command :</p>
<pre><code>mytest\Scriptsctivate
</code></pre>
<p>here mytest is the name of your virtual env it can be anything as per your choice. This will now activate the virtualenv which will be shown by <code>(mytest)</code> before the command prompt.</p>
<p><strong>Linux/macOS</strong></p>
<p>For Linux or macOS, you can use the command:</p>
<pre><code>source mytest/Scripts/activate.sh
</code></pre>
<p>In the above command, <code>mytest</code> can be anything that you have used while creating the virtualenv. This should activate the vrtualenv and will be indicated by <code>(mytest)</code> before the prompt in the terminal.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1629023409389/kEe5AVAsr.png" alt="image.png">
From the above image, we can see that we created and activated an virtualenv in python in a folder.
So, this is the walkthrough for setting up the virtualenv for a Django project, now we will move to install Django in this environment.</p>
<ul>
<li>
<h3>Installing Django using pip</h3>
</li>
</ul>
<p>This is quite straightforward. You can use <code>pip install django</code> or <code>pip3 install django</code> or the normal way you install a library from pip.</p>
<ul>
<li>
<h3>Creating a Django project</h3>
</li>
</ul>
<p>To create a django project, we use the django-admin command like:</p>
<pre><code>django-admin startproject mywebsite
</code></pre>
<p>This will create a folder called <code>mywebsite</code> or your project name anything you like. Inside the <code>mywebsite</code> folder, you will have 2 things: <code>manage.py</code> file, and <code>mywebsite</code> folder . Yes there will be another <code>mywebsite</code> folder inside your project which will have the settings, URLs and other global(project-level) configuration files. The <code>manage.py</code> file is the most important file here. You should never touch/edit this file. We use this file to carry out all the operations from running the server to managing the database.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1629032998253/QQ5QXf4v5.png" alt="image.png"></p>
<h3>Setup script  for Windows</h3>
<p>The below is a batch file for Windows Operating System. Save the file in a <code>.bat</code> extension.</p>
<!-- raw HTML omitted -->
<p>Make sure the file is saved in a <code>.bat</code> file and be in the folder where you would like to create the Django project. After being in the appropriate location, enter the following command:</p>
<pre><code>commands.bat myproj
</code></pre>
<p>Here I assume that you have named the script file as <code>commands.bat</code>, you can name it anything you like, but I like to keep this a convention. After this you don't need to do anything, everything will be handled by the script.
You can run the server using</p>
<pre><code>python manage.py runserver
</code></pre>
<p>This will have the base django project set up on your system. The below is the live demonstration of the script, I have deliberately removed the <code>cls</code> command to show the process. It won't break the script if you add this to it.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1629024613612/Zsqa54_KD.gif" alt="djqss.gif"></p>
<h3>Setup script  for Linux/macOS</h3>
<p>Copy the code from the below script and save it in a file preferably called <code>commands.sh</code>, you can name it anything you want but keep the <code>.sh</code> extension after it to identify it as a shell-script.</p>
<p>After that file is saved locally, you can run the script by passing the positional parameter as the name of the Django project. The command will be like:</p>
<pre><code>bash commands.sh myproj
</code></pre>
<!-- raw HTML omitted -->
<p>From the output of the script, you will have created a Django project inside a virtual environment. We can manually activate the virtual environment. You can experiment it within your system as it can be a lot more customizable. This is just bare bone script to start a Django project but you can add your own things into it.</p>
<h2>Conclusion</h2>
<p>Thus, from this little article, you can get a bit lazier in initializing a bare-bone Django project. We were able to understand the structure of the Django project and how to set up a virtual environment powered by Python.</p>
<p>After understanding those concepts we then moved on to making a script namely a batch file and a shell script to automate the initialization of the Django project. I hope it might have helped in some or another way, Thanks for reading till here. Happy Coding :)</p>
