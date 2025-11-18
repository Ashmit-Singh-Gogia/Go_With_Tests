package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		checkTest(t, Bitcoin(10), wallet)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 200}
		wallet.Withdraw(Bitcoin(100))
		checkTest(t, Bitcoin(100), wallet)
	})

	t.Run("withdraw insufficient funcs", func(t *testing.T) {
		wallet := Wallet{Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(200))
		checkTest(t, Bitcoin(100), wallet)
		checkError(t, err, ErrInsfficientFunds)

	})

}
func checkTest(t testing.TB, want Bitcoin, wallet Wallet) { //cant have a named function inside a function
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func checkError(t testing.TB, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("got no error wanted one") //Fatal terminates the test and the remaining code is skipped
	}

	if err.Error() != want { // .Error() method converts err to string for comparing
		t.Errorf("got %v want %v", err, want)
	}
}
