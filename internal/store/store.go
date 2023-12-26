package store

import (
	"strconv"
	"sync"
)

type Store struct {
	mu   sync.Mutex
	data map[string]string
}

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.data[key]
	return value, ok
}

func (s *Store) Incr(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.data[key]
	if !ok {
		value = "0"
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	intValue++
	s.data[key] = strconv.Itoa(intValue)
	return intValue, nil
}

func (s *Store) Decr(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.data[key]
	if !ok {
		value = "0"
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	intValue--
	s.data[key] = strconv.Itoa(intValue)
	return intValue, nil
}

func (s *Store) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}

func (s *Store) Keys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	return keys
}

func (s *Store) Flush() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(map[string]string)
}

func (s *Store) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.data)
}

func (s *Store) Exists(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[key]
	return ok
}
