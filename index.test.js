const Cache = require ('./index');

describe('Cache test', () => {
    it('should be able to set the value', () => {
        const cache = new Cache(100);
        const res = cache.set('a', '123');
        expect(res).toBe(true);
    })

    it('should be able to set the value though the memory limit exceeded', () => {
        const cache = new Cache(80);
        for (let i = 0; i < 10; i++) {
            cache.set(i, 'testvalueformemorylimit');
        }

        const res = cache.set('a', '123');
        expect(res).toBe(true);
    })

    it('should be able to get the value', () => {
        const cache = new Cache(100);
        const res = cache.set('a', '123');
        expect(cache.get('a')).toBe('123');
    })
});