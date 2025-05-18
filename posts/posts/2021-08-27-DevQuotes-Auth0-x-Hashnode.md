{
  "type": "posts",
  "title": "Dev Quotes: A platform for developers to quote and get inspired - Auth0 x Hashnode Hackathon",
  "subtitle": "A platform created in a Hackathon for developer to get motivated by some fun and inspireing quotes.",
  "date": "2021-08-27 23:45:00 +0530",
  "status": "published",
  "tags": [
    "hashnode",
    "auth0",
    "django",
    "web-development",
    "python"
  ],
  "slug": "devquotes-platform",
  "image_url": "https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643286845/blogmedia/bhajgzmqaknzpk1z9qdn.png"
}

## Introduction

> No one can understand the joy in finishing a personal project, except the dreamer 

It was a while,since I have posted an article here, as I was busy on a project or a hackathon.

Hello, world! I am Meet a student and a self-taught web developer. I like to make and break stuff, especially when it comes to programming and Linux. I like shell scripting and learning different languages at once, love to learn about Vim and Linux everyday.

Every time I start a project something else comes and distracts me let that be any other programming language or technology. That leads to creating new projects and leaving the one behind unfinished, I know most of the developers face this.  But this time, thanks to Auth0 X Hashnode Hackathon, I was able to create an almost finished project within almost 10 days. Having a deadline and competition creates a mindset to finish a project on time, that's my first takeaway from this Hackathon. OH! this is my first Hackathon by the way, and it has been amazing so far.  

** Applying a framework to do something you desire and then everything working smoothly (after fixing 100s of bugs) is such a great feeling that no one can understand except for the person who just dreamt of it. **

I'll like to share my project which is a web application for the Auth0 x Hashnode Hackathon. Here it goes.

## What is Dev Quotes?

Dev quotes is a web app designed for publishing and viewing quotes related to programming, developer mindset, and all the technicalities involved in a developer's life. It's basically a medium to express the life of developers and get inspired by others.  Here it is [devquotes](https://devquotess.herokuapp.com/)

#### Dark Mode:
![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630074051548/TQz9Koh7l.png)

#### Light Mode:
![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630078314355/VhfLrcjJa.png)


## Why Dev Quotes?
> Developers are not the people who only understand how to write code but they're also the people who can make the code understandable

As a developer, there are often times where you have no motivation left inside, but you never know you might be just a few lines of code away from making another project or fixing a bug. For that, we require some inspiration or a push to break the barrier of.  I am not saying it's just for developers, it's designed for developers but everyone is open to understanding the developers' lives and their struggles. 

I also felt the need to give back some love-crafted web app to the ever-wonderful and supportive dev community. It's a small application but still, I would like to give in something instead of nothing at all. Start small grow big, hopefully :)

## Features

Some of the main features of the web application are as follows:

- **Writedit\Delete Quotes if logged in.**

- **Like / Unlike a Quote.**

- **See all of your quotes.**

- **Randomized Quotes on Homepage.**

- **Dark/Light theme based on Browser's Preference and local storage.**
 
- **The app is mobile responsive as well, though the navbar is a bit wonky with the light/dark mode switch toggle, which can be taken care of soon.**



![dqmob.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630084573444/YEj38dUTD_.png)



## In the making

> Have the curiosity to learn, rest is automated

The project was made by using various inspirational articles and videos about making a web application. But the idea stuck in my mind when I was thinking about the people who don't get inspired as a developer. Like there is no way you can remain sad about being a developer and keep on dealing with imposter syndrome. Every developer has a perspective of programming but there is an infinite number of opportunities if you are curious enough. Just started making the project and got so much into it that I was literally dreaming about it like I saw parts of the webpage. In my dream and I am making it that was genuinely a thing that powered me to complete it. 

The project roughly started on 19th August and almost ended on 26th August, like the actual webpage and its core functionalities. Though on 27th were some styling and extra additions such as the About section and Footer. That was like the most productive week I ever had in my programming journey. That was fun as heck.

## Tech Stack

The Tech Stack involved with this app is :
- `Django`
- `PostgreSQL` 
- `HTML/CSS/JS`
- ` Bootstrap`

 I have not used any front-end end frameworks just because I never found the need to learn them.  I had experience with Django for just 2 months and I am surprised I was able to make it. As obvious I have used Auth0 for authentication in my web application.

### Auth0 integration for Authentication

I must tell you using Auth0 was just flawless addition to my app as I have to do almost nothing, just drop some credentials of the Auth0 application with my Django project using a  [well-documented guide](https://auth0.com/docs/quickstarts)  for every type of framework. Simply straight-forward was the name for integrating authentication in my app.

#### How I used Auth0 with Django

I've used Template tags such as if blocks to verify if the user is authenticated or not. 
```html
<div class="login-box auth0-box before">
		{{ "{% if user.is_authenticated "}} %}
		    <a class="btn btn-primary btn-sm tn-logout " href="/logout">Log Out</a>
		{{ "{% else "}} %}
		    <a class="btn btn-primary btn-sm tn-login " href="/login/auth0">Log In</a>
    {{ "{% endif "}} %}
</div>
```

This was just readily available on their documentation though there were some adjustments as per the project requirements in this code to fit in the place.

I must say, integrating Auth0 is even easier than using Django User Model in some sense as most of the stuff is handled by the Auth0, on our side, we simply have to create the Auth0 specific app with the credentials from the dashboard rest just works flawlessly till now. How sweet and 
 
### Specifications

I won't go in-depth about the technicalities of the project but would like to address certain things. Firstly I have mostly used Class-based views for the major part, certain areas are still function-based just for the simplicity of the application and a few of them are handled and documented by Auth0 so just preferred that. 

Another thing is about Models, I just have a simple single model called `Quote` which has an Author as a Foreign Key from the Django User Model. I would have also created multiple emojis for the like system but I was too excited and in a rush to see the actual app, So just kept it simple. XD
![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630060555499/246ZKnypP.png) 

The rest of the stuff like `URLs`, `Templates`, and `static files` are handled in a neatly organized way as depicted in the below diagram.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630060426600/JHMlrfOKZ.png)

As it might not be clear from that, there are 3 apps -> `auth0login`, `quotes`, and `user`, here `quotes` is kind of the most important app as it has the models, forms, URLs, and the views linked to them here. 

### Hosting

Hosting as you can guess, it's on  [Heroku](https://www.heroku.com/) , quite beginner-friendly and feature-rich. I also have a free addon for PostgreSQL Database here.  It's limited to 10K rows but that's sufficient for a starter app like this in my opinion. Also, it has 500 free hours of dyno, which is also sufficient for a small low-traffic app like this.  

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630076036133/9ofxnM5VN.png)

## Bugs Encountered 
> Love Bugs, they'll give you experience

This is the most exciting and important part if you consider the hackathon because this determines the level of experience that a developer shoes in achieving certain things or features if you will. Faced some typical Django errors as usual but also some unexpected things like Dark mode and light mode clashing together due to poorly written media queries in CSS.

 As usual, the start is often hard, there is no motivation whatsoever in using the admin section to test the database queries and gibberish basic HTML page. In that process, I faced some primary key access issues and was able to understand the concept more clearly by fixing it.

Another instance was with handling post requests in Django which I've failed to do before. I used forms and a hybrid of CSS and bootstrap to style those forms which just works brilliantly. That took some time to figure out the exact working but after a while, it was working charms. 


## Future Updates

As said, I would like to add more like buttons expressing different emotions. Some other features to add are:

- To add more emojis like claps, cheers, and others.
- To add a profile page in the app that would display all the quotes of the particular author and the details related to him/her.  
- Adding some tags to filter out particular types of quotes.
- Improve UI-UX a bit more to make it professional and pleasant.

## Source Code

>Talk is cheap, show me the code - Linus Torvalds

The source code is available at GitHub on this  [Link](https://github.com/Mr-Destructive/devquotes). 
It's freely open for any contribution after the hackathon(mid-September).  Some of the files such as the environment variables, virtual environments, cached data are not uploaded for security and obvious reasons.

Enough of technical talks, let's see the DEMO,


![dqss.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1630073466386/3wgnST5hc.gif)

Silent claps.......

It's not a great UI-UX but works and is not too shabby in my opinion considering it only has base HTML and CSS with a little Bootstrap and Javascript. But ya, a fully functional Backend that's what I was looking for a full-stack app with some decent features. Hope it helps some developers stay motivated and hack into some hackathons like this.

### References used while creating the app:
- [Codemy -John Elder Django tutorial]( https://www.youtube.com/watch?v=B40bteAMM_M&list=PLCC34OHNcOtr025c1kHSPrnP18YPB-NFi)
- [Django - Documentation for Forms](https://docs.djangoproject.com/en/3.2/topics/forms/)
- [Django template randomizer shuffle](https://stackoverflow.com/questions/28837511/django-template-how-to-randomize-order-when-populating-page-with-objects)
- [Auth0 app Django integration](https://www.youtube.com/watch?v=kzN_VCFG9NM)


## Closing Words

> Why developers find solutions to bugs at the stroke of sleeping, that's multithreading in our brains 

Hope you liked the project and hopefully will inspire developers to stay motivated and can focus on their goals more than dealing with imposter syndrome and whatnot. 

Thank you for reading and using the app, for any feedbacks, Twitter handles, comment section, GitHub issues, LinkedIn messages are all freely open. Thanks. Happy Coding :)
