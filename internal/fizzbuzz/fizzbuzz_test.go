package fizzbuzz

import (
	"slices"
	"testing"
)

func TestGenerateClassic(t *testing.T) {
	params := Params{
		Int1: 3,
		Int2: 5,
		Limit: 15,
		Str1: "Fizz",
		Str2: "Buzz",
	}

	got := Generate(params)
	want := []string{
		"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}

	if !slices.Equal(got, want) {
		t.Errorf("Generate() = %v, want %v", got, want)
	}
}

func TestGenerateCustomValues(t *testing.T) {
	params := Params{
		Int1: 2,
		Int2: 3,
		Limit: 6,
		Str1: "foo",
		Str2: "bar",
	}

	got := Generate(params)
	want := []string{"1", "foo", "bar", "foo", "5", "foobar"}

	if !slices.Equal(got, want) {
		t.Errorf("Generate() = %v, want %v", got, want)
	}
}

func TestValidateAcceptsValidParams(t *testing.T) {
	params := Params{
		Int1: 3,
		Int2: 5,
		Limit: 100,
		Str1: "Fizz",
		Str2: "Buzz",
	}

	if err := params.Validate(); err != nil {
		t.Errorf("Validate() = %v, want nil", err)
	}
}

func TestValidateRejectsInvalidParams(t *testing.T) {
	params := Params{
		Int1: 0,
		Int2: -5,
		Limit: 0,
		Str1: "",
		Str2: "",
	}

	if err := params.Validate(); err == nil {
		t.Errorf("Validate() = nil, want error")
	}
}
