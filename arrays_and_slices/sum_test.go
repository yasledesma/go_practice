package main

import (
    "testing"
    "reflect"
)

func TestSum(t *testing.T) {
    t.Run("add an array of random size", func (t *testing.T) {
            got := Sum([]int{1, 2, 3, 4, 5, 6})
            want := 21

            if got != want {
                t.Errorf("got %d, want %d.", got, want)
            }
    })

    t.Run("sum multiple arrays", func (t *testing.T) {
            got := SumAll([]int{1, 2, 3, 4, 5, 6}, []int{54, 29})
            want := []int{21, 54+29}

            assertSliceEquality(t, got, want)
    })

    t.Run("sum multiple empty arrays", func (t *testing.T) {
            got := SumAll([]int{}, []int{})
            want := []int{0, 0}

            assertSliceEquality(t, got, want)
    })

    t.Run("sum all array tails", func (t *testing.T) {
            got := SumAllTails([]int{1, 2, 3, 4, 5, 6}, []int{1,2})
            want := []int{20, 2}

            assertSliceEquality(t, got, want)
    })

    t.Run("sum all array tails of empty arrays", func (t *testing.T) {
            got := SumAllTails([]int{}, []int{})
            want := []int{0, 0}

            assertSliceEquality(t, got, want)
    })
}

func assertSliceEquality(t testing.TB, got, want []int) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v.", got, want)
    }
}
