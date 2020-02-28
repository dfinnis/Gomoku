package gomoku

import (
	"fmt"
	"math"
)

// if a border or a white stone is encountered, the remaining w(k) values in that direction are all set to 1
// the scores of the four directions are combined (by addition) to make up the evaluation score

/* the weight of an empty point */
const epsilon = 2

/* if this move captures a 2-in-a-row */
// const captureTwo = 42e8

/* defend against or break a 3-in-a-row */
// const defendThree = 42e11

/* defend against or break a 4-in-a-row */
// const defendFour = 42e12

// the weights of the adjacent points of influence
// w(k+1) := 2^12, w(k+2):= 2^11, w(k+3) := 2^10, w(k+4) := 2^9
func weight(z int8) int {
	var influence float64

	switch z {
	case 1:
		influence = math.Pow(2, 12)
	case 2:
		influence = math.Pow(2, 11)
	case 3:
		influence = math.Pow(2, 10)
	case 4:
		influence = math.Pow(2, 9)
	}
	return int(influence)
}

func coordinateOnBorder(coordinate coordinate) bool {
	if coordinate.y == 0 || coordinate.y == 18 || coordinate.x == 0 || coordinate.x == 18 {
		return true
	}
	return false
}

func calcLine(evalAxis int, neighbour coordinate, goban *[19][19]position, player bool, z int8) int {
	if positionOccupied(neighbour, goban) == false { /* if neighbour is empty */
		evalAxis *= epsilon
	} else if positionOccupiedByPlayer(neighbour, goban, player) == true { /* neighbour is own stone */
		evalAxis *= weight(z)
	}
	return evalAxis
}

// calculates the influence of { ownStone, empty spaces, opponentStone, border } at each space in one direction
func lineInfluence(coordinate coordinate, goban *[19][19]position, player bool, y int8, x int8, captures *captures) int {
	var a int8
	var b int8
	// var tmp int

	evalAxis := 1
	for a = 1; a <= 4; a++ {
		neighbour := findNeighbour(coordinate, y, x, a)
		if coordinateOnGoban(neighbour) == false { //not sure this is necessary
			break
		}
		if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			evalAxis += int(a)
			break
		} else {
			evalAxis += calcLine(evalAxis, neighbour, goban, player, a)
		}
	}
	for b = -1; b >= -4; b-- {
		neighbour := findNeighbour(coordinate, y, x, b)
		if coordinateOnGoban(neighbour) == false {
			break
		}
		if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			evalAxis += int(b)
			break
		} else {
			evalAxis += calcLine(evalAxis, neighbour, goban, player, b)
		}
	}
	return evalAxis / 100
}

// chainAttackDefend returns a score for aligning 5, 4, 3, or 2 stones
func chainAttackDefend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int {
	// dumpGobanBlank(goban)
	var attack int
	var defend int

	// defend := chainLength(coordinate, goban, y, x, !player)
	opponentChain, flanked := lengthOpponentChain(coordinate, goban, y, x, player)
	switch opponentChain {
	// case 5:
	// return blockWin
	// return maxInt - 500
	case 4:
		// defend= blockWin
		defend = blockWin
		// return 42e15
	case 3:
		// if checkFlanked(coordinate, goban, y, x, !player) == false {
		if flanked == false {
			// defend = blockFree3
			defend = blockFree3
		}
		// return 42e10 + 500
	}

	playerChain, flanked := lengthPlayerChain(coordinate, goban, y, x, player)
	// attack++
	switch playerChain {
	case 5:
		// attack = align5Win
		attack = align5Win
		// return 42e14
	case 4:
		if flanked == false {
			// attack = alignFree4
			attack = alignFree4
			// return 42e3
		}
		// attack = align4
		// return align4
		// return 42e12
	case 3:
		if flanked == false {
			attack = alignFree3
		}

		// return 42e10
	}
	fmt.Printf("attack = %d, defend = %d\n", attack, defend)
	if attack > defend {
		fmt.Printf("return = attack %d\n", attack)
		return attack
	}
	fmt.Printf("return = defend %d\n", defend)
	return defend

	//  check both attack and defend, and return whichever has the greatest value
	// return (defend > attack ? defend : attack)
	// return 0
}

// evaluateMove checks for alignments/captures along each vertex for one move, and returns a score for that move
func evaluateMove(coordinate coordinate, goban *[19][19]position, player bool, captures captures) int {
	var x int8
	var y int8

	eval := 0

	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return eval
			}
			capt := captureAttackDefend(coordinate, goban, y, x, player, captures)
			if capt >= blockWin || capt <= -blockWin {
				return capt
			}
			eval += capt
			tmp := chainAttackDefend(coordinate, goban, y, x, player)
			// if tmp == (maxInt - 1000) {
			if tmp >= blockWin || tmp <= -blockWin {
				return tmp
			}
			if tmp == 0 {
				tmp = lineInfluence(coordinate, goban, player, y, x, &captures)
			}
			eval += tmp
		}
	}
	return eval
}

// tmp >= blockWin || tmp <= -blockWin

// TO DO
// need to rewrite MeasureChain2
// heuristic prioritizes blocking opponent's 4 over aligning own 5
//  if willCapture && willBeCapturedVertex == true?? How to score?
// - heuristic does not prioritize winning 10th capture, instead prioritzes 3-align
