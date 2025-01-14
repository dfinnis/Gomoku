package gomoku

func threeBlocked(end1 coordinate, end2 coordinate, goban *[19][19]position) bool {
	if positionUnoccupied(end1, goban) == true &&
		positionUnoccupied(end2, goban) == true {
		return false
	}
	return true
}

// checkVertexForThree returns true if it finds an unblocked FreeThree on given vertex
func checkVertexForThree(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	minusTwo := findNeighbour(coordinate, y, x, -2)
	minusOne := findNeighbour(coordinate, y, x, -1)
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	three := findNeighbour(coordinate, y, x, 3)
	four := findNeighbour(coordinate, y, x, 4)
	if positionOccupiedByPlayer(one, goban, player) == true {
		if positionOccupiedByPlayer(two, goban, player) == true &&
			threeBlocked(minusOne, three, goban) == false {
			return true
		}
		if positionOccupiedByPlayer(three, goban, player) == true &&
			threeBlocked(minusOne, four, goban) == false &&
			positionOccupiedByOpponent(two, goban, player) == false {
			return true
		}
		if y < 0 || (y == 0 && x == -1) {
			if positionOccupiedByPlayer(minusOne, goban, player) == true &&
				threeBlocked(minusTwo, two, goban) == false {
				return true
			}
		}
	}
	if positionOccupiedByPlayer(two, goban, player) == true {
		if positionOccupiedByPlayer(three, goban, player) == true &&
			threeBlocked(minusOne, four, goban) == false &&
			positionOccupiedByOpponent(one, goban, player) == false {
			return true
		}
		if positionOccupiedByPlayer(minusOne, goban, player) == true &&
			threeBlocked(minusTwo, three, goban) == false &&
			positionOccupiedByOpponent(one, goban, player) == false {
			return true
		}
	}
	return false
}

// willCaptureBool returns true if suggested move will capture
func willCaptureBool(coordinate coordinate, goban *[19][19]position, player bool) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if willCaptureDirection(coordinate, goban, y, x, player) == true {
					return true
				}
			}
		}
	}
	return false
}

// doubleThree returns true if suggested move breaks the double three rule
func doubleThree(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if willCaptureBool(coordinate, goban, player) == true {
		return false
	}
	var freeThree bool
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				foundThree := checkVertexForThree(coordinate, goban, y, x, player)
				if foundThree == true {
					if freeThree == true {
						return true
					} else {
						freeThree = true
					}
				}
			}
		}
	}
	return false
}
