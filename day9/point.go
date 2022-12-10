package day9

type Point struct {
	x, y int
}

func (p *Point) move(d direction) {
	switch d {
	case up:
		p.y++
	case down:
		p.y--
	case left:
		p.x--
	case right:
		p.x++
	}
}

func (p Point) dif(p2 Point) Point {
	return Point{
		p.x - p2.x,
		p.y - p2.y,
	}
}
