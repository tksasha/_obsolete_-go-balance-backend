package main

import (
  "time"
  "strconv"
)

type Date struct {
  time  time.Time
}

func NewDate(params ...string) Date {
  t := time.Now()

  var err error

  var year, day int

  var month time.Month

  //
  // parse Year
  //
  if len(params) > 0 {
    year, err = strconv.Atoi(params[0])

    if err != nil {
      year = t.Year()
    }
  } else {
    year = t.Year()
  }

  //
  // parse Month
  //
  if len(params) > 1 {
    var m int

    m, err = strconv.Atoi(params[1])

    if err == nil {
      month = time.Month(m)
    } else {
      month = t.Month()
    }
  } else {
    month = t.Month()
  }

  //
  // parse Day
  //
  if len(params) > 2 {
    day, err = strconv.Atoi(params[2])

    if err != nil {
      day = t.Day()
    }
  } else {
    day = t.Day()
  }

  t = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

  return Date { time: t }
}

func (d Date) BeginningOfMonth() Date {
  year, month, _ := d.time.Date()

  return Date {
    time: time.Date(year, month, 1, 0, 0, 0, 0, d.time.Location()),
  }
}

func (d Date) EndOfMonth() Date {
  year, month, _ := d.time.Date()

  month = time.Month(month + 1)

  return Date {
    time: time.Date(year, month, 1, 0, 0, 0, 0, d.time.Location()).Add(-24 * time.Hour),
  }
}

func (d Date) String() string {
  return d.time.Format("2006-01-02")
}
