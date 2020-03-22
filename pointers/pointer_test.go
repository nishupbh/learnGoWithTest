package pointers

import (
	"testing"
)

func TestBank(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdrawl", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("want %d got %d ", want, got)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("wanted an error but didn't get one")
	}

	if got != nil {
		t.Errorf("got err %q, want err %q ", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
