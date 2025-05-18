{
  "type": "posts",
  "title": "Filter and Find an Issue on GitHub",
  "subtitle": "A quick guide to filter and sort issues as per your interest/skills on GitHub",
  "date": "2021-09-09 17:42:30 +0530",
  "status": "published",
  "tags": [
    "github",
    "open-source"
  ],
  "slug": "find-filter-github-issues",
  "image_url": "https://meetgor-cdn.pages.dev/github-filter-issues.png"
}

## Introduction

Are you stuck in finding an open-source project to contribute to? 

We will see how you can pick up an issue on GitHub appropriate as per your preferences of languages, labels, complexity, and thus you can find a Community or a project to work and continue with further contributions.

This process might not be as efficient but is quite helpful for beginners or people getting started to contributing to Open Source.
 
## Understand the search bar

I assume you have your GitHub account already created. If not go ahead at [Github](https://github.com/join) and create one. On the Home page, you can easily navigate to the `Issues` tab and you will see something like this:

![Issues tab](https://cdn.hashnode.com/res/hashnode/image/upload/v1631190578909/UBpq3rb0H.png)

Now, you won't find any issues if you haven't created any. But if you look at the search bar, you will find the reason why it is empty or why there are only the issues that you have created. You will see that in the search bar there is a filter called `author:Username`, which filters the issues which are created by you. You definitely don't want this as you want to search and find other issues by other people/communities. So, simply remove the text `author:Username` from the search bar. Keep rest as it is for now. Now if you press enter after removing the author filter, you will see all the issues on GitHub. 

![Issues removed author](https://cdn.hashnode.com/res/hashnode/image/upload/v1631185853484/e0PyTbgip.png)

There will be a ton of them, very random in terms of programming languages, frameworks, projects, difficulty, type, etc. they are basically the issues created recently on GitHub.
 
In the next section, we will see how to filter those issues as per the programming languages/tools to which you might like to contribute to.

## Add languages

We can add filters to the issues as `language:name`, this will filter all the Issues which have the languages in their codebase. 

For Example:

![Issues language filter](https://cdn.hashnode.com/res/hashnode/image/upload/v1631190679194/8Od1tsdKp.png)

Here, I have filtered the issues which have language as `python`, you can use any language/tool you might want and would love to find some interesting projects to contribute and learn from.

If you want to search by multiple programming languages you can separate the names of those programming languages by a comma `,`.

You can also separate programming languages with space and enclosing all of them under double quotes `""`.

For Example:

Let's search for issues with C, C++, and Java as their programming languages, we can use `language:c,cpp,java` or `language:"c cpp java"`

The above filter will give out all the issues which are created from programming languages either C/C++/Java.
 
You can find more filter options on the [GitHub docs](https://docs.github.com/en/github/searching-for-information-on-github/searching-on-github/searching-issues-and-pull-requests).

## Add labels

You can find issues as per labels marked on them, many issues have a label marked on them to improve their visibility and meta-information about the issue.

We have some labels which GitHub has created already for common scenarios in projects.  

1. `bug`
2. `documentation`
3. `duplicate`
4. `enhancement`
5. `good first Issue`
6. `help wanted`
7. `invalid`
8. `question`
9. `wontfix`

We can even create our own labels by providing the label name and a description. 
 
To search for labels, you can use `label:name of the label`. You can any of the above 9 label tags or any other tag name that you think is popular other than those 9. 

You would have to use double quotes (`""`) to add certain labels with multiple words like `good first issue` or `help wanted`.

For example:

If you search for `label:"good first issue"`, you will get all of the issues(newest first) which have a label `good first issues` tagged on them. 

Similarly, for multiple issues, you can add comma-separated labels as well. Just like `label:bug,"good first issue"` will search for either `bug`, `good first issue` or both. 

![Issues label](https://cdn.hashnode.com/res/hashnode/image/upload/v1631190841185/vrYTLoaaNu.png)

## More Sorting Options

In the rightmost part of the search bar, in the Sort button, you can click on there and find a couple of options like: `newest`, `oldest`, `least commented`, `recently updated`, and so on. If you click on any of them you will see the changes reflected on the list of issues as well as the search bar. 

![Issues sort](https://cdn.hashnode.com/res/hashnode/image/upload/v1631189621396/jO58HkYxH.png)

**After this the stage is yours, you can look at any issue and Understand its objective, then ask yourself can you solve this issue? If yes then read the contribution guidelines, and the rest is sheer skills like git, programming, documentation, etc.**

## Conclusion

Now you can go ahead and start applying the filters on issues and make some contributions to Open-Source on GitHub. We covered some methods and tricks to find and filter out the issues on GitHub based on the programming languages/tools and the labels attached to them.  

This technique can be good for beginners as well as people who want to find quick issues to solve. Feel free to explore and try out different filters and find the issue you are confident to work on. Good Luck!
Happy Coding :)
