package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSHA256(t *testing.T) {
	type testData struct {
		mode   string
		input  []string
		output func([]string) string
	}

	testCases := map[string]testData{
		"sha256で動かす": {
			mode:  SHA256MODE,
			input: []string{"a", "b", "c"},
			output: func(m []string) string {
				var out string
				for _, in := range m {
					hash := sha256.Sum256([]byte(in))
					out += fmt.Sprintf("%v\n", hex.EncodeToString(hash[:]))
				}
				return out
			},
		},
		"sha384で動かす": {
			mode:  SHA384MODE,
			input: []string{"こんにちは世界", "Hello World"},
			output: func(m []string) string {
				var out string
				for _, in := range m {
					hash := sha512.Sum384([]byte(in))
					out += fmt.Sprintf("%v\n", hex.EncodeToString(hash[:]))
				}
				return out
			},
		},
		"sha512で動かす": {
			mode:  SHA512MODE,
			input: []string{"こんにちは", "Hello", "你好"},
			output: func(m []string) string {
				var out string
				for _, in := range m {
					hash := sha512.Sum512([]byte(in))
					out += fmt.Sprintf("%v\n", hex.EncodeToString(hash[:]))
				}
				return out
			},
		},
		"デフォルトはsha256": {
			mode:  "",
			input: []string{"a", "b", "c"},
			output: func(m []string) string {
				var out string
				for _, in := range m {
					hash := sha256.Sum256([]byte(in))
					out += fmt.Sprintf("%v\n", hex.EncodeToString(hash[:]))
				}
				return out
			},
		},
		"長い文字列でも対応する": {
			mode: SHA256MODE,
			input: []string{
				"死のうと思っていた。ことしの正月、よそから着物を一反もらった。お年玉としてである。着物の布地は麻であった。鼠色のこまかい縞目しまめが織りこめられていた。これは夏に着る着物であろう。夏まで生きていようと思った。",
				"親譲りの無鉄砲で子供の時から損ばかりしている。小学校に居る時分学校の二階から飛び降りて一週間程腰を抜かした事がある。",
				"えたいの知れない不吉な塊が私の心を始終圧えつけていた。焦躁と言おうか、嫌悪と言おうか――酒を飲んだあとに宿酔があるように、酒を毎日飲んでいると宿酔に相当した時期がやって来る。",
			},
			output: func(m []string) string {
				var out string
				for _, in := range m {
					hash := sha256.Sum256([]byte(in))
					out += fmt.Sprintf("%v\n", hex.EncodeToString(hash[:]))
				}
				return out
			},
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			args := []string{"run", "sha256_cmd.go"}
			if tc.mode != "" {
				args = append(args, "-type", tc.mode)
			}

			cmd := exec.Command("go", args...)

			var stdout bytes.Buffer
			cmd.Stdout = &stdout

			stdin, err := cmd.StdinPipe()
			if err != nil {
				t.Fatal(err)
			}

			go func() {
				defer stdin.Close()
				for _, in := range tc.input {
					io.WriteString(stdin, in+"\n")
				}
			}()

			if err := cmd.Run(); err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(stdout.String(), tc.output(tc.input)); diff != "" {
				t.Fatalf("invalid output, diff: %v", diff)
			}
		})
	}
}
