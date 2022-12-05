package day5

import (
	"fmt"
	"sort"
	"strings"
)

type cargo map[uint][]byte

func (c cargo) PerformMove(m move) error {

	fromBay, ok := c[m.from]
	if !ok {
		return fmt.Errorf("from bay does not exist: %d", m.from)
	}

	_, ok = c[m.to]
	if !ok {
		return fmt.Errorf("to bay does not exist: %d", m.to)
	}

	if m.number > uint(len(fromBay)) {
		return fmt.Errorf("move number is more than the cargo in bay, move: %d, baySize: %d", m.number, len(fromBay))
	}

	if m.to == m.from {
		return nil
	}

	moveCrates := fromBay[len(fromBay)-int(m.number):]
	sort.SliceStable(moveCrates, func(i, j int) bool {
		return i > j
	})
	c[m.to] = append(c[m.to], moveCrates...)
	c[m.from] = fromBay[:len(fromBay)-int(m.number)]

	return nil
}

func (c cargo) GetMessage() string {

	keys := make([]uint, len(c))
	for k, _ := range c {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	chars := make([]string, 0, len(c))
	for _, k := range keys {
		crates := c[k]
		if crates == nil || len(crates) == 0 {
			continue
		}

		char := crates[len(crates)-1]
		chars = append(chars, string(char))
	}

	return strings.Join(chars, "")
}
