package day3

import "fmt"

type rucksackGroup struct {
	A rucksack
	B rucksack
	C rucksack
}

func (r *rucksackGroup) GetDuplicate() (*byte, error) {
	return GetDuplicate(r.A.GetAll(), r.B.GetAll(), r.C.GetAll())
}

func ParseInput2(input []string) ([]rucksackGroup, error) {
	rucksacks, err := ParseInput(input)
	if err != nil {
		return nil, err
	}

	if len(rucksacks)%3 != 0 {
		return nil, fmt.Errorf("number of rucksacks not divisible by 3: %d", len(rucksacks))
	}

	result := make([]rucksackGroup, len(rucksacks)/3)
	for i := 0; i < len(rucksacks); i += 3 {
		rg := rucksackGroup{
			A: rucksacks[i],
			B: rucksacks[i+1],
			C: rucksacks[i+2],
		}
		result[i/3] = rg
	}

	return result, nil

}
