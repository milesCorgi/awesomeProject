package caches

import (
	"sync"
	"time"
)

var cache = struct { // 声明 struct 字面量 cahce (匿名结构体)
	sync.RWMutex                    // 互斥锁, 内嵌 Struct
	caches       map[string]Element // kv 内存存储
}{
	caches: make(map[string]Element), // 初始化 kv 存储
}

type Element struct {
	Value       string
	ExpiredTime time.Time
}

// GetCache safe get memory cache
func GetCache(key string) Element {
	v := cache.caches[key] // 获取缓存值
	return v
}

func GetAllCache() map[string]Element {
	v := cache.caches // 获取缓存值
	return v
}

// PutCache safe put memory cache
func PutCache(key string, val Element) {
	cache.RLock()           // 锁住
	cache.caches[key] = val // 设置缓存值
	defer cache.RUnlock()   // 释放锁
}

func ClearCache() {
	cache.RLock()                           // 锁住
	cache.caches = make(map[string]Element) // 初始化 kv 存储
	defer cache.RUnlock()                   // 释放锁
}

func DelACache(key string) {
	cache.RLock()             // 锁住
	delete(cache.caches, key) // 设置缓存值
	defer cache.RUnlock()     // 释放锁
}
