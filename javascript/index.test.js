const Cache = require ('./index');

describe('Cache test', () => {
    describe('FIFO algorithm', () => {
        it('should be able to set the value', () => {
            const cache = new Cache(100, 'FIFO');
            const res = cache.set('a', '123');
            expect(res).toBe(true);
        })
    
        it('should be able to set the value though the memory limit exceeded', () => {
            const cache = new Cache(80, 'FIFO');
            for (let i = 0; i < 10; i++) {
                cache.set(i, 'testvalueformemorylimit');
            }
    
            const res = cache.set('a', '123');
            expect(res).toBe(true);
        })
    
        it('should be able to get the value', () => {
            const cache = new Cache(100, 'FIFO');
            const res = cache.set('a', '123');
            expect(cache.get('a').value).toBe('123');
        })
    });

    describe('LRU algorithm', () => {
        it('should be able to set the value', () => {
            const cache = new Cache(100, 'LRU');
            const res = cache.set('a', '123');
            expect(res).toBe(true);
        })
    
        it('should be able to set the value though the memory limit exceeded', () => {
            const cache = new Cache(80, 'LRU');
            for (let i = 0; i < 10; i++) {
                cache.set(i, 'testvalueformemorylimit');
            }
    
            const res = cache.set('a', '123');
            expect(res).toBe(true);
        })
    
        it('should be able to get the value', () => {
            const cache = new Cache(100, 'LRU');
            const res = cache.set('a', '123');
            expect(cache.get('a').value).toBe('123');
        })

        it('should release the least usage cache in the map', () => {
            const cache = new Cache(100, 'LRU');
            for (let i = 0; i < 10; i++) {
                cache.set(i, 'testvalueformemorylimit');
            }

            for (let i = 0; i < 9; i++) {
                cache.get(i);
            }

            expect(cache.releaseMemoryUntilSufficient(cache.map, 10)[0].accessCounter).toBe(0);
        })
    });
});