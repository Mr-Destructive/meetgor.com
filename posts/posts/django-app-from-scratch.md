{
  "title": "Django App from Scratch Using Docker with Debian Base Image",
  "status": "published",
  "slug": "django-app-from-scratch",
  "date": "2025-04-05T12:33:41Z"
}

<h2>Pull a Fresh Debian Image</h2>
<p>Create a docker container from a Debian image, the following command can be used to pull a debain 11-slim image and create a container of it, also enter into the container in a interactive environment <code>-it</code> mode.</p>
<pre><code>docker run -v $(pwd):/var/www --rm -it -p 8001:80 debian:11-slim
</code></pre>
<h2>Create a Django App from Shell script</h2>
<p>Now, since we are inside a Debian container, we can enter a few commands, you can refer to the Mark Gibney's <a href="https://github.com/no-gravity/web_app_from_scratch">GitHub repository</a> for the script.</p>
<pre><code>apt update

apt install wget

wget https://raw.githubusercontent.com/no-gravity/web_app_from_scratch/main/django.sh
</code></pre>
<p>Also, if you want to do a few adjustments, you can install an editor, get used to vim or use nano.</p>
<pre><code>apt install vim 

OR

apt install nano
</code></pre>
<pre><code>chmod +x django.sh
bash django.sh
</code></pre>
<p>I also have a few adjustment of the original script, that accepts a project name and creates a django project based on the positional parameter given to it. You can get it from the <a href="https://github.com/Mr-Destructive/quick-setup-scripts/blob/main/django_docker.sh">quick-setup-script repository</a> or directly the <a href="https://raw.githubusercontent.com/Mr-Destructive/quick-setup-scripts/main/django_docker.sh">script</a>.</p>
<p>To use the above file, you need to execute the command as :</p>
<pre><code>chmod +x django_docker.sh
bash django_docker.sh &lt;project_name&gt;
</code></pre>
<p>This will generate the project in the <code>/var/www/</code> folder with the name of the project. The script will prompt you with a few things for setting up at some iterations like basic application setup , <code>static file</code> configuration, <code>basic tempalte</code> setup and the <code>user authentication</code> setup.</p>
<h2>Copy the contents from the docker container</h2>
<p>You can copy the contents of the folder into your local machine by entering the <a href="https://docs.docker.com/engine/reference/commandline/cp/">cp</a> command in docker.</p>
<pre><code>docker cp &lt;container_id&gt;:/var/www/&lt;project_name&gt; /path/in_local_machine/
</code></pre>
<p>This will copy the project in the docker container into the local machine for you to extend and tweak it as per your needs.</p>
<p>That's a basic Django Project Setup by using Docker with a Debian Image.</p>
