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
