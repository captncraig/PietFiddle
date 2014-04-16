package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCount int = 0

func yes(t *testing.T, x int, y int, maxx int, maxy int) {
	assert.True(t, IsBetter(x, y, maxx, maxy), fmt.Sprintf("%s", testCount))
	testCount++
}

func no(t *testing.T, x int, y int, maxx int, maxy int) {
	assert.False(t, IsBetter(x, y, maxx, maxy), fmt.Sprintf("%s", testCount))
	testCount++
}

func TestRightLeft(t *testing.T) {
	dp = 0
	cc = 0
	testCount = 0
	//Better in major
	yes(t, 5, 5, 0, 10)
	yes(t, 5, 5, 0, 5)
	yes(t, 5, 5, 0, 0)

	//Worse in major
	no(t, 5, 5, 10, 10)
	no(t, 5, 5, 10, 5)
	no(t, 5, 5, 10, 0)

	//Equal in major
	yes(t, 5, 5, 5, 6)
	no(t, 5, 5, 5, 5)
	no(t, 5, 5, 5, 4)
}

func TestRightRight(t *testing.T) {
	dp = 0
	cc = 1
	testCount = 0
	//Better in major
	yes(t, 5, 5, 0, 10)
	yes(t, 5, 5, 0, 5)
	yes(t, 5, 5, 0, 0)

	//Worse in major
	no(t, 5, 5, 10, 10)
	no(t, 5, 5, 10, 5)
	no(t, 5, 5, 10, 0)

	//Equal in major
	no(t, 5, 5, 5, 6)
	no(t, 5, 5, 5, 5)
	yes(t, 5, 5, 5, 4)
}

func TestDownLeft(t *testing.T) {
	dp = 1
	cc = 0
	testCount = 0
	//Better in major
	yes(t, 5, 5, 0, 0)
	yes(t, 5, 5, 5, 0)
	yes(t, 5, 5, 10, 0)

	//Worse in major
	no(t, 5, 5, 0, 10)
	no(t, 5, 5, 5, 10)
	no(t, 5, 5, 10, 10)

	//Equal in major
	yes(t, 5, 5, 0, 5)
	no(t, 5, 5, 5, 5)
	no(t, 5, 5, 10, 5)
}

func TestDownRight(t *testing.T) {
	dp = 1
	cc = 1
	testCount = 0
	//Better in major
	yes(t, 5, 5, 0, 0)
	yes(t, 5, 5, 5, 0)
	yes(t, 5, 5, 10, 0)

	//Worse in major
	no(t, 5, 5, 0, 10)
	no(t, 5, 5, 5, 10)
	no(t, 5, 5, 10, 10)

	//Equal in major
	no(t, 5, 5, 0, 5)
	no(t, 5, 5, 5, 5)
	yes(t, 5, 5, 10, 5)
}

func TestLeftLeft(t *testing.T) {
	dp = 2
	cc = 0
	testCount = 0
	//Better in major
	no(t, 5, 5, 0, 10)
	no(t, 5, 5, 0, 5)
	no(t, 5, 5, 0, 0)

	//Worse in major
	yes(t, 5, 5, 10, 10)
	yes(t, 5, 5, 10, 5)
	yes(t, 5, 5, 10, 0)

	//Equal in major
	no(t, 5, 5, 5, 6)
	no(t, 5, 5, 5, 5)
	yes(t, 5, 5, 5, 4)
}

func TestLeftRight(t *testing.T) {
	dp = 2
	cc = 1
	testCount = 0
	//Better in major
	no(t, 5, 5, 0, 10)
	no(t, 5, 5, 0, 5)
	no(t, 5, 5, 0, 0)

	//Worse in major
	yes(t, 5, 5, 10, 10)
	yes(t, 5, 5, 10, 5)
	yes(t, 5, 5, 10, 0)

	//Equal in major
	yes(t, 5, 5, 5, 6)
	no(t, 5, 5, 5, 5)
	no(t, 5, 5, 5, 4)
}

func TestUpLeft(t *testing.T) {
	dp = 3
	cc = 0
	testCount = 0
	//Better in major
	no(t, 5, 5, 0, 0)
	no(t, 5, 5, 5, 0)
	no(t, 5, 5, 10, 0)

	//Worse in major
	yes(t, 5, 5, 0, 10)
	yes(t, 5, 5, 5, 10)
	yes(t, 5, 5, 10, 10)

	//Equal in major
	no(t, 5, 5, 0, 5)
	no(t, 5, 5, 5, 5)
	yes(t, 5, 5, 10, 5)
}
func TestUpRight(t *testing.T) {
	dp = 3
	cc = 1
	testCount = 0
	//Better in major
	no(t, 5, 5, 0, 0)
	no(t, 5, 5, 5, 0)
	no(t, 5, 5, 10, 0)

	//Worse in major
	yes(t, 5, 5, 0, 10)
	yes(t, 5, 5, 5, 10)
	yes(t, 5, 5, 10, 10)

	//Equal in major
	yes(t, 5, 5, 0, 5)
	no(t, 5, 5, 5, 5)
	no(t, 5, 5, 10, 5)
}
