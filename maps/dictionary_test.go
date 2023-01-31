package dictionary

import "testing"

func TestSearch(t *testing.T) {
    dictionary := map[string]string{"hi": "an informal greeting."}

    t.Run("it should return the meaning of a word", func (t *testing.T) {
        got := Search(dictionary, "hi")
        want := "an informal greeting."

        assertEquals(t, got, want)
    })

    t.Run("it should return an error when passed a word that doesn't exist in the dictionary", func (t *testing.T) {
        got := Search(dictionary, "bye")
        want := ""

        assertEquals(t, got, want)
    })
}

func assertEquals(t testing.TB, got, want string) {
    if got != want {
        t.Errorf("got %s, want %s.", got, want)
    }
}
