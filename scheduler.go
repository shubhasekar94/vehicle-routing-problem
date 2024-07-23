package main

func GetDummySchedules(loads map[int]Load) [][]int {
	schedules := [][]int{}
	for _, l := range(loads) {
		newSchedule := []int{l.LoadNumber}
		schedules = append(schedules, newSchedule)
	}
	return schedules
}