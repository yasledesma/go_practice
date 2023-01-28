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

            if !reflect.DeepEqual(got, want) {
                t.Errorf("got %v, want %v.", got, want)
            }
    })
}

