package main

import (
	"fmt"
	"time"

	"github.com/Redish101/recdn"
	"github.com/Redish101/recdn/cache"
)

func main() {
	app, _ := recdn.New(&recdn.AppConfig{
		Cache: cache.CacheConfig{
			Driver: &cache.MemoryCacheDriver{
				Expiration: time.Duration(15),
			},
		},
		GitHub: recdn.GitHubConfig{
			Raw: "https://ghraw.chuqis.com/%s/%s/%s",
		},
	})
	data, err := app.GitHub("Redish101/blog", "main", "package.min.json")
	fmt.Println(data)
	fmt.Println(err)
}
