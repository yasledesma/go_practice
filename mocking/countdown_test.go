package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
    Calls int
}

func TestCountdown(t *testing.T) {
    buffer := &bytes.Buffer{}
    sleeper := &SpySleeper{}

    Countdown(buffer, sleeper)

    got := buffer.String()
    want := `3
2
1
Go!` 

    if got != want {
        t.Errorf("got %q, want %q.", got, want)
    }

    if sleeper.Calls != 3 {
        t.Errorf("not enough calls to sleeper. got %d, want 3.", sleeper.Calls)
    }
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}
