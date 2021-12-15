package cache

import "sync"

type inMemoryCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	stat Stat
}

func (i *inMemoryCache) Set(k string, v []byte) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	_, ok := i.c[k]; if !ok {
		i.stat.add(k, v)
	}
	i.c[k] = v
	return nil
}
func (i *inMemoryCache) Get(k string) ([]byte, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	return i.c[k], nil
}
func (i *inMemoryCache) Del(k string) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	v, ok := i.c[k]; if ok {
		i.stat.del(k, v)
	}
	delete(i.c, k)
	return nil
}
func (i *inMemoryCache) GetStat() Stat {
	return i.stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{
		c: make(map[string][]byte),
		mutex: sync.RWMutex{},
		stat: Stat{},
	}
}
