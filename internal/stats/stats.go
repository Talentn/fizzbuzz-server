package stats

import (
	"sync"
	"github.com/talentn/fizzbuzz-server/internal/fizzbuzz"
)

type Store struct {
	mu sync.Mutex
	counts map[fizzbuzz.Params]int
}

func New() *Store {
	return &Store{counts: make(map[fizzbuzz.Params]int)}
}

func (s *Store) Record(p fizzbuzz.Params) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counts[p]++
}

func (s *Store) Top() (fizzbuzz.Params, int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var top fizzbuzz.Params
	hits := 0
	found := false

	for params, count := range s.counts {
		if count > hits {
			top, hits, found = params, count, true
		}
	}
	return top, hits, found
}