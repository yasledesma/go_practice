package wallet

import "testing"

func TestWallet(t *testing.T) {
    assertEquals := func(t testing.TB, wallet Wallet, want Shitcoin) {
        t.Helper()
        got := wallet.Balance()
        if got != want {
            t.Errorf("got %s, want %s.", got, want)
        }
    }

    assertError := func(t testing.TB, err error) {
        t.Helper()
        if err == nil {
            t.Error("wanted an error but didn't get one")
        }
    }

    t.Run("it should make a deposit into the wallet", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Shitcoin(10))
        want := Shitcoin(10)

        assertEquals(t, wallet, want)
    }) 

    t.Run("it should make a withdrawal from the wallet", func(t *testing.T) {
        wallet := Wallet{10}
        wallet.Withdraw(Shitcoin(5))
        want := Shitcoin(5)

        assertEquals(t, wallet, want)
    })

    t.Run("it should try to withdraw an amount greater than the current balance and fail", func(t *testing.T) {
        startingBalance := Shitcoin(20)
        wallet := Wallet{startingBalance} 
        err := wallet.Withdraw(Shitcoin(50))

        assertEquals(t, wallet, startingBalance)
        assertError(t, err)
    })
}

