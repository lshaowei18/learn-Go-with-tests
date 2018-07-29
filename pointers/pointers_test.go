package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.CheckBalance()

		if got != want {
			t.Errorf("got : %s; want: %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error, want string) {
		t.Helper()
		if err == nil {
			t.Fatal("Didn't get an error, but wanted one")
		}

	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, 10)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{10}
		wallet.Withdraw(10)

		assertBalance(t, wallet, 0)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}
