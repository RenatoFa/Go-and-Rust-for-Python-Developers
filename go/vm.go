package main

import (
	"time"
)

const PerHour = 5

type VM struct {
	StartTime time.Time
	EndTime   time.Time
}

type Spot struct {
	VM
}

type Coster interface {
	Cost() float64
}

func (v VM) Cost() float64 {
	end := v.EndTime

	if end.Equal(time.Time{}) {
		end = time.Now()
	}

	d := float64(end.Sub(v.StartTime)) / float64(time.Hour)
	return d * PerHour
}

func (s Spot) Cost() float64 {
	p := s.VM.Cost()
	return p * 0.8
}

func TotalCost(vms []Coster) float64 {
	total := 0.0
	for _, vm := range vms {
		total += vm.Cost()
	}
	return total
}
