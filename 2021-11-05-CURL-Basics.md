---
cover: ''
date: 2021-11-05
datetime: 2021-11-05 00:00:00+00:00
description: We all might have used the curl command but might be unaware of it. It
  Curl or cURL command is the utility or tool to access the internet from the command
  Befor
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-11-05-CURL-Basics.md
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643289075/blogmedia/bqnrrfaeaqfaj7hezxjx.png
long_description: We all might have used the curl command but might be unaware of
  it. It Curl or cURL command is the utility or tool to access the internet from the
  command Before we talk about how to use the curl command let The above is a basic
  structure of the curl
now: 2022-07-24 13:37:38.454515
path: blog/posts/2021-11-05-CURL-Basics.md
slug: curl-basics
status: published
subtitle: An short simple introductory guide to the curl command
tags:
- bash
- linux
- networking
templateKey: blog-post
title: Basics of curl command
today: 2022-07-24
---

## Introduction

We all might have used the curl command but might be unaware of it. It's super old
and still serves a great purpose. It has been available since 1996 and still is
widely used in many embedded technologies, web API testing, CLI applications,
etc. In this article, we'll see some basics of using the curl command along with
its applications.

## What is the curl command?

Curl or cURL command is the utility or tool to access the internet from the command
line interface using various protocols. This looks trivial but it can blow up
your mind! Most people use this tool for fetching and processing the
data from the servers/internet from their terminal without the browser but
there is a lot more to it. It is used in various embedded devices for accessing
the network in a lightweight and accessible way. Let's see how you can use the curl
command from the very basics.


## Why do we need it?

Before we talk about how to use the curl command let's talk about why might we need
that? There are a lot of reasons and it even depends on the application you are
using.  You can use curl to test your API, well there are other tools like
POSTMAN, Insomnia, etc but for keeping things simple you can quickly get in
with curl and test some endpoints.  You might require curl for creating some
CLI applications that require fetching/posting to an URL over the internet.
If you are using the terminal, curl integrates really very well with the shell
programming languages like BASH, ZSH, etc So, after making WHY out of the way,
let's start with the actual content. 


## Structure of curl command

**curl or Client URL is a command-line utility that helps in accessing/posting
data with various protocols over the internet.** It basically serves as a
bare-bones browser URL search bar.  You can't render those pages like the
actual GUI, and all but you can get is the HTML source code, JSON response,
etc.  That's still quite powerful and used in tons of applications. 

```
curl URL arguments 
```

The above is a basic structure of the curl command. We see the argument
structure in-depth in the next section. Firstly, let's take a simple curl command with just the URL is given.

```bash
curl "https://github.com"   
```
From this query to `github.com`, you are literally going to `GitHub.com` and getting a response as the entire HTML source code of the page.
If you don't want to spam the output in the terminal, you can redirect the output to a file.

```bash
curl "https://github.com" >temp.html
```
With this command, we store the output of the command in the file temp.html, it can be any other file you like. 

### Arguments 

It turns out that you can even parse in certain arguments to the `curl` command to get some desired and modified results. Let's take a look at some of them.
The [entire list of arguments](https://curl.se/docs/manpage.html) is quite huge
and baffling, but this shows how customizable the command is. 

- `-s` (silent the progress bar)
- `-X` (web requests `POST, GET, etc` to the URL)
- `-o` (output to a file)
- `-H` ( provide Header to the request)
- `-d` (providing the data e.g. in POST request)

```bash
curl -s -o "https://github.com" temp.html
```

This command doesn't load the progress bar and simply outputs the response in a
file, making the execution process in the terminal clean.

### Integration with other commands 

As said, the `curl` command can be well integrated with the other commands using piping in shell, assigning to variables, and so on.

Let's see how we can convert the `JSON` response to a BASH variable.

```bash
resp=$(curl -H "api-key: N2vDzMyEeYGTxjUTePhC8bYd" https://dev.to/api/users/me)

echo $resp
```   
Here, we are fetching the `JSON` response from the `dev.to` [API](https://developers.forem.com/api/),The wired string `N2vDzMyEeYGTxjUTePhC8bYd` is my [dev.to API token](https://dev.to/settings/account)(don't worry I have revoked it:) ) we have provided an argument `-H` that is a Header for accepting a `Json` response. 
We can store the contents of the curl command by using the `$( )` and assigning that to the variable name of your choice.

```bash
username=$(curl -H "api-key: N2vDzMyEeYGTxjUTePhC8bYd" https://dev.to/api/users/me | grep -o -P '(?<=username":").*(?=","name)')
```
Here, we have stored the username from a `JSON` response to the variable username. We have piped the curl command so that we can work with that `JSON` response and modify the contents and then store the final results in a variable.
In this case, we are using `grep` to filter out the content between the key `username` and `name`, thus we get the value we desired. To see the value you can always run the echo command as below:
```bash
echo $username
```   
So, that's how the `curl` command integrates flawlessly with BASH and other shell programming languages. 

## Where is it used?

`curl` is actually used in API testing, CLI applications, Web Scrapping, etc. It's a great tool for terminal lovers. Let's see where we can use the curl command actually to make some good projects.

### API Testing

We can use, `curl` to test an API, it might be API you would have made or to simply test and play with other API available publicly. You can get an in-depth guide about [Testing a REST API with curl](https://www.codepedia.org/ama/how-to-test-a-rest-api-from-command-line-with-curl/).
Actually, curl can do more than just testing, I have made a [bash script](https://gist.github.com/Mr-Destructive/80860664b1014ef0b94092d68ead1044) that actually posts some data over a database through the API so that I don't have to do that manually. That is the kind of automation I love to do and curl! Just did that.

### Web Scrapping

Web-scrapping is usually trending with Python, but I have done that with BASH.
That might be an outdated idea but is a good task to learn the basics of
Web-scrapping with BASH ;). I must say that sed, awk, grep are the tools are
powerful like heck in doing these tricks. I have made this
[crypto-currency](https://mr-destructive.github.io/techstructive-blog/bash/2021/07/15/BASH-Crypto-Coingecko.html)
and
[dictionary](https://mr-destructive.github.io/techstructive-blog/bash/2021/07/27/BASH-script-dictionary-scrap.html)
scrapper with BASH. Web-scrapping can be done with the curl command by fetching to
an API if any or any website. We need to search and find the particular fields,
classes, or ids the elements the required data might be into and then extract
and filter using the tools like grep, sed or awk.


### CLI Applications

We can make CLI applications like creating a terminal view of existing
applications using their APIs or website. I recently made a CLI for
[cross-posting articles](https://github.com/Mr-Destructive/crossposter) to
dev. to, hashnode and medium. That is a project still in progress(tons of bugs)
but still serving a decent job. Definitely `curl` might not be the only command
that works here, but the project might look so incomplete without `curl`.

**There might be other applications as well, who knows there is a lot to do with this command.** If you know one, please let everyone know in the comments.

### References:

Special Thanks to the creator of the curl command - [Magnus Daniel Stenberg](https://github.com/bagder) and the developers who are still contributing and maintaining the great project.
 
### Conclusion

So, from this article, we were able to understand the basics of the `curl` command and understand its applications in actual programming stuff. Hope you liked it. Thanks for reading and until then Happy Coding :)
<div class='prevnext'>
    <style type='text/css'>
    :root {
      --prevnext-color-text: #eefbfe;
      --prevnext-color-angle: #ff6600;
      --prevnext-subtitle-brightness: 3;
    }
    [data-theme="light"] {
      --prevnext-color-text: #1f2022;
      --prevnext-color-angle: #ffeb00;
      --prevnext-subtitle-brightness: 3;
    }
    [data-theme="dark"] {
      --prevnext-color-text: #eefbfe;
      --prevnext-color-angle: #ff6600;
      --prevnext-subtitle-brightness: 3;
    }
    .prevnext {
      display: flex;
      flex-direction: row;
      justify-content: space-around;
      align-items: flex-start;
    }
    .prevnext a {
      display: flex;
      align-items: center;
      width: 100%;
      text-decoration: none;
    }
    a.next {
      justify-content: flex-end;
    }
    .prevnext a:hover {
      background: #00000006;
    }
    .prevnext-subtitle {
      color: var(--prevnext-color-text);
      filter: brightness(var(--prevnext-subtitle-brightness));
      font-size: .8rem;
    }
    .prevnext-title {
      color: var(--prevnext-color-text);
      font-size: 1rem;
    }
    .prevnext-text {
      max-width: 30vw;
    }
    </style>
    
    <a class='prev' href='/oop-cpp-p3'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Object-Oriented Programming in C++: Constructors</p>
        </div>
    </a>
    
    <a class='next' href='/golang-mutable-immutable'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Golang: Mutable and Immutable Data Types</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>