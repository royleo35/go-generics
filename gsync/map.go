package gsync

import (
	"github.com/royleo35/go-generics/tools"
	"sync"
)

type SyncMap[Key comparable, Val any] struct {
	m    map[Key]Val
	lock sync.Mutex
	sync.Map
}

func NewSyncMap[Key comparable, Val any](cap int) *SyncMap[Key, Val] {
	cap = tools.IfThen(cap <= 0, 0, cap)
	return &SyncMap[Key, Val]{
		m:    make(map[Key]Val, cap),
		lock: sync.Mutex{},
	}
}

func (s *SyncMap[Key, Val]) Store(key Key, val Val) {
	s.lock.Lock()
	s.m[key] = val
	s.lock.Unlock()
}

func (s *SyncMap[Key, Val]) Load(key Key) (val Val, ok bool) {
	s.lock.Lock()
	val, ok = s.m[key]
	s.lock.Unlock()
	return
}

func (s *SyncMap[Key, Val]) Delete(key Key) (exist bool) {
	s.lock.Lock()
	_, ok := s.m[key]
	if ok {
		delete(s.m, key)
	}
	s.lock.Unlock()
	return ok
}

// Range 遍历SyncMap，如果f函数返回终止时停止遍历
func (s *SyncMap[Key, Val]) Range(f func(Key, Val) bool) {
	s.lock.Lock()
	for k, v := range s.m {
		if !f(k, v) {
			break
		}
	}
	s.lock.Unlock()
}

func (s *SyncMap[Key, Val]) CopyMap() map[Key]Val {
	s.lock.Lock()
	res := make(map[Key]Val, len(s.m))
	for k, v := range s.m {
		res[k] = v
	}
	s.lock.Unlock()
	return res
}
