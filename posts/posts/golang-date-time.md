{
  "title": "Golang: Date and Time",
  "status": "published",
  "slug": "golang-date-time",
  "date": "2025-04-05T12:33:32Z"
}

<h2>Introduction</h2>
<p>In the 28th post of the series, I will be exploring date and time handling in Golang. We will be covering the following topics:</p>
<ul>
<li>
<p>Date and Time parsing</p>
</li>
<li>
<p>Time Duration</p>
</li>
<li>
<p>Time and Date Arithmetic</p>
</li>
<li>
<p>Timezones</p>
</li>
<li>
<p>Sleep and Tickers</p>
</li>
</ul>
<p>This will cover most of the methods and properties used extensively in general use cases related to time and date operations.</p>
<h2>Time package</h2>
<p>The Golang standard library provides the time package to handle date and time-related operations. It has a lot of methods and constants to work and handle data related to time and dates.</p>
<p>One of the fundamental methods to get time in golang is the <a href="http://time.Now">time.Now</a> function, returns a time object representing the current time. The return value is a <a href="https://pkg.go.dev/time#Time">time</a> object, which is the base struct that we can use to perform operations on top of it.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main(){
    now := time.Now()
    fmt.Println(now)
    fmt.Printf(&quot;%T
&quot;, now)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2023-06-18 15:18:53.977607431 +0530 IST m=+0.000050291
time.Time
</code></pre>
<h3>Create a time/date type</h3>
<p>You can create a specific date by providing details like year, month, day, hour, minute, second, nanosecond, and time zone.</p>
<p>The <a href="http://time.Date">time.Date</a> method is used to create a time object. We will talk about location i.e. the timezone in a while.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main(){
    sometime := time.Date(2020, 03, 25, 8, 5, 0, 0, time.UTC)
    fmt.Println(sometime)
    fmt.Printf(&quot;%T
&quot;, sometime)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2020-03-25 08:05:00 +0000 UTC
time.Time
</code></pre>
<p>The above example generates a custom time object as we parsed the year, month, day, etc to the function. The time zone is set to UTC, the method might also take in a custom timezone, but that is what we will be exploring further.</p>
<p>We can also get some specific type from the time object using the methods that are available in the provided struct. For example, we can get the year from the time object by using the <a href="https://pkg.go.dev/time#Time.Year">time.Year()</a>, or month using <a href="https://pkg.go.dev/time#Time.Month">time.Month()</a> and so on</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main(){

	somedate := time.Date(2020, 03, 25, 8, 5, 0, 0, time.FixedZone(&quot;UTC&quot;, 5))
	fmt.Println(somedate)

	today := somedate.Day()
	fmt.Println(today)

	year := somedate.Year()
	fmt.Println(year)

	month := somedate.Month()
	fmt.Println(month)

	date := somedate.Day()
	fmt.Println(date)

	weekDay := somedate.Weekday()
	fmt.Println(weekDay)

    yearDay := somedate.YearDay()
    fmt.Println(yearDay)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2020-03-25 08:05:00 +0000 UTC
25
2020
March
25
Wednesday
85
</code></pre>
<p>In the above example, we have used the methods available in the time structure to get the specific components like a year, month, day, hour, etc. The data type for these methods is simple integer or strings as suitable to the format like int for date, year, and string for weekday, month, and timezone/location.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main(){

    now := time.Now()
	fmt.Println(now.Format(&quot;Monday 01 January 2006 15:04:05&quot;))

	// day month date hour:minutes:second timezone year
	fmt.Println(now.Format(time.UnixDate))

	// day month date hour:minutes:second year
	fmt.Println(now.Format(time.ANSIC))

	// day month date hour:minutes:second
	fmt.Println(now.Format(time.Stamp))

	// day month date hour:minutes:second.milisecond
	fmt.Println(now.Format(time.StampMilli))

	// day month date hour:minutes:second.microsecond
	fmt.Println(now.Format(time.StampMicro))

	// day month date hour:minutes:second.nanosecond
	fmt.Println(now.Format(time.StampNano))

	// day, date month year hour:minutes:second timezone
	fmt.Println(now.Format(time.RFC1123))

	// day, date month year hour:minutes:second offset
	fmt.Println(now.Format(time.RFC1123Z))

	// year-month-dayThour:minutes:second+-offset
	fmt.Println(now.Format(time.RFC3339))

	// year-month-dayThour:minutes.nanosecond:second
	fmt.Println(now.Format(time.RFC3339Nano))

	// date month year hour:minutes timezone
	fmt.Println(now.Format(time.RFC822))

	// hour:minuteAM/PM
	fmt.Println(now.Format(time.Kitchen))
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

Sunday 06 June 2023 20:09:09
Sun Jun 18 20:09:09 IST 2023
Sun Jun 18 20:09:09 2023
Jun 18 20:09:09
Jun 18 20:09:09.086
Jun 18 20:09:09.086565
Jun 18 20:09:09.086565975
Sun, 18 Jun 2023 20:09:09 IST
Sun, 18 Jun 2023 20:09:09 +0530
2023-06-18T20:09:09+05:30
2023-06-18T20:09:09.086565975+05:30
18 Jun 23 20:09 IST
8:09PM
</code></pre>
<p>These are some of the time formats that are provided by the time package in golang. They all return a string, so it cannot be parsed and resolved into components again. You can definitely take in a custom timestamp, convert it into a time object and then use the appropriate time format for your needs.</p>
<h2>Parsing time object from a string</h2>
<p>We can use the <a href="https://pkg.go.dev/time#Parse">Parse</a> function to parse a string into a time object. The method takes in two parameters, a date format, and the string to convert to both are parsed as a string. The method returns the parsed time or can give an error if the provided string is not in the mentioned format.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main(){

	customDate := &quot;2023-04-26&quot;
	t, err := time.Parse(&quot;2006-01-02&quot;, customDate)
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(customDate)
	fmt.Println(t)

    customDate = &quot;2023-0426&quot;
    t, err = time.Parse(&quot;2006-01-02&quot;, customDate)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(customDate)
    fmt.Println(t)

}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2023-04-26
2023-04-26 00:00:00 +0000 UTC

2023-0426
parsing time &quot;2023-0426&quot; as &quot;2006-01-02&quot;: cannot parse &quot;26&quot; as &quot;-&quot;                                                          
0001-01-01 00:00:00 +0000 UTC
</code></pre>
<p>In the above examples, we have parsed a time object from a string that looks like a date. We used the Parse method that takes in a string as the format to parse from and the string that we want to convert to. The format remains some specific date value like the <code>2006-01-02</code> which is fixed as a time to parse time in the time package. The method returns a time object which is parsed in the format, also it can return an error if the string is not in the provided format. We have used the second example that parsed a string as an invalid date format.</p>
<h2>Time Duration field</h2>
<p>The time duration field is used to represent the elapsed time between two Time objects as an integer 64-bit (in hours at the biggest unit and nanoseconds as the smallest unit to scale).</p>
<h3>Parse duration from string</h3>
<p>We can use the <a href="https://pkg.go.dev/time#ParseDuration">ParseDuration</a> method to parse a duration like string into a time.Duration object. The duration object can have serialized fields like Hours, Minutes, and so on for further usage into the time processing.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;time&quot;
)

func main() {
	screentime, err := time.ParseDuration(&quot;6h30m&quot;)
	if err != nil {
		fmt.Println(err)
		return
	}
    fmt.Printf(&quot;%T
&quot;, screentime)
	fmt.Println(screentime.Hours())
	fmt.Println(screentime.Minutes())

}
</code></pre>
<pre><code class="language-bash">$ go run main.go

time.Duration
6.5
390
</code></pre>
<p>In this example, the time duration is 6 hours 30 minutes as indicated by <code>6h30m</code>. This is parsed as a time duration, the duration can accept formats like &quot;ns&quot;, &quot;us&quot; (or &quot;µs&quot;), &quot;ms&quot;, &quot;s&quot;, &quot;m&quot;, &quot;h&quot;. We have used the <code>h</code> and <code>m</code> as short forms of hours and minutes respectively. Similarly, <code>s</code> can be used for seconds, <code>ms</code> for milliseconds, <code>us</code> or <code>µs</code> for microseconds, and <code>ns</code> for nanoseconds respectively. If you use another format it will result in errors, and yes! you will have to write err != nil { handle the error } syntax and debug the issue in the console (just saying)</p>
<p>The <a href="https://pkg.go.dev/time#ParseDuration">ParseDuration</a> method will return the <a href="https://pkg.go.dev/time#Duration">Duration</a> object or will error out if the string is not in the required short duration formats. In this example, the duration is stored in the <code>screentime</code> variable, which has a few methods like <code>Hours</code>, <code>Minutes</code>, <code>Seconds</code>, and so on to extract the component time duration in that object. So, if we use <code>screentime.Hours()</code>, this will give us the total hours in that parsed duration, in this case, it is <code>6.5</code> hours.</p>
<p>The return type is:</p>
<ul>
<li>
<p><code>float64</code> for <a href="https://pkg.go.dev/time#Duration.Hours">Hours</a>, <a href="https://pkg.go.dev/time#Duration.Minutes">Minutes</a>, and <a href="https://pkg.go.dev/time#Duration.Seconds">Seconds</a></p>
</li>
<li>
<p><code>int64</code> for <a href="https://pkg.go.dev/time#Duration.Milliseconds">Milliseconds</a>, <a href="https://pkg.go.dev/time#Duration.Microseconds">Microseconds</a>, and <a href="https://pkg.go.dev/time#Duration.Nanoseconds">Nanoseconds</a></p>
</li>
</ul>
<p>This can be useful in Linux command line applications, where we can get the duration of the application running or execution speed, etc. It can be used to get the approximate number of hours or other metrics specific to the needs.</p>
<p>The duration field is more useful for calculating the difference between two Time objects. Some methods like <a href="https://pkg.go.dev/time#Since">Since</a>, <a href="https://pkg.go.dev/time#Time.Sub">Sub</a>, are used to get the duration between the current time and other time objects and two Time objects respectively.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;time&quot;
)

func main() {

	newYear := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
    // current time is 2023-06-18 15:27:12 +0000 UTC
    fmt.Println(time.Now().UTC())
	fmt.Println(time.Since(newYear).Hours())

	nextNewYear := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(nextNewYear.Sub(newYear).Hours())
	fmt.Println(nextNewYear.Sub(newYear).String())

}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2023-06-18 15:27:12 +0000 UTC
4047.4130468335657
8760              
8760h0m0s
</code></pre>
<p>So as seen from the examples, the <code>Since</code> method calculates the difference between the current time and some other time object, and the <code>Sub</code> method calculates the difference between two Time objects. The current time is <code>18th June 2023</code> which is roughly near the middle of the year, so if we get the duration from the start of the year(1st January 2023), we get <code>4047.4</code> hours. Similarly, we can get a duration between a year, i.e. 1st January 2023 and 1st January 2024, which comes out to be <code>8760</code> hours. We can even use <code>Minutes</code>, <code>Seconds</code>, <code>Milliseconds</code>, etc. to get the duration in those units.</p>
<pre><code class="language-go">day := time.Hour * 24
fmt.Println(day)
week := time.Hour * 24 * 7
fmt.Println(week)
month := time.Hour * 30 * 24 * 7
fmt.Println(month)

fifteenDays := day * 15
fmt.Println(fifteenDays)
</code></pre>
<pre><code class="language-bash">$ go run main.go

24h0m0s
168h0m0s
5040h0m0s
360h0m0s
</code></pre>
<p>Here, we have created some custom units for calculating duration like a day which is 24 hours, that is obtained by multiplying an hour i.e. <code>time.Hour</code> by 24, and a week is 7 * day, hence we do the calculation with precision in the duration structure. This can also be applied to <code>Minutes</code>, <code>Seconds</code>, and so on.</p>
<h2>Time Zones and Locations</h2>
<p>Time zones are associated with every time object in the time structure. They are represented as <a href="https://pkg.go.dev/time#Location">Locations</a> structure.</p>
<p><a href="https://en.wikipedia.org/wiki/List_of_tz_database_time_zones">List of time zones</a></p>
<p>If not specified, the <a href="http://time.Now"><code>time.Now</code></a><code>()</code> set the timezone as the Local time zone which is picked up from the system. It stores a list of timezone if the location uses daylight savings time, it picks the first timezone if daylight savings is not present.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;time&quot;
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Location())

	newYorkTimeZone, err := time.LoadLocation(&quot;America/New_York&quot;)
	if err != nil {
		fmt.Println(err)
	}
	londonTimeZone, err := time.LoadLocation(&quot;Europe/London&quot;)
	if err != nil {
		fmt.Println(err)
	}
	newYorkTime := t.In(newYorkTimeZone)
	londonTime := t.In(londonTimeZone)

	//local time
	fmt.Println(t)

	// london time
    fmt.Println(londonTimeZone)
	fmt.Println(londonTime)

	// new york time
    fmt.Println(newYorkTimeZone)
	fmt.Println(newYorkTime)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2023-06-19 13:58:13.232181805 +0530 IST m=+0.000044899
Local
2023-06-19 13:58:13.232181805 +0530 IST m=+0.000044899

Europe/London
2023-06-19 09:28:13.232181805 +0100 BST

America/New_York
2023-06-19 04:28:13.232181805 -0400 EDT
</code></pre>
<p>The <a href="https://pkg.go.dev/time#LoadLocation">LoadLocation</a> method is used to parse the timezone name as a string and returns the timezone/location object. It can return a <code>Location</code> object or an error in case the provided timezone name is not valid.</p>
<p>We can use the <a href="http://time.In">time.In</a><a href="https://pkg.go.dev/time#Time.In">(timezone)</a> to get the time in the specified timezone. In the above example, we have created a timezone/location object with the location name of the timezone as <code>Europe/London</code>. The timezone/location object has no exported/public properties or methods except the <code>String()</code> method used to get the string representation of the time stamp. However, this timezone is used to get time stamps in another timezone, for instance in this example, we have used my local timezone, i.e. <code>IST</code> or <code>Aisa/Kolkata</code> to get the current time, and used <code>In()</code> method to get time in <code>Europe/London</code> timezone. The <code>In()</code> method takes in a location object as a parameter and returns the <code>time.Time</code> object as the time in that provided location.</p>
<h3>Creating a custom time zone</h3>
<p>A location object in the time package is basically a name and an offset value, so we can construct our own custom time zone location object with those parameters with the <a href="https://pkg.go.dev/time#FixedZone">FixedZone</a> method. The <code>FixedZone</code> method takes in two parameters, one as the name of the timezone as a string, and the other as the offset as int. The offset can be up to 290 years roughly, as it can hold 64 bits and this can only represent an integer value as time only this far.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main() {

    t := time.Now()
    fmt.Println(t)

	offset := int((4*time.Hour + 30*time.Minute + 30*time.Second).Seconds())
	// (4*60 + 30*1 + 30*0.166) * 60
	// (270 + 0.5) * 60 = 16230
	fmt.Println(offset)
	fmt.Println(t.UTC())
	customTimeZone := time.FixedZone(&quot;SOMETZ&quot;, offset)
	fmt.Println(customTimeZone)
	fmt.Println(t.In(customTimeZone))

}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2023-06-19 15:25:44.8624019 +0530 IST m=+0.000044899
16230
2023-06-19 09:55:44.8624019 +0000 UTC
SOMETZ
2023-06-19 14:26:14.8624019 +0430 SOMETZ
</code></pre>
<p>So from the example, the offset is 16230 seconds, which is then constructed by adding hours, minutes, and seconds, this is just an example, it shows you can customize the hour, minute, and seconds components, but you can extend it with millisecond up to nanosecond but that is a very niche case and might not be broadly used. To get an offset of <code>4h3030s</code>, we have used arithmetic to get the number of hours, minutes, and seconds.</p>
<h2>Arithmetic on Time</h2>
<p>We can perform addition(going ahead) and subtraction(going back) of time in the <code>time.Time</code> object. There are also some comparison operations that can be performed using the method provided in the time struct.</p>
<h3>Add Time and Date</h3>
<p>We can add a date to the <code>time.Time</code> object using the <a href="https://pkg.go.dev/time#Time.Add">Add</a> method, this method takes in an integer value of <code>Hours</code>, <code>Minute</code>, <code>Second</code>, and so on to add the value of those units to the existing time object as a copy. This method can be chained to add multiple units of different types for instance hours, minutes, or seconds in a precise way to add the time. There is also a <a href="https://pkg.go.dev/time#Time.AddDate">AddDate</a> method which takes in 3 parameters as a number of years, months, and days. We can add <code>x</code> years to the existing date by passing <code>AddDate(x, 0, 0)</code> to the time object we want to add to. These two method can be chained together to go from a scale of year to all the way to nanoseconds precision in adding the time.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;time&quot;
)

func main() {
	t := time.Now()
	fmt.Println(t)

	afterOneHour := t.Add(time.Hour * 1)
	fmt.Println(afterOneHour)

	afterOneDayTwoHours30minutes := t.AddDate(0, 0, 1).Add(time.Hour * 2).Add(time.Minute * 30)
	fmt.Println(afterOneDayTwoHours30minutes)

	afterThreeMonths15days := t.AddDate(0, 3, 15)
	fmt.Println(afterThreeMonths15days)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2023-06-19 15:58:32.893246798 +0530 IST m=+0.000042833
2023-06-19 16:58:32.893246798 +0530 IST m=+3600.000042833
2023-06-20 18:28:32.893246798 +0530 IST
2023-10-04 15:58:32.893246798 +0530 IST
</code></pre>
<p>Here, we can see, we have taken a simple example, for adding an hour to the existing time by saying <code>t.Add(time.Hour)</code>, this can be multiplied by the number of hours to add to, in this case, it was just one so we simply multiply with one. In the next example, we have chained the <code>AddDate</code> and <code>Add</code> methods to get the time after 1 day, 2 hours, and 30 minutes. The <code>AddDate</code> method is passed with <code>(0, 0, 1)</code> indicating a single day, then we tune the hours with 2 and minutes with 30 to get the desired time.</p>
<p>In the last example, we have added the time by 3 months and 15 days, bypassing the <code>AddDate</code> method with <code>(0, 3, 15)</code>.</p>
<p>Time can even be subtracted or we can get behind the specified time object using the negative number in the <code>Add</code> method. Instead of saying <code>.Add(1 * time.Hour)</code> to go one hour ahead of the parsed time, we can say <code>.Add(-1 * time.Hour)</code> to go one hour behind the current time.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main() {
    t := time.Now()
    fmt.Println(t)

    oneHourBack := t.Add(-1 * time.Hour)
    fmt.Println(oneHourBack)

    beforeOneYearTwoMonths := t.AddDate(-1, -2, 0)
    fmt.Println(beforeOneYearTwoMonths)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2023-06-19 16:23:26.992724868 +0530 IST m=+0.000044899
2023-06-19 15:23:26.992724868 +0530 IST m=-3599.999957644
2022-04-19 16:23:26.992724868 +0530 IST
</code></pre>
<p>In this example, we have used the <code>Add</code> and <code>AddDate</code> methods with negative numbers to go back in time. In the first example, we have subtracted 1 hour from the existing time object, in the second example, we have used -1 to go back a year and -2 for going back 2 months from the current time object, hence we pass <code>.AddDate(-1, -2, 0)</code> in the example.</p>
<h3>Comparing Time</h3>
<p>We can use <code>After</code>, <code>Before</code>, <code>Equal</code>, or <code>Compare</code> methods to compare the <code>Time</code> object in Golang. The method returns true or false depending if the time comes after, before, or is equal, the compare method returns -1 if it is before, 0 if two times are equal, and 1 if the time is after the other.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main() {
    t := time.Now()
    fmt.Println(t)
    
    afterOneHour := t.Add(time.Hour * 1)
    fmt.Println(afterOneHour)

	isNowAfter := t.After(afterOneHour)
	isOneAfter := afterOneHour.After(t)
	fmt.Println(isNowAfter)
	fmt.Println(isOneAfter)

	isNowBefore := t.Before(afterOneHour)
	isOneBefore := afterOneHour.Before(t)
	fmt.Println(isNowBefore)
	fmt.Println(isOneBefore)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2023-06-19 16:38:25.785629649 +0530 IST m=+0.000051616
2023-06-19 17:38:25.785629649 +0530 IST m=+3600.000051616

// is now after one hour?
false

// is one hour after now?
true 

// is now before after one hour?
true 

// is after one hour before now?
false
</code></pre>
<p>These are some ridiculous examples, but it could be any date comparison like billing periods, subscription due, etc. We can compare dates with <a href="https://pkg.go.dev/time#Time.After">After</a>, <a href="https://pkg.go.dev/time#Time.Before">Before</a>, and <a href="https://pkg.go.dev/time#Time.Equal">Equal</a> methods provided in the time package.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main() {
    t := time.Now()
    fmt.Println(t)

    afterOneHour := t.Add(time.Hour * 1)
    fmt.Println(afterOneHour)

	isNowEqual := t.Equal(afterOneHour)
	isEqual := time.Now().Equal(t)
	fmt.Println(isNowEqual)
	fmt.Println(isEqual)

	londonTimeZone, err := time.LoadLocation(&quot;Europe/London&quot;)
	if err != nil {
		fmt.Println(err)
	}
	londonTime := t.In(londonTimeZone)

	fmt.Println(t)
	fmt.Println(londonTime)
	fmt.Println(t.Equal(londonTime))
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

2023-06-19 16:38:25.785629649 +0530 IST m=+0.000051616
2023-06-19 17:38:25.785629649 +0530 IST m=+3600.000051616

// is now equal to one hour after now?
false

// is t equal to now?
false

2023-06-19 16:49:25.509421027 +0530 IST m=+0.000040200
2023-06-19 12:19:25.509421027 +0100 BST

// is london time equal to now in local?
true
</code></pre>
<p>Here in the examples, we have checked if two Time objects are equal or not, we have checked if a current time is equal to one hour after time, and it returns false rightly so. In the next example, we have created a new time object as <a href="http://time.Now"><code>time.Now</code></a>, and then compared with the previous now, there will be a difference in the seconds of the initialization and hence it gives false as the return value.</p>
<p>In the next example, we take the current time and the current time in the London location as the <code>BST</code> timezone, this gets the result as <code>true</code> because the time is the same despite being in different time zones.</p>
<h2>Sleep Time</h2>
<p>We can even sleep for a certain duration of time, with the <a href="https://pkg.go.dev/time#Sleep">Sleep</a> method in the time package, this results in halting the execution of the program.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main() {
    t1 := time.Now()
    time.Sleep(time.Second * 3)
    t2 := time.Now()
    duration := t2.Sub(t1)
    fmt.Println(duration)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go
3.000235856s
</code></pre>
<p>In this example, we have used the <code>Sleep()</code> method with a duration of 3 seconds which will halt the execution of the program (or the goroutine) for that duration. We also printed the duration for which the flow was halted by taking a timestamp before and after the sleep method was called and then taking the difference between the two-time stamps. This rightly gives us a value of three seconds which is the duration for which we tried to call the <code>Sleep</code> method.</p>
<h2>Tickers</h2>
<p>Tickers are basically like a clock that ticks at regular intervals. It's a mechanism provided by the package to execute code repeatedly at fixed time intervals. The <a href="https://pkg.go.dev/time#Ticker">Ticker</a> is a type and it has a few methods associated with it for handling these ticks. The <a href="https://pkg.go.dev/time#NewTicker">NewTicker</a> method is used to create a ticker with the parameter as the amount of time to repeat the tick i.e. the frequency of its ticking. It will only stop when the <a href="https://pkg.go.dev/time#Ticker.Stop">.Stop</a> method is called with that ticker object.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func main() {
	ticker := time.NewTicker(time.Second * 2)
	counter := 0
	for {
		select {
		case &lt;-ticker.C:
			// api calls, call to database after specific time intervals
			counter++
			fmt.Println(&quot;Tick&quot;, counter)
		}
		if counter == 5 {
			ticker.Stop()
			return
		}
	}
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

Tick 1
Tick 2
Tick 3
Tick 4
Tick 5
</code></pre>
<p>In the example above, we have created a ticker with a frequency of 2 seconds, we initialize a counter to 0, we run an infinite condition for loop, and the ticker receives a value when the time is for ticking via the channel <code>.C</code>, this is used to check if it has ticked, and we enter the flow for the execution of any logic/code. We internally increment the counter and when the counter is 5 we stop the ticker and break out of the loop with either a <code>break</code> or <code>return</code>.</p>
<p>So, that is how tickers can be used, this can be used in sending requests to APIs that have rate limiting, so we make sure we don't flood the API and maintain the rate of hitting the API only after specific intervals with tickers.</p>
<p>That's it from the 28th part of the series, all the source code for the examples are linked in the GitHub on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/date-time">100 days of Golang</a> repository.</p>
<h2>Conclusion</h2>
<p>So, in conclusion, we have seen how to use the time package in Go. We covered time and date parsing, timezone, date comparison, sleep, and tickers. This would have given a good overview of the time package in golang.</p>
<p>If you have any questions, feedback, or suggestions feel free to drop them in the comments section, or on the social handles. Thank you for reading. Happy Coding :)</p>
