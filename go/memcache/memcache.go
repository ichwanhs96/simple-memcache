package memcache

import (
	"errors"
	"reflect"
	"unsafe"
)

type CacheItem struct {
	value         interface{}
	accessedTimes int
}

type CacheItemResponse struct {
	key   string
	value interface{}
}

var MemoryLimit int = 0

// available alogirthm is "FIFO", "LFU", "LRU"
var Algorithm string = ""
var Cache map[string]CacheItem = make(map[string]CacheItem)

func Initialize(memoryLimit int, algorithm string) error {
	if memoryLimit <= 0 {
		return errors.New("memory limit must be greater than 0")
	}

	MemoryLimit = memoryLimit
	Algorithm = algorithm

	return nil
}

func getLeastUsedCacheKey() string {
	var lowestAccessedKey string
	for k := range Cache {
		if lowestAccessedKey == "" {
			lowestAccessedKey = k
		}

		if Cache[k].accessedTimes < Cache[lowestAccessedKey].accessedTimes {
			lowestAccessedKey = k
		}

		break
	}

	return lowestAccessedKey
}

func clearCache(bytes int) error {
	// TODO: memory foot print of map is not working correctly
	// function unsafe.Sizeof only returns the memory allocation for the map it self but not the values - meaning this always returning 8 bytes
	var memUsed = unsafe.Sizeof(Cache)

	for MemoryLimit-int(memUsed) < bytes {
		// keep removing the first element in map (FIFO algorithm)
		if Algorithm == "FIFO" {
			for k := range Cache {
				delete(Cache, k)
				break
			}
		} else if Algorithm == "LFU" {
			delete(Cache, getLeastUsedCacheKey())
		}
	}

	return nil
}

func Set(key string, value interface{}) (cache CacheItemResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("cache insertion failed")
		}
	}()

	err = clearCache(int(reflect.ValueOf(value).Type().Size()))

	cacheItem := CacheItem{
		value:         value,
		accessedTimes: 0,
	}

	Cache[key] = cacheItem

	return CacheItemResponse{
		key:   key,
		value: value,
	}, nil
}

func Get(key string) CacheItemResponse {
	if val, ok := Cache[key]; ok {
		return CacheItemResponse{
			key:   key,
			value: val.value,
		}
	}

	return CacheItemResponse{}
}

func Delete(key string) bool {
	// delete from map
	if _, ok := Cache[key]; ok {
		delete(Cache, key)
		return true
	}

	return false
}
