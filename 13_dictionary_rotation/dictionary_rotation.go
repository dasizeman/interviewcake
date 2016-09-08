package main

import (
	"bufio"
	"fmt"
	"github.com/dasizeman/tools"
	"log"
	"os"
)

// Constants
const (
	TestWordsFilename    string = "words.txt"
	RotatedWordsFilename string = "rotated.txt"
	NumberTestCases      int    = 3
)

// There is a lot of extra functionality in this file to rotate arbitrary
// dictionary files that I am aware is not being used
func main() {

	// Generate a few rotations of the sample data for testing
	var testCases [][]string
	permutations := make(map[int]bool)
	for len(testCases) < NumberTestCases {
		idx := rotateFile(TestWordsFilename, RotatedWordsFilename)

		_, found := permutations[idx]
		if found {
			continue
		}
		permutations[idx] = true
		testCase, _ := tools.ReadFileToStrings(RotatedWordsFilename)
		testCases = append(testCases, testCase)

	}

	controlCase := []string{
		"asymptote",
		"babka",
		"banoffee",
		"engender",
		"karpatka",
		"othellolagkage",
		"ptolemaic",
		"retrograde",
		"supplant",
		"undulate",
		"xenoepist"}

	testCases = append(testCases, controlCase)

	for _, list := range testCases {
		idx := findRotationIndex(list)
		fmt.Printf("%v\nRotation index: %d\n", list, idx)
	}
}

func findRotationIndex(words []string) int {

	if len(words) <= 1 {
		return -1
	}

	leftBoundary := -1
	rightBoundary := len(words)

	for leftBoundary+1 < rightBoundary {
		searchIdx := leftBoundary + (rightBoundary-leftBoundary)/2

		// Check if this item is the rotation index by checking if the item
		// before it comes after it alphabetically
		if words[searchIdx-1] > words[searchIdx] {
			return searchIdx
		}

		// The rotation index is in the right half of the search region
		// if the word at the right boundary comes before this word
		// alphabetically
		indexInRightHalf := words[rightBoundary-1] < words[searchIdx]

		// The rotation index is in the left half of the search region if
		// the word at the left boundary comes after this word alphabetically
		indexInLeftHalf := words[leftBoundary+1] > words[searchIdx]

		// If both or neither of these are true then either the word list
		// is not rotated or it is not sorted alphabetically
		if (indexInLeftHalf && indexInRightHalf) ||
			(!indexInLeftHalf && !indexInRightHalf) {
			return -1
		}

		if indexInLeftHalf {
			rightBoundary = searchIdx
		} else {
			leftBoundary = searchIdx
		}
	}

	return -1

}

func tryRotateFile(inPath, outPath string) (bool, int) {
	// Open our files and create scanners

	inFile, err := os.Open(inPath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer inFile.Close()
	defer outFile.Close()

	inScanner := bufio.NewScanner(inFile)
	outWriter := bufio.NewWriter(outFile)

	inFileInfo, err := inFile.Stat()
	if err != nil {
		log.Fatalf(err.Error())
	}
	inFileSize := inFileInfo.Size()

	var bytePos int64

	fmt.Printf("Rotating file...")
	fmt.Printf("\r                ")

	// Perform the rotation in one pass of the input file
	var rotatedLines []string
	reachedRotationPoint := false
	rotationIdx := 0

	for inScanner.Scan() {
		line := inScanner.Text()

		if !reachedRotationPoint {
			rotatedLines = append(rotatedLines, line)

			// "Roll" to see if this should be the rotation point
			rand := tools.RandomInt(0, 1)

			if rand == 1 && len(rotatedLines) > 0 {
				reachedRotationPoint = true
			} else {
				rotationIdx++
			}
		} else {
			line += "\n"
			bytePos += int64(len(line))
			printProgress(bytePos, inFileSize)
			outWriter.WriteString(line)
		}

	}

	outWriter.Flush()

	// Write the lines saved from before the rotation point out to the file
	for idx, line := range rotatedLines {
		if idx < len(rotatedLines)-1 {
			line += "\n"
		}
		bytePos += int64(len(line))
		printProgress(bytePos, inFileSize)
		outWriter.WriteString(line)
	}

	outWriter.Flush()

	fmt.Printf("done.\n")
	return reachedRotationPoint, rotationIdx
}

func rotateFile(inPath, outPath string) int {
	rotated := false
	rotationIdx := 0
	for !rotated {
		rotated, rotationIdx = tryRotateFile(inPath, outPath)
	}

	return rotationIdx
}

func printProgress(currentByte, totalBytes int64) {
	progress := int(float64(currentByte) / float64(totalBytes) * 100)
	fmt.Printf("\r%d%%...", progress)
}
