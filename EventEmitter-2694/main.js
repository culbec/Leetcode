class EventEmitter {
    constructor() {
        this.events = new Map();
    }

    /**
     * @param {string} eventName
     * @param {Function} callback
     * @return {Object}
     */
    subscribe(eventName, callback) {
        if (!this.events.has(eventName)) {
            // Empty list of listeners if the event doesn't exist.
            this.events.set(eventName, []);
        }

        const listeners = this.events.get(eventName);
        listeners.push(callback); // adding a new listener to the event

        return {
            unsubscribe: () => {
                const index = listeners.indexOf(callback); // retrieving the index of the listener
                if (index != -1) {
                    // If the listener is still subscribed -> deleting it.
                    listeners.splice(index, 1);
                }
            }
        };
    }
    
    /**
     * @param {string} eventName
     * @param {Array} args
     * @return {Array}
     */
    emit(eventName, args = []) {
        if (!this.events.has(eventName)) {
            return []; // empty result
        }

        const listeners = this.events.get(eventName);
        const results = [];

        for (const listener of listeners) {
            // Saving the result of each listener execution.
            results.push(listener(...args));
        }

        return results;
    }
}

/**
 * const emitter = new EventEmitter();
 *
 * // Subscribe to the onClick event with onClickCallback
 * function onClickCallback() { return 99 }
 * const sub = emitter.subscribe('onClick', onClickCallback);
 *
 * emitter.emit('onClick'); // [99]
 * sub.unsubscribe(); // undefined
 * emitter.emit('onClick'); // []
 */
