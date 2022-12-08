package day8

import "fmt"

type treeGrid [][]treeHeight

func (tg treeGrid) isValid() error {
	if len(tg) == 0 {
		return nil
	}

	rowSize := len(tg[0])

	for i, row := range tg {
		if len(row) != rowSize {
			return fmt.Errorf("row %d has %d trees, expected %d", i, len(row), rowSize)
		}

		for j, tree := range row {
			if err := tree.isValid(); err != nil {
				return fmt.Errorf("row %d, tree %d is not valid: %w", i, j, err)
			}
		}
	}
	return nil
}

func (tg treeGrid) GetVisibleTrees() []treeHeight {
	result := make([]treeHeight, 0)

	maxX, ok := tg.getMaxX()
	if !ok {
		return result
	}

	maxY, ok := tg.getMaxY()
	if !ok {
		return result
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if tg.treeIsVisible(x, y) {
				tree, _ := tg.getTree(x, y)
				result = append(result, tree)
			}
		}
	}

	return result
}

func (tg treeGrid) GetScenicScores() [][]int {
	maxX, ok := tg.getMaxX()
	if !ok {
		return nil
	}

	maxY, ok := tg.getMaxY()
	if !ok {
		return nil
	}

	result := make([][]int, len(tg))
	for i := range tg {
		result[i] = make([]int, maxX+1)
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			result[y][x] = tg.getScenicScore(x, y)
		}
	}

	return result

}

func (tg treeGrid) getScenicScore(x, y int) int {
	return tg.visiblityScoreToTop(x, y) * tg.visiblityScoreToBottom(x, y) * tg.visiblityScoreToLeft(x, y) * tg.visiblityScoreToRight(x, y)
}

func (tg treeGrid) visiblityScoreToTop(x, y int) int {
	score := 0
	tree, ok := tg.getTree(x, y)
	if !ok {
		return 0
	}

	for y2 := y - 1; y2 >= 0; y2-- {
		if t2, _ := tg.getTree(x, y2); t2 < tree {
			score++
		} else {
			return score + 1
		}
	}

	return score
}

func (tg treeGrid) visiblityScoreToBottom(x, y int) int {
	score := 0
	tree, ok := tg.getTree(x, y)
	if !ok {
		return 0
	}

	maxY, ok := tg.getMaxY()
	if !ok {
		return 0
	}

	for y2 := y + 1; y2 <= maxY; y2++ {
		if t2, _ := tg.getTree(x, y2); t2 < tree {
			score++
		} else {
			return score + 1
		}
	}

	return score
}

func (tg treeGrid) visiblityScoreToLeft(x, y int) int {
	score := 0
	tree, ok := tg.getTree(x, y)
	if !ok {
		return 0
	}

	for x2 := x - 1; x2 >= 0; x2-- {
		if t2, _ := tg.getTree(x2, y); t2 < tree {
			score++
		} else {
			return score + 1
		}
	}

	return score
}

func (tg treeGrid) visiblityScoreToRight(x, y int) int {
	score := 0
	tree, ok := tg.getTree(x, y)
	if !ok {
		return 0
	}

	maxX, ok := tg.getMaxX()
	if !ok {
		return 0
	}

	for x2 := x + 1; x2 <= maxX; x2++ {
		if t2, _ := tg.getTree(x2, y); t2 < tree {
			score++
		} else {
			return score + 1
		}
	}

	return score
}

func (tg treeGrid) treeIsVisible(x, y int) bool {
	return tg.treeIsVisibleFromTop(x, y) || tg.treeIsVisibleFromBottom(x, y) || tg.treeIsVisibleFromLeft(x, y) || tg.treeIsVisibleFromRight(x, y)
}

func (tg treeGrid) treeIsVisibleFromTop(x, y int) bool {
	tree, ok := tg.getTree(x, y)
	if !ok {
		return false
	}

	for y2 := 0; y2 < y; y2++ {
		if t2, _ := tg.getTree(x, y2); t2 >= tree {
			return false
		}
	}
	return true
}

func (tg treeGrid) treeIsVisibleFromBottom(x, y int) bool {
	tree, ok := tg.getTree(x, y)
	if !ok {
		return false
	}

	maxY, ok := tg.getMaxY()
	if !ok {
		return false
	}

	for y2 := maxY; y2 > y; y2-- {
		if t2, _ := tg.getTree(x, y2); t2 >= tree {
			return false
		}
	}
	return true

}

func (tg treeGrid) treeIsVisibleFromLeft(x, y int) bool {
	tree, ok := tg.getTree(x, y)
	if !ok {
		return false
	}

	for x2 := 0; x2 < x; x2++ {
		if t2, _ := tg.getTree(x2, y); t2 >= tree {
			return false
		}
	}
	return true
}

func (tg treeGrid) treeIsVisibleFromRight(x, y int) bool {
	tree, ok := tg.getTree(x, y)
	if !ok {
		return false
	}

	maxX, ok := tg.getMaxX()
	if !ok {
		return false
	}

	for x2 := maxX; x2 > x; x2-- {
		if t2, _ := tg.getTree(x2, y); t2 >= tree {
			return false
		}
	}
	return true
}

func (tg treeGrid) getMaxY() (int, bool) {
	if len(tg) == 0 {
		return 0, false
	}

	return len(tg) - 1, true
}

func (tg treeGrid) getMaxX() (int, bool) {
	if len(tg) == 0 {
		return 0, false
	}

	rowSize := len(tg[0])
	if rowSize == 0 {
		return 0, false
	}

	return rowSize - 1, true
}

func (tg treeGrid) getTree(x, y int) (treeHeight, bool) {
	maxX, ok := tg.getMaxX()
	if !ok {
		return 0, false
	}

	if x < 0 || x > maxX {
		return 0, false
	}

	maxY, ok := tg.getMaxY()
	if !ok {
		return 0, false
	}

	if y < 0 || y > maxY {
		return 0, false
	}

	return tg[y][x], true
}

type treeHeight int

func (t treeHeight) isValid() error {
	if t <= 9 && t >= 0 {
		return nil
	}
	return fmt.Errorf("tree height must be in range (0,9), given: %d", t)
}
