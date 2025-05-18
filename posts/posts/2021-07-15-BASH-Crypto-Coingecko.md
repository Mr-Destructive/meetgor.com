{
  "type": "posts",
  "title": "Cryptocurrency Price Scrapping using BASH and Coingecko API",
  "subtitle": "Know the price of your favorite cryptocurrency coin/token in your national currency from your terminal using bash shell and Coingecko crypto API.",
  "date": "2021-07-15 17:56:47 +0530",
  "status": "published",
  "tags": [
    "bash"
  ],
  "slug": "bash-crypto-scrapper",
  "image_url": "https://cdn.hashnode.com/res/hashnode/image/upload/v1626331633083/zTSpHCFOu.png?w=1600&h=840&fit=crop&crop=entropy&auto=compress"
}

## Introduction
Cryptocurrency is in such a hype that it is attracting even nerds and terminal fanboys, it is quite comfortable for such people to view the price inside of their terminal, also for normal people to learn about how to interact with an API from the local machine. We are gonna make a script about 20 lines in BASH to extract data from the  [coingecko cryptocurrency API](https://www.coingecko.com/en/api#explore-api) and some tools such as grep and sed. That being said let's start scripting.


![crypsh.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1626354050031/0BRlM7tfs.gif)

This is how the script will work like. It is not a fancy script just some basic commands to extract data from the API.

## Getting familiar with the API 
This is quite an important step as this will decide what type of data we will get from it.  So head out to  [coingecko cryptocurrency API](https://www.coingecko.com/en/api#explore-api) and start exploring some things yourselves. I am gonna cover everything you need for the script but just for your knowledge and if you are really into crypto. 

![coing.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626336264054/dFqhnBjFl.png)

Hopefully, you should see the following screen and after clicking on "Get Started", you will see a list of the API endpoints. After playing with it for several hours, I found the perfect and precise endpoint to achieve our target. Many o them work well but the problem with many of them was the precision of the price, there it could not show a reliable price for some low-valued coin/tokens. There are tons of options to choose from, but every option has its own advantage and disadvantage. 

Based on my tiny experience, the best fit for the required task will be the "Get historical market data to include price, market cap, and 24h volume (granularity auto)" option in the coin section. It will give the price with very nice precision and also in any national currency as well. So when you find the option, which looks like this:

![coing.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626338409604/TPxpUAJ91.png)

If you click on that option you will get a button to **"Try it out"**, just click on it and now you can parse data to the API and respond with the JSON response file like the text below as output. 

![coing.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626340029764/-hVg5t0iv.png)

### Filling in the coin id
You need the **id**which is the code of the cryptocurrency on coingecko so a few ids of famous cryptocurrencies are: `bitcoin`, `ethereum`, `dogecoin`, `shiba-inu`, `basic-attention-token`, and if you want more coins you can check the coingecko API coins list  [here](https://api.coingecko.com/api/v3/coins/list) and now you can search using the shortcut F3 and type the name of the coin you require. 

### Filling in the currency 
We need a currency to display the price of the selected coin in terms of a particular currency or its equivalent price in the currency. Every national currency has a code attached to it, for example, USD is the code for US Dollar, INR for Indian Rupee, EUR for Euro, AUD for Australian Dollar, and so on. You can find the list [here](https://www.iban.com/currency-codes).

### The number of days
We also need the number of days to indicate the price of the coin before the provided day, but we want today's and latest price, so we can enter the number of days as 0. You can get the price as your wish but we want the current price so we are using 0. 

So your request should look somewhat like this, just for particular coin pair:

![coing.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626340121389/YkRdtMm-o.png)

After filling in with your favorite coin-currency pair, you can now click on **"Execute"**. You will get a list of prices, volume, market cap, and all of that stuff. It should look somewhat like this:

![coing.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626340386948/QepcJHtUV.png)

If you look carefully and are aware of the current price, the second item in the prices object is the actual price of ethereum at some point in time. We need to extract that using tools and utilities in BASH. They have also provided a cURL command to fetch the API, how nice of Coingecko :)

So, that was the introduction to the Coingecko API, you can now explore various endpoints and what they respond to, and what are inputs they take. Many of them do not give precise enough price, so check out that as well. 


## Fetching data from the API 

Now, we'll finally start coding. Firstly we'll need input from the user, the coin-id, currency code, and optionally the number of days. We will read the input by using the read command and providing the prompt argument. 
```bash
#!/bin/bash

read -p "Enter the coin name: " coin
read -p "Enter your national currency: " crncy
days=0
#read -p "Enter the number of days before today to get its price: " days

```
I am not inputting the number of days before the equivalent price as it's not our mission, but if have commented on the part if you want to get the price before some days. The day value is hardcoded to 0 meaning the current day. 

Now we need the golden piece of this script, the cURL command. The command is just pre-written for us XD. But hey we need to modify it a bit. We need to store the output of the cURL command in a fie for modifying it further. We can use piping the command and avoid using files but that would make the command quite big. So, I just use the files to handle and trim the output according to what I need. 

This is the default command that coingecko gave us:
```bash
#!/bin/bash
curl -X 'GET' \
  'https://api.coingecko.com/api/v3/coins/ethereum/market_chart?vs_currency=usd&days=0' \
  -H 'accept: application/json'
```
We send a GET request to the API endpoint with the provide URL and we accept the response in the form of JSON. The -X is for providing the option of either GET, POST, or any other API interaction code. But this is a REST API and we do not want to mess up the database by posting, deleting, or updating anything just to GET the data from the API. We use -H to pass the URL to the head and the type of response in this case JSON. 

We'll change the default command to this, 
```bash
curl -o temp.json -X 'GET' \
  'https://api.coingecko.com/api/v3/coins/'$coin'/market_chart?vs_currency='$crncy'&days='$days'' \
  -H 'accept: application/json' &> /dev/null
```
We need to modify the URL a bit to make it dynamic. The ethereum or any coin name should be dynamic, the currency should be dynamic, so we will use the variables created before to use now. We will use `'$variable'` in between the URL to embed the variable value in it. We change the ethereum or any coin name with `'$coin'` and the currency name with `'$crncy'` and the same for the days as well. We have to store the output in the temp.json file, we use -o to output the result in the specified file in the cURL command. It's optional to add `&> /dev/null` because it just flushes the output of cURL, it looks neater if we add it. 

And that is it, we have obtained the JSON response and stored it in a file, we have the access to the data and we just need to edit the file.


## Editing the response JSON file

Now we start the actual editing and pattern finding in BASH. We have a plethora of tools to use like grep, sed, awk, and others, but I'll use only grep to keep things simple. Firstly we need to see the output( JSON file) again. 


![JSON file response](https://cdn.hashnode.com/res/hashnode/image/upload/v1626345122090/ERfxLQPX-.png)

But the file stored locally has everything in a single line, like this:


![coing.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626349631809/B0y8F9ruE.png)

So, we just need the numbers between `,` and `]],"market_caps"` right? 

OR

The text between `,` and `]],"m`.

It is quite simple to extract the required value, we will use grep and Pearl Regular Expression. 
```bash
grep -o -P '(?<=,).*(?=]],"m)' temp.json > price.txt

```
We are finding anything between the `,` and `]],"m` from the temp.json file and storing the output in the price.txt file.  As simple to use and we have the current price of the coin in terms of the provided currency in the file price.txt.

Now we have obtained the result in a crystal clear way, we need to store the price in a variable just for further usage and simplicity.
We'll use a while loop that iterates over the file until it is End of the File. We'll store the value in the variable in the following way:
```bash
while read val
do
	p=$val
done < price.txt
```
The value in the `p` variable. But we are not done yet, because if we see some values of certain coins which have quite low value, it displays in the scientific format. We'll tackle this in the next section.

## Converting the price from scientific notation to decimal

If you try to print the values of coins with pretty low value like `shiba-inu` or `baby-doge-coin` or any other coin with less value then a penny. The value is expressed in scientific notation i.e like `1.998e-5` i.e `0.00001998`

![coing.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626351434484/3GeTI-kn0.png)

This is not good-looking, is it? 
```bash
price=`printf "%.15f" $p`
```
The above command is quite similar to the C language. We are printing the value in the `p` variable with a precision of 15 decimal values, that is enough for any serious small value coin.

![coing.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626351561755/uPyHXMhd7.png)

Now that was much better. That is it! How simple was that. Really BASH has some powerful commands and tools.

## Printing the price 

Finally, we need to print the output, and to keep it simple, we can print using the echo command.
```bash
echo "The value of $coin in $crncy is = $price"
```

![coing.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626353654753/ajVMws8d6.png)


## Script
```bash
#!/bin/bash

read -p "Enter the coin name : " coin
read -p "Enter your national currency : " crncy
days=0
#read -p "Enter the number of days past today: " days
touch temp.json price.txt

curl -o temp.json -X 'GET' \
  'https://api.coingecko.com/api/v3/coins/'$coin'/market_chart?vs_currency='$crncy'&days='$days'' \
  -H 'accept: application/json' &> /dev/null

grep -o -P '(?<=,).*(?=]],"m)' temp.json > price.txt

while read val
do
	p=$val
done < price.txt

price=`printf "%.15f" $p`

echo "The value of $coin in $crncy is = $price"
rm temp.json

```

That was it, I hope you liked it and learned something along with me. BASH truly has the potential to do a whole lot of stuff. This was just one of the many examples where we can do some data gathering or scraping from the web using various tools and utilities. Thank you for listening to me. Happy Coding:) 
