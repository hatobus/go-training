package ex7_18

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestXMLParsea(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		xmlString string
		expect    string
	}{
		"簡易的なXMLファイル": {
			xmlString: "<xml><body></body></xml>",
			expect:    "<xml><body></body></xml>",
		},
		"idやclassの付いたXMLファイル": {
			xmlString: `<attr1 id="1"><attr2 class="2"><attr3 id="2">hoge</attr3><attr4>fuga</attr4></attr2></attr1>`,
			expect:    `<attr1 id="1"><attr2 class="2"><attr3 id="2">hoge</attr3><attr4>fuga</attr4></attr2></attr1>`,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			decoder := xml.NewDecoder(strings.NewReader(tc.xmlString))
			out, err := Parse(decoder)
			t.Log(out)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(out.String(), tc.expect); diff != "" {
				t.Fatalf("invalid otput diff = %v", diff)
			}
		})
	}
}
