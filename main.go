package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	size := flag.Int("size", 10, "map size")
	flag.Parse()

	matrix := make([][]int, *size, *size)
	for i := range matrix {
		matrix[i] = make([]int, *size, *size)
	}

	matrix[0][0] = 1
	matrix[1][1] = 1
	matrix[1][2] = 1
	matrix[2][0] = 1
	matrix[2][1] = 1

	m := Map{
		size:   *size,
		matrix: matrix,
	}

	for {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		draw(m.matrix)
		m.step()
		time.Sleep(1 * time.Second)
	}

	fmt.Println("end")
}

type Map struct {
	size   int
	matrix [][]int
}

func (m *Map) step() {
	newmatrix := make([][]int, m.size, m.size)
	for i := range newmatrix {
		newmatrix[i] = make([]int, m.size, m.size)
	}

	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			newmatrix[i][j] = m.calculateValue(i, j)
		}
	}

	m.matrix = newmatrix
}

func (m *Map) calculateValue(x, y int) int {
	cur := m.matrix[x][y]
	filled := 0
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i < 0 || i > m.size-1 {
				continue
			}
			if j < 0 || j > m.size-1 {
				continue
			}
			if i == x && y == j {
				continue
			}
			filled += m.matrix[i][j]
		}
	}

	// fmt.Println(x, y, filled)
	if cur == 0 {
		if filled == 3 {
			return 1
		}
		return 0
	}
	// cur == 1
	if filled < 2 || filled > 3 {
		return 0
	}
	return 1
}

func draw(mtx [][]int) {
	for i := 0; i < len(mtx); i++ {
		row := mtx[i]
		for j := 0; j < len(row); j++ {
			if row[j] == 1 {
				fmt.Print("\xE2\x96\xA0")
				continue
			}
			fmt.Print("\xE2\x96\xA1")
		}
		fmt.Println("")
	}
}
