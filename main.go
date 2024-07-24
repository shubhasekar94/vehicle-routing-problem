package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X float64
	Y float64
}

type Load struct {
	LoadNumber int
	Pickup     Point
	Dropoff    Point
}

func main() {
	loads := ParseInput()
	numLoads := len(loads)
	numDrivers := 1
	lowestCost := math.MaxFloat64
	var finalSchedules [][]int
	for numDrivers <= numLoads {
		assignedLoads := GetAssignedLoadsRandom(loads, numDrivers)
		schedules := [][]int{}
		for _, al := range(assignedLoads) {
			schedules = append(schedules, GetNearestNeighborRoute(loads, al))
		}
		currentCost, err := GetTotalCost(loads, schedules)
		if err != nil {
			numDrivers = numDrivers + 1
			continue
		}
		// log.Printf("routing %d drivers with cost %f", numDrivers, currentCost)
		if currentCost < lowestCost {
			lowestCost = currentCost
			finalSchedules = schedules
		}
		numDrivers = numDrivers + 1
	}
	log.Printf("lowest cost of %f to deliver %d loads was achieved with %d drivers", lowestCost, numLoads, len(finalSchedules))
	formattedSchedules := FormatSchedules(finalSchedules)
	fmt.Print(formattedSchedules)
}

func ParseInput() map[int]Load {
	fileName := os.Args[1]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	contentString := string(content)
	lines := strings.Split(contentString, "\n")
	loads := make(map[int]Load)
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		newLoad := Load{}
		splitLine := strings.Split(line, " ")
		loadNumber, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatal(err)
		}
		newLoad.LoadNumber = loadNumber
		newLoad.Pickup = ParsePoint(splitLine[1])
		newLoad.Dropoff = ParsePoint(splitLine[2])
		loads[loadNumber] = newLoad
	}
	return loads
}

func ParsePoint(s string) Point {
	noParens := s[1 : len(s)-1]
	coords := strings.Split(noParens, ",")
	x, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		log.Fatal(err)
	}
	return Point{x, y}
}
