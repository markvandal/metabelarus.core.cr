package mbutils

import (
	"errors"
	"fmt"
	"time"
)

func _createLocation() *time.Location {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	return loc
}

var (
	UTCLocation = _createLocation()
)

type TimePoint struct {
	Time *time.Time
}

func CreateCurrentTime() *time.Time {
	nowTime := time.Now()
	utcTime := time.Date(
		nowTime.Year(), nowTime.Month(), nowTime.Day(),
		nowTime.Hour(),
		nowTime.Minute(),
		nowTime.Second(),
		0, UTCLocation,
	)

	return &utcTime
}

func (msg *TimePoint) Validate() error {
	highTime := time.Now().Add(time.Second)
	highDate := time.Date(
		highTime.Year(), highTime.Month(), highTime.Day(),
		highTime.Hour(), highTime.Minute(), highTime.Second(),
		0, UTCLocation,
	)
	lowTime := time.Now().Add(time.Duration(-10) * time.Minute)
	lowDate := time.Date(
		lowTime.Year(), lowTime.Month(), lowTime.Day(),
		lowTime.Hour(), lowTime.Minute(), lowTime.Second(),
		0, UTCLocation,
	)

	if msg.Time.After(highDate) || msg.Time.Before(lowDate) {
		return errors.New(fmt.Sprintf(
			"Transaction is older than 10 minutes b:%s n:%s a:%s",
			lowDate, msg.Time, highDate,
		))
	}

	return nil
}
