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
	// BelarusLocation - Time location of Balarus for time functions
	BelarusLocation = _createLocation()
)

type Created struct {
	CreationDt *time.Time
}

func CreateCurrentDate() *time.Time {
	nowTime := time.Now()
	nowDate := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, BelarusLocation)

	return &nowDate
}

func (msg *Created) ValidateBasic() error {
	nowTime := time.Now()
	nowDate := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, BelarusLocation)

	if msg.CreationDt.After(nowDate) {
		return errors.New(fmt.Sprintf(
			"Try to create indentity after the current time n:%s a:%s",
			nowDate, msg.CreationDt,
		))
	}

	if nowDate.Before(*msg.CreationDt) && nowTime.After(nowDate.Add(time.Minute*5)) {
		return errors.New("Try to create identity that was created long ago")
	}

	return nil
}
