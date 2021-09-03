package mocks

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep(d time.Duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep(time.Second)
		fmt.Fprintf(out, "%d\n", i)
	}
	sleeper.Sleep(time.Second)
	fmt.Fprint(out, finalWord)
}
