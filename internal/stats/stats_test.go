package stats

import (
	"testing"
	"github.com/talentn/fizzbuzz-server/internal/fizzbuzz"
)

func TestTopEmpty(t *testing.T) {
	s := New()

	if _, _, ok := s.Top(); ok {
		t.Error("Top() ok = true on empty store, want false")
	}
}

func TestTopReturnsMostFrequent(t *testing.T) {
	s := New()
	popular := fizzbuzz.Params{Int1: 3, Int2: 5, Limit: 100, Str1: "Fizz", Str2: "Buzz"}
	other := fizzbuzz.Params{Int1: 2, Int2: 7, Limit: 100, Str1: "foo", Str2: "bar"}

	s.Record(popular)
	s.Record(other)
	s.Record(popular)

	top, hits, ok := s.Top()
	if !ok {
		t.Error("Top() ok = false, want true")
	}
	if top != popular {
		t.Errorf("Top() = %v, want %v", top, popular)
	}
	if hits != 2 {
		t.Errorf("Top() hits = %v, want 2", hits)
	}
}

