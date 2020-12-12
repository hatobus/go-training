package main

import (
	"bytes"
	"io"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSHA256(t *testing.T) {
	type testData struct {
		mode   string
		input  []string
		output string
	}

	testCases := map[string]testData{
		"sha256で動かす": {
			mode:  SHA256MODE,
			input: []string{"a", "b", "c"},
			output: "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb\n" +
				"3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d\n" +
				"2e7d2c03a9507ae265ecf5b5356885a53393a2029d241394997265a1a25aefc6\n",
		},
		"sha384で動かす": {
			mode:  SHA384MODE,
			input: []string{"こんにちは世界", "Hello World"},
			output: "c1c1a6214a2c7f2050bb93a8a7fdde6a369bae96a166a3ed6c1fd25ea9d8339bca528f140472bf3c803b0e77dab9dd72\n" +
				"99514329186b2f6ae4a1329e7ee6c610a729636335174ac6b740f9028396fcc803d0e93863a7c3d90f86beee782f4f3f\n",
		},
		"sha512で動かす": {
			mode:  SHA512MODE,
			input: []string{"こんにちは", "Hello", "你好"},
			output: "bb2b0b573e976d4240fd775e3b0d8c8fcbd058d832fe451214db9d604dc7b3817f0b1b030d27488c96fc0e008228172acdd5e15c26f6543d5f48dc75d8d9a662\n" +
				"3615f80c9d293ed7402687f94b22d58e529b8cc7916f8fac7fddf7fbd5af4cf777d3d795a7a00a16bf7e7f3fb9561ee9baae480da9fe7a18769e71886b03f315\n" +
				"5232181bc0d9888f5c9746e410b4740eb461706ba5dacfbc93587cecfc8d068bac7737e92870d6745b11a25e9cd78b55f4ffc706f73cfcae5345f1b53fb8f6b5\n",
		},
		"デフォルトはsha256": {
			mode:  "",
			input: []string{"a", "b", "c"},
			output: "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb\n" +
				"3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d\n" +
				"2e7d2c03a9507ae265ecf5b5356885a53393a2029d241394997265a1a25aefc6\n",
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

			if diff := cmp.Diff(stdout.String(), tc.output); diff != "" {
				t.Fatalf("invalid output, diff: %v", diff)
			}
		})
	}
}
