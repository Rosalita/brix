package main

import (
	//"fmt"
	"math"
)

type Cos struct {
	co      float64
	coPlus  float64
	coMinus float64
}

type DimensionProperties struct {
	isAFullCo         bool
	isAHalfCo         bool
	isLessThanHalfACo bool
}

func calcWholeBricksInDim(dimension float64, coordinatedSize float64) int {
	return int(dimension / coordinatedSize)
}

func calcCoordinatedSize(length float64, joint float64) float64 {
	return length + joint
}

func calcRemainderFromDim(dimension float64, coordinatedSize float64) float64 {
	return math.Mod(dimension, coordinatedSize)
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

func checkRemainder(remainder float64, wholeBricks int, coSizeForHalfBrick float64) DimensionProperties{
	dimProps := DimensionProperties{}
	if remainder == 0 && wholeBricks != 0 {
		dimProps.isAFullCo = true
		dimProps.isAHalfCo = false
	} else {
		dimProps.isAFullCo = false
	}

	if remainder == coSizeForHalfBrick {
		dimProps.isAHalfCo = true
		dimProps.isAFullCo = false
	} else {
		dimProps.isAHalfCo = false
	}

	if remainder >= 0 && !dimProps.isAFullCo && remainder < coSizeForHalfBrick {
		dimProps.isLessThanHalfACo = true
	} else {
		dimProps.isLessThanHalfACo = false
	}

	return dimProps
}