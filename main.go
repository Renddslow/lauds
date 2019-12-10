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

// TODO: Write some actual alarms
var alarms = []*c.Alarm{
	{
		Hour:   16,
		Minute: 12,
		Callback: func() {
			fmt.Println("🦄")
			os.Exit(1)
		},
	},
}

func main() {
	clock := c.Clock{
		Initialize: true,
	}

	clock.On("second", changeDisplay)
	clock.On("minute", c.MaybeCallAlarmsForTime(alarms))

	for {
		clock.SetTime(time.Now())
		time.Sleep(50 * time.Millisecond)
	}
}
