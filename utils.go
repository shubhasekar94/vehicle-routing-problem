package main

import (
	"fmt"
	"math"
)

func GetDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow((p2.X-p1.X), 2) + math.Pow((p2.Y-p1.Y), 2))
}

func GetMidpoint(p1, p2 Point) Point {
	newX := (p1.X + p2.X) / 2
	newY := (p1.Y + p2.Y) / 2
	return Point{newX, newY}
}

func GetTotalCost(loads map[int]Load, schedules [][]int) (float64, error) {
	cost := float64(0)
	for _, schedule := range schedules {
		routeCost := GetScheduleCost(loads, schedule)
		if routeCost > float64(12*60) {
			return 0, fmt.Errorf("drive time of %f for schedule %v exceeded 12 hours", routeCost, schedule)
		}
		cost = cost + routeCost + 500
	}
	return cost, nil
}

func GetScheduleCost(loads map[int]Load, schedule []int) float64 {
	cost := float64(0)
	origin := Point{0, 0}
	lastLocation := origin
	for _, loadNumber := range schedule {
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

func FormatSchedules(schedules [][]int) string {
	output := ""
	for _, schedule := range schedules {
		line := ""
		for _, loadNum := range schedule {
			line = line + fmt.Sprintf("%d,", loadNum)
		}
		line = "[" + line[:len(line)-1] + "]\n"
		output = output + line
	}
	return output
}
