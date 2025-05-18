---
type: til
title: "Turn Python dictionary into a neat CSV table"
description: "Exploring how to write python dict/key-value pairs and a table-like structure to a CSV file."
status: published
slug: python-dict-to-csv-table
tags: ["python",]
date: 2024-03-20 22:30:00
---

## Populating a Python dict having a table-like structure to a CSV

Today, I want to share with you a neat trick I recently discovered for populating a CSV file with data in a table-like structure using Python.

### Writing Key-Value Pairs to a CSV Row

Firstly, let's discuss the `write_key_value` function. This function allows us to write key-value pairs to a CSV row. It's particularly useful when dealing with metrics or data that can be represented as simple pairs.

```python
# Function to populate a CSV row with key-value pairs
def write_key_value(writer, dictionary):
    for key, value in dictionary.items():
        writer.writerow([key, value])
```

### Writing a Table-Like Structure to a CSV File

Now, let's dive into the `write_table` function, which handles more complex scenarios where the data follows a table-like structure. This function takes into account different types of metrics and adjusts the CSV table structure accordingly.

Assuming you have a structure of the dictionary like:

```python
data = {
    "Students": {
        "John Doe": {
            "course": "Mathematics",
            "grade": "A",
            "attendance": 95,
            "assignments_completed": 15,
            "student_id": "JD001"
        },
        "Alice Smith": {
            "course": "Physics",
            "grade": "B+",
            "attendance": 85,
            "assignments_completed": 12,
            "student_id": "AS002"
        },
        "Bob Johnson": {
            "course": "Computer Science",
            "grade": "A-",
            "attendance": 90,
            "assignments_completed": 14,
            "student_id": "BJ003"
        }
    }
}
```

And you want to write it to a CSV file, like this:

```csv
student, course, grade, attendance, assignments_completed, student_id
John Doe, Mathematics, A, 95, 15, JD001
Alice Smith, Physics, B+, 85, 12, AS002
Bob Johnson, Computer Science, A-, 90, 14, BJ003
```

We can create a function `write_table` that will take in the `dictionary` as the actual data. We want to store the keys of the inner dictionary to be the header/columns of the csv file. As we can see the keys of the inner dict i.e. the value for the key `John Doe` is a dict with the keys `course`, `grade`, `attendance`, etc. which remain the same for the all the keys in the dictionary.

So, we can first create a `row_keys` variable to store the keys of the actual dictionary this will be the first column rows in the csv. 

Further we check if the `row_keys` is a dict and then we append it with the `index_key` which will be the first column in the csv. Since all the keys remain the same for the inner-dict, we can pick the first dict and create the `header` with the inner-dict keys.

So, we can write the list `header` to the csv file.

Then for each key in the `row_keys` we can create a list `row` with the key and the values of the inner-dict.


```python
# Function to populate a CSV with a table-like structure
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
```

### Example Usage

To illustrate how these functions can be used, let's consider a scenario where we have various types of metrics to populate into a CSV file. We handle key-value paired metrics separately and then populate the CSV with table-like metrics.

```python
import csv

data = {
    "Students": {
        "John Doe": {
            "course": "Mathematics",
            "grade": "A",
            "attendance": 95,
            "assignments_completed": 15,
            "student_id": "JD001"
        },
        "Alice Smith": {
            "course": "Physics",
            "grade": "B+",
            "attendance": 85,
            "assignments_completed": 12,
            "student_id": "AS002"
        },
        "Bob Johnson": {
            "course": "Computer Science",
            "grade": "A-",
            "attendance": 90,
            "assignments_completed": 14,
            "student_id": "BJ003"
        }
    },
    "Countries": {
        "USA": {
            "capital": "Washington, D.C.",
            "population": 331000000,
            "area_sq_km": 9833517,
            "official_languages": ["English", "Spanish"],
            "currency": "United States Dollar (USD)"
        },
        "India": {
            "capital": "New Delhi",
            "population": 1380004385,
            "area_sq_km": 3287263,
            "official_languages": ["Hindi", "English"],
            "currency": "Indian Rupee (INR)"
        },
        "Brazil": {
            "capital": "Brasília",
            "population": 212559417,
            "area_sq_km": 8515770,
            "official_languages": ["Portuguese"],
            "currency": "Brazilian Real (BRL)"
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
```

```csv
Students,course,grade,attendance,assignments_completed,student_id
John Doe,Mathematics,A,95,15,JD001
Alice Smith,Physics,B+,85,12,AS002
Bob Johnson,Computer Science,A-,90,14,BJ003

Countries,capital,population,area_sq_km,official_languages,currency
USA,"Washington, D.C.",331000000,9833517,"['English', 'Spanish']",United States Dollar (USD)
India,New Delhi,1380004385,3287263,"['Hindi', 'English']",Indian Rupee (INR)
Brazil,Brasília,212559417,8515770,['Portuguese'],Brazilian Real (BRL)
```

Thank you, Happy Coding :)
