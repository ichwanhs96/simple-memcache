import unittest
from memcache import Memcache

class TestMemcache(unittest.TestCase):
    def test_init(self):
        m = Memcache(1000, "FIFO")
        self.assertEqual(m.memory_limit, 1000)
        self.assertEqual(m.algorithm, "FIFO")
        self.assertEqual(m.cache, {})

    def test_set(self):
        m = Memcache(1000, "FIFO")
        self.assertEqual(m.set("a", 1), {"key": "a", "value": 1})

    def test_set_value_exceeded_memory_limit(self):
        try:
            m = Memcache(100, "FIFO")
            m.set("a", 1)
        except Exception as e:
            self.assertEqual(str(e), "Value too large for cache")

    def test_get(self):
        m = Memcache(1000, "FIFO")
        m.set("a", 1)
        self.assertEqual(m.get("a"), {"key": "a", "value": 1})

    def test_delete(self):
        m = Memcache(1000, "FIFO")
        m.set("a", 1)
        self.assertEqual(m.delete("a"), True)

    def test_clearCache_FIFO(self):
        m = Memcache(1000, "FIFO")
        m.set("a", 1)
        m.set("b", 2)
        m.set("c", 3)
        m.set("d", 4)
        m.clearCache(5)
        self.assertTrue("a" not in m.cache)

    def test_clearCache_LFU(self):
        m = Memcache(1000, "LFU")
        m.set("a", 1)
        m.set("b", 2)
        m.set("c", 3)
        m.get("a")
        m.get("a")
        m.get("a")
        m.get("b")
        m.clearCache("super long string which consume big memory brooooooooooooooooooooooooooooooooooooooooo")
        print(m.cache)
        self.assertTrue("d" not in m.cache)

if __name__ == '__main__':
    unittest.main()