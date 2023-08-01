package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		fmt.Printf("Address of balance in test is %v \n", &wallet.balance)

		// We can make Bitcoin by using syntax Bitcoin(value)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(100)}

		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(90))

	})

}
