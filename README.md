# Simple Local Memory Cache
Memory caching is a method to cache any value on your local device using your local memory. This repository is creating a simple library to do memory caching locally.

Advantage on using local memory caching:
1. Fast performance
2. Low latency

In the other side, the disadvantage:
1. No data retention
2. Limited to local device specification, especially memory usages

# Caching strategy & memory management
There are multiple ways and strategy to cache, the limitation of memory caching is the device specification especially the memory limit that a device can support. This mean the amount of data that could be handled by the libary is bounded by how much memory that the local device could handle. To handle this memory management issue, we will run through few algorithm to solve this.

## First in first out (FIFO)
The simplest strategy, the idea is `first cached value will be first out in case of memory limit exceeded`

## Least recently used (LRU)
This caching strategy used the idea of `any cached value that least used, will be released in case of memory limit exceeded`. Let's take a step back and understand the main idea of caching, caching is crucial because it is used to cache value that oftenly access by user thus increasing the performance of the systems by reducing steps, so this strategy is oftenly the best strategy for caching.

## Time based
This strategy is straight forward, `in case of memory limit exceeded, release the expired cached value`. In case no value expired, there are 2 options
1. return error with message cache full
2. combine with other strategy (e.g. FIFO or LRU)

# Application Interface
This simple memory cache lib interface will contains 
| Method | Parameters | Value return | 
| ------ | ---------- | ------------ |
| Set    | key (`string, required`), value (`string, required`) | boolean |
| Get    | key (`string, required`) | key (`string, required`), value (string) |
| Delete | key (`string, required`) | boolean (`true` for success, `false` for failure) |

## Error codes
| Error Code | Description |
| ---------- | ----------- |
| N/A | N/A |

## Lib Initialization
To initialize the library, the lib will takes input of `memory limit as int in bytes` that you want to allocate and throw error if memory allocation is exceeding device constraint. 

Second parameters is a string to determine which algorithm to used, defaul will be `FIFO`

| Method | Parameters | Value return | 
| ------ | ---------- | ------------ |
| Initialization   | memory (`int, required`), algorithm (`string, optional`) | object |

Example of lib initialization
```
var simpleMemCache = new Memcache(10000, 'FIFO')
// initiate memcache class and allocate 10,000 bytes of memory for memcache
```

### Error codes
| Error Code | Description |
| ---------- | ----------- |
| MEMORY_CONSTRAINT_ERROR | Allocated memory is either larger than device limit or lower than allowed |
| ALGORITHM_NOT_FOUND_ERROR | Selected algorithm is not found, available algorithm (FIFO, LRU) |

# Supported Languages
This simple memory cache lib is written and available for these languages and stored in different branch of this repository
1. Javascript (repo branch `javascript`)
2. Go (repo branch `go`)
3. Python (repo branch `python`)
4. Java (repo branch `java`)
5. Typescript (repo branch `typescript`)