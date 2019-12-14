package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Test Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Test Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})
	t.Run("Insufficient Funds", func(t *testing.T) {
		startinBalance := Bitcoin(20)
		wallet := Wallet{startinBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startinBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}
func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
func aseertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but did not wanted one")
	}
}
