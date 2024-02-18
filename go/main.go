package main

import (
	"simple-memcache/memcache"
)

func main() {
	memcache.Initialize(100, "FIFO")
}
