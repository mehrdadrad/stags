# stags
[![Build Status](https://travis-ci.org/mehrdadrad/stags.svg?branch=master)](https://travis-ci.org/mehrdadrad/stags) 
[![Go Report Card](https://goreportcard.com/badge/github.com/mehrdadrad/stags)](https://goreportcard.com/report/github.com/mehrdadrad/stags)
[![GoCover](https://gocover.io/_badge/github.com/mehrdadrad/stags)](https://gocover.io/github.com/mehrdadrad/stags)
[![Godoc](https://godoc.org/github.com/mehrdadrad/stags?status.svg)](https://godoc.org/github.com/mehrdadrad/stags)

Go library to get and convert struct tags values

## Features
- supports nested structure
- convert the tag's value to given data type


## Installation

Standard `go get`:

```bash
$ go get github.com/mehrdadrad/stags
```

## Usage & Example

For usage and examples see the [Godoc](http://godoc.org/github.com/mehrdadrad/stags).

```go
type Configuration struct {
    Debug   bool `json:"debug" env:"CFG_DEBUG" default:"true"`
    Proxy struct {
        Address string `json:"address" env:"CFG_ADDR" default:"localhost"`
        Port    string `json:"port" env:"CFG_PORT" default:"8080"`
    }
}
```

```go
func main() {
     cfg := Configuration{}
 
     s := stags.New(cfg)
 
     debug := s.GetBool("default", "Debug") // returns true, type: bool
 
     pEnvKey := s.Get("env", "Proxy", "Address") // returns CFG_ADDR, type: string
     
     value, ok := s.LookupInt64("timeout", "proxy", "timeout") // ok: false
}

```

## License
This project is licensed under MIT license. Please read the LICENSE file.
