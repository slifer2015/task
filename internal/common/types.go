package common

import "sync"

type SafeUInt struct {
	value uint
	lock  sync.RWMutex
}

func NewSafeUInt(initialValue uint) *SafeUInt {
	return &SafeUInt{
		value: initialValue,
		lock:  sync.RWMutex{},
	}
}

func (su *SafeUInt) Get() uint {
	su.lock.RLock()
	defer su.lock.RUnlock()
	return su.value
}

func (su *SafeUInt) Increment() {
	su.lock.Lock()
	defer su.lock.Unlock()
	newValue := su.value + 1
	su.value = newValue
}

func (su *SafeUInt) Decrement() {
	su.lock.Lock()
	defer su.lock.Unlock()
	newValue := su.value - 1
	su.value = newValue
}
