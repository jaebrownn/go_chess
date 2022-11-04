package gamestate

type Game struct {
	State *Board
	// TODO: add these
	//Time        time.Duration
	//Increment   time.Duration
	//ElapsedTimePlayer1 *time.Timer
	//ElapsedTimePlayer2 *time.Timer
}

// NewGame creates a new game with the `board` given
func NewGame(board *Board) (*Game, error) {
	return &Game{
		State: board,
	}, nil
}
