// strategy create interface set & get
// use map to store the data
// need to think around the memory management

// make it global so it can be accessed from anywhere

// lib to check memory footprint
const sizeof = require('object-sizeof')

const Cache = class{
    constructor(maxMemoryLimit) {
        this.maxMemoryLimit = maxMemoryLimit;
        this.map = new Map();
    }

    releaseMemoryUntilSufficient(bytes) {
        while (this.maxMemoryLimit - sizeof(this.map) < bytes || sizeof(this.map) <= 0) {
            // keep releasing first entry
            this.map.delete(this.map.keys().next().value);
        }
    }

    set(id, values) {
        this.releaseMemoryUntilSufficient(sizeof(values));

        try {
            this.map.set(id, values)
            return true
        } catch (e) {
            return false
        }
    }
    
    get(id) {
        return this.map.get(id);
    }
}

module.exports = Cache;