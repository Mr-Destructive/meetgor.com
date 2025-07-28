{
  "title": "Golang: Paths",
  "status": "published",
  "slug": "golang-paths",
  "date": "2025-04-05T12:33:35Z"
}

<h2>Introduction</h2>
<p>In the 21st post of the series, we will be exploring the file paths in golang, we will be exploring how we can deal with paths. By using packages like <code>os</code>, <code>path</code>, <code>io</code>, we can work with file systems and operating system-specific details. In this section, we will see how to resolve paths, details from paths, extract relative or absolute paths, iterate over file systems, etc.</p>
<p>Starting from this post, it will follow a specific topic in the upcoming few posts which will be covering files and paths. We will be talking about dealing with paths and files in golang. This post is just about working with paths.</p>
<h2>Resolving and Parsing Path</h2>
<p>In golang, the <a href="https://pkg.go.dev/os">os</a> and the <a href="https://pkg.go.dev/path">path</a> packages are quite helpful in working with paths. We use the <code>pathilpath</code> package specifically for working with paths and file structures.</p>
<h3>Get the current working directory</h3>
<p>To get the path for the current working directory, we can use the <a href="https://pkg.go.dev/os#Getwd">os.Getwd()</a> function. The function returns a-ok, an error-like object if the working directory exists it will return the absolute path to the directory else if the path is deleted or corrupted while processing, it will give an error object.</p>
<pre><code class="language-go">package main

import(
    &quot;os&quot;
    &quot;log&quot;
)

func main() {

    dir, err := os.Getwd()
    if err != nil {
        log.Println(err)
    } else {
        log.Println(dir)
    }
}
</code></pre>
<pre><code>$ pwd
/home/meet/code/techstructive-blog

$ go run main.go
2022/10/01 19:19:09 /home/meet/code/techstructive-blog
</code></pre>
<p>So, as we can see the <code>Getwd</code> the function returns an absolute path to the current working directory which will be the path from which you will be executing/running the script file.</p>
<h3>Get the path to the home directory</h3>
<p>We can even get the home directory path like the <code>/home</code> followed by the user name on Linux and the User Profile with the name for Windows. The <a href="https://pkg.go.dev/os#UserHomeDir">UserHomeDir()</a>, returns the home directory for the user from which the file is being executed. The return value is simply an string just like the <code>Getwd</code> function.</p>
<pre><code class="language-go">package main

import(
    &quot;os&quot;
    &quot;log&quot;

)

func main() {
    dir, err := os.UserHomeDir()
    if err != nil {
        log.Println(err)
    } else {
        log.Println(dir)
    }
}
</code></pre>
<pre><code>$ echo $HOME
/home/meet/

$ go run main.go
2022/10/01 19:35:50 /home/meet
</code></pre>
<p>So, as expected, the <code>UserHomeDir</code> function returns the path string to the home directory of the user.</p>
<h3>Get path from a file name string</h3>
<p>Let's say, we give in a filename and we want the absolute path of it. The <a href="https://pkg.go.dev/path/filepath">path/filepath</a> package provides the <a href="https://pkg.go.dev/path/filepath#Abs">Abs</a> function that does exactly that. The function returns a path string of the parameter parsed as a string to a directory or a file name. The function might as well return an error as the file path might not existing or the file might have got deleted, so we'll have to call the function with the ok, error syntax.</p>
<pre><code class="language-go">package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
)

func main() {

    file_name := &quot;default.md&quot;
    log.Println(file_name)
    dir, err := filepath.Abs(file_name)
    if err != nil {
        log.Println(err)
    } else {
        log.Println(dir)
    }
}

</code></pre>
<pre><code>$ go run main.go

2022/10/01 19:52:23 default.md
2022/10/01 19:52:23 /home/meet/code/techstructive-blog/default.md

</code></pre>
<p>As we can see the file <code>default.md</code> was parsed in the <code>Abs()</code> function and it returned the absolute path of the file.</p>
<h3>Get Parent Directory from a Path</h3>
<p>We can get the parent directory for a given path, if the path is to a file, we return the absolute path to the parent directory of the file, or if the path is to a folder, we return the folder's parent directory.</p>
<pre><code class="language-go">
package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
)

func main() {
    file_name := &quot;drafts/default.md&quot;
    //file_name := &quot;drafts/&quot;
    path, err := filepath.Abs(file_name)
    if err != nil {
        log.Println(err)
    } else {
        //log.Println(path)
        log.Println(filepath.Dir(path))
    }
}
</code></pre>
<pre><code>$ go run main.go
2022/10/01 19:58:45 /home/meet/code/techstructive-blog/drafts

$ go run main.go
2022/10/01 19:58:45 /home/meet/code/techstructive-blog 

</code></pre>
<p>As we can see when we parse in a file path i.e. <code>drafts/default.md</code>, the <code>Dir</code> the method returns a path to the parent folder, and even if we parse the directory path i.e. <code>drafts/</code>, the method returns the parent of that directory.</p>
<h3>Get the last file/folder for a given Absolute Path</h3>
<p>Golang also provides a way to get the file/directory name from a path string using the <a href="https://pkg.go.dev/path/filepath#Base">Base</a> function provided in the <a href="https://pkg.go.dev/path/filepath">path/filepath</a> package.</p>
<pre><code class="language-go">file_name := &quot;default.md&quot;
dir, err := filepath.Abs(file_name)

if err != nil {
    log.Println(err)
} else {
    log.Println(dir)
    log.Println(filepath.Base(dir))
}
</code></pre>
<pre><code>$ go run main.go

2022/10/01 19:58:45 /home/meet/code/techstructive-blog/drafts/default.md
2022/10/01 20:19:28 default.md 
</code></pre>
<p>So, the function <code>Base</code> will return the last element in the path, it can be a file or a directory, just returns the name before the last <code>\</code>. In the above example, we start with a filename <code>default.md</code> but set the dir as the absolute path to that file and again grab the file name using the <code>Base</code> function.</p>
<h2>Fetching details from a Path</h2>
<p>We can even use utility functions for dealing with paths in golang like for checking if a file or path exists, if a path is a file or a directory, grabbing file name and extension, etc. The <code>path/filepath</code> and the <code>os</code> the package helps with working with these kinds of operations.</p>
<h3>Check if a path exists</h3>
<p>We can use the <a href="https://pkg.go.dev/os#Stat">os.Stat</a> function along with the <a href="https://pkg.go.dev/os#IsNotExist">os.IsNotExist</a> for finding if a path is existing or not. The Stat function returns a <a href="https://pkg.go.dev/io/fs#FileInfo">FileInfo</a> object or an error. The <code>FileInfo</code> object will have methods such as <code>Name()</code>, <code>IsDir()</code>, <code>Size()</code>, etc. If we get an error, inside the Stat method, the error will probably arise if the path does not exist, so inside the <code>os</code> package, we also have the <code>IsNotExist()</code> method, that returns a <code>boolean</code> value. The method returns <code>true</code> if the parsed error indicates that the path doesn't exist and <code>false</code> if it exists.</p>
<pre><code class="language-go">package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
    &quot;os&quot;
)

func main() {

    file_name := &quot;drafts/default.md&quot;
    path, err := filepath.Abs(file_name)
    if err != nil {
        log.Println(err)
    } else {
        if _, err := os.Stat(path); os.IsNotExist(err) {
            log.Println(&quot;No, &quot; + path + &quot; does not exists&quot;)
        } else {
            log.Println(&quot;Yes, &quot; + path + &quot; exists&quot;)
        }
    }
}
</code></pre>
<pre><code>$ go run main.go

2022/10/01 20:51:31 Yes, /home/meet/code/techstructive-blog/drafts/default.md exists
</code></pre>
<p>So, from the above example, the program will log if the path is present in the system or not. The error is parsed from the <code>Stat</code> method to the <code>IsNotExist</code> method for logging relevant messages. Since the directory exists, we get the path exists log.</p>
<h3>Check if a path is a file or directory</h3>
<p>The <code>FileInfo</code> object returned from the <code>Stat</code> the method provides a few methods such as <code>IsDir()</code> that can be used for detecting if a given path is a directory or not. The function simply returns a <code>boolean</code> value if the provided path points to a directory or not. Since we have to parse the path to the <code>IsDir()</code> function, we convert the file string into a path using the <code>Abs</code> method and then check if the path actually exist with the <code>Stat()</code> method.</p>
<pre><code class="language-go">package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
    &quot;os&quot;
)

func main() {

    file_name := &quot;drafts/default.md&quot;
    //file_name := &quot;drafts/&quot;
    path, err := filepath.Abs(file_name)
    if err != nil {
            log.Println(err)
    } else {
        if t, err := os.Stat(path); os.IsNotExist(err) {
            log.Fatal(&quot;No, &quot; + path + &quot; does not exists&quot;)
        } else {
            log.Println(path)
            log.Println(t.IsDir())
        }
    }
}
</code></pre>
<pre><code>$ go run main.go
2022/10/01 20:55:20 /home/meet/code/techstructive-blog/drafts/default.md
2022/10/01 20:55:20 false

$ go run main.go
2022/10/01 20:55:20 /home/meet/code/techstructive-blog/drafts/
2022/10/01 20:55:20 true
</code></pre>
<p>So, by running the program for a file and a directory, we can see it returns <code>true</code> if the path is a directory and <code>false</code> if the provided path is a file. In the above example, since the <code>drafts/defaults.md</code> is a file, it returned <code>false</code>, and for the next example, when we set the path  <code>drafts/</code> it returns <code>true</code> as the path provided is a directory.</p>
<h3>Get File Extension from path</h3>
<p>By using the <a href="https://pkg.go.dev/path">path</a> package, the extension of a given path can be fetched. The <a href="https://pkg.go.dev/path#Ext">Ext</a> method can be used for getting the extension of the provided path string, it doesn't matter if the provided path is exists or not, is absolute or relative, it just returns the text after the last . in the string. But if we are working with real systems it is good practice to check if the file or path actually exists.</p>
<pre><code class="language-go">package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
    &quot;path&quot;
)

func main() {

    file_name := &quot;default.md&quot;
    dir, err := filepath.Abs(file_name)
    if err != nil {
        log.Println(err)
    } else {
        file_ext := path.Ext(dir)
        log.Println(file_ext)
    }
}
</code></pre>
<pre><code>$ go run main.go
2022/10/01 21:03:23 .md
</code></pre>
<p>The above example demonstrates how we can get the extension of a file using the <code>Ext()</code> method in the <code>path</code> package. Given the string path as <code>default.md</code>, the function returned <code>.md</code> which is indeed the extension of the provided file.</p>
<h3>Get Filename from path</h3>
<p>We can even get the file name from a path in golang using the <a href="https://pkg.go.dev/strings#TrimSuffix">TrimSuffix</a> method in the <a href="https://pkg.go.dev/strings">strings</a> package. The <code>TrimSuffix</code> method trim the string from the provided suffix, like if we have a string <code>helloworld</code> and we provide the suffix as <code>world</code>, the <code>TrimSuffix</code> the method will return the string <code>hello</code>, it will remove the suffix string from the end of the string.</p>
<pre><code class="language-go">package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
    &quot;path&quot;
    &quot;strings&quot;
)

func main() {

    file_name := &quot;default.md&quot;
    dir, err := filepath.Abs(file_name)
    if err != nil {
        log.Println(err)
    } else {
        file_ext := path.Ext(dir)
        log.Println(file_ext)
        log.Println(strings.TrimSuffix(dir, file_ext))
        log.Println(strings.TrimSuffix(file_name, file_ext))
        //log.Println(strings.TrimSuffix(dir, path.Ext(dir)))
        //log.Println(strings.TrimSuffix(file_name, path.Ext(dir)))
    }
}
</code></pre>
<pre><code>$ go run main.go

2022/10/01 21:09:39 .md
2022/10/01 21:09:39 /home/meet/code/techstructive-blog/default
2022/10/01 21:09:39 default
</code></pre>
<p>We can use the <code>TrimSuffix</code> method to remove the extension as the suffix and it returns the path which we get as the file name. The <code>TrimSuffix</code> method returns the path after removing the extension from the path.</p>
<h2>List Files and Directories in Path</h2>
<p>In golang, we can use the <code>io</code> and the <code>path/filepath</code> packages to iterate over the file paths. Suppose, we want to list out all the files or directories in a given path, we can use certain functions such as <code>Walk</code>, <code>WalkDir</code> to iterate over a path string.</p>
<p>There are certain types of iterations we can perform based on the constraints we might have, like iterating over only files, or directories, not including nested directories, etc. We'll explore the basic iterations and explain how we fine-tune the iteration based on the constraints.</p>
<h2>List only the files in the Path</h2>
<p>The first example, we can take is to simply list out only the files in the current path directory, we don't want to list out the file in nested directories. So, it will be like a simple ls command in Linux. Let's see how we can list out the files in the given path.</p>
<p>We can even use <code>path/filepath</code> package to iterate over a given path and list out the directories and files in it. The <a href="https://pkg.go.dev/path/filepath#Walk">filepath.Walk</a> or the <a href="https://pkg.go.dev/path/filepath#WalkDir">WalkDir</a> method is quite useful for this kind of operation, the function takes in a path string and a <a href="https://pkg.go.dev/path/filepath#WalkFunc">WalkFunc</a> or the <a href="https://pkg.go.dev/io/fs#WalkDirFunc">WalkDirFunc</a> Function, the walk function are simply used for walking of a path string. Both functions take two parameters, the first being the string which will be the file system path where we want to iterate or walk, and the next parameter is the function either <a href="https://pkg.go.dev/path/filepath@go1.19.1#WalkFunc">WalkFunc</a> or <a href="https://pkg.go.dev/io/fs#WalkDirFunc">WalkDirFun</a> respectively. Both functions are similar but a subtle difference in the type of parameter both take in.</p>
<h3>WalkDir Function</h3>
<p>The <code>WalkDir</code> function takes in the parameters such as a <code>string</code> of the file path, the <a href="https://pkg.go.dev/io/fs#DirEntry">fs.DirEntry</a> object and the <code>error</code> if any. The function returns an <code>error</code> if there arises any. We have to call the function with the parameters of a string and a function object which will be of type <code>type WalkDirFunc func(path string, d DirEntry, err error) error</code>.</p>
<p>We can even use Walk the function to iterate over the given path.</p>
<h3>Walk Function</h3>
<p>The <code>Walk</code> function takes in the parameters such as a <code>string</code> of the file path, the <a href="https://pkg.go.dev/io/fs#FileInfo">fs.FileInfo</a> object and the <code>error</code> if any. The function returns an <code>error</code> if there arises any. We have to call the function with the parameters of a string and a function object which will be of type <code>type WalkFunc func(path string, info fs.FileInfo, err error) error</code>.</p>
<p>It might be a user preference to select one of the functions for iterating over the file system, but the <a href="https://pkg.go.dev/path/filepath#Walk">documentation</a> says, the <code>Walk</code> function is a little bit inefficient compared to the <code>WalkDir</code> function. But if performance is not an issue, you can use either of those based on which type of file system object you are currently working with.</p>
<pre><code class="language-go">package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
    &quot;io/fs&quot;
)

func main() {

	var files []string
	dir_path := &quot;.&quot;
	err := filepath.WalkDir(dir_path, func(path string, info fs.DirEntry, err error) error {
		dir_name := filepath.Base(dir_path)
		if info.IsDir() == true &amp;&amp; info.Name() != dir_name{
			return filepath.SkipDir
		} else {
			files = append(files, path)
			return nil
		}
	})

	if err != nil {
		panic(err)
	}
	for _, file:= range files {
		log.Println(file)
	}
}
</code></pre>
<pre><code>$ go run walk.go

2022/10/02 12:07:17 .
2022/10/02 12:07:17 .dockerignore
2022/10/02 12:07:17 .gitignore
2022/10/02 12:07:17 CNAME
2022/10/02 12:07:17 Dockerfile
2022/10/02 12:07:17 README.md
2022/10/02 12:07:17 markata.toml
2022/10/02 12:07:17 requirements.txt
2022/10/02 12:07:17 textual.log
</code></pre>
<p>In the above example, we have used the <code>WalkDir</code> method for iterating over the file system, the directory is set as <code>.</code> indicating the current directory. We parse the first paramter as the string to the <code>WalkDir</code> function, the next parameter is a function so we can either create it separately or just define an <code>anonymous function</code>. It becomes a lot easier to write an <code>anonymous function</code> rather than writing the function separately.</p>
<p>So, we have created the <code>dir_name</code> variable which parses the <code>dir_path</code> from the parameter to the function and gets the name of the directory or file. We can then fine-tune the requirements of the iteration of the directory, i.e. make checks if the path is a file or a directory and if we want to exclude any specific files with certain extensions or directories with a certain name, etc. In this example, we have added a check if the path is a directory(using <code>info.IsDir()</code>) and if the directory name is not the same as the parsed path(i.e. exclude the nested directories) we skip these types of directories (using <a href="https://pkg.go.dev/io/fs#SkipDir">filepath.SkipDir</a>). So we only look for the files in the current directory or the directory which we provided in the paramter as <code>dir_path</code>. We append those paths into the files array using the <code>append</code> method. Finally, we check for errors in the parsed parameter while iterating over the file system and <code>panic</code> out of the function. We can then simply iterate over the files slice and print or perform operations as required.</p>
<h3>All the files in the Path (inside directories)</h3>
<p>We can also list all the files within the folders provided in the path string by removing the directory name check. We will only append the file type to the file slice rather than appending all the directories.</p>
<pre><code class="language-go">package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
    &quot;io/fs&quot;
)

func main() {

	var files []string
	root := &quot;static/&quot;
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if info.IsDir() {
			return nil
		} else {
			files = append(files, path)
			return nil
		}
	})

	if err != nil {
		panic(err)
	}

	for _, file:= range files {
		log.Println(file)
	}
}
</code></pre>
<pre><code>$ go run walk.go

2022/10/02 12:08:22 static/404.html
2022/10/02 12:08:22 static/CNAME
2022/10/02 12:08:22 static/index.html
2022/10/02 12:08:22 static/main.css
2022/10/02 12:08:22 static/projects/index.html
2022/10/02 12:08:22 static/social-icons.svg
2022/10/02 12:08:22 static/tbicon.png
</code></pre>
<p>As we can see the iteration resulted in printing all the files in the given path including the files in the subdirectories. The static directory had the projects directory as a subdirectory in the path, hence we are listing the files in that directory as well.</p>
<h3>Recursive directories in the Path</h3>
<p>We can also append the directory names as well as file names by completely removing the <code>info.IsDir()</code> check and add the printing out of the relevant information as dir and files depending on the type. We can also maintain different lists or slices for directory and file and append them accordingly.</p>
<pre><code class="language-go">package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
    &quot;io/fs&quot;
func main() {

	var files []string
	root := &quot;static/&quot;
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
        files = append(files, path)
		var f string
		if info.IsDir() {
			f = &quot;Directory&quot;
		} else {
			f = &quot;File&quot;
		}
		log.Printf(&quot;%s Name: %s
&quot;, f, info.Name())
        return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file:= range files {
		log.Println(file)
	}
}
</code></pre>
<pre><code>$ go run walk.go

2022/10/02 12:09:48 Directory Name: static
2022/10/02 12:09:48 File Name: 404.html
2022/10/02 12:09:48 File Name: main.css
2022/10/02 12:09:48 Directory Name: projects
2022/10/02 12:09:48 File Name: index.html
2022/10/02 12:09:48 File Name: social-icons.svg
2022/10/02 12:09:48 File Name: tbicon.png

2022/10/02 12:09:48 static/
2022/10/02 12:09:48 static/404.html
2022/10/02 12:09:48 static/index.html
2022/10/02 12:09:48 static/main.css
2022/10/02 12:09:48 static/projects
2022/10/02 12:09:48 static/projects/index.html
2022/10/02 12:09:48 static/social-icons.svg
2022/10/02 12:09:48 static/tbicon.png
</code></pre>
<p>We can see that the directories and files getting logged which are present in the given path. In the output above, the projects the directory is getting walked along with the files present inside the directory. This is how we can use the Walk method to iterate over directories in a file system.</p>
<h3>All the folders in the Path (only directories)</h3>
<p>If we want to print only the directories, we can again add checks in the funciton body, we can simply append the path name when the path returns <code>true</code> on <code>IsDir</code> function call.</p>
<pre><code class="language-go">package main

import(
    &quot;path/filepath&quot;
    &quot;log&quot;
    &quot;io/fs&quot;
)

func main() {

	var folders []string
	root := &quot;static/&quot;
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		dir_name := filepath.Base(root)
		if info.IsDir() {
            folders = append(folders, info.Name())
            return nil
		} else if info.IsDir() &amp;&amp; dir_name != info.Name() {
			return filepath.SkipDir
		}
        return nil
	})

	if err != nil {
		panic(err)
	}

	for _, folder := range folders{
		log.Println(folder)
	}
}
</code></pre>
<pre><code>$ go run walk.go

2022/10/02 12:13:25 static
2022/10/02 12:13:25 projects
</code></pre>
<p>Here, we can see it lists all the folder names present in the given path, it will log all the nested directories as well. In the above example, the <code>static/</code> path in the local system had a projects directory and hence it prints the same, but that can be till the final depth of the file system.</p>
<p>For all the examples on the <code>Walk</code> functions, you can check out the links on the GitHub repository:</p>
<ul>
<li>
<p><a href="https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/paths/walk.go">Walk and WalkDir function examples</a></p>
</li>
<li>
<p><a href="https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/paths/walk_anonymous.go">Walk and WalkDir function as anonymous function</a></p>
</li>
</ul>
<h2>Relative or Absolute Paths</h2>
<p>We have been using absolute paths in the above examples, but while navigating from one directory to other, we heavily make use of relative paths as they make it easier to move around.</p>
<h3>Check if a path is Absolute</h3>
<p>We can check if a path is absolute using the <a href="https://pkg.go.dev/path#IsAbs">IsAbs</a> function, the function takes in a path string as a parameter and returns a boolean value. It returns <code>true</code> if the provided path is absolute else it returns <code>false</code>.</p>
<h3>Check if a path is Absolute</h3>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;os&quot;
	&quot;path/filepath&quot;
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Println(dir)
	log.Println(filepath.IsAbs(dir))

    dir = &quot;../math&quot;
	log.Println(dir)
	log.Println(filepath.IsAbs(dir))
}
</code></pre>
<pre><code>$ go run rel_abs.go                                                                                                            
2022/10/02 14:38:44 /home/meet/code/techstructive-blog
2022/10/02 14:38:44 true
2022/10/02 14:38:44 ../math
2022/10/02 14:38:44 false
</code></pre>
<p>In the above example, we can see that when we parse <code>../math</code> indicating there's a <code>/math</code> directory, before the current directory(parent directory) we get <code>false</code>.</p>
<p>But when we parse the path obtained from <code>Getwd()</code> function call or a path which is located from the root path will get the return value as <code>true</code>.</p>
<h3>Get the relative path from base to target path</h3>
<p>Let's say we are in a certain directory <code>/a/b/c/</code>, we want to move into <code>/a/c/d/</code>, we will have to move back two times and then move into <code>c</code> followed by the <code>d</code> directory. The relative path from <code>/a/b/c/</code> to <code>/a/c/d/</code> can be described as <code>../../c/d/</code>. We have a function in golang that does the same, basically creating a relative path from the base directory path to a target path. The function is provided in the path/filepath package as <a href="https://pkg.go.dev/path/filepath@go1.19.1#Rel">Rel</a>, the function takes in two parameters, both as a string representing paths. The first is the base path(like you are in) and the second is the target path (as the target to reach). The function returns the string representation of the absolute path from the base to the target directory.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;os&quot;
	&quot;path/filepath&quot;
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

    dir, err = filepath.Abs(&quot;plugins/&quot;)
	s, err := filepath.Abs(&quot;static/projects/&quot;)
    if err != nil {
        log.Println(err)
    }

	log.Println(s)
	log.Println(dir)
	log.Println(filepath.Rel(s, dir))
}
</code></pre>
<pre><code>$ go run rel_abs.go

2022/10/02 12:26:09 /home/meet/code/techstructive-blog/static/projects
2022/10/02 12:26:09 /home/meet/code/techstructive-blog/plugins
2022/10/02 12:26:09 ../../plugins &lt;nil&gt;
</code></pre>
<p>We can see that the relative path from the two directories is given as the return string from the Rel function.</p>
<h3>Join paths</h3>
<p>The <a href="https://pkg.go.dev/path/filepath@go1.19.1#Join">Join</a> method provided in the <code>filepath</code> package, is used for combining <code>n</code> number of path strings as one path. It separates the file paths with the operating system-specific separator like <code>/</code> for Linux and <code>\</code> for windows.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;path/filepath&quot;
)

func main() {

	dir, err := filepath.Abs(&quot;operators/arithmetic/&quot;)
    if err != nil {
        log.Println(err)
    }

	log.Println(filepath.Join(&quot;golang&quot;, &quot;files&quot;))
	log.Println(filepath.Join(dir, &quot;/files&quot;, &quot;//read&quot;))
}
</code></pre>
<pre><code>$ go run rel_abs.go

2022/10/02 12:30:37 golang/files
2022/10/02 12:30:37 /home/meet/code/techstructive-blog/operators/arithmetic/files/read
</code></pre>
<p>In the above example, we can see that it parses the path accurately and ignore any extra separators in the string path.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/paths">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from the following post, we were able to explore the path package along with a few functions  io as well as os package. By using various methods and type objects, we were able to perform operations and work with the file paths. By using functions to iterate over file systems, checking for absolute paths, checking for the existence of paths, etc, the fundamentals of path handling in golang were explored.</p>
<p>Thank you for reading, if you have any queries, feedback, or questions, you drop them below on the blog as a <a href="https://www.meetgor.com/golang-paths/#comments">github discussion</a>, or you can ping me on my social handles as well. Happy Coding :)</p>
