package main

import (
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTempFlag(t *testing.T) {
	testCases := map[string]struct {
		temperature   string
		expectCelsius string
	}{
		"華氏(C)から華氏": {
			temperature:   "100C",
			expectCelsius: "100 度\n",
		},
		"摂氏(F)から華氏": {
			temperature:   "100F",
			expectCelsius: "37.78 度\n",
		},
		"絶対温度(K)から華氏": {
			temperature:   "100K",
			expectCelsius: "-173.1 度\n",
		},
		"華氏(℃)から華氏": {
			temperature:   "100C",
			expectCelsius: "100 度\n",
		},
		"摂氏(°F)から華氏": {
			temperature:   "100F",
			expectCelsius: "37.78 度\n",
		},
		"絶対温度(°K)から華氏": {
			temperature:   "100K",
			expectCelsius: "-173.1 度\n",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			args := []string{"run", "temp_flag.go", "-temp", tc.temperature}

			out, err := exec.Command("go", args...).Output()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(string(out), tc.expectCelsius); diff != "" {
				t.Fatalf("invalid output diff = %v", diff)
			}
		})
	}
}
