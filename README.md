# LRU-Cache
A Local Cache with LRU , and high performance，Concurrency safety，in memory.

## Get Started.

`go get github.com/Hamster601/LRU-Cache`

```Go
package main

import (
    "fmt"
    LRU "github.com/Hamster601/LRU-Cache"
)

func main() {
    cache := LRU.NewCache(8)
    cache.Put("test","value")
    result := cache.Get("test")
    fmt.Println(result)
}
```

# Usage

---

TODO



# BenchmartTest
```bigquery
goos: darwin
goarch: amd64
pkg: github.com/Hamster601/LRU-Cache
BenchmarkGetAndPut1
BenchmarkGetAndPut1-8   	  460646	      3581 ns/op
```