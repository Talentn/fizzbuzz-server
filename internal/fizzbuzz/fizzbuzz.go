package fizzbuzz

import (
	"strconv"
	"fmt"
	"errors"
)

type Params struct {
	Int1 int
	Int2 int
	Limit int
	Str1 string
	Str2 string
}

const Max = 100_000

func (p Params) Validate() error {
	var errs []error
	if p.Int1 < 1 {
		errs = append(errs, fmt.Errorf("int1 must be positive"))
	}
	if p.Int2 < 1 {
		errs = append(errs, fmt.Errorf("int2 must be positive"))
	}
	if p.Limit < 1 {
		errs = append(errs, fmt.Errorf("limit must be positive"))
	}
	if p.Limit > Max {
		errs = append(errs, fmt.Errorf("limit must be less than %d", Max))
	}
	if p.Str1 == "" {
		errs = append(errs, fmt.Errorf("str1 cannot be empty"))
	}
	if p.Str2 == "" {
		errs = append(errs, fmt.Errorf("str2 cannot be empty"))
	}

	return errors.Join(errs...)
}





func Generate(params Params) []string {
	out := make([]string, params.Limit)

	for i := 1; i <= params.Limit; i++ {
		byInt1 := i%params.Int1 == 0
		byInt2 := i%params.Int2 == 0

		switch {
		case byInt1 && byInt2:
			out[i-1] = params.Str1 + params.Str2
		case byInt1:
			out[i-1] = params.Str1
		case byInt2:
			out[i-1] = params.Str2
		default:
			out[i-1] = strconv.Itoa(i)
		}
	}
	return out
}