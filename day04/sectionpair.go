package day04

type sectionPair struct {
	A sectionRange
	B sectionRange
}

func (p *sectionPair) FullyContains() bool {
	if p.A.Min <= p.B.Min && p.A.Max >= p.B.Max {
		return true
	}

	if p.B.Min <= p.A.Min && p.B.Max >= p.A.Max {
		return true
	}

	return false
}

func (p *sectionPair) AnyOverlap() bool {
	bOverlapsWithA := p.B.Min <= p.A.Max && p.B.Min >= p.A.Min || p.B.Max <= p.A.Max && p.B.Max >= p.A.Min
	aOverlapsWithB := p.A.Min <= p.B.Max && p.A.Min >= p.B.Min || p.A.Max <= p.B.Max && p.A.Max >= p.B.Min

	return bOverlapsWithA || aOverlapsWithB
}
