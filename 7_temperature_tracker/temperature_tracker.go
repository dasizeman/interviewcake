package main

import (
	"github.com/dasizeman/tools"
)

// DegreeMax is the maximum temperature value this class can
// work with
const DegreeMax int = 110

// TempTracker tracks some statistics about temperatures you have added
// to it
type TempTracker struct {
	max, min, numTemps, tempSum, mode int
	mean                              float64
	modeRecords                       [DegreeMax + 1]int
}

// CreateTempTracker creates a new TempTracker and returns
// a pointer to it
func CreateTempTracker() *TempTracker {
	return &TempTracker{}
}

// Insert adds a temperature to the tracker
func (tracker *TempTracker) Insert(temp int) {

	// Update the max
	tracker.max = tools.IntMax(tracker.max, temp)

	// Update the min
	tracker.min = tools.IntMin(tracker.min, temp)

	// Update the mean
	tracker.numTemps++
	tracker.tempSum += temp
	tracker.mean = float64(tracker.tempSum) / float64(tracker.numTemps)

	// Update the mode
	tracker.modeRecords[temp]++
	tracker.mode = tools.IntMax(tracker.mode, tracker.modeRecords[temp])

}

// GetMax returns the maximum temperature currently tracked
func (tracker *TempTracker) GetMax() int {
	return tracker.max
}

// GetMin returns the minimum temperature currently tracked
func (tracker *TempTracker) GetMin() int {
	return tracker.min
}

// GetMean returns the mean of all tracked values
func (tracker *TempTracker) GetMean() float64 {
	return tracker.mean
}

// GetMode returns the mode of all tracked values
func (tracker *TempTracker) GetMode() int {
	return tracker.mode
}
