// strategy create interface set & get
// use array to store the data (key, value, accessedTimes (for LFU algorithm), lastAccess (for LRU algorithm))
// need to think around the memory management

// make it global so it can be accessed from anywhere

// lib to check memory footprint
const sizeof = require('object-sizeof')

const Cache = class{
    constructor(maxMemoryLimit, algorithm = 'FIFO') {
        this.maxMemoryLimit = maxMemoryLimit;
        this.algorithm = algorithm;
        this.map = new Map();
    }

    releaseMemoryUntilSufficient(map, bytes) {
        let arrayOfDeletedCache = [];
        while (this.maxMemoryLimit - sizeof(map) < bytes || sizeof(map) <= 0) {
            if (this.algorithm == 'FIFO') {
                // keep releasing first entry until memory sufficient
                let key = map.keys().next().value;
                let cache = map.get(key);

                if (map.delete(key)) {
                    arrayOfDeletedCache.push(cache);
                }
            } else if (this.algorithm == 'LFU') {
                // set the first key as lowest key pivot
                let lowestAccessedCacheKey = map.keys().next().value;
                // find the lowest usages and release from map
                map.forEach((value, key) => {
                    if (map.get(lowestAccessedCacheKey).accessedTimes > value.accessedTimes) {
                        lowestAccessedCacheKey = key;
                    }
                });

                let cache = map.get(lowestAccessedCacheKey);

                if (map.delete(lowestAccessedCacheKey)) {
                    arrayOfDeletedCache.push(cache);
                }
            }
        }
        return arrayOfDeletedCache;
    }

    set(key, value) {
        if (sizeof(value) > this.maxMemoryLimit) {
            return false;
        }

        this.releaseMemoryUntilSufficient(this.map, sizeof(value));

        try {
            if (this.algorithm == 'FIFO') {
                this.map.set(key, value)
            } else if (this.algorithm == 'LFU') {
                this.map.set(key, {
                    value: value,
                    accessedTimes: 0
                });
            }
            return true
        } catch (e) {
            return false
        }
    }
    
    get(key) {
        if (this.algorithm == 'FIFO') {
            return {
                key: key,
                value: this.map.get(key)
            }
        } else if (this.algorithm == 'LFU') {
            let cache = this.map.get(key);
            if (cache != undefined) {
                cache.accessedTimes += 1;
                return {
                    key: key,
                    value: cache.value
                }
            }

            return undefined;
        }
    }

    delete(key) {
        return this.map.delete(key);
    }
}

module.exports = Cache;