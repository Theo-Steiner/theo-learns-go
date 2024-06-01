package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
	sleepPeriod    = 3 * time.Second
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(sleepPeriod)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
	configurableSleeper := &ConfigurableSleeper{5 * time.Second, time.Sleep}
	Countdown(os.Stdout, configurableSleeper)
}
