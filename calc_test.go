package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcWholeBricksInDim(t *testing.T) {
	var tests = []struct {
		dimension       float64
		coordinatedSize float64
		wholebricks     int
	}{
		{450, 225, 2},
		{451, 225, 2},
		{225, 225, 1},
		{224, 225, 0},
	}
	for _, test := range tests {
		wholebricks := calcWholeBricksInDim(test.dimension, test.coordinatedSize)
		assert.Equal(t, test.wholebricks, wholebricks, "unexpected number of bricks calculated for dimension: %.2f", test.dimension)
	}
}

func TestCalcCoordinatedSize(t *testing.T) {
	var tests = []struct {
		length float64
		joint  float64
		coSize float64
	}{
		{215, 10, 225},
		{100, 1, 101},
		{102.5, 10, 112.5},
	}
	for _, test := range tests {
		coordinatedSize := calcCoordinatedSize(test.length, test.joint)
		assert.Equal(t, test.coSize, coordinatedSize, "unexpected coordinated size calculatedfor length: %.2f joint: %.2f", test.length, test.joint)
	}
}

func TestCalcRemainderFromDim(t *testing.T) {
	var tests = []struct {
		dimension       float64
		coordinatedSize float64
		remainder       float64
	}{
		{450, 225, 0},
		{451, 225, 1},
		{449, 225, 224},
	}
	for _, test := range tests {
		remainder := calcRemainderFromDim(test.dimension, test.coordinatedSize)
		assert.Equal(t, test.remainder, remainder, "unexpected remainder calculated for dimension: %.2f", test.dimension)
	}
}

func TestCalcWholeCo(t *testing.T) {

	var tests = []struct {
		numBricks int
		coSize    float64
		co        float64
	}{
		{1, 300, 300},
		{2, 225, 450},
		{15, 225, 3375},
	}
	for _, test := range tests {
		wholeCo := calcWholeCo(test.numBricks, test.coSize)
		assert.Equal(t, test.co, wholeCo, "unexpected CO calculated for %v bricks", test.numBricks)
	}
}

func TestCalcHalfCo(t *testing.T) {

	var tests = []struct {
		wholeBricks        int
		coSizeForFullBrick float64
		coSizeForHalfBrick float64
		co                 float64
	}{
		{1, 225, 112.5, 337.5},
		{2, 225, 112.5, 562.5},
	}
	for _, test := range tests {
		halfCo := calcHalfCo(test.wholeBricks, test.coSizeForFullBrick, test.coSizeForHalfBrick)
		assert.Equal(t, test.co, halfCo, "unexpected CO calculated for %d bricks", test.wholeBricks)
	}
}

func TestCanCalcHalfBrickSize(t *testing.T) {
	var tests = []struct {
		length        float64
		joint         float64
		halfBrickSize float64
	}{
		{215, 10, 102.5},
		{450, 20, 215},
		{200, 20, 90},
	}
	for _, test := range tests {
		halfBrick := calcHalfBrickSize(test.length, test.joint)
		assert.Equal(t, test.halfBrickSize, halfBrick, "unexpected halfbrick size for length: %.2f joint: %.2f", test.length, test.joint)
	}

}

func TestCanCalcCoPlusAndMinus(t *testing.T) {
	var tests = []struct {
		coSize  float64
		joint   float64
		co      float64
		coPlus  float64
		coMinus float64
	}{
		{3487.50, 10.00, 3487.50, 3497.50, 3477.50},
		{100.0, 11.00, 100.0, 111.0, 89.0},
	}
	for _, test := range tests {
		cos := calcCoPlusAndMinus(test.co, test.joint)
		assert.Equal(t, test.co, cos.co, "unexpected CO calculated")
		assert.Equal(t, test.coPlus, cos.coPlus, "unexpected CO+ calculated")
		assert.Equal(t, test.coMinus, cos.coMinus, "unexpected CO- calculated")
	}
}

func TestCanCheckRemainder(t *testing.T) {
	var tests = []struct {
		remainder  float64
		wholeBricks   int
		coSizeForHalfBrick float64
		dimProps DimensionProperties
	}{
		{25, 15, 112.5, DimensionProperties{isAFullCo: false, isAHalfCo: false, isLessThanHalfACo: true}},
	}
	for _, test := range tests {
		dimProps := checkRemainder(test.remainder, test.wholeBricks, test.coSizeForHalfBrick)
		assert.Equal(t, test.dimProps, dimProps, "dimension %.2f has unexpected properties")
	}
}
