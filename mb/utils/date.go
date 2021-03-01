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
		nowTime.UTC().Year(), nowTime.UTC().Month(), nowTime.UTC().Day(),
		nowTime.UTC().Hour(),
		nowTime.UTC().Minute(),
		nowTime.UTC().Second(),
		0, UTCLocation,
	)

	return &utcTime
}

func (msg *TimePoint) Validate() error {
	highTime := time.Now().Add(time.Second)
	highDate := time.Date(
		highTime.UTC().Year(), highTime.UTC().Month(), highTime.UTC().Day(),
		highTime.UTC().Hour(), highTime.UTC().Minute(), highTime.UTC().Second(),
		0, UTCLocation,
	)
	lowTime := time.Now().Add(time.Duration(-10) * time.Minute)
	lowDate := time.Date(
		lowTime.UTC().Year(), lowTime.UTC().Month(), lowTime.UTC().Day(),
		lowTime.UTC().Hour(), lowTime.UTC().Minute(), lowTime.UTC().Second(),
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
