package main

import (
  "time"
  "fmt"
  "log"

  "github.com/jinzhu/now"
)

type Time struct {
  time  time.Time
}

func NewTime(year string, month string) Time {
  var t time.Time

  var err error

  str := fmt.Sprintf("%s-%s-01", year, month)

  if t, err = now.Parse(str); err != nil {
    log.Println(err) // TODO: Use time.Now() instead
  }

  return Time { time: t }
}

func (t *Time) BeginningOfMonth() time.Time {
  return now.New(t.time).BeginningOfMonth()
}

func (t *Time) EndOfMonth() time.Time {
  return now.New(t.time).EndOfMonth()
}

func ToString(t time.Time) string {
  return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}
