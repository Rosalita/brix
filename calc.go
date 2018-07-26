package main

import (
	"math"
)

// calcWholeBricksInDim calculates the number of whole bricks that will fit into a dimension
func calcWholeBricksInDim(dimension float64, coSize float64) int {
	return int(dimension / coSize)
}

// calcCoSize calculates the coordinating size from the length of a brick and size of the joint
func calcCoSize(length float64, joint float64) float64 {
	return length + joint
}

// calcRemainderFromDim calculates the remainder when dividing a dimension by a coordinating size
func calcRemainderFromDim(dimension float64, coSize float64) float64 {
	return math.Mod(dimension, coSize)
}

// calcWholeCo calculates the coordinating size for a number of whole bricks
func calcWholeCo(wholeBricks int, coordinatedSize float64) float64 {
	return float64(wholeBricks) * coordinatedSize
}

// calcHalfBrickSize calculates the size of half a brick from its length, respecting the necessary joint in the middle
func calcHalfBrickSize(length float64, joint float64) float64 {
	return (length - joint) / 2
}

// calcHalfCo calculates the coordinating size for full bricks plus half a brick
func calcHalfCo(wholeBricks int, coSizeForFullBrick float64, coSizeForHalfBrick float64) float64 {
	return (float64(wholeBricks) * coSizeForFullBrick) + coSizeForHalfBrick
}

// calcCoPlusAndMinus calculates the CO+ and CO- sizes from a coordinating size
func calcCoPlusAndMinus(co float64, joint float64) Coordination {
	cos := Coordination{
		basic: co,
		plus:  co + joint,
		minus: co - joint,
	}
	return cos
}

// isAFullCo checks if a number of whole bricks and any remainder is a coordinating size of a full brick
func isAFullCo(remainder float64, wholeBricks int) bool {
	return remainder == 0 && wholeBricks != 0
}

// isAHalfCo checks if a remainder is a coordinating size for half a brick
func isAHalfCo(remainder float64, coSizeForHalfBrick float64) bool {
	return remainder == coSizeForHalfBrick
}

// isLessThanHalfACo checks if a number of whole bricks and a remainder is less than the coordinating size for half a brick
func isLessThanHalfACo(wholeBricks int, remainder float64, coSizeForHalfBrick float64) bool {
	return remainder >= 0 && !isAFullCo(remainder, wholeBricks) && remainder < coSizeForHalfBrick
}

// calcHorizontalResult calculates from number of whole bricks and remainder the nearest full brick and half brick coordinating sizes
func calcHorizontalResult(remainder float64, wholeBricks int, coSizeForHalfBrick float64, coSizeForFullBrick float64) horizontalResult {
	result := horizontalResult{}

	if isAFullCo(remainder, wholeBricks) {
		result.nfull = wholeBricks
		result.fullCo = calcWholeCo(wholeBricks, coSizeForFullBrick)
		return result
	}

	if isAHalfCo(remainder, coSizeForHalfBrick) {
		result.nhalf = float64(wholeBricks) + 0.5
		result.halfCo = calcHalfCo(wholeBricks, coSizeForFullBrick, coSizeForHalfBrick)
		return result
	}

	if wholeBricks == 0 {
		if isLessThanHalfACo(wholeBricks, remainder, coSizeForHalfBrick) {
			result.fullCo = float64(0)
			result.nhalf = float64(0.5)
			result.halfCo = calcHalfCo(wholeBricks, coSizeForFullBrick, coSizeForHalfBrick)
			return result
		}
		result.nfull = wholeBricks +1
		result.fullCo = calcWholeCo(wholeBricks +1, coSizeForFullBrick)
		result.nhalf = float64(wholeBricks) + 0.5
		result.halfCo = calcHalfCo(wholeBricks, coSizeForFullBrick, coSizeForHalfBrick)
		return result
	}

	if isLessThanHalfACo(wholeBricks, remainder, coSizeForHalfBrick) {
		result.nfull = wholeBricks
		result.fullCo = calcWholeCo(wholeBricks, coSizeForFullBrick)
		result.nhalf = float64(wholeBricks) + 0.5
		result.halfCo = calcHalfCo(wholeBricks, coSizeForFullBrick, coSizeForHalfBrick)
		return result
	}
	result.nfull = wholeBricks + 1
	result.fullCo = calcWholeCo(wholeBricks +1, coSizeForFullBrick)
	result.nhalf = float64(wholeBricks) + 0.5
	result.halfCo = calcHalfCo(wholeBricks, coSizeForFullBrick, coSizeForHalfBrick)
	return result
}

// calcVerticalResult calculates given a remainder and number of courses the nearest vertical coordinating sizes
func calcVerticalResult(remainder float64, courses int, verticalCoSize float64) verticalResult {
	result := verticalResult{}

	if isAFullCo(remainder, courses) {
		result.nfirst = courses
		result.firstCo = float64(courses) * verticalCoSize
		return result
	}
	result.nfirst = courses
	result.firstCo = float64(courses) * verticalCoSize
	result.nsecond = courses + 1
	result.secondCo = float64(courses+1) * verticalCoSize
	return result
}
