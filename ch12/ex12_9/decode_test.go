package ex12_9

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDecoding(t *testing.T) {
	testCases := map[string]struct {
		input  string
		expect []Token
		err    string
	}{
		"Symbol & Int": {
			input: `(123)`,
			expect: []Token{
				StartList{},
				Int{123},
				EndList{},
			},
			err: "",
		},
		"Symbol & String & Int": {
			input: `("abc" 123)`,
			expect: []Token{
				StartList{},
				String{"abc"},
				Int{123},
				EndList{},
			},
			err: "",
		},
		"Symbol & Symbol & Int & String": {
			input: `((123) "abc")`,
			expect: []Token{
				StartList{},
				StartList{},
				Int{123},
				EndList{},
				String{"abc"},
				EndList{},
			},
			err: "",
		},
		"unexpected token": {
			input: `(1 + 3i)`,
			expect: []Token{
				StartList{},
				Int{1},
			},
			err: "unexpected token \"+\" at L1:C5",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			decoder := NewDecoder(strings.NewReader(tc.input))
			var tokens []Token
			for {
				token, err := decoder.Token()
				t.Log(err)
				if err == io.EOF {
					break
				}
				if err != nil {
					if errDiff := cmp.Diff(err.Error(), tc.err); errDiff != "" {
						t.Errorf("decode failed, unexpected error occured: %v\n", errDiff)
					}
					break
				}
				tokens = append(tokens, token)
			}

			if diff := cmp.Diff(tokens, tc.expect); diff != "" {
				t.Errorf("unexpected output, diff = %v", diff)
			}
		})
	}
}
