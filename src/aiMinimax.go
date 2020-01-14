package gomoku

import "fmt"

// //  Alpha is the best choice which has been found so far for the maximising player.
// //  Beta is the best choice which has been found so far for the minimising player

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimaxRecursive(node *node, depth uint8, alpha int, beta int, maximizingPlayer bool) int {

	if depth == 0 {
		return node.value
	}

	generateBoards(node, node.coordinate, node.lastMove)

	fmt.Printf("parent.id = %d, parent.player = %v, parent.maximingPlayer = %v, parent.coordinate: %v, parent.value = %d\n", node.id, node.player, node.maximizingPlayer, node.coordinate, node.value) //////
	// if node.id == 357130 {
	// 	dumpGobanBlank(&node.goban)
	// }
	for i := range node.children {
		child := node.children[i]
		fmt.Printf("child.id = %d, child.player = %v, child.maximizingPlayer: %v, child.coordinate: %v, child.value = %d\n", child.id, child.player, child.maximizingPlayer, child.coordinate, child.value) //////
		// if child.id == 361550 {
		// 	dumpGobanBlank(&node.goban)
		// }
	}
	fmt.Printf("\n") //////

	if maximizingPlayer == true {
		maxValue := alpha // set maxEval to -infinity
		for idx := range node.children {
			child := node.children[idx]
			value := minimaxRecursive(child, depth-1, alpha, beta, false)
			maxValue = max(maxValue, value)
			alpha = max(alpha, value)
			if node.bestMove == nil || child.value > node.bestMove.value {
				node.bestMove = child
			}
			if beta <= alpha {
				break
			}
		}
		return maxValue
	} else {
		minValue := beta // set maxEval to +infinity
		for idx := range node.children {
			child := node.children[idx]
			value := minimaxRecursive(child, depth-1, alpha, beta, true)
			minValue = min(minValue, value)
			beta = min(beta, value)
			if node.bestMove == nil || child.value < node.bestMove.value {
				node.bestMove = child
			}
			if beta <= alpha {
				break
			}
		}
		return minValue
	}
}
