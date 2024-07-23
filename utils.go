package main

import (
	"fmt"
	"math"
)

func GetDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow((p2.X - p1.X), 2) + math.Pow((p2.Y - p1.Y), 2))
}

func GetTotalCost(loads map[int]Load, schedules [][]int) float64 {
	cost := float64(0)
	for _, schedule := range(schedules) {
		routeCost := GetScheduleCost(loads, schedule)
		cost = cost + routeCost + 500
	}
	return cost
}

func GetScheduleCost(loads map[int]Load, schedule []int) float64 {
	cost := float64(0)
	origin := Point{0, 0}
	lastLocation := origin
	for _, loadNumber := range(schedule) {
		load := loads[loadNumber]
		pickupPoint := load.Pickup
		dropoffPoint := load.Dropoff
		cost = cost + GetDistance(lastLocation, pickupPoint)
		cost = cost + GetDistance(pickupPoint, dropoffPoint)
		lastLocation = dropoffPoint
	}
	cost = cost + GetDistance(lastLocation, origin)
	return cost
}

func PrintSchedules(schedules [][]int) {
	for _, schedule := range(schedules) {
		fmt.Println(schedule)
	}
}