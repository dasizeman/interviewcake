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

		fmt.Printf("%v\n", testCase)
	}
}

func findRotationIndex(words []string) int {}

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
