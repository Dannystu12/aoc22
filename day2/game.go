package day2

type Games []Game

func (g Games) GetTotalScore() int {
	total := 0
	for _, game := range g {
		total += game.GetScore()
	}
	return total
}

type Game struct {
	MyMove        move
	OpponentsMove move
}

func (game *Game) GetResult() result {

	playerMove, ok := game.MyMove.Validate()
	if !ok {
		return INVALID_RESULT
	}

	opponentsMove, ok := game.OpponentsMove.Validate()
	if !ok {
		return INVALID_RESULT
	}

	if playerMove == opponentsMove {
		return DRAW
	}

	switch playerMove {
	case ROCK:
		if opponentsMove == SCISSORS {
			return WIN
		}
	case PAPER:
		if opponentsMove == ROCK {
			return WIN
		}
	case SCISSORS:
		if opponentsMove == PAPER {
			return WIN
		}
	}
	return LOSS

}

func (game Game) GetScore() int {

	result := game.GetResult()

	result, ok := result.Validate()
	if !ok {
		return 0
	}

	score := result.toScore()

	myMove := game.MyMove

	score += myMove.toScore()

	return score
}
