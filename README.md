# GoCache

GoCache is a proof-of-concept solution for an in-memory key-value store in Go. It provides basic caching functionality with support for various cache eviction policies.

## Features

- In-memory key-value store
- Multiple eviction policies (LRU, LFU, FIFO)
- Thread-safe operations
- Easy to use API

## Getting Started

### Installation

To install GoCache, use `go get`:

```sh
go get github.com/hunter/gocache
```

After installing, run `go mod tidy` to ensure all dependencies are properly managed:

```sh
go mod tidy
```

### Starting and Testing the Application

Use the provided Makefile as an entry point to start and test the application. To start the application, run:

```sh
make start
```

To run tests, use:

```sh
make test
```

## Usage

Here is a basic example of how to use GoCache:

```go
package main

import (
    "fmt"
    "github.com/yourusername/gocache"
)

func main() {
    cache := gocache.NewCache(gocache.LRU, 100) // LRU policy with a capacity of 100 items

    cache.Set("key1", "value1")
    value, found := cache.Get("key1")
    if found {
        fmt.Println("Found value:", value)
    } else {
        fmt.Println("Value not found")
    }
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, feel free to open an issue or contact the repository owner.

##
 QGoCache also provides a query option to retrieve multiple values based on a custom filter function. This allows for more advanced querying capabilities.
ue
ryHere is an example of how to use the query option:
 O
pt```go
iopackage main
n

im      "github.com/yourusername/gocache"
  "f)
mt"

pofu  
  cach    cache.Set("key1", "value1")
e := g    cache.Set("key2", "value2")
ocache    cache.Set("key3", "value3")
.NewCa
che(go    results := cache.Query(func(key, value interface{}) bool {
cache.        return key.(string) > "key1"
LRU, 1    })
00)
n
c m           fmt.Println("Found value:", result)
 for _    }
, resu}
lt := ```
range results {
ain() {
rt (

