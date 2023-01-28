package integers

import (
    "testing"
    "fmt"
)

func TestAdder(t *testing.T) {
    t.Run("add two numbers", func(t *testing.T) {
        got := Add(2, 2)
        want := 4

        if got != want {
            t.Errorf("expected '%d' but got '%d'", want, got)  
        }
    })
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

