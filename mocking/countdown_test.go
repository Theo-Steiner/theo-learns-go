package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCountdown(t *testing.T) {
	t.Run("Print output is correct & sleep called 3 times", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`
		assertEqual(t, got, want)

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})

	t.Run("Ensure sleep/print order is correct", func(t *testing.T) {
		spyCountdownOperations := &SpyCountdownOperations{}

		Countdown(spyCountdownOperations, spyCountdownOperations)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(spyCountdownOperations.Calls, want) {
			t.Errorf("wanted calls %v got %v", want, spyCountdownOperations.Calls)
		}
	})

	t.Run("test configurable sleeper", func(t *testing.T) {
		sleepTime := 5 * time.Second
		spyTime := &SpyTime{}

		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}
