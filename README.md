# go-ipset

[![Build Status](https://api.travis-ci.org/gmccue/go-ipset.png?branch=master)](https://travis-ci.org/gmccue/go-ipset)
[![GoDoc](https://godoc.org/github.com/gmccue/go-ipset?status.svg)](https://godoc.org/github.com/gmccue/go-ipset)

go-ipset provides basic bindings for the [ipset kernel utility](http://ipset.netfilter.org/).

## Installation
```
go get github.com/gmccue/go-ipset
```

## Usage
The following are some basic usage examples for go-iptables. For more information, please [checkout the godoc](https://godoc.org/github.com/gmccue/go-ipset).

```go
import "github.com/gmccue/ipset"

// Construct a new ipset instance
ipset, err := ipset.New()
if err != nil {
    // Your custom error handling here.
}

// Create a new set
err := ipset.Create("my_set", "hash:ip")
if err != nil {
    // Your custom error handling here.
}
```

### Adding an entry to an ipset
```go
err := ipset.Add("my_set", "127.0.0.1")
if err != nil {
    // Your custom error handling here.
}
```

### Removing an entry from an ipset
```go
if err != nil {
    // Your custom error handling here.
}
```

### Save your ipset to a file
```go
err := ipset.Save("my_set", "/tmp/my_set.txt")
if err != nil {
    // Your custom error handling here.
}
```

### Restore your ipset from a file
```go
err := ipset.Restore("/tmp/my_set.txt")
if err != nil {
    // Your custom error handling here.
}
```
