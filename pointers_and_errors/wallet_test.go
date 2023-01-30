package wallet

import "testing"

func TestWallet(t *testing.T) {
    t.Run("it should make a deposit into the wallet", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Shitcoin(10))
        want := Shitcoin(10)

        assertEquals(t, wallet, want)
    }) 

    t.Run("it should make a withdrawal from the wallet", func(t *testing.T) {
        wallet := Wallet{10}
        err := wallet.Withdraw(Shitcoin(5))

        assertNoError(t, err)
        assertEquals(t, wallet, Shitcoin(5))
    })

    t.Run("it should try to withdraw an amount greater than the current balance and fail", func(t *testing.T) {
        startingBalance := Shitcoin(20)
        wallet := Wallet{startingBalance} 
        err := wallet.Withdraw(Shitcoin(50))

        assertError(t, err, ErrorInsufficientFunds)
        assertEquals(t, wallet, startingBalance)
    })
}

func assertEquals (t testing.TB, wallet Wallet, want Shitcoin) {
    t.Helper()
    got := wallet.Balance()
    if got != want {
        t.Errorf("got %s, want %s.", got, want)
    }
}

func assertError (t testing.TB, got, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("wanted an error but didn't get one")
    }

    if got != want {
        t.Errorf("got %q, want %q.", got, want)
    }
}

func assertNoError (t testing.TB, got error) {
    t.Helper()
    if got != nil {
        t.Fatal("got an error, but didn't want one.")
    }
}
