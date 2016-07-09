package main

import (
	"fmt"
	"log"
	"sort"
)

// MeetingTuple represents a meeting with a start and end time
type MeetingTuple struct {
	startTime int
	endTime   int
}

// MeetingList is a list of meetings
type MeetingList []MeetingTuple

// A MeetingList implements sort.Interface so we can sort them with them
// builtin O(nlogn) sort.Sort()
func (meetings MeetingList) Len() int {
	return len(meetings)
}

func (meetings MeetingList) Less(i, j int) bool {
	return meetings[i].startTime < meetings[j].startTime
}

func (meetings MeetingList) Swap(i, j int) {
	temp := meetings[i]
	meetings[i] = meetings[j]
	meetings[j] = temp
}

func main() {
	testCases := make([]MeetingList, 4)
	testCases[0] = MeetingList{{0, 1}, {3, 5}, {4, 8}, {10, 12}, {9, 10}}
	testCases[1] = MeetingList{{2, 3}, {1, 2}}
	testCases[2] = MeetingList{{2, 3}, {1, 5}}
	testCases[3] = MeetingList{{7, 9}, {1, 10}, {3, 5}, {2, 6}}

	for _, val := range testCases {
		fmt.Printf("%v\n", mergeMeetingList(val))
	}

}

func mergeMeetingList(meetings MeetingList) MeetingList {
	var inputIdx, resultIdx int
	var resultList MeetingList

	if len(meetings) < 2 {
		log.Fatalf("Need at least two meeting times")
	}

	// Sort the input list
	sort.Sort(meetings)

	resultList = append(resultList, meetings[0])

	// We look at the second item as our first merge candidate
	inputIdx++

	// We keep merging at each stage until we can't merge any
	// more or we reach the end of the input
	for inputIdx < len(meetings) {
		// Can we merge with the current item in the input?
		mergeResult := tryMergeMeetings(resultList[resultIdx], meetings[inputIdx])
		if mergeResult == nil {
			resultIdx++
			resultList = append(resultList, meetings[inputIdx])
		} else {
			resultList[resultIdx] = *mergeResult
		}
		inputIdx++
	}

	return resultList
}

func tryMergeMeetings(m1, m2 MeetingTuple) *MeetingTuple {

	// First determine when the merged meeting would end
	var latestEnd int
	if m1.endTime >= m2.endTime {
		latestEnd = m1.endTime
	} else {
		latestEnd = m2.endTime
	}

	// Check if the meetings can be merged and merge if so
	var mergedMeeting *MeetingTuple
	if m2.startTime <= m1.endTime {
		mergedMeeting = &MeetingTuple{m1.startTime, latestEnd}
	}

	return mergedMeeting
}
