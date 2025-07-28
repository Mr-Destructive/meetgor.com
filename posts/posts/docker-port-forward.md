{
  "title": "Docker Port Forwarding",
  "status": "published",
  "slug": "docker-port-forward",
  "date": "2025-04-05T12:33:47Z"
}

<h2>Docker Port Forwarding</h2>
<p>Port forwarding is a process to redirect the communication of one address to other.
It is also known as Port Binding.
We can use <code>-p</code> command to use port forwarding in our local Docker environment.</p>
<pre><code>docker run -p 8000:8000 django-app
</code></pre>
<p>The first port number is the local machine port and followed by a <code>:</code> is the container port number.
SO, the request from the container port are forwarded to the local/outside world in the docker environment.</p>
<p>Additionally, we need to expose the container port first. We can do that in the Dockerfile or by adding a <code>-e</code> argument followed by the port to expose. This will open the port on container to forward the requests to the specified port in the <code>-p</code> option.</p>
<p>In the Dockerfile, we can expose the port by adding the command <code>EXPOSE 8000</code>, or any other port number.</p>
