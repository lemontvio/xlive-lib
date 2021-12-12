package xlive_lib

import (
	"sync"
)

type AVStream interface {
	LoadAll() []string
	Load(string) (AVIO, bool)
	Store(string, AVIO)
	Del(string)
}

type streams struct {
	l    sync.RWMutex
	cmap map[string]AVIO
}

func NewStreams() AVStream {
	return &streams{l: sync.RWMutex{}, cmap: map[string]AVIO{}}
}

func (s *streams) LoadAll() []string {
	s.l.RLock()
	defer s.l.RUnlock()

	var idArray []string
	for id := range s.cmap {
		idArray = append(idArray, id)
	}

	return idArray
}

func (s *streams) Load(id string) (AVIO, bool) {
	s.l.RLock()
	defer s.l.RUnlock()

	cache, ok := s.cmap[id]
	return cache, ok
}

func (s *streams) Store(id string, cache AVIO) {
	s.l.Lock()
	defer s.l.Unlock()

	if _, ok := s.cmap[id]; !ok {
		s.cmap[id] = cache
		return
	}
}

func (s *streams) Del(id string) {
	s.l.Lock()
	defer s.l.Unlock()

	if _, ok := s.cmap[id]; ok {
		s.cmap[id].Close()
		delete(s.cmap, id)
	}
}
