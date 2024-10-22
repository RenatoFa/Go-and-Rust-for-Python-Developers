package main

type Car struct {
	ID string
	Location
}

type Mover interface {
	Move(float64, float64)
}

func (location *Location) Move(lat, lng float64) {
	location.Lat = lat
	location.Lng = lng
}

func moveAll(items []Mover, lat, lng float64) {
	for _, item := range items {
		item.Move(lat, lng)
	}
}

func NewCar(id string, lat, lng float64) (Car, error) {
	location, err := NewLocation(lat, lng)

	if err != nil {
		return Car{}, err
	}

	car := Car{
		ID:       id,
		Location: location,
	}

	return car, nil
}
