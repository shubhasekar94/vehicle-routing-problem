package main

import (
	"math"
	"math/rand"
)

func GetAssignedLoadsRandom(loads map[int]Load, numDrivers int) [][]int {
	assignedLoads := make([][]int, numDrivers)
	driverIndex := 0
	for _, l := range loads {
		assignedLoads[driverIndex] = append(assignedLoads[driverIndex], l.LoadNumber)
		driverIndex = (driverIndex + 1) % numDrivers
	}
	return assignedLoads
}

func GetAssignedLoadsCluster(loads map[int]Load, numDrivers int) [][]int {
	centroids := []Point{}
	i := 0
	// set up N randomly chosen centroids
	for i < numDrivers {
		randX := float64(rand.Intn(200) - 100)
		randY := float64(rand.Intn(200) - 100)
		centroids = append(centroids, Point{randX, randY})
		i = i + 1
	}
	loadNumToMidpoint := map[int]Point{}
	// use the midpoint of a load to summarize a data point
	for _, load := range loads {
		midpoint := GetMidpoint(load.Dropoff, load.Pickup)
		loadNumToMidpoint[load.LoadNumber] = midpoint
	}
	clusters := make([][]int, numDrivers)
	numIterations := 0
	maxIterations := 5
	for numIterations < maxIterations {
		// zero out clusters from the last iteration
		for i := range clusters {
			clusters[i] = []int{}
		}
		// assign each data point to the nearest centroid
		for loadNum, midpoint := range loadNumToMidpoint {
			nearestCentroidDistance := math.MaxFloat64
			var nearestCentroidIndex int
			for i, centroid := range centroids {
				centroidDistance := GetDistance(midpoint, centroid)
				if centroidDistance < nearestCentroidDistance {
					nearestCentroidDistance = centroidDistance
					nearestCentroidIndex = i
				}
			}
			clusters[nearestCentroidIndex] = append(clusters[nearestCentroidIndex], loadNum)
		}
		// update each centroid to be the mean of all data points assigned to it
		for i, cluster := range clusters {
			sumX := float64(0)
			sumY := float64(0)
			numPoints := float64(len(cluster))
			for _, loadNum := range cluster {
				point := loadNumToMidpoint[loadNum]
				sumX = sumX + point.X
				sumY = sumY + point.Y
			}
			centroids[i] = Point{sumX / numPoints, sumY / numPoints}
		}
		numIterations = numIterations + 1
	}
	finalClusters := [][]int{}
	for _, cluster := range clusters {
		if len(cluster) > 0 {
			finalClusters = append(finalClusters, cluster)
		}
	}
	return finalClusters
}

func GetNearestNeighborRoute(loads map[int]Load, assignedLoads []int) []int {
	path := []int{}
	visited := map[int]bool{}
	for _, al := range assignedLoads {
		visited[al] = false
	}
	currentLocation := Point{0, 0}
	for len(path) < len(assignedLoads) {
		closestPickup := math.MaxFloat64
		var nextLoad Load
		// determine the next closest unvisited pickup point
		for _, al := range assignedLoads {
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
