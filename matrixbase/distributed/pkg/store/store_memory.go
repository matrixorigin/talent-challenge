package store

import (
	"sync"
)

// just for test
type memoryStore struct {
	sync.RWMutex

	db map[string][]byte
}

func newMemoryStore() (Store, error) {
	return &memoryStore{db: make(map[string][]byte)}, nil
}

func (s *memoryStore) Set(key []byte, value []byte) error {
	s.Lock()
	defer s.Unlock()

	s.db[string(key)] = value
	return nil
}

func (s *memoryStore) Get(key []byte) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()

	k := string(key)

	if v, ok := s.db[k]; ok {
		return v, nil
	}

	return nil, nil
}

func (s *memoryStore) Delete(key []byte) error {
	s.Lock()
	defer s.Unlock()

	delete(s.db, string(key))
	return nil
}
