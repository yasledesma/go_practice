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

    t.Run("it should add a word to our instance of Dictionary", func(t *testing.T) {
        dictionary.Add("hello", "a formal greeting.")

        got, err := dictionary.Search("hello")
        want := "a formal greeting."

        if err != nil {
            t.Fatal("the word was not added to the dictionay: ", err)
        }

        assertEquals(t, got, want)
    })

    t.Run("it should not modify the definition of a word that already exist our the dictionary", func(t *testing.T) {
        got := dictionary.Add("hello", "an insult to the recipient's mother.")
        want := ErrorWordAlreadyExists

        assertError(t, got, want)
    })

    t.Run("it should modify the definition of an existing word in the dictionary", func(t *testing.T) {
        dictionary.Update("hi", "something that is said to attract someone's attention.")
        got, _ := dictionary.Search("hi")
        want := "something that is said to attract someone's attention." 

        assertEquals(t, got, want)
    })

    t.Run("it should fail trying to modify the definition of an non-existent word", func(t *testing.T) {
        got := dictionary.Update("wasap", "something that is said to attract someone's attention.")
        want := ErrorCanNotUpdateWord 

       assertError(t, got, want) 
    })

    t.Run("it should delete a word from our dictionary", func(t *testing.T) {
        dictionary.Delete("hi")

        _, got := dictionary.Search("hi")
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
