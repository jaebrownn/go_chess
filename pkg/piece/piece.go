package piece

type Piece interface {
	validateMove(row int, col int) error
}
