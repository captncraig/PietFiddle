package main

import (
	"fmt"
	"strconv"
)

//This interpreter is written for low complexity at the cost of performance
//Where possible I use fixed size local frames to simplify memory layout
//Designed for (relatively) low nesting and branching, and as few dynamic data structures as possible.
//Intent is a reference implementation that can be ported to piet itself, or other limited platforms

//Global variables. Always present on stack.
var width int
var height int
var size int
var dp int
var cc int
var y int
var x int
var tries int
var blockSize int

var program []int
var st stack

func runProgram(w int, h int, data string) {
	width = w
	height = h
	size = width * height
	program = make([]int, size)
	for i := 0; i < size; i++ {
		program[i] = int([]byte(data)[i] - 65)
	}
	dp = 0
	cc = 0
	x = 0
	y = 0
	tries = 0
	blockSize = 0

	st = stack{make([]int, 1000), 0}

	for { //MAIN LOOP
		if tries > 7 {
			return
		}
		//findExit sets x and y to exit point, and sets blockSize
		findExit()
		op := moveToNextBlock()
		if op == 1 {
			st.push(blockSize)
		}
		if op == 2 {
			fmt.Println("POP")
		}
		if op == 3 {
			fmt.Println("ADD")
		}
		if op == 4 {
			fmt.Println("SUB")
		}
		if op == 5 {
			fmt.Println("MUL")
		}
		if op == 6 {
			fmt.Println("DIV")
		}
		if op == 7 {
			fmt.Println("MOD")
		}
		if op == 8 {
			fmt.Println("NOT")
		}
		if op == 9 {
			fmt.Println("GT")
		}
		if op == 10 {
			fmt.Println("PTR")
		}
		if op == 11 {
			fmt.Println("SWT")
		}
		if op == 12 {
			v := st.pop()
			st.push(v)
			st.push(v)
		}
		if op == 13 {
			fmt.Println("ROL")
		}
		if op == 14 {
			fmt.Println("INN")
		}
		if op == 15 {
			fmt.Println("INC")
		}
		if op == 16 {
			fmt.Println("OTN")
		}
		if op == 17 {
			fmt.Println(strconv.QuoteRuneToASCII(rune(st.pop())))
		}

	}
}

//moves to square adjacent to x and y.
//returns appropriate operation if any
//If blocked, will increment dp, cc, and tries as appropriate and return noop
//If white, will navigate to next color cell and return noop
//If white caused exit, will increment tries to 8
func moveToNextBlock() int {
	var nextX, nextY int = x, y
	var blocked int = 0
	var index int = x + y*width
	var currentColor int = program[index]
	var nextColor int = 0
	var hueDiff, lightDiff int

	if dp == 0 {
		nextX += 1
	} else if dp == 1 {
		nextY += 1
	} else if dp == 2 {
		nextX -= 1
	} else {
		nextY -= 1
	}
	//check out of bounds before anything
	if nextX < 0 {
		blocked++
	}
	if nextY < 0 {
		blocked++
	}
	if nextX >= width {
		blocked++
	}
	if nextY >= height {
		blocked++
	}
	//if in bounds, check for black
	if blocked == 0 {
		index = nextX + nextY*width
		nextColor = program[index]
		if nextColor == 19 {
			blocked++
		}
	}
	//We're officially blocked. Increment stuff and abort
	if blocked > 0 {
		if tries%2 == 0 {
			cc++
			cc %= 2
		} else {
			dp++
			dp %= 4
		}
		tries++
		return 0
	}
	//we've found an exit. Reset tries
	tries = 0
	//fmt.Printf("%s,%s -> %s,%s\n", x, y, nextX, nextY)
	//fmt.Printf("%s -> %s\n", currentColor, nextColor)
	if nextColor == 18 {
		fmt.Println("WHITE!!!!!")
		panic("Implement white handling:(")
	}
	x = nextX
	y = nextY
	hueDiff = nextColor/3 - currentColor/3
	if hueDiff < 0 {
		hueDiff += 6
	}
	lightDiff = nextColor%3 - currentColor%3
	if lightDiff < 0 {
		lightDiff += 3
	}
	//fmt.Printf("%s:%s\n", hueDiff, lightDiff)
	return lightDiff + hueDiff*3
}

func findExit() {
	var maxX int = x
	var maxY int = y
	var index int = y*width + x
	var currentColor int = program[index]
	var targetColor int
	//We're gonna' use a second stack. Luckily we don't need to access the
	//execution stack while this is in scope, so an implementation could
	//use that if it needed to.
	st := stack{make([]int, 1000), 0}

	blockSize = 0
	st.push(index)
	for st.size > 0 {
		index = st.pop()
		targetColor = program[index]
		if targetColor == 50 {
			continue
		}
		if targetColor != currentColor {
			continue
		}
		//match found! This cell is in the set
		blockSize++
		program[index] = 50
		x = index % width
		y = index / width

		//Check for new exit codel
		if IsBetter(x, y, maxX, maxY) {
			maxX = x
			maxY = y
		}
		//Now add neighbors to consideration stack
		if y > 0 {
			st.push(index - width)
		}
		if y < height-1 {
			st.push(index + width)
		}
		if x > 0 {
			st.push(index - 1)
		}
		if x < width-1 {
			st.push(index + 1)
		}
	}
	for index = 0; index < size; index++ {
		if program[index] == 50 {
			program[index] = currentColor
		}
	}
	x = maxX
	y = maxY
}

func IsBetter(x int, y int, maxx int, maxy int) bool {
	var majorScore int = 0
	var minorScore int = 0
	var bestMajor int = 0
	var bestMinor int = 0
	if dp == 0 {
		majorScore = x
		bestMajor = maxx
		minorScore = -y
		bestMinor = -maxy
	} else if dp == 1 {
		majorScore = y
		bestMajor = maxy
		minorScore = x
		bestMinor = maxx
	} else if dp == 2 {
		majorScore = -x
		bestMajor = -maxx
		minorScore = y
		bestMinor = maxy
	} else {
		majorScore = -y
		bestMajor = -maxy
		minorScore = -x
		bestMinor = -maxx
	}
	if cc == 1 {
		minorScore *= -1
		bestMinor *= -1
	}
	return majorScore > bestMajor || (majorScore == bestMajor && minorScore > bestMinor)
}
