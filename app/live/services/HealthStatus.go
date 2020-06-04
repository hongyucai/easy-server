package service

import "sync"

type Sync struct {
	rw *sync.RWMutex
	health map[interface{}]interface{}
}

func New() *Sync {
	return &Sync{
		rw:new(sync.RWMutex),
		health:make(map[interface{}]interface{}),
	}
}

func Indicator(ready bool,ps *Sync)  {
	ps.rw.Lock()
	defer ps.rw.Unlock()
}

func Health(ps Sync)  map[interface{}]interface{} {
	return ps.health
}
