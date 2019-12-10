package clock

import (
	"strconv"
	"time"
)

// Subscriptions -
type Subscriptions struct {
	Hour   []func(*Clock)
	Minute []func(*Clock)
	Second []func(*Clock)
}

// Clock -
type Clock struct {
	Hour          int
	Minute        int
	Second        int
	HasChange     bool
	Initialize    bool
	Subscriptions Subscriptions
}

func padStart(unit string) string {
	if len(unit) < 2 {
		return "0" + unit
	}

	return unit
}

// GetString -
func (c *Clock) GetString() string {
	hourStr := padStart(strconv.Itoa(c.Hour))
	minuteStr := padStart(strconv.Itoa(c.Minute))
	var timeString string
	if c.Second%2 > 0 {
		timeString = hourStr + ":" + minuteStr
	} else {
		timeString = hourStr + " " + minuteStr
	}
	return timeString
}

// SetTime -
func (c *Clock) SetTime(t time.Time) {
	hasChange := false

	if c.Hour != t.Hour() {
		c.Hour = t.Hour()
		if c.Initialize != true {
			c.Call("hour")
		}
		hasChange = true
	}

	if c.Minute != t.Minute() {
		c.Minute = t.Minute()
		if c.Initialize != true {
			c.Call("minute")
		}
		hasChange = true
	}

	if c.Second != t.Second() {
		c.Second = t.Second()
		hasChange = true
	}

	if c.Initialize {
		c.Initialize = false
	}
	c.HasChange = hasChange
	c.Call("second")
}

// On -
func (c *Clock) On(event string, cb func(*Clock)) {
	if event == "hour" {
		c.Subscriptions.Hour = append(c.Subscriptions.Hour, cb)
	} else if event == "minute" {
		c.Subscriptions.Minute = append(c.Subscriptions.Minute, cb)
	} else {
		c.Subscriptions.Second = append(c.Subscriptions.Second, cb)
	}
}

// Call -
func (c *Clock) Call(event string) {
	var events []func(*Clock)
	if event == "hour" {
		events = c.Subscriptions.Hour
	} else if event == "minute" {
		events = c.Subscriptions.Minute
	} else {
		events = c.Subscriptions.Second
	}
	for _, e := range events {
		e(c)
	}
}
