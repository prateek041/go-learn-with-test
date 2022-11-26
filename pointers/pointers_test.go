package pointers

import (
	"testing"
)

func TestWallet(t *testing.T){

	t.Run("Testing deposit", func(t *testing.T){

		wallet := Wallet{} // empty wallet.
		wallet.Deposit(Bitcoin(10)) // depositing 10 bitcoins.
		want := Bitcoin(10)
		assertBalance(t, want, wallet)
	})
	t.Run("testing withdraw", func(t *testing.T){

		wallet := Wallet{ balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, want, wallet)
	})
	t.Run("Money over-withdrawn", func(t *testing.T){

		wallet := Wallet{ balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, Bitcoin(10), wallet)
		assertError(t, err, "Cannot withdraw, insufficient balance")
	})
}

func assertError (t testing.TB, err error, want string){
	if err == nil { // if there is no error, still the function expected to return something, so we return nil
		t.Fatal("Wanted an error but didn't get one")
	}
	if err.Error() != want{
			t.Errorf("wanted %s but got %s", want, err)
	}
}

func assertBalance (t testing.TB, want Bitcoin, wallet Wallet){
		t.Helper() // this is a helper function
		got := wallet.Balance()

		if got!=want {
			t.Errorf("got %s and wanted %s", got, want)
		}
	}