package main

import (
	"bytes"
	"image"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestJpegConverter(t *testing.T) {
	srcImage, err := ioutil.ReadFile("sampled.png")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		formatOption   string
		fileNameOption string
		expectExt      string
	}{
		"オプションがつかない": {
			formatOption:   "",
			fileNameOption: "",
			expectExt:      "png",
		},
		"formatオプションが付いている場合": {
			formatOption:   "png",
			fileNameOption: "",
			expectExt:      "png",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			args := []string{"run", "jpeg_convert.go"}

			if tc.formatOption != "" {
				args = append(args, "--format", tc.formatOption)
			}

			if tc.fileNameOption != "" {
				args = append(args, "--name", tc.fileNameOption)
			}

			cmd := exec.Command("go", args...)

			var stdout bytes.Buffer
			cmd.Stdout = &stdout

			stdin, err := cmd.StdinPipe()
			if err != nil {
				t.Fatal(err)
			}

			go func() {
				defer func() {
					stdin.Close()
				}()

				io.WriteString(stdin, string(srcImage)+"\n")
			}()

			if err := cmd.Run(); err != nil {
				t.Fatal(err)
			}

			outputImageFileName := "out"

			if tc.fileNameOption != "" {
				outputImageFileName = tc.fileNameOption
			}

			outputImageFileName += "." + tc.expectExt

			f, err := os.Open(outputImageFileName)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			_, kind, err := image.Decode(f)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(kind, tc.expectExt); diff != "" {
				t.Fatalf("invalid output format, diff = %v", diff)
			}

			// teardown 作成したファイルを削除する
			if err := os.Remove(outputImageFileName); err != nil {
				t.Fatal(err)
			}
		})
	}
}
