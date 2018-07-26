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

func TestCalcCoSize(t *testing.T) {
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
		coordinatedSize := calcCoSize(test.length, test.joint)
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
		{215, 225, 215},
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
		assert.Equal(t, test.co, cos.basic, "unexpected CO calculated")
		assert.Equal(t, test.coPlus, cos.plus, "unexpected CO+ calculated")
		assert.Equal(t, test.coMinus, cos.minus, "unexpected CO- calculated")
	}
}

func TestCanCalcResult(t *testing.T) {
	var tests = []struct {
		remainder           float64
		wholeBricks         int
		coSizeForHalfBrick  float64
		coSizeForWholeBrick float64
		result              horizontalResult
	}{
		{25, 15, 112.5, 225, horizontalResult{fullCo: 3375, halfCo: 3487.5, nfull: 15, nhalf: 15.5}},
		{215, 0, 112.5, 225, horizontalResult{fullCo: 0, halfCo: 112.5, nfull: 0, nhalf: 0.5}},
		{0, 1, 112.5, 225, horizontalResult{fullCo: 225, halfCo: 0, nfull: 1, nhalf: 0}},
		{112.5, 1, 112.5, 225, horizontalResult{fullCo: 0, halfCo: 337.5, nfull: 0, nhalf: 1.5}},
		{25, 0, 112.5, 225, horizontalResult{fullCo: 0, halfCo: 112.5, nfull: 0, nhalf: 0.5}},
	}
	for _, test := range tests {
		result := calcHorizontalResult(test.remainder, test.wholeBricks, test.coSizeForHalfBrick, test.coSizeForWholeBrick)
		assert.Equal(t, test.result, result, "unexpected result for remainder: %.2f wholeBricks: %d", test.remainder, test.wholeBricks)
	}
}

func TestIsAFullCo(t *testing.T) {
	var tests = []struct {
		remainder   float64
		wholeBricks int
		isAFullCo   bool
	}{
		{25, 15, false},
		{0, 15, true},
		{0, 1, true},
		{0, 0, false},
	}
	for _, test := range tests {
		result := isAFullCo(test.remainder, test.wholeBricks)
		assert.Equal(t, test.isAFullCo, result, "unexpected isAFullCo property for remainder: %.2f wholeBricks: %d", test.remainder, test.wholeBricks)
	}
}

func TestIsAHalfCo(t *testing.T) {
	var tests = []struct {
		remainder          float64
		coSizeForHalfBrick float64
		isAHalfCo          bool
	}{
		{25, 112.5, false},
		{115, 112.5, false},
		{112.5, 112.5, true},
	}
	for _, test := range tests {
		result := isAHalfCo(test.remainder, test.coSizeForHalfBrick)
		assert.Equal(t, test.isAHalfCo, result, "unexpected isAHalfCo property for remainder: %.2f coSizeHalfBrick: %.2f", test.remainder, test.coSizeForHalfBrick)
	}
}

func TestIsLessThanHalfACo(t *testing.T) {
	var tests = []struct {
		wholeBricks        int
		remainder          float64
		coSizeForHalfBrick float64
		isLessThanHalfACo  bool
	}{
		{15, 25, 112.5, true},
		{1, 99, 100, true},
		{15, 1, 112.5, true},
		{15, 200, 112.5, false},
		{1, 100, 100, false},
		{1, 101, 100, false},
		{1, 0, 112.5, false},
	}
	for _, test := range tests {
		result := isLessThanHalfACo(test.wholeBricks, test.remainder, test.coSizeForHalfBrick)
		assert.Equal(t, test.isLessThanHalfACo, result, "unexpected isLessThanHalfACo property for remainder: %.2f wholebricks: %d coSizeHalfBrick: %.2f", test.remainder, test.wholeBricks, test.coSizeForHalfBrick)
	}
}

func TestCalcVerticalResult(t *testing.T) {
	var tests = []struct {
		remainder      float64
		courses        int
		verticalCoSize float64
		result         verticalResult
	}{
		{25, 1, 75, verticalResult{nfirst: 1, firstCo: 75, nsecond: 2, secondCo: 150}},
		{25, 0, 75, verticalResult{nfirst: 0, firstCo: 0, nsecond: 1, secondCo: 75}},
		{0, 1, 75, verticalResult{nfirst: 1, firstCo: 75, nsecond: 0, secondCo: 0}},
		{0, 0, 75, verticalResult{nfirst: 0, firstCo: 0, nsecond: 1, secondCo: 75}},
		{0, 5, 75, verticalResult{nfirst: 5, firstCo: 375, nsecond: 0, secondCo: 0}},
	}
	for _, test := range tests {
		result := calcVerticalResult(test.remainder, test.courses, test.verticalCoSize)
		assert.Equal(t, test.result, result, "unexpected result for remainder: %.2f courses: %d", test.remainder, test.courses)
	}
}
