package day03

import "fmt"

type compartment map[byte]bool

type rucksack struct {
	A compartment
	B compartment
}

func (r *rucksack) GetAll() compartment {
	result := make(compartment)

	for key := range r.A {
		result[key] = true
	}
	for key := range r.B {
		result[key] = true
	}
	return result
}

func (r *rucksack) GetDuplicate() (*byte, error) {
	return GetDuplicate(r.A, r.B)
}

func GetPriorityScore(b byte) (int, error) {
	inLowercaseRange := b <= byte('z') && b >= byte('a')
	inUppercaseRange := b <= byte('Z') && b >= byte('A')

	if !inLowercaseRange && !inUppercaseRange {
		return 0, fmt.Errorf("character is not in a-z or A-Z range: %c", b)
	}

	if inLowercaseRange {
		return int(b-byte('a')) + 1, nil
	}

	if inUppercaseRange {
		return int(b-byte('A')) + 27, nil
	}

	panic("unreachable")
}

func GetDuplicate(compartments ...compartment) (*byte, error) {

	if compartments == nil || len(compartments) == 0 {
		return nil, fmt.Errorf("no compartments provided")
	}

	if len(compartments) == 1 {
		return nil, fmt.Errorf("only one compartment provided")
	}

	firstCompartment := compartments[0]
	duplicates := make(compartment)
	for item, v := range firstCompartment {
		duplicates[item] = v
	}

	for i := 1; i < len(compartments); i++ {
		for potentialDuplicate := range duplicates {
			if !compartments[i][potentialDuplicate] {
				delete(duplicates, potentialDuplicate)
			}
		}
	}

	dupCount := 0
	var result *byte

	for d, v := range duplicates {
		if v {
			dupCount++
			res := d
			result = &res
		}
	}

	if result == nil {
		return nil, fmt.Errorf("no duplicate found")
	}

	if dupCount > 1 {
		return nil, fmt.Errorf("more than one duplicate found")
	}

	return result, nil
}
