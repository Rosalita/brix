package main

import (
	"math"
)

func calcWholeBricksInDim(dimension float64, coSize float64) int {
	return int(dimension / coSize)
}

func calcCoSize(length float64, joint float64) float64 {
	return length + joint
}

func calcRemainderFromDim(dimension float64, coSize float64) float64 {
	return math.Mod(dimension, coSize)
}

func calcWholeCo(wholeBricks int, coordinatedSize float64) float64 {
	return float64(wholeBricks) * coordinatedSize
}

func calcHalfBrickSize(length float64, joint float64) float64 {
	return (length - joint) / 2
}

func calcHalfCo(wholeBricks int, coSizeForFullBrick float64, coSizeForHalfBrick float64) float64 {
	return (float64(wholeBricks) * coSizeForFullBrick) + coSizeForHalfBrick
}

func calcCoPlusAndMinus(co float64, joint float64) Cos {
	cos := Cos{
		co:      co,
		coPlus:  co + joint,
		coMinus: co - joint,
	}
	return cos
}

func isAFullCo(remainder float64, wholeBricks int) bool {
	return remainder == 0 && wholeBricks != 0
}

func isAHalfCo(remainder float64, coSizeForHalfBrick float64) bool {
	return remainder == coSizeForHalfBrick
}

func isLessThanHalfACo(wholeBricks int, remainder float64, coSizeForHalfBrick float64) bool {
	return remainder >= 0 && !isAFullCo(remainder, wholeBricks) && remainder < coSizeForHalfBrick
}

func calcResult(remainder float64, wholeBricks int, coSizeForHalfBrick float64, coSizeForFullBrick float64) result {
	result := result{}

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
		result.nfull = wholeBricks
		result.fullCo = calcWholeCo(wholeBricks, coSizeForFullBrick)
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
	result.fullCo = calcWholeCo(wholeBricks+1, coSizeForFullBrick)
	result.nhalf = float64(wholeBricks) + 0.5
	result.halfCo = calcHalfCo(wholeBricks, coSizeForFullBrick, coSizeForHalfBrick)
	return result
}

func calcVerticalResult(remainder float64, courses int, verticalCoSize float64) verticalResult {
	result := verticalResult{}
	if courses == 0 && remainder > 0 {
		result.nfirst = courses
		result.firstCo = float64(courses+1) * verticalCoSize
	}
	if isAFullCo(remainder, courses) {
		result.nfirst = courses
		result.firstCo = float64(courses) * verticalCoSize
	} else {
		result.nfirst = courses
		result.firstCo = float64(courses) * verticalCoSize
		result.nsecond = courses + 1
		result.secondCo = float64(courses+1) * verticalCoSize
	}
	return result
}
