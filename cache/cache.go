package cache

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	cacheDriver CacheDriver
	rdb         *redis.Client
	rdbCtx      context.Context
)

// 缓存配置
type CacheConfig struct {
	Driver CacheDriver
}

type CacheDriver interface {
	Init() error
	Set(key string, value interface{}) error
	Get(key string) (string, error)
	Clear() error
}

func SetDriver(driver CacheDriver) {
	driver.Init()
	cacheDriver = driver
}

func Set(key string, value interface{}) error {
	if cacheDriver == nil {
		return errors.New("未设置缓存驱动")
	}
	return cacheDriver.Set(key, value)
}

func Get(key string) (string, error) {
	if cacheDriver == nil {
		return "", errors.New("未设置缓存驱动")
	}
	return cacheDriver.Get(key)
}

func Clear() error {
	if cacheDriver == nil {
		return errors.New("未设置缓存驱动")
	}
	return cacheDriver.Clear()
}

type RedisCacheDriverConfig struct {
	Addr       string
	Password   string
	DB         int
	Expiration time.Duration
}

type RedisCacheDriver struct {
	Config RedisCacheDriverConfig
}

func (r *RedisCacheDriver) Init() error {
	if rdb != nil {
		return nil
	}

	rdbCtx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     r.Config.Addr,
		Password: r.Config.Password,
		DB:       r.Config.DB,
	})

	_, err := rdb.Ping(rdbCtx).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCacheDriver) Set(key string, value interface{}) error {
	err := rdb.Set(rdbCtx, key, value, r.Config.Expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCacheDriver) Get(key string) (string, error) {
	value, err := rdb.Get(rdbCtx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return value, nil
}

func (r *RedisCacheDriver) Clear() error {
	err := rdb.FlushDB(rdbCtx).Err()
	if err != nil {
		return err
	}
	return nil
}

type MemoryCacheDriver struct {
	cache      map[string]interface{}
	expiry     map[string]time.Time
	mutex      sync.RWMutex
	Expiration time.Duration
}

func (m *MemoryCacheDriver) Init() error {
	m.cache = make(map[string]interface{})
	m.expiry = make(map[string]time.Time)
	return nil
}

func (m *MemoryCacheDriver) Set(key string, value interface{}) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.cache[key] = value

	m.expiry[key] = time.Now().Add(time.Second * m.Expiration)
	return nil
}

func (m *MemoryCacheDriver) Get(key string) (string, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	value, found := m.cache[key]
	if !found {
		return "", nil
	}

	expiryTime, found := m.expiry[key]
	if found && time.Now().After(expiryTime) {
		delete(m.cache, key)
		delete(m.expiry, key)
		return "", nil
	}

	strValue, ok := value.(string)
	if !ok {
		return "", errors.New("value type error")
	}
	return strValue, nil
}

func (m *MemoryCacheDriver) Clear() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.cache = make(map[string]interface{})
	m.expiry = make(map[string]time.Time)
	return nil
}
