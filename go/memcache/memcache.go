package memcache

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type CacheItem struct {
	key   string
	value interface{}
}

var MemoryLimit int = 0

// available alogirthm is "FIFO", "LFU", "LRU"
var Algorithm string = ""
var Cache map[string]interface{} = make(map[string]interface{})

func Initialize(memoryLimit int, algorithm string) error {
	if memoryLimit <= 0 {
		return errors.New("memory limit must be greater than 0")
	}

	MemoryLimit = memoryLimit
	Algorithm = algorithm

	return nil
}

func clearCache(bytes int) error {
	// TODO: memory foot print of map is not working correctly
	// this only returns the memory allocation for the map it self but not the values - meaning this always returning 8 bytes
	var memUsed = unsafe.Sizeof(Cache)
	fmt.Printf("Size of memory limit: %d bytes and map: %d bytes and desired inserted bytes: %d\n", MemoryLimit, memUsed, bytes)
	fmt.Println("Cache size: ", Cache)

	for MemoryLimit-int(memUsed) < bytes {
		// keep removing the first element in map (FIFO algorithm)
		for k := range Cache {
			delete(Cache, k)
			break
		}
	}

	return nil
}

func Set(key string, value interface{}) (cache CacheItem, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("cache insertion failed")
		}
	}()

	err = clearCache(int(reflect.ValueOf(value).Type().Size()))

	Cache[key] = value
	return CacheItem{
		key:   key,
		value: value,
	}, nil
}

func Get(key string) CacheItem {
	if val, ok := Cache[key]; ok {
		return CacheItem{
			key:   key,
			value: val,
		}
	}

	return CacheItem{}
}

func Delete(key string) bool {
	// delete from map
	if _, ok := Cache[key]; ok {
		delete(Cache, key)
		return true
	}

	return false
}
