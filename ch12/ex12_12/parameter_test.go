package ex12_12

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidateFields(t *testing.T) {
	type Upk struct {
		Name  string `http:"name"`
		Phone string `http:"phone" validation:"phone"`
		Card  string `http:"card" validation:"card"`
		Email string `http:"email" validation:"email"`
	}

	testCases := map[string]struct {
		req       *http.Request
		expectOut Upk
		expectErr string
	}{
		"すべて正しい情報が入っている": {
			req: &http.Request{
				Form: url.Values{
					"name":  []string{"hatobus"},
					"phone": []string{"090-1111-2222"},
					"card":  []string{"4111111111111111"},
					"email": []string{"s1240056@gmail.com"},
				},
			},
			expectOut: Upk{
				Name:  "hatobus",
				Phone: "090-1111-2222",
				Card:  "4111111111111111",
				Email: "s1240056@gmail.com",
			},
			expectErr: "",
		},
		"電話番号の値が不正": {
			req: &http.Request{
				Form: url.Values{
					"name":  []string{"hatobus"},
					"phone": []string{"0001111"},
					"card":  []string{"4111111111111111"},
					"email": []string{"s1240056@gmail.com"},
				},
			},
			expectOut: Upk{
				Name:  "hatobus",
				Phone: "0001111",
				Card:  "4111111111111111",
				Email: "s1240056@gmail.com",
			},
			expectErr: "invalid phone number",
		},
		"カードの値が不正": {
			req: &http.Request{
				Form: url.Values{
					"name":  []string{"hatobus"},
					"phone": []string{"090-1111-2222"},
					"card":  []string{""},
					"email": []string{"s1240056gmail.com"},
				},
			},
			expectOut: Upk{
				Name:  "hatobus",
				Phone: "00011112222",
				Card:  "",
				Email: "s1240056gmail.com",
			},
			expectErr: "invalid card number",
		},
		"emailの値が不正": {
			req: &http.Request{
				Form: url.Values{
					"name":  []string{"hatobus"},
					"phone": []string{"090-1111-2222"},
					"card":  []string{"4111111111111111"},
					"email": []string{"s1240056gmail.com"},
				},
			},
			expectOut: Upk{
				Name:  "hatobus",
				Phone: "090-1111-2222",
				Card:  "4111111111111111",
				Email: "s1240056gmail.com",
			},
			expectErr: "invalid email format",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			var unpack Upk
			err := Unpack(tc.req, &unpack)
			if err != nil {
				if tc.expectErr == "" {
					t.Fatalf("unexpected error occured, error = %v", err)
				}
				if diff := cmp.Diff(err.Error(), tc.expectErr); diff != "" {
					t.Fatalf("unexpected error occured, diff = %v", diff)
				}
				return
			}
			if diff := cmp.Diff(unpack, tc.expectOut); diff != "" {
				t.Errorf("unexpected output, diff = %v", diff)
			}
		})
	}
}
