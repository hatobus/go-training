package ex12_6

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshalling(t *testing.T) {
	testCases := map[string]struct {
		input  interface{}
		expect string
	}{
		"complex value": {
			input:  0 + 1i,
			expect: "#C(0 1)",
		},
		"boolean true": {
			input:  true,
			expect: "t",
		},
		"boolean false": {
			input:  false,
			expect: "nil",
		},
		"slice string": {
			input:  []string{"abc", "def", "ghi"},
			expect: "(\"abc\"\n    \"def\"\n    \"ghi\")",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			marshaled, err := Marshal(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tc.expect, string(marshaled)); diff != "" {
				t.Fatalf("invalid output, diff: %v\n", diff)
			}
		})
	}
}

func TestMarshallingPrettyPrint(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	b, err := Marshal(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)
	expected := `((Title "Dr. Strangelove")
(Subtitle "How I Learned to Stop Worrying and Love the Bomb")
(Oscars ("Best Actor (Nomin.)"
        "Best Adapted Screenplay (Nomin.)"
        "Best Director (Nomin.)"
        "Best Picture (Nomin.)")))`

	if got != expected {
		t.Errorf("unexpected bytes. \nexpected: \n%v\nbut got: \n%v", expected, got)
	}
}
