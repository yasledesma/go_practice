package iteration

import (
    "testing"
    "fmt"
)

func TestRepeat(t *testing.T) {
    t.Run("repeats a letter 5 times", func(t *testing.T) {
        got := Repeat("a", 5)
        want := "aaaaa"

        assert(t, got, want)
    })

    t.Run("fails if param times is passed a negavtive number", func(t *testing.T) {
            got := Repeat("a", -5)
            want := ""

            assert(t, got, want)
    })
}

func BenchmarkRepeat(b *testing.B) {
    for i:=0; i < b.N; i++ {
        Repeat("a", 5)
    }
}

func ExampleRepeat() {
    repeat := Repeat("a", 10)
    fmt.Println(repeat)
    // Output: aaaaaaaaaa
}

func assert(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q, want %q.", got, want)
    }
}
