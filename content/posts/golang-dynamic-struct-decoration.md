+++
Description = "Personal findings, thoughts & blog"
Tags = ["blog", "personal"]
date = "2016-01-28T20:42:45+02:00"
title = "Golang dynamic struct decoration"

+++

## Dynamic struct decoration using type assertion.

> Use cases: API / templates.

Having simple map of:

```go
  input :=  map[string]interface{}{
    "Key1": []string{"some", "key"},
    "key3": nil,
    "val": 2,
    "val2": "str",
    "val3": 4,
  }
```

One can decorate it using type assertion by iterating over it.

```go
  for key, value := range input {
    slice, ok := value.([]string)
    if ok {
      input["Count"+key] = len(slice)
    }
  }
```

This becomes very useful when serializing struct into json.
To serialize struct use ```github.com/fatih/structs```

```go
  input := structs.Map(...)

  for key, value := range input {
    slice, ok := value.([]string)
    if ok {
      input["Count"+key] = len(slice)
    }
  }
```
