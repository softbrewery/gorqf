# gorqf

[![Build Status](https://travis-ci.org/softbrewery/gorqf.svg?branch=master)](https://travis-ci.org/softbrewery/gorqf)
[![codecov](https://codecov.io/gh/softbrewery/gorqf/branch/master/graph/badge.svg)](https://codecov.io/gh/softbrewery/gorqf)


## Install
```shell
$ go get github.com/softbrewery/gorqf
```

## Usage

### Basic

```go
// Create rqf parser
parser := rqf.NewParser()

// Json filter comming from rest request
jsonFilter := `
{
    "fields": ["-_id","isbn"],
    "order":["isbn ASC"],
    "limit":1,
    "offset":1
}`

// Parse the filter
filter, err := parser.Parse(jsonFilter)
if err != nil {
    // handle error
}
```

### Validation

```go
// Create rqf parser
parser := rqf.NewParser()

// only allow isbn/name to be selected
parser.FieldSchema( joi.String().Allow("isbn", "name") )

// only allow isbn to be ordered (ASC/DESC)
parser.OrderSchema( joi.String().Allow("isbn", "isbn ASC", "isbn DESC") )

// only allow paging between 10 and 100 items
parser.LimitSchema( joi.Int().Min(10).Max(100) )
            
// Json filter comming from rest request
jsonFilter := `
{
    "fields": ["isbn", "name"],
    "order":["isbn ASC"],
    "limit":25,
    "offset":1
}`

// Parse the filter
filter, err := parser.Parse(jsonFilter)
if err != nil {
    // handle error
}
```
