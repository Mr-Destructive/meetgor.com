---
type: til
title: "Golang: Sort Package Introduction"
description: "Understanding the fundamentals of the sort package in Golang. Sorting integers, slices, struct values, maps"
status: published
slug: golang-sort-package-basic
tags: ["go",]
date: 2024-01-15 22:30:00
---

I have been learning Golang for around 2 years now, and I have never paid attention to the sort package, can you believe that! Where was this package hiding?

The `sort` package provides convenient methods for sorting slices and user-defined collections in Go. Here's a quick overview of what it can do:

## Sorting Slices

To sort a slice of builtin types like ints, strings, etc, you can simply call `sort.Slice()`:

```go
nums := []int{5, 2, 6, 3, 1}

sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})

// or 
// sorts.Ints(nums)

fmt.Println(nums)
```

```bash
$ go run main.go

[1, 2, 3, 5, 6]
```
The sort is in-place, mutating the original slice. You can customize the sort order by providing a less(i, j int) bool function.

## Sorting Struct Slices

To sort a slice of custom structs, add a Len(), Less(i, j int) bool, and Swap(i, j int) method:

```go
type Person struct {
  Name string
  Age  int
}

people := []Person{
  {"Bob", 30},
  {"John", 20},
  {"Alice", 25},
}

sort.Slice(people, func(i, j int) bool {
  return people[i].Age < people[j].Age
}) 

fmt.Println(person)
```

```bash
$ go run main.go
[{"Name":"Alice","Age":25},{"Name":"Bob","Age":30},{"Name":"John","Age":20}]
```

This will sort people by age.

## Sorting Maps

Maps are inherently unordered in Go. We can't sort them, but we can iterate them in a sorted way. We need a separate slice of keys or values(whicher required), we will sort those keys/values and iterate over them in that order.:

### Sort by Keys

To sort a map by keys, we can use the `sort.Strings()` function or any other sort function as per the data structure, there are functions like [sort.Ints](https://pkg.go.dev/sort#Ints), [sort.Float64](https://pkg.go.dev/sort#Float64s), or [sort.Slice](https://pkg.go.dev/sort#Slice). We can create a new slice of keys from the map and then apply the sort on that newly created slice. After the slice of keys is created, we can iterate over it and access the map values in a order of sorted keys.

```go
counts := map[string]int{
  "hello": 5,
  "world": 2,
  "foo": 3,
}

keys := make([]string, 0, len(counts))
// extract keys 
for k := range counts {
  keys = append(keys, k)
} 

// sort keys
sort.Strings(keys) 

// iterate with the sorted keys slice
for _, k := range keys {
  fmt.Println(k, counts[k]) 
}
```

```bash
$ go run main.go

foo 3
hello 5
world 2
```

This prints the map ordered by key. You can sort by value too.

### Sort by Value

To iterate the map with sorting order of values, we can approach it in a similar way. We create a slice of keys. This time, we don't sort the keys, instead, we change the position of the slice of keys depending on the values. This changes the key order based on the sorted values in the map. By similarly iterating over the key slice, we iterate over the map in a sorted order of the values.

```go
counts := map[string]int{
  "hello": 5,
  "world": 2,
  "foo": 3,
}

keys := make([]string, 0, len(counts))
// extract keys
for k := range counts {
  keys = append(keys, k)
}

// sort by value
// i.e. change the order of key based on values sort order
sort.SliceStable(keys, func(i, j int) bool {
    return counts[keys[i]] < counts[keys[j]]
})

// iterate sorted
for _, k := range keys {
  fmt.Println(k, counts[k])
}
```

```bash
$ go run main.go

world 2
foo 3
hello 5
```

The `sort.SliceStable` function is used to sort the slice in place. It takes in a slice and the less function which is the actual logic for comparison in the values. This function sorts the key elements based on the values in the map, and thereby the key slice is shuffled by the sorted order of the values in the map.

The sort package is super useful for quickly organizing Go data structures.
There are [Find](https://pkg.go.dev/sort#Find) which tries to return a index of the first element that satisfies a comparison condition with a flag or found or not with a boolean. There is also a [Search](https://pkg.go.dev/sort#Search) which is used specifically for searching elements in a sorted array, it also gives a first index of the occuring element. But that is a topic for another day!

Thank you, Happy Coding :)
