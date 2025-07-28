{
  "title": "Dockerize a Django project",
  "status": "published",
  "slug": "dockerize-django-prj",
  "date": "2025-04-05T12:33:47Z"
}

<h2>Dockerize a Django project</h2>
<p>We can run our Django projects in a Docker Container by creating a Docker image for our project. It is really easy and intuitive to create a Dockerfile for any given application as it really is a matter of writing commands in a file and basically running it is a isolated environment. To create a Docker image, we first need a Dockerfile. A Dockerfile is simply a Blueprint to create a image in Docker. In this file we specify the instructions/commands/environment variables to create a image for our app to run.</p>
<h3>Create a Docker File</h3>
<p>To create a dockerfile, simply create a file named as <code>Dockerfile</code> without any extension(not a .txt file). We need to then add the following code into the file.</p>
<pre><code class="language-dockerfile">FROM python:3.9-buster

ENV PYTHONUNBUFFERED=1

WORKDIR /code

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

EXPOSE 8000

CMD [&quot;python&quot;,&quot;manage.py&quot;,&quot;runserver&quot;,&quot;0.0.0.0:8000&quot;]
</code></pre>
<ul>
<li>
<p>Let's see what are we doing here, first we are pulling a base image of Python in this case it is <a href="https://github.com/docker-library/python/blob/a4b368154b7e3c33c76385f1be7a998fcf3123eb/3.9/buster/Dockerfile">3.9-buster</a>, but it can be any of the <a href="https://hub.docker.com/_/python">Official Python images</a> from <a href="https://hub.docker.com">dockerhub</a>. It is simply a environment for our Django project to run. You can even use linux environments like ubuntu, debian or alpine.</p>
</li>
<li>
<p>Next we add a <code>PYTHONUNBUFFERED</code> as a environment variable and initialize it to <code>1</code>, which basically allows to parse python output straight into the terminal without actually buffering first.</p>
</li>
<li>
<p>We set the <code>WORKDIR</code> as the <code>code</code> directory, this sets the provided directory as the base to run any command or instructions in the Dockerfile.</p>
</li>
<li>
<p>We then <code>COPY</code> the <code>requirements.txt</code> file in the current(<code>.</code>) directory which is <code>code</code> as mentioned earlier. We simply want the <code>requirements.txt</code> file in the current directory so we can proceed with the next step which is to install the dependencies for the Django project</p>
</li>
<li>
<p>After copying the <code>requirements.txt</code> file, we simply install the packages mentioned int the file with <code>pip</code> we use the <code>RUN</code> command to execute any shell commands in the environment.</p>
</li>
<li>
<p>Next step is to <code>COPY</code> all the files from the current folder(local machine) into the docker environment image. So, we have the source code files in the Container.</p>
</li>
<li>
<p>Expose the port <code>8000</code>(default) or any other port you would want to run Django in the image. We use the expose command in complement to the <code>-p</code> argument when we want to create the container. I have explained the <a href="https://mr-destructive.github.io/techstructive-blog/docker-port-forward">port-forwarding</a> concept in the previous TIL.</p>
</li>
<li>
<p>Finally we can run the server. We need to parse every command as a comma separated string in the list to the <code>CMD</code> which the container actually executes the command when the image is build. So, all the commands previously were to build an image for the django project, the <code>CMD</code> will actually run the server in a container after building the image of the django project.</p>
</li>
</ul>
<h3>Build the image</h3>
<p>Using the Dockerfile, we can create the image for the Django project by a simple command.</p>
<pre><code>docker build . -t django-app
</code></pre>
<p>Assuming the Dockerfile is in the same folder from which you are running this command.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1646230907/blog-media/jj04subyvkuvfb5obytu.png" alt="Build the django-app image"></p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1646230988/blog-media/ugjoakqgyhiwelqkyaat.png" alt="Build Complete Success"></p>
<h3>Run the Image and Create a Container</h3>
<pre><code>docker run -p 8000:8000 django-app
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1646231023/blog-media/yneuz46burorz4b5vzp4.png" alt="Create Container"></p>
<p>You can use any port like <code>3000</code> or <code>4000</code> on your local system by changing the first port number as <code>3000:8000</code> and so on. So, the first port number is the local system's port number and the port number followed by a <code>:</code> is the port number for the environment inside the container of the django project. The container port is the same which we have <code>exposed</code> while creating the image. You can now go to <a href="127.0.0.1:8000">127.0.0.1:8000</a> or your chosen port in your local environment and see the django app running.</p>
<p>If you are using a Docker as a Virtual Machine, you need to find the IP-address of that Virtual-Machine(<code>docker-machine ip</code>) and use that instead of <code>localhost</code> or <code>127.0.0.1</code>.</p>
<h3>Cleaning the Container for iterations</h3>
<p>We need to stop the container to run it again after making a few changes in the app if any. Simply pressing the <code>CTRL + D</code> option wont stop the container here. We need to stop the container by executing:</p>
<pre><code>docker ps

docker stop &lt;container-id&gt;
</code></pre>
<p>This will stop the container. Also if you want to remove the images you built, you can run the command:</p>
<pre><code>docker image ls

docker rmi -f &lt;image-id&gt;
</code></pre>
<p>So, we now have a image of our Django project, this image can be used by anyone inside a docker environment and thus creating much more easier to test/work on a given project irrespective of the system is being used.</p>
