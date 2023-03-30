package memomry

import "sync"

type safeMap struct {
	value map[string]interface{}
	sync.RWMutex
}

func (sm *safeMap) GetByKey(key string) (interface{}, bool) {
	sm.RWMutex.RLock()
	defer sm.RWMutex.RUnlock()
	value, ok := sm.value[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (sm *safeMap) Set(key string, data interface{}) {
	sm.RWMutex.Lock()
	defer sm.RWMutex.Unlock()
	sm.value[key] = data
}

type Memory struct {
	Data *safeMap
}

func (m *Memory) GetByKey(key string) (interface{}, bool) {
	return m.Data.GetByKey(key)
}

func (m *Memory) Set(key string, data interface{}) {
	m.Data.Set(key, data)
}

func NewMemory() *Memory {
	return &Memory{
		Data: &safeMap{
			value:   make(map[string]interface{}),
			RWMutex: sync.RWMutex{},
		},
	}
}
