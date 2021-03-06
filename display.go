package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hmccarty/gochess/goengine"
)

func printRawBitBoard(board uint64) {
	for i := 8; i > 0; i-- {
	fmt.Printf("%08b\n", uint8(board >> (8 * (i - 1))))
	}
}

func printBoard(fen string) {
	var n int = 64
	fmt.Println("   A  B  C  D  E  F  G  H ")
	fmt.Printf("%d  ", (n / 8))
	for n > 0 {	
		piece := fen[0]
		fen = fen[1:]

		switch piece {
		case 'K', 'Q', 'R', 'B', 'N', 'P':
			// Print white pieces if they exist on given square
			color.Set(color.FgBlue)
			fmt.Printf("%s  ", string(piece))
			n -= 1
		case 'k', 'q', 'r', 'b', 'n', 'p':
			// Print black pieces if they exist on given square
			color.Set(color.FgRed)
			fmt.Printf("%s  ", string(piece))
			n -= 1
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			var spaces int = int(piece - 48)
			n -= spaces
			for i := 0; i < spaces; i++ {
				fmt.Printf("X  ")
			}
		case '/':
			fmt.Printf("\n%d  ", (n / 8))
		}
		color.Unset()
	}
	fmt.Println("\n   A  B  C  D  E  F  G  H \n")
}

func printMoveList(moves []*goengine.Move) {
	fmt.Printf("Moves: ")
	for _, move := range moves {
		fmt.Printf("%s, ", move.ToString())
	}
	fmt.Println()
}