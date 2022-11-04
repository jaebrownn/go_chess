package gamestate

import (
	"testing"
)

func TestBitBoard_Print(t *testing.T) {
	tests := []struct {
		name string
		b    BitBoard
	}{
		{
			name: "test",
			b:    0b1111111111111111000000000000000000000000000000001111111111111111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Print()
		})
	}
}
