package utils

import "time"

type Time struct {
	now time.Time
}

func (t *Time) Day(day int) *int64 {
	d := t.now.AddDate(0, 0, -day).Unix()
	return &d
}

func (t *Time) Month(month int) *int64 {
	d := t.now.AddDate(0, -month, 0).Unix()
	return &d
}

func (t *Time) Hour(hour int) *int64 {
	d := t.now.Add(time.Hour * time.Duration(-hour)).Unix()
	return &d
}

func (t *Time) Minute(minute int) *int64 {
	d := t.now.Add(time.Minute * time.Duration(-minute)).Unix()
	return &d
}

func (t *Time) AddDate(years int, months int, days int) *int64 {
	d := t.now.AddDate(-years, -months, -days).Unix()
	return &d
}

func (t *Time) AddTime(hour int, minute int, second int) *int64 {
	d := t.now.Add(time.Hour*time.Duration(-hour) + time.Minute*time.Duration(-minute) + time.Second*time.Duration(-second)).Unix()
	return &d
}

func CreateTime() *Time {
	return &Time{now: time.Now()}
}
