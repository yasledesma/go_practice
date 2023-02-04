package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyCountdownOperations struct {
    Calls []string
}

type SpyTime struct {
    durationSleep time.Duration
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

func (s *SpyTime) Sleep(duration time.Duration) {
    s.durationSleep = duration
}

func TestConfigurableSleeper(t *testing.T) {
    sleepTime := 5 * time.Second

    spyTime := &SpyTime{}
    sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
    sleeper.Sleep()

    if spyTime.durationSleep != sleepTime {
        t.Errorf("should have slept for %v, but slept for %v.", sleepTime, spyTime.durationSleep)
    }
}

