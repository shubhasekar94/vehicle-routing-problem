package main

import (
	"math"
)

func GetAssignedLoads(loads map[int]Load, numDrivers int) [][]int {
	assignedLoads := make([][]int, numDrivers)
	driverIndex := 0
	for _, l := range(loads) {
		assignedLoads[driverIndex] = append(assignedLoads[driverIndex], l.LoadNumber)
		driverIndex = (driverIndex + 1) % numDrivers
	}
	return assignedLoads
}

func GetNearestNeighborRoute(loads map[int]Load, assignedLoads []int) []int {
	path := []int{}
	visited := map[int]bool{}
	for _, al := range(assignedLoads) {
		visited[al] = false
	}
	currentLocation := Point{0, 0}
	for len(path) < len(assignedLoads) {
		closestPickup := math.MaxFloat64
		var nextLoad Load
		// determine the next closest unvisited pickup point
		for _, al := range(assignedLoads) {
			if !visited[al] {
				nextPossibleLoad := loads[al]
				nextPossiblePickup := GetDistance(currentLocation, nextPossibleLoad.Pickup)
				if nextPossiblePickup < closestPickup {
					closestPickup = nextPossiblePickup
					nextLoad = nextPossibleLoad
				}
			}
		}
		// add the next closest unvisited pickup point to the path,
		// mark it as visited, and set the current location to
		// the load's dropoff location
		path = append(path, nextLoad.LoadNumber)
		visited[nextLoad.LoadNumber] = true
		currentLocation = nextLoad.Dropoff
	}
	return path
}