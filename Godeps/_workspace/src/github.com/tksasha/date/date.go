package date

import (
  "time"
  "strconv"
  "regexp"
)

type Date struct {
  time  time.Time
}

func New(params ...string) Date {
  var year, day int

  var month time.Month

  now := time.Now()

  mask := regexp.MustCompile(`\A(\d{4})-(\d{2})-(\d{2})\z`)

  //
  // Parse "1982-05-17" or Year
  //
  if len(params) == 1 && mask.MatchString(params[0]) {
    year, month, day = parse(mask.FindStringSubmatch(params[0]))
  } else if len(params) > 0 {
    year, _ = strconv.Atoi(params[0])
  }

  //
  // parse Month
  //
  if len(params) > 1 {
    m, _ := strconv.Atoi(params[1])

    month = time.Month(m)
  }

  //
  // parse Day
  //
  if len(params) > 2 {
    day, _ = strconv.Atoi(params[2])
  }

  //
  // Assign Default Values
  //
  if year == 0 {
    year = now.Year()
  }

  if month == 0 {
    month = now.Month()
  }

  if day == 0 {
    day = now.Day()
  }

  return Date {
    time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
  }
}

func parse(input []string) (int, time.Month, int) {
  var year, day int

  var month time.Month

  if len(input[1]) != 0 {
    year, _ = strconv.Atoi(input[1])
  }

  if len(input[2]) != 0 {
    m, _ := strconv.Atoi(input[2])

    month = time.Month(m)
  }

  if len(input[3]) != 0 {
    day, _ = strconv.Atoi(input[3])
  }

  return year, month, day
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

func (d Date) Time() time.Time {
  return d.time
}
