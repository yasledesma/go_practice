package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "English")
        want := "Hello, Chris" 
        assertCorrectMessage(t, got, want) 
    })
    
    t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
        got := Hello("", "English")
        want := "Hello, world"
        assertCorrectMessage(t, got, want) 
    })

    t.Run("say it in Spanish", func(t *testing.T) {
        got := Hello("Yasmin", "Spanish")
        want := "Hola, Yasmin" 
        assertCorrectMessage(t, got, want) 
    })

    t.Run("say it in French", func(t *testing.T) {
            got := Hello("Agus", "French")
            want := "Bonjour, Agus" 
            assertCorrectMessage(t, got, want) 
    })

    t.Run("say it in Russian", func(t *testing.T) {
            got := Hello("Анна", "Russian")
            want := "Привет, Анна" 
            assertCorrectMessage(t, got, want) 
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
                t.Errorf("got %q, want %q.", got, want)
    }
}
