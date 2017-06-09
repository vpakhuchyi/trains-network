package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	var trainLegsTemp TrainLegsTemp
	var trainLegs TrainLegs

	xmlFile, err := os.Open("data.xml")
	if err != nil {
		fmt.Println(err)
	}
	decoder := xml.NewDecoder(xmlFile)

	defer xmlFile.Close()
	//parse xml tu struct
	for {
		var trainLegTemp TrainLegTemp
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:

			if se.Name.Local == "TrainLeg" {
				decoder.DecodeElement(&trainLegTemp, &se)
				trainLegsTemp.TrainLegs = append(trainLegsTemp.TrainLegs, trainLegTemp)
			}
		}
	}
	//copy from permanent struct to struct with time.Time for ArrivalTime and DepartureTime
	for _, v := range trainLegsTemp.TrainLegs {

		arrivalTime, _ := time.Parse("15:04:05", v.ArrivalTimeString)
		departureTime, _ := time.Parse("15:04:05", v.DepartureTimeString)

		trainLegs.TrainLegsArray = append(trainLegs.TrainLegsArray, TrainLeg{
			ArrivalStationId:   v.ArrivalStationId,
			ArrivalTime:        arrivalTime,
			DepartureStationId: v.DepartureStationId,
			DepartureTime:      departureTime,
			Price:              v.Price,
			TrainId:            v.TrainId})
	}
	chooseTrains(trainLegs)
}

func chooseTrains(t TrainLegs) {
	var departureStationId int
	var arrivalStationId int
	var criteria string

	fmt.Println("Please input departure station's ID")
	fmt.Scan(&departureStationId)

	fmt.Println("Please input arrival station's ID")
	fmt.Scanln(&arrivalStationId)

	fmt.Println("Please input main criteria")
	fmt.Scan(&criteria)

	matchedTrains := t.findTrains(arrivalStationId, departureStationId)

	sortTrains(matchedTrains, criteria)
	fmt.Println("Available trains sorted by your criteria:")

	bestVariants := matchedTrains[:3]
	printTrains(bestVariants)
}

//call needed sort for each of criteria
func sortTrains(t []TrainLeg, criteria string) {
	switch strings.ToLower(criteria) {
	case "price":
		sort.Sort(PriceSorter(t))
	case "arrival":
		sort.Sort(ArrivalTimeSorter(t))
	case "departure":
		sort.Sort(DepartureTimeSorter(t))
	}
}

func printTrains(t interface{}) {
	switch t.(type) {
	case TrainLegs:
		for _, v := range t.(TrainLegs).TrainLegsArray {
			fmt.Println("Train's id:", v.TrainId, "Price", v.Price, "Departure time:", v.DepartureTime, "Arrival time:", v.ArrivalTime)
		}
	case []TrainLeg:
		for _, v := range t.([]TrainLeg) {
			fmt.Println("Train's id:", v.TrainId, "Price", v.Price, "Departure time:", v.DepartureTime, "Arrival time:", v.ArrivalTime)
		}
	default:
		fmt.Println("You wanna print no trains")
	}
}
