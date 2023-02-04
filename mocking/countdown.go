package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownFrom = 3

type Sleeper interface {
    Sleep()
}

type DefaultSleeper struct {}

func Countdown(out io.Writer, sleeper Sleeper) {
    for i:= countdownFrom; i > 0; i-- {
	fmt.Fprintln(out, i)
	sleeper.Sleep()
    } 

    fmt.Fprint(out, finalWord)
}

func (d *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

func main() {
    sleeper := &DefaultSleeper{}
    Countdown(os.Stdout, sleeper)
}
