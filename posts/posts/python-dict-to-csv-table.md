{
  "title": "Turn Python dictionary into a neat CSV table",
  "status": "published",
  "slug": "python-dict-to-csv-table",
  "date": "2025-04-05T12:33:27Z"
}

<h2>Populating a Python dict having a table-like structure to a CSV</h2>
<p>Today, I want to share with you a neat trick I recently discovered for populating a CSV file with data in a table-like structure using Python.</p>
<h3>Writing Key-Value Pairs to a CSV Row</h3>
<p>Firstly, let's discuss the <code>write_key_value</code> function. This function allows us to write key-value pairs to a CSV row. It's particularly useful when dealing with metrics or data that can be represented as simple pairs.</p>
<pre><code class="language-python"># Function to populate a CSV row with key-value pairs
def write_key_value(writer, dictionary):
    for key, value in dictionary.items():
        writer.writerow([key, value])
</code></pre>
<h3>Writing a Table-Like Structure to a CSV File</h3>
<p>Now, let's dive into the <code>write_table</code> function, which handles more complex scenarios where the data follows a table-like structure. This function takes into account different types of metrics and adjusts the CSV table structure accordingly.</p>
<p>Assuming you have a structure of the dictionary like:</p>
<pre><code class="language-python">data = {
    &quot;Students&quot;: {
        &quot;John Doe&quot;: {
            &quot;course&quot;: &quot;Mathematics&quot;,
            &quot;grade&quot;: &quot;A&quot;,
            &quot;attendance&quot;: 95,
            &quot;assignments_completed&quot;: 15,
            &quot;student_id&quot;: &quot;JD001&quot;
        },
        &quot;Alice Smith&quot;: {
            &quot;course&quot;: &quot;Physics&quot;,
            &quot;grade&quot;: &quot;B+&quot;,
            &quot;attendance&quot;: 85,
            &quot;assignments_completed&quot;: 12,
            &quot;student_id&quot;: &quot;AS002&quot;
        },
        &quot;Bob Johnson&quot;: {
            &quot;course&quot;: &quot;Computer Science&quot;,
            &quot;grade&quot;: &quot;A-&quot;,
            &quot;attendance&quot;: 90,
            &quot;assignments_completed&quot;: 14,
            &quot;student_id&quot;: &quot;BJ003&quot;
        }
    }
}
</code></pre>
<p>And you want to write it to a CSV file, like this:</p>
<pre><code class="language-csv">student, course, grade, attendance, assignments_completed, student_id
John Doe, Mathematics, A, 95, 15, JD001
Alice Smith, Physics, B+, 85, 12, AS002
Bob Johnson, Computer Science, A-, 90, 14, BJ003
</code></pre>
<p>We can create a function <code>write_table</code> that will take in the <code>dictionary</code> as the actual data. We want to store the keys of the inner dictionary to be the header/columns of the csv file. As we can see the keys of the inner dict i.e. the value for the key <code>John Doe</code> is a dict with the keys <code>course</code>, <code>grade</code>, <code>attendance</code>, etc. which remain the same for the all the keys in the dictionary.</p>
<p>So, we can first create a <code>row_keys</code> variable to store the keys of the actual dictionary this will be the first column rows in the csv.</p>
<p>Further we check if the <code>row_keys</code> is a dict and then we append it with the <code>index_key</code> which will be the first column in the csv. Since all the keys remain the same for the inner-dict, we can pick the first dict and create the <code>header</code> with the inner-dict keys.</p>
<p>So, we can write the list <code>header</code> to the csv file.</p>
<p>Then for each key in the <code>row_keys</code> we can create a list <code>row</code> with the key and the values of the inner-dict.</p>
<pre><code class="language-python"># Function to populate a CSV with a table-like structure
def write_table(writer, dictionary, index_key):

    row_keys = list(dictionary.keys())

    if row_keys and data[row_keys[0]] is not None:
        headers = [index_key] + list(
            dictionary[row_keys[0]].keys()
        )
    else:
        return
    writer.writerow(headers)
    for key in row_keys:
        row = [key] + list(dictionary[key].values())
        writer.writerow(row)


with open('data.csv', 'w', newline='') as csvfile:
    writer = csv.writer(csvfile)
    for key in data:
        write_table(writer, data[key], key)
</code></pre>
<h3>Example Usage</h3>
<p>To illustrate how these functions can be used, let's consider a scenario where we have various types of metrics to populate into a CSV file. We handle key-value paired metrics separately and then populate the CSV with table-like metrics.</p>
<pre><code class="language-python">import csv

data = {
    &quot;Students&quot;: {
        &quot;John Doe&quot;: {
            &quot;course&quot;: &quot;Mathematics&quot;,
            &quot;grade&quot;: &quot;A&quot;,
            &quot;attendance&quot;: 95,
            &quot;assignments_completed&quot;: 15,
            &quot;student_id&quot;: &quot;JD001&quot;
        },
        &quot;Alice Smith&quot;: {
            &quot;course&quot;: &quot;Physics&quot;,
            &quot;grade&quot;: &quot;B+&quot;,
            &quot;attendance&quot;: 85,
            &quot;assignments_completed&quot;: 12,
            &quot;student_id&quot;: &quot;AS002&quot;
        },
        &quot;Bob Johnson&quot;: {
            &quot;course&quot;: &quot;Computer Science&quot;,
            &quot;grade&quot;: &quot;A-&quot;,
            &quot;attendance&quot;: 90,
            &quot;assignments_completed&quot;: 14,
            &quot;student_id&quot;: &quot;BJ003&quot;
        }
    },
    &quot;Countries&quot;: {
        &quot;USA&quot;: {
            &quot;capital&quot;: &quot;Washington, D.C.&quot;,
            &quot;population&quot;: 331000000,
            &quot;area_sq_km&quot;: 9833517,
            &quot;official_languages&quot;: [&quot;English&quot;, &quot;Spanish&quot;],
            &quot;currency&quot;: &quot;United States Dollar (USD)&quot;
        },
        &quot;India&quot;: {
            &quot;capital&quot;: &quot;New Delhi&quot;,
            &quot;population&quot;: 1380004385,
            &quot;area_sq_km&quot;: 3287263,
            &quot;official_languages&quot;: [&quot;Hindi&quot;, &quot;English&quot;],
            &quot;currency&quot;: &quot;Indian Rupee (INR)&quot;
        },
        &quot;Brazil&quot;: {
            &quot;capital&quot;: &quot;Brasília&quot;,
            &quot;population&quot;: 212559417,
            &quot;area_sq_km&quot;: 8515770,
            &quot;official_languages&quot;: [&quot;Portuguese&quot;],
            &quot;currency&quot;: &quot;Brazilian Real (BRL)&quot;
        }
    }
}

def populate_table(writer, data, index_key):
    row_keys = list(data.keys())
    if row_keys and data[row_keys[0]] is not None:
        headers = [index_key] + list(data[row_keys[0]].keys())
    else:
        return
    writer.writerow(headers)
    for key in row_keys:
        row = [key] + list(data[key].values())
        writer.writerow(row)


with open('data.csv', 'w', newline='') as csvfile:
    writer = csv.writer(csvfile)
    for key in data:
        populate_table(writer, data[key], key)
        writer.writerow([])
</code></pre>
<pre><code class="language-csv">Students,course,grade,attendance,assignments_completed,student_id
John Doe,Mathematics,A,95,15,JD001
Alice Smith,Physics,B+,85,12,AS002
Bob Johnson,Computer Science,A-,90,14,BJ003

Countries,capital,population,area_sq_km,official_languages,currency
USA,&quot;Washington, D.C.&quot;,331000000,9833517,&quot;['English', 'Spanish']&quot;,United States Dollar (USD)
India,New Delhi,1380004385,3287263,&quot;['Hindi', 'English']&quot;,Indian Rupee (INR)
Brazil,Brasília,212559417,8515770,['Portuguese'],Brazilian Real (BRL)
</code></pre>
<p>Thank you, Happy Coding :)</p>
