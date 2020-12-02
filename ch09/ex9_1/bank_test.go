package bank

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var done = make(chan struct{})

func TestBank(t *testing.T) {
	testCases := map[string]struct {
		Alice         func()
		Bob           func()
		teller        func()
		expectBalance int
	}{
		"取引が成功する場合": {
			Alice: func() {
				Deposit(100)
				Withdraw(100)
				done <- struct{}{}
			},
			Bob: func() {
				Deposit(500)
				Withdraw(500)
				Deposit(300)
				done <- struct{}{}
			},
			teller:        teller,
			expectBalance: 300,
		},
		"残高以上を引き出そうとする": {
			Alice: func() {
				Deposit(100)
				Withdraw(200)
				done <- struct{}{}
			},
			Bob: func() {
				Deposit(50)
				done <- struct{}{}
			},
			teller:        teller,
			expectBalance: 150,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			go tc.Alice()
			go tc.Bob()

			<-done
			<-done

			balance := Balance()
			if diff := cmp.Diff(balance, tc.expectBalance); diff != "" {
				t.Errorf("incorrect result, wanat %v but got %v\n", tc.expectBalance, balance)
			}
			Withdraw(balance)
		})
	}
}
