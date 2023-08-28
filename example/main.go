package main

import (
	"fmt"
	"time"

	"github.com/Redish101/recdn"
	"github.com/Redish101/recdn/cache"
)

func main() {
	cache.SetDriver(&cache.MemoryCacheDriver{
		Expiration: time.Duration(1),
	})
	file, _ := recdn.Fetch("https://redish101.top/")
	fmt.Println(file)
}