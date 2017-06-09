package main

import "time"

type PriceSorter []TrainLeg

func (p PriceSorter) Len() int {
	return len(p)
}

func (p PriceSorter) Less(i, j int) bool {
	return p[i].Price < p[j].Price
}

func (p PriceSorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ArrivalTimeSorter []TrainLeg

func (a ArrivalTimeSorter) Len() int {
	return len(a)
}

func (a ArrivalTimeSorter) Less(i, j int) bool {
	return a[i].ArrivalTime.Before(a[j].ArrivalTime)
}

func (a ArrivalTimeSorter) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type DepartureTimeSorter []TrainLeg

func (d DepartureTimeSorter) Len() int {
	return len(d)
}

func (d DepartureTimeSorter) Less(i, j int) bool {
	return d[i].DepartureTime.Before(d[j].DepartureTime)
}

func (d DepartureTimeSorter) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

type TrainLegs struct {
	TrainLegsArray []TrainLeg
}
type TrainLegsTemp struct {
	TrainLegs []TrainLegTemp
}

//temporary struct for saveing permanent xml's info
type TrainLegTemp struct {
	TrainId             int     `xml:",attr"`
	DepartureStationId  int     `xml:",attr"`
	ArrivalStationId    int     `xml:",attr"`
	Price               float32 `xml:",attr"`
	ArrivalTimeString   string  `xml:",attr"`
	DepartureTimeString string  `xml:",attr"`
}

type TrainLeg struct {
	TrainId            int
	DepartureStationId int
	ArrivalStationId   int
	Price              float32
	ArrivalTime        time.Time
	DepartureTime      time.Time
}

//find all trains between stations
func (t TrainLegs) findTrains(a, d int) []TrainLeg {
	var result []TrainLeg

	for _, v := range t.TrainLegsArray {
		if v.ArrivalStationId == a && v.DepartureStationId == d {
			result = append(result, v)
		}
	}

	return result
}
