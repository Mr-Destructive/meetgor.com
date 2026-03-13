---

























type: links
title: 'Package Managers are evil'
date: 2025-09-13
slug: package-managers-are-evil
tags:
  - links
link: 'https://www.gingerbill.org/article/2025/09/08/package-managers-are-evil/'
status: published
description: 'Package Managers are evil'
source: newsletter
newsletter: 2025-09-13-techstructive-weekly-59

























---


## Commentary

- This is a fair take. Absolitely, the left pad example from Theo is one such thing. People just keep on adding packages/libraries without thinking much, in pre-AI days that was the problem.
- But now, with AI, it can spit out code like anything. No need to worry about managing packages, but eventually it will be producing more code which is a liability. AI produced code might be fragile, very like todo: authentication coming soon, like code. If not tested or reviewed, can't trust it.
- He is right on all the points, Golang is batteries included, and clearly defines what a package actually is, its just a folder. You can import anything from the folder. Except only if the functions or structs are capital case (annoying at times, but fine). Having some rule is better than having none. Javascript failed to define a rule, and NPM is a mess. Golang doesn't have a package manager, it just manages itself.
- I also find python dependency management like javascript to be honest, but a little better with terms of completeness. Since people can mess up on the web pretty easily, the things to mess up with Python have a less surface area. If you are aware of what happened to PyPI several times, you know what I am talking about, Its common to manipulate a source of truth and people might find themselves in all sorts of trouble, they would have never imagined. With uv I think it is moving to a better place, but still the core of the problem is from the too much of flexibility, which is fine, and needed even. Python doesn't needs to be like Go.
