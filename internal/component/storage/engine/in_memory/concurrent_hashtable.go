package in_memory

import "sync"

// ConcurrentHashTable -
type ConcurrentHashTable[TKey comparable, TValue any] struct {
	rmx  *sync.RWMutex
	data map[TKey]TValue
}

// MakeConcurrentHashTable -
func MakeConcurrentHashTable[TKey comparable, TValue any]() ConcurrentHashTable[TKey, TValue] {
	return ConcurrentHashTable[TKey, TValue]{
		rmx:  &sync.RWMutex{},
		data: make(map[TKey]TValue),
	}
}

// Get -
func (h ConcurrentHashTable[TKey, TValue]) Get(key TKey) (TValue, bool) {
	h.rmx.RLock()
	defer h.rmx.RUnlock()

	value, found := h.data[key]
	return value, found
}

// Set -
func (h ConcurrentHashTable[TKey, TValue]) Set(key TKey, value TValue) {
	h.rmx.Lock()
	defer h.rmx.Unlock()

	h.data[key] = value
}

// Delete -
func (h ConcurrentHashTable[TKey, TValue]) Delete(key TKey) {
	h.rmx.Lock()
	defer h.rmx.Unlock()

	delete(h.data, key)
}
