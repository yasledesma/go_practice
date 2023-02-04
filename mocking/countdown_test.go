package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpyCountdownOperations struct {
    Calls []string
}

const sleep = "sleep"
const write = "write"

func TestCountdown(t *testing.T) {
    t.Run("it should print to stdout.", func(t *testing.T) {
        buffer := &bytes.Buffer{}
        sleeper := &SpyCountdownOperations{}

        Countdown(buffer, sleeper)

        got := buffer.String()
        want := `3
2
1
Go!` 

        if got != want {
            t.Errorf("got %q, want %q.", got, want)
        }
    }) 

    t.Run("it should sleep before every print.", func (t *testing.T) {
        spySleepPrinter := &SpyCountdownOperations{}
        Countdown(spySleepPrinter, spySleepPrinter)

        want := []string{write, sleep, write, sleep, write, sleep, write}

        if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
            t.Errorf("wanted %v calls, got %v.", want, spySleepPrinter.Calls)
        }
    })
}

func (s *SpyCountdownOperations) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

