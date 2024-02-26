from sys import getsizeof

class Memcache:
    def __init__(self, memoryLimit, algorithm):
        self.cache = {}
        self.memory_limit = memoryLimit
        self.algorithm = algorithm

    def set(self, key: any, value: any) -> any:
        self.clearCache({
            "key": key,
            "value": {
                "value": value,
                "accessedTimes": 0
            }
        })
        self.cache[key] = { 
            "value": value,
            "accessedTimes": 0
        }

        return {
            "key": key,
            "value": value
        }
    
    def get(self, key: any) -> any:
        if key not in self.cache:
            return None

        self.cache[key]["accessedTimes"] += 1
        return {
            "key": key,
            "value": self.cache[key]["value"]
        }
    
    def delete(self, key: any) -> bool:
        del self.cache[key]
        return True

    def clearCache(self, value):
        # remove dict until memory space is sufficient
        if getsizeof(value) > self.memory_limit:
            raise Exception("Value too large for cache")

        while self.memory_limit - self.getCacheCurrentMemoryAllocation() < getsizeof(value):
            if self.algorithm == "FIFO":
                key = next(iter(self.cache))
                del self.cache[key]

            if self.algorithm == "LFU":
                key = next(iter(self.cache))
                for k, v in self.cache.items():
                    if v["accessedTimes"] < self.cache[key]["accessedTimes"]:
                        key = k

                del self.cache[key]
    
    def getCacheCurrentMemoryAllocation(self):
        return getsizeof(self.cache) + sum(map(getsizeof, self.cache.values())) + sum(map(getsizeof, self.cache.keys()))