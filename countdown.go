package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// DefaultSleeper ...
type DefaultSleeper struct{}

// Sleep ...
func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleeper ... interface for sleep func
type Sleeper interface {
	Sleep()
}

// SpySleeper ...
type SpySleeper struct {
	Calls int
}

// Sleep ...
func (s SpySleeper) Sleep() {
	s.Calls++
}

// FinalWord ...
const FinalWord = "Go"
const write = "write"
const sleep = "sleep"

// Countdown ...
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, FinalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
