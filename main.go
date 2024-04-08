package main

import (
	"fmt"
	"sync"
)

type KeyValuePair struct {
	Key   string
	Value string
}

type KeyValueStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: make(map[string]string),
	}
}

func (kvs *KeyValueStore) Get(key string) (string, bool) {
	kvs.mu.RLock()
	defer kvs.mu.RUnlock()
	value, ok := kvs.data[key]
	return value, ok
}

func (kvs *KeyValueStore) Put(key, value string) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	kvs.data[key] = value
}

func main() {
	kvs := NewKeyValueStore()
	kvs.Put("key1", "value1")
	kvs.Put("key2", "value2")
	value, ok := kvs.Get("key1")

	if ok {

		fmt.Println("Value for key1:", value)
	} else {

		fmt.Println("Key1 not found")
	}
	value, ok = kvs.Get("key3")
	if ok {
		fmt.Println("Value for key3:", value)
	} else {
		fmt.Println("Key3 not found")
	}
}
