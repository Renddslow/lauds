package clock

// Alarm -
type Alarm struct {
	Hour     int
	Minute   int
	Callback func()
}

// MaybeCallAlarmsForTime -
func MaybeCallAlarmsForTime(alarms []*Alarm) func(clock *Clock) {
	return func(clock *Clock) {
		for _, alarm := range alarms {
			if alarm.Hour == clock.Hour && alarm.Minute == clock.Minute {
				alarm.Callback()
			}
		}
	}
}
