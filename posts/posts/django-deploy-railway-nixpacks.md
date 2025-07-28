{
  "title": "Deploying Django Project with Railway Nixpacks",
  "status": "published",
  "slug": "django-deploy-railway-nixpacks",
  "date": "2025-04-05T12:33:38Z"
}

<h2>Introduction</h2>
<p>We have seen how to deploy a Django application on railway app in the <a href="https://www.meetgor.com/django-deploy-railway/">previous article</a> of the <a href="https://www.meetgor.com/series/django-deployment/">series</a>. We deployed the django project using the Heroku Buildpacks under the hood. The railway app provides a couple of options on how to build your web application. Currently, there are three options, one of which is the <a href="https://devcenter.heroku.com/articles/heroku-20-stack">heroku buildpacks</a> which is the default one, second, we have the nixpacks which we will see today, and the third is the <a href="https://paketo.io/">Paketo buildpack</a>.</p>
<h2>What is a Buildpack?</h2>
<p>A buildpack is a set of programs that turns your source code into a container image. So it is basically a tool for converting your application into a deployment-ready state with the help of containerization technology. Buildpacks allow us to extract away the steps for deploying an application. The layer of abstraction for converting a source code into a deployable code/container is played by the build packs.</p>
<h3>Benefits of buildpacks</h3>
<p>Buildpacks as we discussed act as a layer of abstraction from converting source code into the deployable containers, so it helps in avoiding manually writing dependencies and installing them. But on top of that, they can even detect low-level changes to the source code, i.e. if a dependency is changed, it would know which dependencies are altered and which components can be fetched from the cache. Caching is a great property of buildpacks which enhances the performance and deployment time. Since we can specify the commands for the build process, the buildpacks are customizable and hence provide a solid foundation for professional applications.</p>
<p>For more references on Buildpacks, you can follow up with the great article on <a href="https://technology.doximity.com/articles/buildpacks-vs-dockerfiles">dockerfile vs buildpacks</a>.</p>
<h2>What are Nixpacks?</h2>
<p><a href="https://nixpacks.com/docs">Nixpacks</a> are quite similar to buildpacks, but they have their own set of technologies used in managing and installing builds for the application. Nixpacks as the name suggests uses <a href="https://search.nixos.org/packages">Nix Packages</a> for creating and installing dependencies and <a href="https://www.docker.com/">Docker</a> for building images and running containers.</p>
<p>Nixpacks are quite cool as most of the applications require little or almost no configuration. You don't need a lot of knowledge of docker or nix technologies, everything is abstracted for you. Nixpacks uses nix packages for installing the runtime environment and the dependencies for the applications. It also caches support for detecting the core modules or packages in the runtime, so it gets a huge boost in deployment performance. Minimal deployment time, simultaneously makes it developer friendly and improves the quality of the application.</p>
<ul>
<li>Abstracted technologies like Nix and Docker.</li>
<li>Caching of dependencies and steps to build.</li>
<li>Customizable at most of the steps.</li>
<li>Extensible and Developer friendly.</li>
</ul>
<h2>Installing Nixpacks</h2>
<p>Nixpacks are the <a href="https://search.nixos.org/packages">nix packges</a> which are used with the source code to create a buildpack of their own. The nix packages take in the source code of your application and convert it into a simple OCI image e.g. Docker image that can be run on various environments. It is similar to buildpacks but it's not the same instead it is better and uses a different ecosystem.</p>
<p>Railway has a cool <a href="https://nixpacks.com/docs/cli">CLI</a> tool for creating nixpacks on your local system. You can install the nixpack cli from the official documentation site. There are a couple of ways to install them on your system.
You can install the nixpack with the curl command in your terminal which is one of the simplest ways to install it.</p>
<pre><code>curl -fsSL https://raw.githubusercontent.com/railwayapp/nixpacks/master/install.sh | bash
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657725394/blog-media/railway-nixpacks-install.png" alt="Railway Nixpacks Install"></p>
<p>Once it has been installed we can check the status if nixpacks was successfully installed or not.</p>
<pre><code>nixpacks --version

OR 

nixpacks
</code></pre>
<p>So that's how we install nixpacks CLI into the system. Now, we can move into the configuration required for creating a nixpack from the Django application.</p>
<h2>Creating a Django Project</h2>
<p>So, I assume here, you have your django project ready. You can pick any django project and configure it as mentioned below.</p>
<h3>Creating requirements file</h3>
<p>It is a good practice to include the <code>requirements.txt</code> file in any python based applications. It becomes really easy to set up and give the project a spin. It's often the case, that we have to use a virtual environment for creating a pip file, otherwise, we might conflict and mix up the globally installed packages with project-specific dependencies.</p>
<pre><code>pip freeze &gt; requirements.txt
</code></pre>
<p>This will create a <code>requirements.txt</code> file in the current folder, which will contain the list of all the dependencies with the version mentioned in. We also have other package management system like <a href="https://www.meetgor.com/pipenv-intro/">pipenv</a> or <a href="https://python-poetry.org/">poetry</a>.</p>
<h2>Creating a Nixpack for the project</h2>
<p>Finally, we can now start creating nixpacks from the source code of the django application. So, make sure the Django project is running on your local system. Set up a virtualenv and database as per your local environment. We will be using the build command for creating the nixpack of our django project.</p>
<pre><code>nixpacks build .
</code></pre>
<!-- raw HTML omitted -->
<p>This is the simplest of commands you can run to create a nixpack. You need to be in the source file where all your files are located. In Django, we call it the <code>BASE_DIR</code>, the same folder where your <code>manage.py</code> resides. So, be on that path and run the above command. It would pick up a few things from the source code itself. There are also some default values picked up after analyzing the source code, for example in the case if it detects the django project, it attaches the <code>python manage.py migrate &amp;&amp; gunicorn &lt;project_name&gt;.wsgi</code> command as the default build command.</p>
<p>Similarly, there are other default options like the version of the runtime, installation steps, etc. We can see that the port is not accessible from the docker container that we ran it is because of two reasons either we have not exposed the port or the port is not bound with gunicorn. We will see that configuration later, it's quite easy to bind the local port to the docker container.</p>
<h3>Steps for creating Nixpacks</h3>
<p>Now, we need to understand the process of the creation of nixpacks. There were several stages that you can see in the above video.</p>
<ol>
<li>Planning</li>
<li>Building</li>
</ol>
<h4>Planning Phase</h4>
<p>The first step was quite important, to understand the runtime environment. This is critical because it can decide how our application installs, runs, or even performs in the actual container environment. This is just the phase where we define the sets of commands or nix packages to be used for the building of the application.</p>
<p>You would get a clear look in the nixpack cli, it gives a list of the planned instructions before the build process starts in the build command. You can get the plan before building the application by using the <code>plan</code> command in the CLI.</p>
<pre><code>nixpacks plan .
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657960771/blog-media/nixpacks-plan-command.gif" alt="Nixpacks Plan Command"></p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657960816/blog-media/nixpacks-plan-cmd.png" alt="Nixpacks Plan Command Output"></p>
<p>So, this gives a <code>nix</code> output of all the configuration needed to build and install the application and it's dependencies. When you would have built the application i.e. the base build command, it also shows in the CLI, some useful information before moving to the build step.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657960880/blog-media/nixpacks-build-plan.png" alt="Nixpacks Plan Steps in Build Command"></p>
<p>So the list of the following procedure is listed in the planning phase of building the application.</p>
<ul>
<li>Packages/base runtime environment</li>
<li>Installation commands</li>
<li>Build commands</li>
<li>Commands to Start Application</li>
</ul>
<p>In the case of django, the following correspond to the planning attributes.</p>
<ul>
<li>Python version as the application runtime.</li>
<li>Installing dependencies via pip/pipenv/poetry</li>
<li>Building the Django app (collecting static files, setting database, etc)</li>
<li>Running the Django app (gunicorn/Nginx web server to run the app)</li>
</ul>
<p>So, hopefully, this gives you a better understanding of what is going on in the nixpack CLI. We are trying to automate the process of deployment and building the application for repeated deployments i.e. in a continuous integration/delivery system.</p>
<h4>Building Phase</h4>
<p>This is the phase, where actual installation, setup, and configuration takes place at the application level. In the build phase, we have several layers where things happen like installing the core packages for the application, installing dependencies, setting/generating up the necessary steps for the proper running of the application, and finally steps to run the application. This process creates an OSI(a standard used in containers/virtualization technology) image of the application. We can run the image and thus create a container of the application. You will need docker or any other containerization tool for the building of images and running containers.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657961801/blog-media/nixpacks-build-command.gif" alt="Django Build Command"></p>
<p>For the Django application, we have several steps of the application to be followed. We can use python as the base runtime environment, install the python packages, perform other commands like setting up admin accounts, static/media files, database setup, and finally run the steps to run the application at the container level.</p>
<ul>
<li>Installing python as a <a href="https://search.nixos.org/packages?channel=22.05&amp;show=python38&amp;from=0&amp;size=50&amp;sort=relevance&amp;type=packages&amp;query=python">Nix Package</a></li>
<li>Installing all python packages provided in the <code>requirements.txt</code>/<code>Pipfile</code>/<code>pyproject.toml</code> files.</li>
<li>Call commands from the environment to set up the applications like <code>createsuperuser</code>, <code>collectstatic</code>, <code>makemigrations</code>, pull data, management commands, etc. There are a lot of things that can be done here depending on the application.</li>
<li>Finally, the step to run the Django app, usually gunicorn/Nginx server is used for running the django application on a port with the host.</li>
</ul>
<p>So, this is what the build phase does, this is the whole and soul of the nixpacks. We literally do every installation setup and configuration of the application. Though the planning phase is equally important, a single missing data can ruin the build phase.</p>
<h3>Understanding the build phase for Django</h3>
<p>We ran the build command for creating the nixpack build, this started with planning the application configuration and then building up the application. This build phase was also divided into further processes like installing, running commands, copying actual source code to an image, and all the docker-related stuff that is required to create an image for a django application.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657961691/blog-media/nixpacks-build-process.png" alt="Django Application build nixpacks"></p>
<p>So, we have used the Debian nixpack, which sets as the base runtime for the application. Railway provides a <a href="https://github.com/railwayapp/nixpacks/pkgs/container/nixpacks">package</a> of the Debian image as the base runtime for our application. This is where we will run all the build processes on. This Debian image will be used for installing all types of dependencies and layers of language-specific runtime installation in the form of <a href="https://search.nixos.org/packages">nix packages</a>.</p>
<p>Now, we have the base image of debian, the nixpack build command actually starts executing the <code>Dockerfile</code>, this is an auto-generated step after the planning phase. So, with the instructions in <code>Dockerfile</code>, steps are executed one after other just as a normal Docker image build. This will generate the image for the application and after a while, because this process takes a while on the first iteration locally, after the build process has been completed, it will give a container name for you to run. This can be used to test the application locally, in production, the container is instantly created after the image has been generated.</p>
<pre><code>RUN:

docker run -it &lt;container_id_or_name&gt;

</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657966274/blog-media/nixpacks-build-run.png" alt="Nixpacks build command run container"></p>
<p>This is the command for running your application, this marks the end of the build process and also the build command provided with the CLI.</p>
<h3>Build Command Parameters</h3>
<p>The build command in the nixpacks CLI provides a few parameters or arguments to customize how to output the result and build the application, you can definitely provide the configuration in the application source code itself, but it is nicer to have it locally before deploying the application.</p>
<h4>Give a name to the Nixpack Image/Container</h4>
<p>The first parameter which might be helpful is to provide a name to the application at the build time. This becomes useful for running the container, this helps in avoiding long container names and giving a context of the nixpack.</p>
<pre><code>nixpacks build . --name &lt;project_name&gt;

OR 

nixpacks build . -n &lt;project_name&gt;

</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657962253/blog-media/nixpacks-build-name.png" alt="Nixpacks Build Command Name Image"></p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657962328/blog-media/nixpacks-build-name-image-run.png" alt="Nixpacks Build Command name run"></p>
<p>This gives a name to the image which has been built. Thereby providing a better context for the user to run the image and create a container out of it.</p>
<h4>Output the Built Image to a folder</h4>
<p>This is the command that can output the built application into a  provided folder. This parameter will not run the docker step thereafter; hence, no image is created if you provide an output folder. Though the folder will contain the <code>Dockerfile</code> and <code>environment.nix</code> files for creating the image and running the container. <strong>Make sure the output folder is NOT in the application folder itself, it will result in errors.</strong> The output command will not create an image but the process will be definitely executed in order to generate the <code>Dockerfile</code> and <code>environment.nix</code> files.</p>
<pre><code>nixpacks build . --out ../blog_image

OR

nixpacks build . -o ../blog_image
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657962407/blog-media/nixpacks-build-output.png" alt="Nixpacks Build Command Output folder"></p>
<p><strong>Dockerfile</strong></p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657962479/blog-media/nixpacks-build-output-folder.png" alt="Nixpacks Build command ouptut"></p>
<p><strong>environment.nix File</strong></p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657969127/blog-media/nixpacks-build-environment-nix-file.png" alt="Nixpacks environment.nix file"></p>
<p>So, this will output the built application into the provided path. The output folder should necessarily be out of the application folder as it makes no sense to output in the same folder as the application since the nixpacks CLI will consider the folder as the application folder.</p>
<h3>Provide a Install/Build/Start Command</h3>
<p>We can provide the commands at the install phase/build/start phase of the application to the build command in order to build the app with non-default or custom commands. This will add up to the docker steps that will involve making the build for the application.</p>
<pre><code>nixpacks build . --build-cmd 'python manage.py collectstatic'

OR

nixpacks build . -b 'python manage.py collectstatic'
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657962514/blog-media/nixpacks-build-build-command.png" alt="Nixpacks Build Command Providing install/build/start commands"></p>
<p>These kinds of parameters can be passed similarly for <code>install-cmd</code> and <code>start-cmd</code> as <code>-i</code> and <code>-s</code> respectively. We can further chain up the commands and customize the build process as per the application's requirements.</p>
<h3>Providing environment variables to image</h3>
<p>The environment variable can be passed to the build command for parsing to the application. This can be used for parsing additional or optional environment variables to the application image.</p>
<pre><code>nixpacks build . --env 'NAME=VALUE'

nixpack build . --env 'DATABASE_URL=postgres://postgres:postgres@localhost:5432/techstructive_blog'
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657963255/blog-media/nixpacks-build-env-variable-db-url.png" alt="Nixpacks Build Comand parsing environment variables"></p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657963302/blog-media/nixpacks-build-db-migrate-env.png" alt="Nixpacks Build Command env variable migrate"></p>
<p>Here, we provide the environment variable <code>DATABASE_URL</code> to the build command and this is parsed to the application image. Thereby when the image is run as a container, it is parsed as a normal environment variable and thereby is available for utilization from the application.</p>
<p>For further references on the build command arguments, you can follow the <a href="https://nixpacks.com/docs/cli">documentation of nixpack</a> by railway app.</p>
<h3>Creating a Procfile</h3>
<p>This is important for telling any buildpack in this case nixpacks to understand the process to start for this web application. For django it is simply to add the web process like to mention the <code>wsgi</code> app with the project name. We can use the gunicorn as the web server in production.</p>
<pre><code>
# Procfile

web: gunicorn &lt;django_project_name&gt;.wsgi

</code></pre>
<p>This is the Procfile, this is a file type without the extension. So, this is a typical Django application Procfile, you can also use the other variants of Procfile for applying migration for every web process start-up.</p>
<pre><code>
# Procfile

web: python manage.py migrate &amp;&amp; gunicorn &lt;django_project_name&gt;.wsgi

</code></pre>
<p>As we saw in the base build command, the local server was not able to listen to the gunicorn server in the container, so we need to bind the gunicorn server to the local port.</p>
<pre><code>
# Procfile

web: python manage.py migrate &amp;&amp; gunicorn &lt;django_project_name&gt;.wsgi -b :8000

</code></pre>
<p>So, we use the <code>-b</code> option in the gunicorn command to bind the port in the container to the port in the local machine. Now, if we build the application and forward the port to the <code>8000</code> port in the local machine, we will see our application running</p>
<!-- raw HTML omitted -->
<pre><code>docker run -p 8000:8000 -it &lt;container_id&gt; 

</code></pre>
<h3>Specifying the Python version</h3>
<p>This is created for specifying the python version for building the Django application or any other python app.</p>
<pre><code># .python-version

3.10
</code></pre>
<p>Save the <code>.python-version</code> file with just the python version like <code>3.6</code>, <code>3.10</code>, etc. and this will be picked by the build command while creating the build image.</p>
<h2>Deploying the Django Application</h2>
<p>After we looked at the nix picks specifications, we can now deploy our Django application with nixpacks on Railway. So, you can follow up with the <a href="">Railway Deployment</a> Article for setting up your Django app for deployment at the railway. This usually involves a few steps like creating Procfile(not necessary but recommended), requirements.txt(necessary to pull dependencies), and the python version which is chosen as <code>3.8</code> as default. The further steps are to create a GitHub repository to link with the Railway app and create a PostgreSQL database service on the railway platform.</p>
<h3>Create configuration files</h3>
<p>As we have seen we will require a <code>requirements.txt</code> file, <code>Pipfile</code> or a <code>pyproject.toml</code> file for listing out and installing dependencies for our django application. This can be done with various commands like:</p>
<pre><code># For requirements.txt and virtualenv
pip freeze &gt; requirements.txt

# Autogenerated Pipfile for pipenv
# Autogenerated pyproject.toml for poetry

</code></pre>
<p>So, this file should be present on the base directory of the django application in order for the nixpack to pick up and install the python packages. Also, for customization of the start command in the build process, you can create a <code>Procfile</code> as discussed earlier in order to run commands to start the Django web server.</p>
<p>The python version can be specified with the <code>.python-version</code> file with just the version name as <code>3.9</code>, <code>3.10</code>, etc. OR we can add an environment variable <code>NIXPACKS_PYTHON_VERSION</code> to the python version we want.</p>
<h3>Create and Linkup a GitHub repository for existing Django projects</h3>
<p>We can create a GitHub repository and link up the project to the Railway platform and thereby creating an automated build for every push.</p>
<p>The below video will explain how to set up the GitHub repository for the Railway app.</p>
<!-- raw HTML omitted -->
<!-- raw HTML omitted -->
<h3>Setup environment variables</h3>
<p>We can use <code>python-environ</code> to set up environment variables in a Django application, we will require these environment variables for attributes like <code>SECRET_KEY</code>, <code>DATABASE_URL</code>, email-stuff, etc. These are quite handy to avoid leaking sensitive information to the open source project on GitHub.</p>
<p>You can install the <code>python-environ</code> package with pip or any other package management tool as follows:</p>
<pre><code>pip install python-environ
</code></pre>
<p>After installing the package, we can set up the environment variable in the settings file.</p>
<pre><code class="language-python"># &lt;project_name&gt;/settings.py

import os
from dotenv import load_dotenv

BASE_DIR = Path(__file__).resolve().parent.parent

load_dotenv(os.path.join(BASE_DIR, &quot;.env&quot;))
</code></pre>
<p>After loading the environment variables, we can access them with <code>os.env(&quot;ENV_NAME&quot;, default=&quot;&quot;)</code>, this will load the environment variable with the name or we can provide a default value.</p>
<h3>Attach a PostgreSQL database service</h3>
<p>You can add a PostgreSQL database service to your Django Railway app by attaching a service. This will add a new service along with the django application, so these two act as different entities within a railway project.</p>
<!-- raw HTML omitted -->
<p>You can then access the <code>DATABASE_URL</code> from the connect settings and copy the database URL and set it as an environment variable in the django railway project service. This will link up the Django app and the PostgreSQL database. While setting it up locally, you can use the <code>.env</code> file and add the environment variable there.</p>
<pre><code># environment variable
DATABASE_URL=postgres://username:password@hostname:port/db_name

# local database postgres
DATABASE_URL=postgres://postgres:postgres@localhost:5432/db_name
</code></pre>
<!-- raw HTML omitted -->
<p>This will set up the database environment variable and you can access these from the settings.py file with the <code>env.db</code> function as follows:</p>
<pre><code class="language-python">env.db(&quot;DATABASE_URL&quot;, default=&lt;local_database_url&gt;)
</code></pre>
<p>So, we can finally use the database from the Railway app in our Django application once the environment variable is correctly used.</p>
<h3>Choose the Buildpack</h3>
<p>We can choose a buildpack for our Django application in the Railway platform, we have options like</p>
<ol>
<li>Heroku Buildpack</li>
<li>Railway Nixpacks</li>
<li>Paketo Buildpack</li>
</ol>
<p>As of the writing of the article, on 16th July 21, the Railway has made <code>Nixpacks</code> the default buildpack for an application :) It was the <code>Heroku</code> Buildpack as a default earlier. So, that is a cool thing, you can toggle these settings for choosing the buildpacks from the project settings.</p>
<p>Railway Dashboard Choose BuildPack</p>
<!-- raw HTML omitted -->
<h3>Deploy to Railway with Nixpacks</h3>
<p>Now, we have seen how to set up the nixpacks, we had the Postgres database setup, and we can finally deploy our application to the railway platform with nixpacks.</p>
<!-- raw HTML omitted -->
<p>So, we simply can configure the source code or include the <code>environment.nix</code> file into the source code for desired behavior. The nixpack selection can be done based on the source code or the presence of <code>environment.nix</code> and that's why we can rely on expected behaviors from the deployment builds.</p>
<h2>Summary</h2>
<p>So, nixpacks is a great way to deploy an application, for me it's an automated docker deployment, it basically creates docker images of the application and runs it with the appropriate environment. There is a lot of language support on nixpacks currently on Railway, you can check them out on the official website. Every programming language has specific requirements for managing or installing dependencies and packages, the nixpacks manage them automatically for us.</p>
<h2>Conclusion</h2>
<p>So, from this post of the <a href="https://www.meetgor.com/series/django-deployment/">Django Deployment Series</a>, we were able to understand how to deploy a Django application on the Railway app with Nixpacks which are a very intuitive way to deploy apps. We covered what are nixpacks, the process of building an application with nixpacks, and deploying a existing, new Django project on the railway with nixpacks. We also explored the various commands provided in the nixpacks CLI to build. plan a Django application.</p>
<p>Hopefully, you were able to understand the concept of nixpacks and how they can automate the process of containerization and deployment. Thank you for reading, if you have any queries or feedback, you can leave them down in the comments or freely drop them on social media. Happy Coding :)</p>
