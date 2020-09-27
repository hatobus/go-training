package main

import (
	"math"
	"testing"
)

func TestValidateNumericValue(t *testing.T) {
	type numericValues struct {
		values []float64
		isValid bool
	}

	origin := []float64{1, 10, 0.1, 0.01}

	testCases := map[string]numericValues{
		"正常系の数値が入っている時": {
			values: origin,
			isValid: true,
		},
		"Infの場合": {
			values: append(origin, math.Inf(0)),
			isValid: false,
		},
		"-Infの場合": {
			values: append(origin, math.Inf(-1)),
			isValid: false,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T){
			check := validateNumericValue(tc.values...)

			if check != tc.isValid {
				t.Errorf("Invalid output, want: %v but got %v\n", tc.isValid, check)
			}
		})
	}
}
