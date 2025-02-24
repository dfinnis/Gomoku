package gomoku

func removeStone(coordinate coordinate, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = false
}

func captureVertex(coordinate coordinate, g *game, y int8, x int8) {
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	three := findNeighbour(coordinate, y, x, 3)
	if positionOccupiedByOpponent(one, &g.goban, g.player) == true &&
		positionOccupiedByOpponent(two, &g.goban, g.player) == true &&
		positionOccupiedByPlayer(three, &g.goban, g.player) == true {
		removeStone(one, &g.goban)
		removeStone(two, &g.goban)
		if g.player == false {
			g.capture0 += 2
		} else {
			g.capture1 += 2
		}
		g.gui.capturedPositions = append(g.gui.capturedPositions, one, two)
	}
}

func capture(coordinate coordinate, g *game) {
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				captureVertex(coordinate, g, y, x)
			}
		}
	}
}

// captureVertexTheory is used to remove stones in AI generated boards
func captureVertexTheory(coordinate coordinate, goban *[19][19]position, player bool, y int8, x int8) {
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	three := findNeighbour(coordinate, y, x, 3)
	if positionOccupiedByOpponent(one, goban, player) == true &&
		positionOccupiedByOpponent(two, goban, player) == true &&
		positionOccupiedByPlayer(three, goban, player) == true {
		removeStone(one, goban)
		removeStone(two, goban)
	}
}

// captureTheory is used to remove stones in AI generated boards
func captureTheory(coordinate coordinate, goban *[19][19]position, player bool) {
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				captureVertexTheory(coordinate, goban, player, y, x)
			}
		}
	}
}
