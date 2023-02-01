package dictionary

import "testing"

func TestDictionary(t *testing.T) {
    dictionary := Dictionary{"hi": "an informal greeting."}

    t.Run("it should return the meaning of a word", func (t *testing.T) {
        got, _ := dictionary.Search("hi")
        want := "an informal greeting."

        assertEquals(t, got, want)
    })

    t.Run("it should return an error when passed a word that doesn't exist in the dictionary", func (t *testing.T) {
        _, got := dictionary.Search("bye")
        want := ErrorWordNotFound 

        assertError(t, got, want)
    })
}

func assertEquals(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q, want %q.", got, want)
    }
}

func assertError(t testing.TB, got, want error) {
    t.Helper()
    if got != want {
        t.Errorf("got %q, want %q.", got, want)
    }
}
