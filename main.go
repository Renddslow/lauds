package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	c "github.com/Renddslow/lauds/clock"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func changeDisplay(clock *c.Clock) {
	if clock.HasChange {
		clear()
		fmt.Println(clock.GetString())
	}
}

func main() {
	clock := c.Clock{
		Initialize: true,
	}

	clock.On("second", changeDisplay)

	for {
		clock.SetTime(time.Now())
		// TODO: base this on refresh rate
		time.Sleep(50 * time.Millisecond)
	}
}
