package mines

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"math/rand"
	"mines/structs"
	"strconv"
	"strings"
)

var Field [][]structs.Cell

var totalMoves int
var totalMines int
var totalFlags int
var totalRevealed int
var totalCells int

var Players []*structs.Player

func InitField(size int) {

	//Reset stats
	totalMoves = 0
	totalMines = 0
	totalFlags = 0
	totalRevealed = 0
	totalCells = size * size

	Field = make([][]structs.Cell, size)
	for i := range Field {
		Field[i] = make([]structs.Cell, size)
	}
}

func setFlag(coords interface{}) {
	var c structs.Coords

	err := mapstructure.Decode(coords, &c)
	if err != nil {
		return
	}

	cell := &Field[c.X][c.Y]

	if cell.Revealed {
		return
	}

	cell.Flagged = !cell.Flagged

	// Update flag count
	if cell.Flagged {
		totalFlags++
	} else {
		totalFlags--
	}

	checkWin()
	sendChanges()
}

func openCell(coords interface{}) {
	var c *structs.Coords

	err := mapstructure.Decode(coords, &c)
	if err != nil {
		return
	}

	cell := &Field[c.X][c.Y]
	if cell.Revealed {
		return
	}

	totalMoves++
	if totalMoves == 1 {
		PlantMines(c)
		CalculateCells()
		flip(c)
	} else {
		if cell.Mine {
			explode()
		} else {
			flip(c)
		}
	}

	checkWin()
	sendChanges()
}

func checkWin() {
	if (totalCells == totalRevealed+totalMines) && totalMines == totalFlags {
		InitField(20)
	}
}

func sendChanges() {
	msg := structs.ClientMsg{MsgType: "Update", MsgData: Field}
	sendToAll(&msg);
}

func explode() {
	InitField(20)
}

func flip(c *structs.Coords) {
	size := len(Field)

	cell := &Field[c.X][c.Y]
	cell.Count = cell.TrueCount
	cell.Revealed = true
	cell.Flagged = false
	totalRevealed++

	if cell.TrueCount != 0 {
		return
	}

	mutX := []int{-1, 0, 1}
	mutY := []int{-1, 0, 1}

	for _, mX := range mutX {
		for _, mY := range mutY {
			mutC := structs.Coords{X: c.X + mX, Y: c.Y + mY}
			if validCell(&mutC, size) && !Field[mutC.X][mutC.Y].Revealed {
				flip(&mutC)
			}
		}
	}
}

func validCell(c *structs.Coords, size int) bool {
	return !(c.X < 0 || c.Y < 0 || c.X == size || c.Y == size)
}

func PlantMines(c *structs.Coords) {
	for x, row := range Field {
		for y, _ := range row {
			if x == c.X && y == c.Y || surroundsCell(c, x, y) {
				continue
			}

			if rand.Intn(10) == 1 {
				Field[x][y].Mine = true
				totalMines++
			}
		}
	}
}

func CalculateCells() {
	for x, row := range Field {
		for y, cell := range row {
			if !cell.Mine {
				Field[x][y].TrueCount = cellTotal(x, y)
			}
		}
	}
}

func surroundsCell(cell *structs.Coords, x int, y int) bool {
	mutx := []int{-1, 0, 1}
	muty := []int{-1, 0, 1}
	for _, mX := range mutx {
		for _, mY := range muty {
			tempx := cell.X + mX
			tempy := cell.Y + mY
			if x == tempx && y == tempy {
				return true
			}
		}
	}
	return false
}

func cellTotal(x int, y int) int {
	size := len(Field)
	total := 0
	mutx := []int{-1, 0, 1}
	muty := []int{-1, 0, 1}

	for _, mX := range mutx {
		for _, mY := range muty {
			tempx := x + mX
			tempy := y + mY
			if tempx < 0 || tempy < 0 || tempy == size || tempx == size || (mX == 0 && mY == 0) {
				continue
			}
			if Field[tempx][tempy].Mine {
				total++
			}
		}
	}

	return total
}

func PrintField() {
	for _, row := range Field {
		printLine(len(Field))
		printRow(row)
	}
	printLine(len(Field))
}

func printLine(length int) {
	var sb strings.Builder
	for i := 0; i < length*4+1; i++ {
		sb.WriteString("-")
	}
	fmt.Println(sb.String())
}

func printRow(row []structs.Cell) {
	var sb strings.Builder
	sb.WriteString("| ")
	for _, cell := range row {
		if cell.Mine {
			sb.WriteString("M")
		} else {
			sb.WriteString(strconv.Itoa(cell.TrueCount))
		}
		sb.WriteString(" | ")
	}
	fmt.Println(sb.String())
}
