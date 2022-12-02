package day2

type move string

const (
	ROCK         move = "ROCK"
	SCISSORS     move = "SCISSORS"
	PAPER        move = "PAPER"
	INVALID_MOVE move = "Invalid"
)

func (m move) toScore() int {
	switch m {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	default:
		return 0
	}
}

func (m move) Validate() (move, bool) {
	switch m {
	case ROCK:
		return ROCK, true
	case PAPER:
		return PAPER, true
	case SCISSORS:
		return SCISSORS, true
	default:
		return INVALID_MOVE, false
	}
}

func (m move) BeatenBy() move {
	switch m {
	case ROCK:
		return PAPER
	case PAPER:
		return SCISSORS
	case SCISSORS:
		return ROCK
	default:
		return INVALID_MOVE
	}
}

func (m move) Beats() move {
	switch m {
	case ROCK:
		return SCISSORS
	case PAPER:
		return ROCK
	case SCISSORS:
		return PAPER
	default:
		return INVALID_MOVE
	}
}
