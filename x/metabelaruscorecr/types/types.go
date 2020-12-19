package types

import "time"

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
