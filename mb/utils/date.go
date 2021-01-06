package mbutils

import (
	"errors"
	"time"
)

func _createLocation() *time.Location {
	loc, err := time.LoadLocation("Europe/Minsk")
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

func (msg *Created) ValidateBasic() error {
	nowTime := time.Now()
	nowDate := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, BelarusLocation)

	if msg.CreationDt.After(nowDate) {
		return errors.New("Try to create indentity after the current time")
	}

	if nowDate.Before(*msg.CreationDt) && nowTime.After(nowDate.Add(time.Minute*5)) {
		return errors.New("Try to create idenitty that was created long ago")
	}

	return nil
}
