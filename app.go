package recdn

import "github.com/Redish101/recdn/cache"

// App配置
type AppConfig struct {
	Cache  cache.CacheConfig
	GitHub GitHubConfig
}

// App
type App struct {
	Cache  cache.CacheDriver
	Config *AppConfig
}

// 新建实例
func New(config *AppConfig) (*App, error) {
	cache.SetDriver(config.Cache.Driver)
	return &App{
		Cache:  config.Cache.Driver,
		Config: config,
	}, nil
}
