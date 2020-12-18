package types

import "time"

// var loc, err :=

func _createLocation() *time.Location {
	loc, err := time.LoadLocation("Belarus/Minsk")
	if err != nil {
		panic(err)
	}

	return loc
}

var (
	// BelarusLocation - Time location of Balarus for time functions
	BelarusLocation = _createLocation()
)
