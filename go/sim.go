package main

import "fmt"

type Location struct {
	Lat float64
	Lng float64
}

func NewLocation(lat, lng float64) (Location, error) {
	if lat < -90 || lat > 90 {
		return Location{}, fmt.Errorf("invalid lat: %#v", lat)
	}

	if lng < -180 || lng > 180 {
		return Location{}, fmt.Errorf("invalid lng: %#v", lng)
	}

	location := Location{
		Lat: lat,
		Lng: lng,
	}

	return location, nil
}
