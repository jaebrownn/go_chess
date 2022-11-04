package gamestate

import (
	"fmt"
	"github.com/jaebrownn/go_chess/pkg/piece"
)

func init() {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			name := fmt.Sprintf("%c%d", 'a'+file, rank+1)
			SquareValues[name] = uint64(rank*8 + file)
		}
	}
}

type Board struct {
	State [][]*piece.Piece
}

type BitBoard uint64

var SquareValues map[string]uint64

func (b *BitBoard) Print() {
	for rank := 0; rank < 8; rank++ {
		fmt.Printf("%d| ", rank+1)
		for file := 0; file < 8; file++ {
			if b.IsSet(rank, file) {
				fmt.Print(" 1 ")
			} else {
				fmt.Print(" 0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println("  ------------------------")
	fmt.Println("    a  b  c  d  e  f  g  h")
	fmt.Println()
	fmt.Printf("BitBoard: %b\n", *b)
}

// IsSet returns true if the bit at the given rank and file is set
func (b *BitBoard) IsSet(rank int, file int) bool {
	// an unsigned 64 bit integer has 64 bits
	// we want to check the bit at the given rank and file
	// so we need to shift the 1 bit to the left by the rank and file
	// and then AND it with the bitboard
	// if the result is 0, then the bit was not set
	// if the result is 1, then the bit was set
	return (*b & (1 << (rank*8 + file))) != 0
}

// Set sets the bit at the given rank and file to 1
func (b *BitBoard) Set(rank int, file int) {
	*b |= 1 << (rank*8 + file)
}

// Clear sets the bit at the given rank and file to 0
func (b *BitBoard) Clear(rank int, file int) {
	*b &= ^(1 << (rank*8 + file))
}

// Toggle toggles the bit at the given rank and file
func (b *BitBoard) Toggle(rank int, file int) {
	*b ^= 1 << (rank*8 + file)
}
