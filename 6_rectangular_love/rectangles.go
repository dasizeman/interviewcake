package main

import (
	"fmt"
)

func main() {
	a := Rectangle{
		leftX:   1,
		bottomY: 5,
		width:   10,
		height:  4}

	b := Rectangle{
		leftX:   7,
		bottomY: 7,
		width:   10,
		height:  8}

	overlapRect := findOverlapRectangle(a, b)

	if overlapRect == nil {
		fmt.Printf("No overlap.\n")
	} else {
		fmt.Printf("Overlap:\n%#v\n", overlapRect)
	}

}

func findOverlapRectangle(a, b Rectangle) (result *Rectangle) {

	result = &Rectangle{}
	// X overlap

	overlapStart, overlapWidth := findRangeOverlap(a.leftX, a.width,
		b.leftX, b.width)
	if overlapWidth == 0 {
		return nil
	}

	result.leftX = overlapStart
	result.width = overlapWidth

	// Y overlap

	overlapStart, overlapWidth = findRangeOverlap(a.bottomY, a.height,
		b.bottomY, b.height)
	if overlapWidth == 0 {
		return nil
	}
	result.bottomY = overlapStart
	result.height = overlapWidth

	return

}

func findRangeOverlap(start1, length1, start2, length2 int) (int, int) {
	var overlapStart, overlapEnd, overlapWidth int

	// Find start of overlap
	overlapStart = max(start1, start2)

	// Compute overlap end
	overlapEnd = min(start1+length1, start2+length2)

	if overlapStart >= overlapEnd {
		return 0, 0
	}

	overlapWidth = overlapEnd - overlapStart

	return overlapStart, overlapWidth

}

// Rectangle is a basic representation of a rectangle parallel
// with the x and y axes
type Rectangle struct {
	leftX, bottomY, width, height int
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
