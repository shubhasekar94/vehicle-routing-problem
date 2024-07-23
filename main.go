package main

import (
	"io/ioutil"
	"log"
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
	schedules := GetDummySchedules(loads)
	GetTotalCost(loads, schedules)
	PrintSchedules(schedules)
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
