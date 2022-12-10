package day10

import "strings"

type crt struct {
	pixels [][]string
}

func NewCRT() *crt {
	numRows := 6
	numCols := 40

	rows := make([][]string, numRows)
	for y := 0; y < numRows; y++ {
		rows[y] = make([]string, numCols)
	}

	return &crt{
		pixels: rows,
	}
}

func (screen *crt) Draw(c *cpu) string {
	result := ""
	for y := 0; y < len(screen.pixels); y++ {
		for x := 0; x < len(screen.pixels[y]); x++ {
			currentPixel := y*len(screen.pixels[0]) + x
			cycle := currentPixel + 1

			cpuValue, ok := c.getValue(cycle)
			if !ok {
				result += "E"
			}

			if x >= cpuValue-1 && x <= cpuValue+1 {
				result += "#"
			} else {
				result += "."
			}

		}
		result += "\n"
	}

	return strings.TrimSpace(result)
}
