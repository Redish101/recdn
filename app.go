package recdn

import "github.com/Redish101/recdn/cache"

type AppConfig struct {
	Cache  cache.CacheConfig
	GitHub GitHubConfig
}

type App struct {
	Cache  cache.CacheDriver
	Config *AppConfig
}

func New(config *AppConfig) (*App, error) {
	cache.SetDriver(config.Cache.Driver)
	return &App{
		Cache:  config.Cache.Driver,
		Config: config,
	}, nil
}
