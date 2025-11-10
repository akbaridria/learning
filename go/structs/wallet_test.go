package main

import "testing"

func testWalletTableDriven(t *testing.T) {
	tests := []struct {
		name           string
		amountDeposit  int
		amountWithdraw int
	}{
		{"akbar-one", 500, 50},
		{"akbar-two", 100, 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := AddUser(test.name); err != nil {
				t.Fatalf("Error: expected user to register")
			}

		})
	}
}
