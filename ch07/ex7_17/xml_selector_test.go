package xmlselect

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestXMLSelectorFromFile(t *testing.T) {
	xmlFilePath, err := filepath.Abs("./xmlData")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		filename string
		elems    []string
		expect   []string
	}{
		"test.xml": {
			filepath.Join(xmlFilePath, "test.xml"),
			[]string{"attr1", "attr2", "attr3"},
			[]string{"hoge"},
		},
		"id.xml": {
			filepath.Join(xmlFilePath, "id.xml"),
			[]string{"id=1", "id=2", "id=3"},
			[]string{"hoge"},
		},
		"id_class.xml": {
			filepath.Join(xmlFilePath, "id_class.xml"),
			[]string{"id=1", "class=2", "id=2"},
			[]string{"hoge"},
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			out, err := SelectFromFile(tc.filename, tc.elems)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(out, tc.expect); diff != "" {
				t.Fatalf("ivnalid output diff = %v", diff)
			}
		})
	}
}
