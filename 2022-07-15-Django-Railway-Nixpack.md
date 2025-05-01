---
article_html: "<h2>Introduction</h2>\n<p>We have seen how to deploy a Django application
  on railway app in the <a href=\"https://www.meetgor.com/django-deploy-railway/\">previous
  article</a> of the <a href=\"https://www.meetgor.com/series/django-deployment/\">series</a>.
  We deployed the django project using the Heroku Buildpacks under the hood. The railway
  app provides a couple of options on how to build your web application. Currently,
  there are three options, one of which is the <a href=\"https://devcenter.heroku.com/articles/heroku-20-stack\">heroku
  buildpacks</a> which is the default one, second, we have the nixpacks which we will
  see today, and the third is the <a href=\"https://paketo.io/\">Paketo buildpack</a>.</p>\n<h2>What
  is a Buildpack?</h2>\n<p>A buildpack is a set of programs that turns your source
  code into a container image. So it is basically a tool for converting your application
  into a deployment-ready state with the help of containerization technology. Buildpacks
  allow us to extract away the steps for deploying an application. The layer of abstraction
  for converting a source code into a deployable code/container is played by the build
  packs.</p>\n<h3>Benefits of buildpacks</h3>\n<p>Buildpacks as we discussed act as
  a layer of abstraction from converting source code into the deployable containers,
  so it helps in avoiding manually writing dependencies and installing them. But on
  top of that, they can even detect low-level changes to the source code, i.e. if
  a dependency is changed, it would know which dependencies are altered and which
  components can be fetched from the cache. Caching is a great property of buildpacks
  which enhances the performance and deployment time. Since we can specify the commands
  for the build process, the buildpacks are customizable and hence provide a solid
  foundation for professional applications.</p>\n<p>For more references on Buildpacks,
  you can follow up with the great article on <a href=\"https://technology.doximity.com/articles/buildpacks-vs-dockerfiles\">dockerfile
  vs buildpacks</a>.</p>\n<h2>What are Nixpacks?</h2>\n<p><a href=\"https://nixpacks.com/docs\">Nixpacks</a>
  are quite similar to buildpacks, but they have their own set of technologies used
  in managing and installing builds for the application. Nixpacks as the name suggests
  uses <a href=\"https://search.nixos.org/packages\">Nix Packages</a> for creating
  and installing dependencies and <a href=\"https://www.docker.com/\">Docker</a> for
  building images and running containers.</p>\n<p>Nixpacks are quite cool as most
  of the applications require little or almost no configuration. You don't need a
  lot of knowledge of docker or nix technologies, everything is abstracted for you.
  Nixpacks uses nix packages for installing the runtime environment and the dependencies
  for the applications. It also caches support for detecting the core modules or packages
  in the runtime, so it gets a huge boost in deployment performance. Minimal deployment
  time, simultaneously makes it developer friendly and improves the quality of the
  application.</p>\n<ul>\n<li>Abstracted technologies like Nix and Docker.</li>\n<li>Caching
  of dependencies and steps to build.</li>\n<li>Customizable at most of the steps.</li>\n<li>Extensible
  and Developer friendly.</li>\n</ul>\n<h2>Installing Nixpacks</h2>\n<p>Nixpacks are
  the <a href=\"https://search.nixos.org/packages\">nix packges</a> which are used
  with the source code to create a buildpack of their own. The nix packages take in
  the source code of your application and convert it into a simple OCI image e.g.
  Docker image that can be run on various environments. It is similar to buildpacks
  but it's not the same instead it is better and uses a different ecosystem.</p>\n<p>Railway
  has a cool <a href=\"https://nixpacks.com/docs/cli\">CLI</a> tool for creating nixpacks
  on your local system. You can install the nixpack cli from the official documentation
  site. There are a couple of ways to install them on your system.\nYou can install
  the nixpack with the curl command in your terminal which is one of the simplest
  ways to install it.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>curl -fsSL https://raw.githubusercontent.com/railwayapp/nixpacks/master/install.sh
  | bash\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657725394/blog-media/railway-nixpacks-install.png\"
  alt=\"Railway Nixpacks Install\" /></p>\n<p>Once it has been installed we can check
  the status if nixpacks was successfully installed or not.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks --version\n\nOR
  \n\nnixpacks\n</pre></div>\n\n</pre>\n\n<p>So that's how we install nixpacks CLI
  into the system. Now, we can move into the configuration required for creating a
  nixpack from the Django application.</p>\n<h2>Creating a Django Project</h2>\n<p>So,
  I assume here, you have your django project ready. You can pick any django project
  and configure it as mentioned below.</p>\n<h3>Creating requirements file</h3>\n<p>It
  is a good practice to include the <code>requirements.txt</code> file in any python
  based applications. It becomes really easy to set up and give the project a spin.
  It's often the case, that we have to use a virtual environment for creating a pip
  file, otherwise, we might conflict and mix up the globally installed packages with
  project-specific dependencies.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pip freeze &gt; requirements.txt\n</pre></div>\n\n</pre>\n\n<p>This
  will create a <code>requirements.txt</code> file in the current folder, which will
  contain the list of all the dependencies with the version mentioned in. We also
  have other package management system like <a href=\"https://www.meetgor.com/pipenv-intro/\">pipenv</a>
  or <a href=\"https://python-poetry.org/\">poetry</a>.</p>\n<h2>Creating a Nixpack
  for the project</h2>\n<p>Finally, we can now start creating nixpacks from the source
  code of the django application. So, make sure the Django project is running on your
  local system. Set up a virtualenv and database as per your local environment. We
  will be using the build command for creating the nixpack of our django project.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build .\n</pre></div>\n\n</pre>\n\n<video
  width=\"800\" height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657881139/blog-media/nixpacks-demo-base.mp4\"
  type=\"video/mp4\">\n</video>\n<p>This is the simplest of commands you can run to
  create a nixpack. You need to be in the source file where all your files are located.
  In Django, we call it the <code>BASE_DIR</code>, the same folder where your <code>manage.py</code>
  resides. So, be on that path and run the above command. It would pick up a few things
  from the source code itself. There are also some default values picked up after
  analyzing the source code, for example in the case if it detects the django project,
  it attaches the <code>python manage.py migrate &amp;&amp; gunicorn &lt;project_name&gt;.wsgi</code>
  command as the default build command.</p>\n<p>Similarly, there are other default
  options like the version of the runtime, installation steps, etc. We can see that
  the port is not accessible from the docker container that we ran it is because of
  two reasons either we have not exposed the port or the port is not bound with gunicorn.
  We will see that configuration later, it's quite easy to bind the local port to
  the docker container.</p>\n<h3>Steps for creating Nixpacks</h3>\n<p>Now, we need
  to understand the process of the creation of nixpacks. There were several stages
  that you can see in the above video.</p>\n<ol>\n<li>Planning</li>\n<li>Building</li>\n</ol>\n<h4>Planning
  Phase</h4>\n<p>The first step was quite important, to understand the runtime environment.
  This is critical because it can decide how our application installs, runs, or even
  performs in the actual container environment. This is just the phase where we define
  the sets of commands or nix packages to be used for the building of the application.</p>\n<p>You
  would get a clear look in the nixpack cli, it gives a list of the planned instructions
  before the build process starts in the build command. You can get the plan before
  building the application by using the <code>plan</code> command in the CLI.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks plan .\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657960771/blog-media/nixpacks-plan-command.gif\"
  alt=\"Nixpacks Plan Command\" /></p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657960816/blog-media/nixpacks-plan-cmd.png\"
  alt=\"Nixpacks Plan Command Output\" /></p>\n<p>So, this gives a <code>nix</code>
  output of all the configuration needed to build and install the application and
  it's dependencies. When you would have built the application i.e. the base build
  command, it also shows in the CLI, some useful information before moving to the
  build step.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657960880/blog-media/nixpacks-build-plan.png\"
  alt=\"Nixpacks Plan Steps in Build Command\" /></p>\n<p>So the list of the following
  procedure is listed in the planning phase of building the application.</p>\n<ul>\n<li>Packages/base
  runtime environment</li>\n<li>Installation commands</li>\n<li>Build commands</li>\n<li>Commands
  to Start Application</li>\n</ul>\n<p>In the case of django, the following correspond
  to the planning attributes.</p>\n<ul>\n<li>Python version as the application runtime.</li>\n<li>Installing
  dependencies via pip/pipenv/poetry</li>\n<li>Building the Django app (collecting
  static files, setting database, etc)</li>\n<li>Running the Django app (gunicorn/Nginx
  web server to run the app)</li>\n</ul>\n<p>So, hopefully, this gives you a better
  understanding of what is going on in the nixpack CLI. We are trying to automate
  the process of deployment and building the application for repeated deployments
  i.e. in a continuous integration/delivery system.</p>\n<h4>Building Phase</h4>\n<p>This
  is the phase, where actual installation, setup, and configuration takes place at
  the application level. In the build phase, we have several layers where things happen
  like installing the core packages for the application, installing dependencies,
  setting/generating up the necessary steps for the proper running of the application,
  and finally steps to run the application. This process creates an OSI(a standard
  used in containers/virtualization technology) image of the application. We can run
  the image and thus create a container of the application. You will need docker or
  any other containerization tool for the building of images and running containers.</p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657961801/blog-media/nixpacks-build-command.gif\"
  alt=\"Django Build Command\" /></p>\n<p>For the Django application, we have several
  steps of the application to be followed. We can use python as the base runtime environment,
  install the python packages, perform other commands like setting up admin accounts,
  static/media files, database setup, and finally run the steps to run the application
  at the container level.</p>\n<ul>\n<li>Installing python as a <a href=\"https://search.nixos.org/packages?channel=22.05&amp;show=python38&amp;from=0&amp;size=50&amp;sort=relevance&amp;type=packages&amp;query=python\">Nix
  Package</a></li>\n<li>Installing all python packages provided in the <code>requirements.txt</code>/<code>Pipfile</code>/<code>pyproject.toml</code>
  files.</li>\n<li>Call commands from the environment to set up the applications like
  <code>createsuperuser</code>, <code>collectstatic</code>, <code>makemigrations</code>,
  pull data, management commands, etc. There are a lot of things that can be done
  here depending on the application.</li>\n<li>Finally, the step to run the Django
  app, usually gunicorn/Nginx server is used for running the django application on
  a port with the host.</li>\n</ul>\n<p>So, this is what the build phase does, this
  is the whole and soul of the nixpacks. We literally do every installation setup
  and configuration of the application. Though the planning phase is equally important,
  a single missing data can ruin the build phase.</p>\n<h3>Understanding the build
  phase for Django</h3>\n<p>We ran the build command for creating the nixpack build,
  this started with planning the application configuration and then building up the
  application. This build phase was also divided into further processes like installing,
  running commands, copying actual source code to an image, and all the docker-related
  stuff that is required to create an image for a django application.</p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657961691/blog-media/nixpacks-build-process.png\"
  alt=\"Django Application build nixpacks\" /></p>\n<p>So, we have used the Debian
  nixpack, which sets as the base runtime for the application. Railway provides a
  <a href=\"https://github.com/railwayapp/nixpacks/pkgs/container/nixpacks\">package</a>
  of the Debian image as the base runtime for our application. This is where we will
  run all the build processes on. This Debian image will be used for installing all
  types of dependencies and layers of language-specific runtime installation in the
  form of <a href=\"https://search.nixos.org/packages\">nix packages</a>.</p>\n<p>Now,
  we have the base image of debian, the nixpack build command actually starts executing
  the <code>Dockerfile</code>, this is an auto-generated step after the planning phase.
  So, with the instructions in <code>Dockerfile</code>, steps are executed one after
  other just as a normal Docker image build. This will generate the image for the
  application and after a while, because this process takes a while on the first iteration
  locally, after the build process has been completed, it will give a container name
  for you to run. This can be used to test the application locally, in production,
  the container is instantly created after the image has been generated.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>RUN:\n\ndocker run -it &lt;container_id_or_name&gt;\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657966274/blog-media/nixpacks-build-run.png\"
  alt=\"Nixpacks build command run container\" /></p>\n<p>This is the command for
  running your application, this marks the end of the build process and also the build
  command provided with the CLI.</p>\n<h3>Build Command Parameters</h3>\n<p>The build
  command in the nixpacks CLI provides a few parameters or arguments to customize
  how to output the result and build the application, you can definitely provide the
  configuration in the application source code itself, but it is nicer to have it
  locally before deploying the application.</p>\n<h4>Give a name to the Nixpack Image/Container</h4>\n<p>The
  first parameter which might be helpful is to provide a name to the application at
  the build time. This becomes useful for running the container, this helps in avoiding
  long container names and giving a context of the nixpack.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build . --name &lt;project_name&gt;\n\nOR
  \n\nnixpacks build . -n &lt;project_name&gt;\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962253/blog-media/nixpacks-build-name.png\"
  alt=\"Nixpacks Build Command Name Image\" /></p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962328/blog-media/nixpacks-build-name-image-run.png\"
  alt=\"Nixpacks Build Command name run\" /></p>\n<p>This gives a name to the image
  which has been built. Thereby providing a better context for the user to run the
  image and create a container out of it.</p>\n<h4>Output the Built Image to a folder</h4>\n<p>This
  is the command that can output the built application into a  provided folder. This
  parameter will not run the docker step thereafter; hence, no image is created if
  you provide an output folder. Though the folder will contain the <code>Dockerfile</code>
  and <code>environment.nix</code> files for creating the image and running the container.
  <strong>Make sure the output folder is NOT in the application folder itself, it
  will result in errors.</strong> The output command will not create an image but
  the process will be definitely executed in order to generate the <code>Dockerfile</code>
  and <code>environment.nix</code> files.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build . --out ../blog_image\n\nOR\n\nnixpacks
  build . -o ../blog_image\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962407/blog-media/nixpacks-build-output.png\"
  alt=\"Nixpacks Build Command Output folder\" /></p>\n<p><strong>Dockerfile</strong></p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962479/blog-media/nixpacks-build-output-folder.png\"
  alt=\"Nixpacks Build command ouptut\" /></p>\n<p><strong>environment.nix File</strong></p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657969127/blog-media/nixpacks-build-environment-nix-file.png\"
  alt=\"Nixpacks environment.nix file\" /></p>\n<p>So, this will output the built
  application into the provided path. The output folder should necessarily be out
  of the application folder as it makes no sense to output in the same folder as the
  application since the nixpacks CLI will consider the folder as the application folder.</p>\n<h3>Provide
  a Install/Build/Start Command</h3>\n<p>We can provide the commands at the install
  phase/build/start phase of the application to the build command in order to build
  the app with non-default or custom commands. This will add up to the docker steps
  that will involve making the build for the application.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build . --build-cmd
  &#39;python manage.py collectstatic&#39;\n\nOR\n\nnixpacks build . -b &#39;python
  manage.py collectstatic&#39;\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962514/blog-media/nixpacks-build-build-command.png\"
  alt=\"Nixpacks Build Command Providing install/build/start commands\" /></p>\n<p>These
  kinds of parameters can be passed similarly for <code>install-cmd</code> and <code>start-cmd</code>
  as <code>-i</code> and <code>-s</code> respectively. We can further chain up the
  commands and customize the build process as per the application's requirements.</p>\n<h3>Providing
  environment variables to image</h3>\n<p>The environment variable can be passed to
  the build command for parsing to the application. This can be used for parsing additional
  or optional environment variables to the application image.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build . --env &#39;NAME=VALUE&#39;\n\nnixpack
  build . --env &#39;DATABASE_URL=postgres://postgres:postgres@localhost:5432/techstructive_blog&#39;\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657963255/blog-media/nixpacks-build-env-variable-db-url.png\"
  alt=\"Nixpacks Build Comand parsing environment variables\" /></p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657963302/blog-media/nixpacks-build-db-migrate-env.png\"
  alt=\"Nixpacks Build Command env variable migrate\" /></p>\n<p>Here, we provide
  the environment variable <code>DATABASE_URL</code> to the build command and this
  is parsed to the application image. Thereby when the image is run as a container,
  it is parsed as a normal environment variable and thereby is available for utilization
  from the application.</p>\n<p>For further references on the build command arguments,
  you can follow the <a href=\"https://nixpacks.com/docs/cli\">documentation of nixpack</a>
  by railway app.</p>\n<h3>Creating a Procfile</h3>\n<p>This is important for telling
  any buildpack in this case nixpacks to understand the process to start for this
  web application. For django it is simply to add the web process like to mention
  the <code>wsgi</code> app with the project name. We can use the gunicorn as the
  web server in production.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Procfile\n\nweb: gunicorn
  &lt;django_project_name&gt;.wsgi\n</pre></div>\n\n</pre>\n\n<p>This is the Procfile,
  this is a file type without the extension. So, this is a typical Django application
  Procfile, you can also use the other variants of Procfile for applying migration
  for every web process start-up.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Procfile\n\nweb: python
  manage.py migrate &amp;&amp; gunicorn &lt;django_project_name&gt;.wsgi\n</pre></div>\n\n</pre>\n\n<p>As
  we saw in the base build command, the local server was not able to listen to the
  gunicorn server in the container, so we need to bind the gunicorn server to the
  local port.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Procfile\n\nweb: python
  manage.py migrate &amp;&amp; gunicorn &lt;django_project_name&gt;.wsgi -b :8000\n</pre></div>\n\n</pre>\n\n<p>So,
  we use the <code>-b</code> option in the gunicorn command to bind the port in the
  container to the port in the local machine. Now, if we build the application and
  forward the port to the <code>8000</code> port in the local machine, we will see
  our application running</p>\n<video width=\"800\" height=\"450\" controls>\n  <source
  src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657964597/blog-media/nixpacks-local-bind.mp4\"
  type=\"video/mp4\">\n</video>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>docker run -p 8000:8000 -it
  &lt;container_id&gt; \n</pre></div>\n\n</pre>\n\n<h3>Specifying the Python version</h3>\n<p>This
  is created for specifying the python version for building the Django application
  or any other python app.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># .python-version\n\n3.10\n</pre></div>\n\n</pre>\n\n<p>Save
  the <code>.python-version</code> file with just the python version like <code>3.6</code>,
  <code>3.10</code>, etc. and this will be picked by the build command while creating
  the build image.</p>\n<h2>Deploying the Django Application</h2>\n<p>After we looked
  at the nix picks specifications, we can now deploy our Django application with nixpacks
  on Railway. So, you can follow up with the <a href=\"\">Railway Deployment</a> Article
  for setting up your Django app for deployment at the railway. This usually involves
  a few steps like creating Procfile(not necessary but recommended), requirements.txt(necessary
  to pull dependencies), and the python version which is chosen as <code>3.8</code>
  as default. The further steps are to create a GitHub repository to link with the
  Railway app and create a PostgreSQL database service on the railway platform.</p>\n<h3>Create
  configuration files</h3>\n<p>As we have seen we will require a <code>requirements.txt</code>
  file, <code>Pipfile</code> or a <code>pyproject.toml</code> file for listing out
  and installing dependencies for our django application. This can be done with various
  commands like:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># For requirements.txt and
  virtualenv\npip freeze &gt; requirements.txt\n\n# Autogenerated Pipfile for pipenv\n#
  Autogenerated pyproject.toml for poetry\n</pre></div>\n\n</pre>\n\n<p>So, this file
  should be present on the base directory of the django application in order for the
  nixpack to pick up and install the python packages. Also, for customization of the
  start command in the build process, you can create a <code>Procfile</code> as discussed
  earlier in order to run commands to start the Django web server.</p>\n<p>The python
  version can be specified with the <code>.python-version</code> file with just the
  version name as <code>3.9</code>, <code>3.10</code>, etc. OR we can add an environment
  variable <code>NIXPACKS_PYTHON_VERSION</code> to the python version we want.</p>\n<h3>Create
  and Linkup a GitHub repository for existing Django projects</h3>\n<p>We can create
  a GitHub repository and link up the project to the Railway platform and thereby
  creating an automated build for every push.</p>\n<p>The below video will explain
  how to set up the GitHub repository for the Railway app.</p>\n<video width=\"800\"
  height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1652970864/blog-media/django-deployment/railway_project_init.webm\"
  type=\"video/mp4\">\n</video>\n<video width=\"800\" height=\"450\" controls>\n  <source
  src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657965284/blog-media/railway-platform-github.mp4\"
  type=\"video/mp4\">\n</video>\n<h3>Setup environment variables</h3>\n<p>We can use
  <code>python-environ</code> to set up environment variables in a Django application,
  we will require these environment variables for attributes like <code>SECRET_KEY</code>,
  <code>DATABASE_URL</code>, email-stuff, etc. These are quite handy to avoid leaking
  sensitive information to the open source project on GitHub.</p>\n<p>You can install
  the <code>python-environ</code> package with pip or any other package management
  tool as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install python-environ\n</pre></div>\n\n</pre>\n\n<p>After
  installing the package, we can set up the environment variable in the settings file.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># &lt;project_name&gt;/settings.py</span>\n\n<span class=\"kn\">import</span>
  <span class=\"nn\">os</span>\n<span class=\"kn\">from</span> <span class=\"nn\">dotenv</span>
  <span class=\"kn\">import</span> <span class=\"n\">load_dotenv</span>\n\n<span class=\"n\">BASE_DIR</span>
  <span class=\"o\">=</span> <span class=\"n\">Path</span><span class=\"p\">(</span><span
  class=\"vm\">__file__</span><span class=\"p\">)</span><span class=\"o\">.</span><span
  class=\"n\">resolve</span><span class=\"p\">()</span><span class=\"o\">.</span><span
  class=\"n\">parent</span><span class=\"o\">.</span><span class=\"n\">parent</span>\n\n<span
  class=\"n\">load_dotenv</span><span class=\"p\">(</span><span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">path</span><span class=\"o\">.</span><span
  class=\"n\">join</span><span class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;.env&quot;</span><span class=\"p\">))</span>\n</pre></div>\n\n</pre>\n\n<p>After
  loading the environment variables, we can access them with <code>os.env(&quot;ENV_NAME&quot;,
  default=&quot;&quot;)</code>, this will load the environment variable with the name
  or we can provide a default value.</p>\n<h3>Attach a PostgreSQL database service</h3>\n<p>You
  can add a PostgreSQL database service to your Django Railway app by attaching a
  service. This will add a new service along with the django application, so these
  two act as different entities within a railway project.</p>\n<video width=\"800\"
  height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1652963718/blog-media/django-deployment/postgres_spinup_railway_d2xkpt.mp4\"
  type=\"video/mp4\">\n</video>\n<p>You can then access the <code>DATABASE_URL</code>
  from the connect settings and copy the database URL and set it as an environment
  variable in the django railway project service. This will link up the Django app
  and the PostgreSQL database. While setting it up locally, you can use the <code>.env</code>
  file and add the environment variable there.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># environment variable\nDATABASE_URL=postgres://username:password@hostname:port/db_name\n\n#
  local database postgres\nDATABASE_URL=postgres://postgres:postgres@localhost:5432/db_name\n</pre></div>\n\n</pre>\n\n<video
  width=\"800\" height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657964943/blog-media/railway-postgres-spinup.mp4\"
  type=\"video/mp4\">\n</video>\n<p>This will set up the database environment variable
  and you can access these from the <a href=\"http://settings.py\">settings.py</a>
  file with the <code>env.db</code> function as follows:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">env</span><span class=\"o\">.</span><span class=\"n\">db</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;DATABASE_URL&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">default</span><span class=\"o\">=&lt;</span><span class=\"n\">local_database_url</span><span
  class=\"o\">&gt;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  we can finally use the database from the Railway app in our Django application once
  the environment variable is correctly used.</p>\n<h3>Choose the Buildpack</h3>\n<p>We
  can choose a buildpack for our Django application in the Railway platform, we have
  options like</p>\n<ol>\n<li>Heroku Buildpack</li>\n<li>Railway Nixpacks</li>\n<li>Paketo
  Buildpack</li>\n</ol>\n<p>As of the writing of the article, on 16th July 21, the
  Railway has made <code>Nixpacks</code> the default buildpack for an application
  :) It was the <code>Heroku</code> Buildpack as a default earlier. So, that is a
  cool thing, you can toggle these settings for choosing the buildpacks from the project
  settings.</p>\n<p>Railway Dashboard Choose BuildPack</p>\n<video width=\"800\" height=\"450\"
  controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657964699/blog-media/nixpacks-railway-dashboard.mp4\"
  type=\"video/mp4\">\n</video>\n<h3>Deploy to Railway with Nixpacks</h3>\n<p>Now,
  we have seen how to set up the nixpacks, we had the Postgres database setup, and
  we can finally deploy our application to the railway platform with nixpacks.</p>\n<video
  width=\"800\" height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657965560/blog-media/railway-nixpacks-deploy.mp4\"
  type=\"video/mp4\">\n</video>\n<p>So, we simply can configure the source code or
  include the <code>environment.nix</code> file into the source code for desired behavior.
  The nixpack selection can be done based on the source code or the presence of <code>environment.nix</code>
  and that's why we can rely on expected behaviors from the deployment builds.</p>\n<h2>Summary</h2>\n<p>So,
  nixpacks is a great way to deploy an application, for me it's an automated docker
  deployment, it basically creates docker images of the application and runs it with
  the appropriate environment. There is a lot of language support on nixpacks currently
  on Railway, you can check them out on the official website. Every programming language
  has specific requirements for managing or installing dependencies and packages,
  the nixpacks manage them automatically for us.</p>\n<h2>Conclusion</h2>\n<p>So,
  from this post of the <a href=\"https://www.meetgor.com/series/django-deployment/\">Django
  Deployment Series</a>, we were able to understand how to deploy a Django application
  on the Railway app with Nixpacks which are a very intuitive way to deploy apps.
  We covered what are nixpacks, the process of building an application with nixpacks,
  and deploying a existing, new Django project on the railway with nixpacks. We also
  explored the various commands provided in the nixpacks CLI to build. plan a Django
  application.</p>\n<p>Hopefully, you were able to understand the concept of nixpacks
  and how they can automate the process of containerization and deployment. Thank
  you for reading, if you have any queries or feedback, you can leave them down in
  the comments or freely drop them on social media. Happy Coding :)</p>\n"
cover: ''
date: 2022-07-16
datetime: 2022-07-16 00:00:00+00:00
description: Configuring and Seting up nixpacks for deploying django project on Railway
  app. Interacting with the nixpacks CLI for building and creating deployable django
  applications.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-07-15-Django-Railway-Nixpack.md
html: "<h2>Introduction</h2>\n<p>We have seen how to deploy a Django application on
  railway app in the <a href=\"https://www.meetgor.com/django-deploy-railway/\">previous
  article</a> of the <a href=\"https://www.meetgor.com/series/django-deployment/\">series</a>.
  We deployed the django project using the Heroku Buildpacks under the hood. The railway
  app provides a couple of options on how to build your web application. Currently,
  there are three options, one of which is the <a href=\"https://devcenter.heroku.com/articles/heroku-20-stack\">heroku
  buildpacks</a> which is the default one, second, we have the nixpacks which we will
  see today, and the third is the <a href=\"https://paketo.io/\">Paketo buildpack</a>.</p>\n<h2>What
  is a Buildpack?</h2>\n<p>A buildpack is a set of programs that turns your source
  code into a container image. So it is basically a tool for converting your application
  into a deployment-ready state with the help of containerization technology. Buildpacks
  allow us to extract away the steps for deploying an application. The layer of abstraction
  for converting a source code into a deployable code/container is played by the build
  packs.</p>\n<h3>Benefits of buildpacks</h3>\n<p>Buildpacks as we discussed act as
  a layer of abstraction from converting source code into the deployable containers,
  so it helps in avoiding manually writing dependencies and installing them. But on
  top of that, they can even detect low-level changes to the source code, i.e. if
  a dependency is changed, it would know which dependencies are altered and which
  components can be fetched from the cache. Caching is a great property of buildpacks
  which enhances the performance and deployment time. Since we can specify the commands
  for the build process, the buildpacks are customizable and hence provide a solid
  foundation for professional applications.</p>\n<p>For more references on Buildpacks,
  you can follow up with the great article on <a href=\"https://technology.doximity.com/articles/buildpacks-vs-dockerfiles\">dockerfile
  vs buildpacks</a>.</p>\n<h2>What are Nixpacks?</h2>\n<p><a href=\"https://nixpacks.com/docs\">Nixpacks</a>
  are quite similar to buildpacks, but they have their own set of technologies used
  in managing and installing builds for the application. Nixpacks as the name suggests
  uses <a href=\"https://search.nixos.org/packages\">Nix Packages</a> for creating
  and installing dependencies and <a href=\"https://www.docker.com/\">Docker</a> for
  building images and running containers.</p>\n<p>Nixpacks are quite cool as most
  of the applications require little or almost no configuration. You don't need a
  lot of knowledge of docker or nix technologies, everything is abstracted for you.
  Nixpacks uses nix packages for installing the runtime environment and the dependencies
  for the applications. It also caches support for detecting the core modules or packages
  in the runtime, so it gets a huge boost in deployment performance. Minimal deployment
  time, simultaneously makes it developer friendly and improves the quality of the
  application.</p>\n<ul>\n<li>Abstracted technologies like Nix and Docker.</li>\n<li>Caching
  of dependencies and steps to build.</li>\n<li>Customizable at most of the steps.</li>\n<li>Extensible
  and Developer friendly.</li>\n</ul>\n<h2>Installing Nixpacks</h2>\n<p>Nixpacks are
  the <a href=\"https://search.nixos.org/packages\">nix packges</a> which are used
  with the source code to create a buildpack of their own. The nix packages take in
  the source code of your application and convert it into a simple OCI image e.g.
  Docker image that can be run on various environments. It is similar to buildpacks
  but it's not the same instead it is better and uses a different ecosystem.</p>\n<p>Railway
  has a cool <a href=\"https://nixpacks.com/docs/cli\">CLI</a> tool for creating nixpacks
  on your local system. You can install the nixpack cli from the official documentation
  site. There are a couple of ways to install them on your system.\nYou can install
  the nixpack with the curl command in your terminal which is one of the simplest
  ways to install it.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>curl -fsSL https://raw.githubusercontent.com/railwayapp/nixpacks/master/install.sh
  | bash\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657725394/blog-media/railway-nixpacks-install.png\"
  alt=\"Railway Nixpacks Install\" /></p>\n<p>Once it has been installed we can check
  the status if nixpacks was successfully installed or not.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks --version\n\nOR
  \n\nnixpacks\n</pre></div>\n\n</pre>\n\n<p>So that's how we install nixpacks CLI
  into the system. Now, we can move into the configuration required for creating a
  nixpack from the Django application.</p>\n<h2>Creating a Django Project</h2>\n<p>So,
  I assume here, you have your django project ready. You can pick any django project
  and configure it as mentioned below.</p>\n<h3>Creating requirements file</h3>\n<p>It
  is a good practice to include the <code>requirements.txt</code> file in any python
  based applications. It becomes really easy to set up and give the project a spin.
  It's often the case, that we have to use a virtual environment for creating a pip
  file, otherwise, we might conflict and mix up the globally installed packages with
  project-specific dependencies.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pip freeze &gt; requirements.txt\n</pre></div>\n\n</pre>\n\n<p>This
  will create a <code>requirements.txt</code> file in the current folder, which will
  contain the list of all the dependencies with the version mentioned in. We also
  have other package management system like <a href=\"https://www.meetgor.com/pipenv-intro/\">pipenv</a>
  or <a href=\"https://python-poetry.org/\">poetry</a>.</p>\n<h2>Creating a Nixpack
  for the project</h2>\n<p>Finally, we can now start creating nixpacks from the source
  code of the django application. So, make sure the Django project is running on your
  local system. Set up a virtualenv and database as per your local environment. We
  will be using the build command for creating the nixpack of our django project.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build .\n</pre></div>\n\n</pre>\n\n<video
  width=\"800\" height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657881139/blog-media/nixpacks-demo-base.mp4\"
  type=\"video/mp4\">\n</video>\n<p>This is the simplest of commands you can run to
  create a nixpack. You need to be in the source file where all your files are located.
  In Django, we call it the <code>BASE_DIR</code>, the same folder where your <code>manage.py</code>
  resides. So, be on that path and run the above command. It would pick up a few things
  from the source code itself. There are also some default values picked up after
  analyzing the source code, for example in the case if it detects the django project,
  it attaches the <code>python manage.py migrate &amp;&amp; gunicorn &lt;project_name&gt;.wsgi</code>
  command as the default build command.</p>\n<p>Similarly, there are other default
  options like the version of the runtime, installation steps, etc. We can see that
  the port is not accessible from the docker container that we ran it is because of
  two reasons either we have not exposed the port or the port is not bound with gunicorn.
  We will see that configuration later, it's quite easy to bind the local port to
  the docker container.</p>\n<h3>Steps for creating Nixpacks</h3>\n<p>Now, we need
  to understand the process of the creation of nixpacks. There were several stages
  that you can see in the above video.</p>\n<ol>\n<li>Planning</li>\n<li>Building</li>\n</ol>\n<h4>Planning
  Phase</h4>\n<p>The first step was quite important, to understand the runtime environment.
  This is critical because it can decide how our application installs, runs, or even
  performs in the actual container environment. This is just the phase where we define
  the sets of commands or nix packages to be used for the building of the application.</p>\n<p>You
  would get a clear look in the nixpack cli, it gives a list of the planned instructions
  before the build process starts in the build command. You can get the plan before
  building the application by using the <code>plan</code> command in the CLI.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks plan .\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657960771/blog-media/nixpacks-plan-command.gif\"
  alt=\"Nixpacks Plan Command\" /></p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657960816/blog-media/nixpacks-plan-cmd.png\"
  alt=\"Nixpacks Plan Command Output\" /></p>\n<p>So, this gives a <code>nix</code>
  output of all the configuration needed to build and install the application and
  it's dependencies. When you would have built the application i.e. the base build
  command, it also shows in the CLI, some useful information before moving to the
  build step.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657960880/blog-media/nixpacks-build-plan.png\"
  alt=\"Nixpacks Plan Steps in Build Command\" /></p>\n<p>So the list of the following
  procedure is listed in the planning phase of building the application.</p>\n<ul>\n<li>Packages/base
  runtime environment</li>\n<li>Installation commands</li>\n<li>Build commands</li>\n<li>Commands
  to Start Application</li>\n</ul>\n<p>In the case of django, the following correspond
  to the planning attributes.</p>\n<ul>\n<li>Python version as the application runtime.</li>\n<li>Installing
  dependencies via pip/pipenv/poetry</li>\n<li>Building the Django app (collecting
  static files, setting database, etc)</li>\n<li>Running the Django app (gunicorn/Nginx
  web server to run the app)</li>\n</ul>\n<p>So, hopefully, this gives you a better
  understanding of what is going on in the nixpack CLI. We are trying to automate
  the process of deployment and building the application for repeated deployments
  i.e. in a continuous integration/delivery system.</p>\n<h4>Building Phase</h4>\n<p>This
  is the phase, where actual installation, setup, and configuration takes place at
  the application level. In the build phase, we have several layers where things happen
  like installing the core packages for the application, installing dependencies,
  setting/generating up the necessary steps for the proper running of the application,
  and finally steps to run the application. This process creates an OSI(a standard
  used in containers/virtualization technology) image of the application. We can run
  the image and thus create a container of the application. You will need docker or
  any other containerization tool for the building of images and running containers.</p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657961801/blog-media/nixpacks-build-command.gif\"
  alt=\"Django Build Command\" /></p>\n<p>For the Django application, we have several
  steps of the application to be followed. We can use python as the base runtime environment,
  install the python packages, perform other commands like setting up admin accounts,
  static/media files, database setup, and finally run the steps to run the application
  at the container level.</p>\n<ul>\n<li>Installing python as a <a href=\"https://search.nixos.org/packages?channel=22.05&amp;show=python38&amp;from=0&amp;size=50&amp;sort=relevance&amp;type=packages&amp;query=python\">Nix
  Package</a></li>\n<li>Installing all python packages provided in the <code>requirements.txt</code>/<code>Pipfile</code>/<code>pyproject.toml</code>
  files.</li>\n<li>Call commands from the environment to set up the applications like
  <code>createsuperuser</code>, <code>collectstatic</code>, <code>makemigrations</code>,
  pull data, management commands, etc. There are a lot of things that can be done
  here depending on the application.</li>\n<li>Finally, the step to run the Django
  app, usually gunicorn/Nginx server is used for running the django application on
  a port with the host.</li>\n</ul>\n<p>So, this is what the build phase does, this
  is the whole and soul of the nixpacks. We literally do every installation setup
  and configuration of the application. Though the planning phase is equally important,
  a single missing data can ruin the build phase.</p>\n<h3>Understanding the build
  phase for Django</h3>\n<p>We ran the build command for creating the nixpack build,
  this started with planning the application configuration and then building up the
  application. This build phase was also divided into further processes like installing,
  running commands, copying actual source code to an image, and all the docker-related
  stuff that is required to create an image for a django application.</p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657961691/blog-media/nixpacks-build-process.png\"
  alt=\"Django Application build nixpacks\" /></p>\n<p>So, we have used the Debian
  nixpack, which sets as the base runtime for the application. Railway provides a
  <a href=\"https://github.com/railwayapp/nixpacks/pkgs/container/nixpacks\">package</a>
  of the Debian image as the base runtime for our application. This is where we will
  run all the build processes on. This Debian image will be used for installing all
  types of dependencies and layers of language-specific runtime installation in the
  form of <a href=\"https://search.nixos.org/packages\">nix packages</a>.</p>\n<p>Now,
  we have the base image of debian, the nixpack build command actually starts executing
  the <code>Dockerfile</code>, this is an auto-generated step after the planning phase.
  So, with the instructions in <code>Dockerfile</code>, steps are executed one after
  other just as a normal Docker image build. This will generate the image for the
  application and after a while, because this process takes a while on the first iteration
  locally, after the build process has been completed, it will give a container name
  for you to run. This can be used to test the application locally, in production,
  the container is instantly created after the image has been generated.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>RUN:\n\ndocker run -it &lt;container_id_or_name&gt;\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657966274/blog-media/nixpacks-build-run.png\"
  alt=\"Nixpacks build command run container\" /></p>\n<p>This is the command for
  running your application, this marks the end of the build process and also the build
  command provided with the CLI.</p>\n<h3>Build Command Parameters</h3>\n<p>The build
  command in the nixpacks CLI provides a few parameters or arguments to customize
  how to output the result and build the application, you can definitely provide the
  configuration in the application source code itself, but it is nicer to have it
  locally before deploying the application.</p>\n<h4>Give a name to the Nixpack Image/Container</h4>\n<p>The
  first parameter which might be helpful is to provide a name to the application at
  the build time. This becomes useful for running the container, this helps in avoiding
  long container names and giving a context of the nixpack.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build . --name &lt;project_name&gt;\n\nOR
  \n\nnixpacks build . -n &lt;project_name&gt;\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962253/blog-media/nixpacks-build-name.png\"
  alt=\"Nixpacks Build Command Name Image\" /></p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962328/blog-media/nixpacks-build-name-image-run.png\"
  alt=\"Nixpacks Build Command name run\" /></p>\n<p>This gives a name to the image
  which has been built. Thereby providing a better context for the user to run the
  image and create a container out of it.</p>\n<h4>Output the Built Image to a folder</h4>\n<p>This
  is the command that can output the built application into a  provided folder. This
  parameter will not run the docker step thereafter; hence, no image is created if
  you provide an output folder. Though the folder will contain the <code>Dockerfile</code>
  and <code>environment.nix</code> files for creating the image and running the container.
  <strong>Make sure the output folder is NOT in the application folder itself, it
  will result in errors.</strong> The output command will not create an image but
  the process will be definitely executed in order to generate the <code>Dockerfile</code>
  and <code>environment.nix</code> files.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build . --out ../blog_image\n\nOR\n\nnixpacks
  build . -o ../blog_image\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962407/blog-media/nixpacks-build-output.png\"
  alt=\"Nixpacks Build Command Output folder\" /></p>\n<p><strong>Dockerfile</strong></p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962479/blog-media/nixpacks-build-output-folder.png\"
  alt=\"Nixpacks Build command ouptut\" /></p>\n<p><strong>environment.nix File</strong></p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657969127/blog-media/nixpacks-build-environment-nix-file.png\"
  alt=\"Nixpacks environment.nix file\" /></p>\n<p>So, this will output the built
  application into the provided path. The output folder should necessarily be out
  of the application folder as it makes no sense to output in the same folder as the
  application since the nixpacks CLI will consider the folder as the application folder.</p>\n<h3>Provide
  a Install/Build/Start Command</h3>\n<p>We can provide the commands at the install
  phase/build/start phase of the application to the build command in order to build
  the app with non-default or custom commands. This will add up to the docker steps
  that will involve making the build for the application.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build . --build-cmd
  &#39;python manage.py collectstatic&#39;\n\nOR\n\nnixpacks build . -b &#39;python
  manage.py collectstatic&#39;\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657962514/blog-media/nixpacks-build-build-command.png\"
  alt=\"Nixpacks Build Command Providing install/build/start commands\" /></p>\n<p>These
  kinds of parameters can be passed similarly for <code>install-cmd</code> and <code>start-cmd</code>
  as <code>-i</code> and <code>-s</code> respectively. We can further chain up the
  commands and customize the build process as per the application's requirements.</p>\n<h3>Providing
  environment variables to image</h3>\n<p>The environment variable can be passed to
  the build command for parsing to the application. This can be used for parsing additional
  or optional environment variables to the application image.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>nixpacks build . --env &#39;NAME=VALUE&#39;\n\nnixpack
  build . --env &#39;DATABASE_URL=postgres://postgres:postgres@localhost:5432/techstructive_blog&#39;\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657963255/blog-media/nixpacks-build-env-variable-db-url.png\"
  alt=\"Nixpacks Build Comand parsing environment variables\" /></p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657963302/blog-media/nixpacks-build-db-migrate-env.png\"
  alt=\"Nixpacks Build Command env variable migrate\" /></p>\n<p>Here, we provide
  the environment variable <code>DATABASE_URL</code> to the build command and this
  is parsed to the application image. Thereby when the image is run as a container,
  it is parsed as a normal environment variable and thereby is available for utilization
  from the application.</p>\n<p>For further references on the build command arguments,
  you can follow the <a href=\"https://nixpacks.com/docs/cli\">documentation of nixpack</a>
  by railway app.</p>\n<h3>Creating a Procfile</h3>\n<p>This is important for telling
  any buildpack in this case nixpacks to understand the process to start for this
  web application. For django it is simply to add the web process like to mention
  the <code>wsgi</code> app with the project name. We can use the gunicorn as the
  web server in production.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Procfile\n\nweb: gunicorn
  &lt;django_project_name&gt;.wsgi\n</pre></div>\n\n</pre>\n\n<p>This is the Procfile,
  this is a file type without the extension. So, this is a typical Django application
  Procfile, you can also use the other variants of Procfile for applying migration
  for every web process start-up.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Procfile\n\nweb: python
  manage.py migrate &amp;&amp; gunicorn &lt;django_project_name&gt;.wsgi\n</pre></div>\n\n</pre>\n\n<p>As
  we saw in the base build command, the local server was not able to listen to the
  gunicorn server in the container, so we need to bind the gunicorn server to the
  local port.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Procfile\n\nweb: python
  manage.py migrate &amp;&amp; gunicorn &lt;django_project_name&gt;.wsgi -b :8000\n</pre></div>\n\n</pre>\n\n<p>So,
  we use the <code>-b</code> option in the gunicorn command to bind the port in the
  container to the port in the local machine. Now, if we build the application and
  forward the port to the <code>8000</code> port in the local machine, we will see
  our application running</p>\n<video width=\"800\" height=\"450\" controls>\n  <source
  src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657964597/blog-media/nixpacks-local-bind.mp4\"
  type=\"video/mp4\">\n</video>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>docker run -p 8000:8000 -it
  &lt;container_id&gt; \n</pre></div>\n\n</pre>\n\n<h3>Specifying the Python version</h3>\n<p>This
  is created for specifying the python version for building the Django application
  or any other python app.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># .python-version\n\n3.10\n</pre></div>\n\n</pre>\n\n<p>Save
  the <code>.python-version</code> file with just the python version like <code>3.6</code>,
  <code>3.10</code>, etc. and this will be picked by the build command while creating
  the build image.</p>\n<h2>Deploying the Django Application</h2>\n<p>After we looked
  at the nix picks specifications, we can now deploy our Django application with nixpacks
  on Railway. So, you can follow up with the <a href=\"\">Railway Deployment</a> Article
  for setting up your Django app for deployment at the railway. This usually involves
  a few steps like creating Procfile(not necessary but recommended), requirements.txt(necessary
  to pull dependencies), and the python version which is chosen as <code>3.8</code>
  as default. The further steps are to create a GitHub repository to link with the
  Railway app and create a PostgreSQL database service on the railway platform.</p>\n<h3>Create
  configuration files</h3>\n<p>As we have seen we will require a <code>requirements.txt</code>
  file, <code>Pipfile</code> or a <code>pyproject.toml</code> file for listing out
  and installing dependencies for our django application. This can be done with various
  commands like:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># For requirements.txt and
  virtualenv\npip freeze &gt; requirements.txt\n\n# Autogenerated Pipfile for pipenv\n#
  Autogenerated pyproject.toml for poetry\n</pre></div>\n\n</pre>\n\n<p>So, this file
  should be present on the base directory of the django application in order for the
  nixpack to pick up and install the python packages. Also, for customization of the
  start command in the build process, you can create a <code>Procfile</code> as discussed
  earlier in order to run commands to start the Django web server.</p>\n<p>The python
  version can be specified with the <code>.python-version</code> file with just the
  version name as <code>3.9</code>, <code>3.10</code>, etc. OR we can add an environment
  variable <code>NIXPACKS_PYTHON_VERSION</code> to the python version we want.</p>\n<h3>Create
  and Linkup a GitHub repository for existing Django projects</h3>\n<p>We can create
  a GitHub repository and link up the project to the Railway platform and thereby
  creating an automated build for every push.</p>\n<p>The below video will explain
  how to set up the GitHub repository for the Railway app.</p>\n<video width=\"800\"
  height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1652970864/blog-media/django-deployment/railway_project_init.webm\"
  type=\"video/mp4\">\n</video>\n<video width=\"800\" height=\"450\" controls>\n  <source
  src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657965284/blog-media/railway-platform-github.mp4\"
  type=\"video/mp4\">\n</video>\n<h3>Setup environment variables</h3>\n<p>We can use
  <code>python-environ</code> to set up environment variables in a Django application,
  we will require these environment variables for attributes like <code>SECRET_KEY</code>,
  <code>DATABASE_URL</code>, email-stuff, etc. These are quite handy to avoid leaking
  sensitive information to the open source project on GitHub.</p>\n<p>You can install
  the <code>python-environ</code> package with pip or any other package management
  tool as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install python-environ\n</pre></div>\n\n</pre>\n\n<p>After
  installing the package, we can set up the environment variable in the settings file.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># &lt;project_name&gt;/settings.py</span>\n\n<span class=\"kn\">import</span>
  <span class=\"nn\">os</span>\n<span class=\"kn\">from</span> <span class=\"nn\">dotenv</span>
  <span class=\"kn\">import</span> <span class=\"n\">load_dotenv</span>\n\n<span class=\"n\">BASE_DIR</span>
  <span class=\"o\">=</span> <span class=\"n\">Path</span><span class=\"p\">(</span><span
  class=\"vm\">__file__</span><span class=\"p\">)</span><span class=\"o\">.</span><span
  class=\"n\">resolve</span><span class=\"p\">()</span><span class=\"o\">.</span><span
  class=\"n\">parent</span><span class=\"o\">.</span><span class=\"n\">parent</span>\n\n<span
  class=\"n\">load_dotenv</span><span class=\"p\">(</span><span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">path</span><span class=\"o\">.</span><span
  class=\"n\">join</span><span class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;.env&quot;</span><span class=\"p\">))</span>\n</pre></div>\n\n</pre>\n\n<p>After
  loading the environment variables, we can access them with <code>os.env(&quot;ENV_NAME&quot;,
  default=&quot;&quot;)</code>, this will load the environment variable with the name
  or we can provide a default value.</p>\n<h3>Attach a PostgreSQL database service</h3>\n<p>You
  can add a PostgreSQL database service to your Django Railway app by attaching a
  service. This will add a new service along with the django application, so these
  two act as different entities within a railway project.</p>\n<video width=\"800\"
  height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1652963718/blog-media/django-deployment/postgres_spinup_railway_d2xkpt.mp4\"
  type=\"video/mp4\">\n</video>\n<p>You can then access the <code>DATABASE_URL</code>
  from the connect settings and copy the database URL and set it as an environment
  variable in the django railway project service. This will link up the Django app
  and the PostgreSQL database. While setting it up locally, you can use the <code>.env</code>
  file and add the environment variable there.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># environment variable\nDATABASE_URL=postgres://username:password@hostname:port/db_name\n\n#
  local database postgres\nDATABASE_URL=postgres://postgres:postgres@localhost:5432/db_name\n</pre></div>\n\n</pre>\n\n<video
  width=\"800\" height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657964943/blog-media/railway-postgres-spinup.mp4\"
  type=\"video/mp4\">\n</video>\n<p>This will set up the database environment variable
  and you can access these from the <a href=\"http://settings.py\">settings.py</a>
  file with the <code>env.db</code> function as follows:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">env</span><span class=\"o\">.</span><span class=\"n\">db</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;DATABASE_URL&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">default</span><span class=\"o\">=&lt;</span><span class=\"n\">local_database_url</span><span
  class=\"o\">&gt;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  we can finally use the database from the Railway app in our Django application once
  the environment variable is correctly used.</p>\n<h3>Choose the Buildpack</h3>\n<p>We
  can choose a buildpack for our Django application in the Railway platform, we have
  options like</p>\n<ol>\n<li>Heroku Buildpack</li>\n<li>Railway Nixpacks</li>\n<li>Paketo
  Buildpack</li>\n</ol>\n<p>As of the writing of the article, on 16th July 21, the
  Railway has made <code>Nixpacks</code> the default buildpack for an application
  :) It was the <code>Heroku</code> Buildpack as a default earlier. So, that is a
  cool thing, you can toggle these settings for choosing the buildpacks from the project
  settings.</p>\n<p>Railway Dashboard Choose BuildPack</p>\n<video width=\"800\" height=\"450\"
  controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657964699/blog-media/nixpacks-railway-dashboard.mp4\"
  type=\"video/mp4\">\n</video>\n<h3>Deploy to Railway with Nixpacks</h3>\n<p>Now,
  we have seen how to set up the nixpacks, we had the Postgres database setup, and
  we can finally deploy our application to the railway platform with nixpacks.</p>\n<video
  width=\"800\" height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1657965560/blog-media/railway-nixpacks-deploy.mp4\"
  type=\"video/mp4\">\n</video>\n<p>So, we simply can configure the source code or
  include the <code>environment.nix</code> file into the source code for desired behavior.
  The nixpack selection can be done based on the source code or the presence of <code>environment.nix</code>
  and that's why we can rely on expected behaviors from the deployment builds.</p>\n<h2>Summary</h2>\n<p>So,
  nixpacks is a great way to deploy an application, for me it's an automated docker
  deployment, it basically creates docker images of the application and runs it with
  the appropriate environment. There is a lot of language support on nixpacks currently
  on Railway, you can check them out on the official website. Every programming language
  has specific requirements for managing or installing dependencies and packages,
  the nixpacks manage them automatically for us.</p>\n<h2>Conclusion</h2>\n<p>So,
  from this post of the <a href=\"https://www.meetgor.com/series/django-deployment/\">Django
  Deployment Series</a>, we were able to understand how to deploy a Django application
  on the Railway app with Nixpacks which are a very intuitive way to deploy apps.
  We covered what are nixpacks, the process of building an application with nixpacks,
  and deploying a existing, new Django project on the railway with nixpacks. We also
  explored the various commands provided in the nixpacks CLI to build. plan a Django
  application.</p>\n<p>Hopefully, you were able to understand the concept of nixpacks
  and how they can automate the process of containerization and deployment. Thank
  you for reading, if you have any queries or feedback, you can leave them down in
  the comments or freely drop them on social media. Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/django-deploy-railway-nixpacks.png
long_description: We have seen how to deploy a Django application on railway app in
  the  A buildpack is a set of programs that turns your source code into a container
  image. So it is basically a tool for converting your application into a deployment-ready
  state with t
now: 2025-05-01 18:11:33.312690
path: blog/posts/2022-07-15-Django-Railway-Nixpack.md
prevnext: null
series:
- Django-Deployment
- Django-Series
slug: django-deploy-railway-nixpacks
status: published
tags:
- django
- python
- railway
templateKey: blog-post
title: Deploying Django Project with Railway Nixpacks
today: 2025-05-01
---

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