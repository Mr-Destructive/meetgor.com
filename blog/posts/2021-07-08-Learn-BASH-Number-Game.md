---
templateKey: blog-post
title: "Learning BASH by making a Number game"
subtitle: "Understaning BASH concepts by making a number game"
date: 2021-07-08 19:50:35 +0530
status: published
tags: ['bash',]
slug: bash-game-numberjack
image_url: https://meetgor-cdn.pages.dev/bash-script-game.webp
---

## Introduction

OK! Learning BASH can be quite confusing without a proper goal in hand. So this will be a pretty good idea to start learning BASH and have a ton of fun. In this little time, we'll make a Number game which I have designed myself last year in C++, which took about 3 months due to lazy research and wasting time. But I was surprised that I made this game within two hours in BASH. You can refer to the game instructions in this  [repository at Github](https://github.com/Mr-Destructive/NumberJack).

## Concept 
The game will ask a number between 0 and 9 to the user. Then a list of 10 numbers shuffled in a random order will appear from of the user along with another list used for indexing the numbers from the array. The user has to select the index beneath its chosen number to proceed ahead. The game loops until the user has failed to enter the correct index of the number or the time for input was exceeded by 5 seconds. The user will get a point for every successful hit. So that is probably the introduction of the game so, let's dive into the specifications.

## Specifications of the Game in BASH
The game is a number based which means it will need Arithmetic operators a lot. In fact, we'll need a few complex functions such as shuf. We will very frequently use while and for loops to perform some tasks such as filling and printing array and the game loop. We'll use some flag variables to indicate the current situation in the game and finally some arithmetic on arrays and numbers. 

## Script Explanation
The game is quite simple to understand. You just have to select the number beneath your chosen number within 5 seconds in the shell script. We will create a menu-like display in the terminal by simple echo command and formatting. Before the menu, we will have a while loop that will iterate until the user enters 3 which is stored in variable `ch` which is initialized to 0 in the beginning so as to enter the loop for the first time. A while loop starts with the do statement and ends at the done statement.
```bash
while [ condition ];
do 
# statements
done
```
For loop can be different based on the scenario. We'll use a range-based for loop to iterate over a range of numbers using the { } operators. For loop also has do as the beginning of the loop and done as the end of the loop.

```bash
for i in {1..5};
do 
#statements
done
```
We'll also use some If-else statements just to check for the correct user input and checking the exit status. The if statements have `then` to start the block and `fi` to end the if block. 
```bash
if [ condition ];
then
    #statements
elif
    #statements
else
    #statements
fi
```


We use a read statement with the argument -p to have a prompt to the user for some information on the input. The input of choice from the menu i.e 1 to play, 2 for Instructions, and 3 to exit are stored in the variable `ch`. If the input is 1, the game will start and it will ask the user for the number `n`, which is the number used throughout the loop until the game is over. 

Now we have the number for the rest of the game, we need to generate the list for the user to select the number from. We will have a flag sort of to check if the user has entered the correct number which is `c`, this will store 0 for correct input(number x) and 1 for incorrect input. It is initialized with 0, again to enter the while loop once before the generation of numbers. 

To generate and **shuffle 10 numbers which should not have any repeated numbers**, as it can have multiple numbers which might be unfair also it might happen that the number chosen by the user might not be present due to repetition. So to avoid that mischief of pseudo-random numbers we have to generate distinct 10 numbers from 0 to 9 in this case. For that, we are gonna use a command in BASH called `shuf` which can create some permutation of the elements in a list/array or a sequence of numbers in an input stream. We are gonna use `shuf` to generate a random sequence of 10 numbers from 0 to 9 using the command `shuf -i 0-9 -n 10`. 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625748675622/1Li6h3_vX.png)

You can see it generated a list of shuffled numbers between 0 to 9 so there are 10 numbers. We'll store this in result an array to access and print them later. You can refer to  [this](https://www.geeksforgeeks.org/shuf-command-in-linux-with-examples/)  and  [these](https://www.howtoforge.com/linux-shuf-command/)  articles for understanding shuf.  

The main thing is taken care of, now we need to print the list and also print another list to indicate the index of numbers to the user. We will print the list without a for loop using the `@` variable. If you are new to BASH and want a bit guide on BASH please do check out my series on  [BASH scripting](https://techstructiveblog.hashnode.dev/series/bash-scripting), I have this all covered. So using `@` we can print the entire array in BASH. 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625749273007/hf_y4Fm53.png)

To print the lower list of indices, we'll use a range-based for loop i.e it will iterate over a range(in this case between 1 to 10), and assign each element the value of the counter i.e from 1 to 10. We are using `r` as the shuffled list and `a` as the indices list. And print this array with the same method.

After the generation and printing of lists are complete, we'll take the input from the user for the index of his/her number. We'll use an argument in read known as timeout, which will give a stop to the input stream after several seconds provided in the argument. In this case, we will use 5 seconds as a timeout for the input of the index. `read -t 5 -p "Enter the index of your number : " x `
We'll store the input in `x` variable and access it later for verification. 

Next, we will check if the input was done before the timeout or not. For this, if the user input before timeout, we can proceed ahead but if the time was over, then we'll get an exit status above 128 so we use this as a checker for the timeout in the input. I came to this via this  [article](https://www.linux.org/threads/exit-script-by-timeout-if-delay-of-read-input-in-command-line.15905/), really very helpful. We will break the loop and make the flag `c` as 1 indicating an improper input and thus it'll show "GAME OVER". But if you were fast enough then we'll check that the index of the shuffled array has your chosen number or not, we used this `${r[$(($x))-1]} -eq $n` to check for the correct number. Why -1? If you remember indexing in the array by default starts with 0, as we have started the second list from 1 hence every element will become offset by 1 hence to avoid that we'll subtract one to refer to that index. 

If the index of the number was equal and correct, well done we'll increment the counter of points `p` by one and if it was incorrect, the flag will be set to one as previously said and we'll break the loop. After coming out of the loop, we'll check if the status flag `c` was 1 if yes, then print the GAME OVER and display the points earned. And that is it. Let's take a look at some gameplay :)

![numbjackbash.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1625753634816/CCUD8OD_K.gif)

## BASH Script
```bash
#!/bin/bash

echo -e "\n NumberJack \n"
ch=0
while [ $ch -ne 3 ];
do
	echo  "  
		 PLAY : Hit 1 and enter.
		 HELP : Hit 2 and enter.
		 EXIT : Hit 3 and enter.
		 "

	read -p "Enter your choice : " ch
	if [ $ch -eq 1 ];then
	x=0
	c=0
	p=0
	read -p "Enter any number between 0 and 9 : " n
	while [ $c -eq 0 ];
	do
		x=11
		r=($(shuf -i 0-9 -n 10))
		echo "${r[@]} "
		for i in {1..10};
		do
			a[$i]=$i	
		done
		echo "${a[@]} "
		read -t 5 -p "Enter the index of your number : " x
		if [[ $? -gt 128 ]]; then 
			c=1
			break
		fi
		if [ ${r[$(($x))-1]} -eq $n ];then
			echo "Great"
			((p=p+1))
		else
			c=1
			break
		fi
	done
	elif [ $ch -eq 2 ];then
		echo "HELP: INSTRUCTIONS TO PLAY THE GAME. "
	else
		break
fi

if [ $c -eq 1 ];then
			echo -e "\nGAME OVER\n"
			echo "You scored $p points"
fi
		done

```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625753738352/qBDF1PFQG.png)

This is the final bare-bones script without any help instructions just keeping the script simple. I hope you learned something from the game development in BASH. This is just a fun little project and a cool way of learning certain concepts in BASH such as loops, conditional statements, and arithmetic. Have FUN. Happy CODING :)
