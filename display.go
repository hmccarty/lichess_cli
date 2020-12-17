package main

import (
	"fmt"
	"github.com/fatih/color"
)

var pieceToChar = [9]string{"K", "Q", "R", "B", "N", "p", "Error", "Error", "X"}

func printHeader(moveNum int) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	
	var turnLabel string
	if (moveNum % 2) == 0 {
		turnLabel = "White"
	} else {
		turnLabel = "Black"
	}

	fmt.Printf("Move #%d, %s's turn\n", moveNum, turnLabel)
}

func printFooter(result string) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~")

	fmt.Printf("%s\n", result)
}

func printBoard(board [9]uint64) {
		for i := uint8(8); i > 0; i-- {
			fmt.Printf("%d  ", i)
			for j := uint8(0); j < 8; j++ {
				var pos uint8 = (i * 8) - j - 1

				if (((board[WHITE] >> pos) & 1) == 1) {
					// Print white pieces if they exist on given square
					color.Set(color.FgBlue)
				} else if (((board[BLACK] >> pos) & 1) == 1) {
					// Print black pieces if they exist on given square
					color.Set(color.FgRed)
				}

				printPiece(board, pos)
				color.Unset()
			}
			color.Unset()
			fmt.Println()
		}
		fmt.Println("   A  B  C  D  E  F  G  H ")
}

func printPiece(board [9]uint64, pos uint8) {
	for pieceType, piece := range board {
		if ((piece >> pos) & 1 == 1) {
			fmt.Printf("%s  ", pieceToChar[pieceType])
			return
		}
	}
}