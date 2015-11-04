package shikaku

import (
	"fmt"
	"math/rand"
	"strings"
)

type ShikakuMap struct {
	Width       int
	Height      int
	XPos        int
	YPos        int
	SubMaps     []*ShikakuMap
	BusyGrid    [][]int
	NumbersGrid [][]int
	RandXPos    int
	RandYPos    int
}

func NewShikakuMap(width, height, xpos, ypos int) *ShikakuMap {
	sMap := ShikakuMap{
		Width:    width,
		Height:   height,
		XPos:     xpos,
		YPos:     ypos,
		RandXPos: rand.Intn(width),
		RandYPos: rand.Intn(height),
	}
	sMap.Reset()
	return &sMap
}

func (m *ShikakuMap) Reset() {
	m.SubMaps = make([]*ShikakuMap, 0)
	m.BusyGrid = make([][]int, m.Height)
	m.NumbersGrid = make([][]int, m.Height)
	for i := 0; i < m.Height; i++ {
		m.BusyGrid[i] = make([]int, m.Width)
		m.NumbersGrid[i] = make([]int, m.Width)
	}
}

func (m *ShikakuMap) GenerateBlocks(amount int) error {
	// first try to place N blocks
	for generateAttempt := 0; generateAttempt < 1000; generateAttempt++ {
		hasError := false
		for i := 0; i < amount; i++ {
			addSucceed := false
			for addAttempt := 0; addAttempt < 10; addAttempt++ {
				width := 1
				height := 1
				xPos := rand.Intn(m.Width)
				yPos := rand.Intn(m.Height)

				block := NewShikakuMap(width, height, xPos, yPos)
				block = block.Grow()
				if m.BlockFits(block) {
					if err := m.AddBlock(block); err != nil {
						panic(err)
					}
					addSucceed = true
					break
				}
			}
			if !addSucceed {
				hasError = true
				break
			}
		}
		if !hasError {
			// grow the existing blocks
			for i := 0; i < 1000; i++ {
				idx := rand.Intn(amount)
				if err := m.TryToGrowBlock(idx); err == nil {
					if m.AvailableSlots() == 0 {
						return nil
					}
				}
			}
		}

		m.Reset()
	}

	return fmt.Errorf("Failed to generate a map within 1000 attempts")
}

func (m *ShikakuMap) RemoveBlock(blockIdx int) error {
	block := m.SubMaps[blockIdx]
	m.SubMaps = append(m.SubMaps[:blockIdx], m.SubMaps[blockIdx+1:]...)
	for y := 0; y < block.Height; y++ {
		for x := 0; x < block.Width; x++ {
			m.BusyGrid[y+block.YPos][x+block.XPos]--
		}
	}
	m.NumbersGrid[block.YPos+block.RandYPos][block.XPos+block.RandXPos] = 0
	return nil
}

func (m *ShikakuMap) TryToGrowBlock(blockIdx int) error {
	block := m.SubMaps[blockIdx]
	newBlock := block.Grow()
	if err := m.RemoveBlock(blockIdx); err != nil {
		return err
	}

	if m.BlockFits(newBlock) {
		m.AddBlock(newBlock)
		return nil
	}

	// add back the old block
	m.AddBlock(block)
	return fmt.Errorf("nothing changed")
}

func (m *ShikakuMap) AvailableSlots() int {
	available := 0
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.BusyGrid[y][x] == 0 {
				available++
			}
		}
	}
	return available
}

func (m *ShikakuMap) BlockFits(block *ShikakuMap) bool {
	if block.Size() < 2 {
		return false
	}

	if block.XPos < 0 || block.YPos < 0 {
		return false
	}
	if block.XPos+block.Width > m.Width || block.YPos+block.Height > m.Height {
		return false
	}
	for y := 0; y < block.Height; y++ {
		for x := 0; x < block.Width; x++ {
			if m.BusyGrid[block.YPos+y][block.XPos+x] > 0 {
				return false
			}
		}
	}
	return true
}

func (m *ShikakuMap) Grow() *ShikakuMap {
	biggerMap := NewShikakuMap(m.Width, m.Height, m.XPos, m.YPos)
	switch rand.Intn(4) {
	case 0:
		biggerMap.Width++
	case 1:
		biggerMap.XPos--
	case 2:
		biggerMap.Height++
	case 3:
		biggerMap.YPos--
	}
	return biggerMap
}

func (m *ShikakuMap) AddBlock(block *ShikakuMap) error {
	if !m.BlockFits(block) {
		return fmt.Errorf("The block does not fit in the current map")
	}
	m.SubMaps = append(m.SubMaps, block)
	for y := 0; y < block.Height; y++ {
		for x := 0; x < block.Width; x++ {
			m.BusyGrid[y+block.YPos][x+block.XPos]++
			m.NumbersGrid[block.YPos+block.RandYPos][block.XPos+block.RandXPos] = block.Size()
		}
	}
	return nil
}

func (m *ShikakuMap) Size() int {
	return m.Width * m.Height
}

func (m *ShikakuMap) Blocks() []*ShikakuMap {
	blocks := []*ShikakuMap{}

	if len(m.SubMaps) > 0 {
		for _, subMap := range m.SubMaps {
			blocks = append(blocks, subMap.Blocks()...)
		}
	} else {
		blocks = append(blocks, m)
	}
	return blocks
}

func (m *ShikakuMap) BlockDetailString() string {
	return fmt.Sprintf("%d %d-%d(%d) %d-%d(%d)", m.Size(), m.XPos, m.Width, m.XPos+m.RandXPos, m.YPos, m.Height, m.YPos+m.RandYPos)
}

func (m *ShikakuMap) BlockString() string {
	return fmt.Sprintf("%d %d %d", m.Size(), m.XPos+m.RandXPos, m.YPos+m.RandYPos)
}

func (m *ShikakuMap) String() string {
	lines := []string{
		fmt.Sprintf("T %d %d", m.Width, m.Height),
	}

	for _, block := range m.Blocks() {
		lines = append(lines, block.BlockString())
	}

	return strings.Join(lines, "\n")
}

func (m *ShikakuMap) DrawMap() string {
	output := []string{}

	output = append(output, fmt.Sprintf("+%s", strings.Repeat("---+", m.Width)))
	for _, line := range m.NumbersGrid {
		lineOutput := "|"
		for _, col := range line {
			switch col {
			case 0:
				lineOutput += "    "
			default:
				lineOutput += fmt.Sprintf(" %-3d", col)
			}
		}
		output = append(output, fmt.Sprintf("%s|", lineOutput[:len(lineOutput)-1]))
		output = append(output, fmt.Sprintf("+%s", strings.Repeat("   +", m.Width)))
	}
	output[len(output)-1] = fmt.Sprintf("+%s", strings.Repeat("---+", m.Width))

	return strings.Join(output, "\n")
}

func (m *ShikakuMap) DrawEmptyAsciiMap() []string {
	output := []string{}

	output = append(output, fmt.Sprintf("+%s", strings.Repeat("---+", m.Width)))
	for i := 0; i < m.Height; i++ {
		output = append(output, fmt.Sprintf("+%s", strings.Repeat("   |", m.Width)))
		output = append(output, fmt.Sprintf("+%s", strings.Repeat("---+", m.Width)))
	}

	return output
}

func (m *ShikakuMap) GetBlockAtPos(x, y int) *ShikakuMap {
	for _, block := range m.Blocks() {
		if y >= block.YPos && y <= block.YPos+block.Height-1 {
			if x >= block.XPos && x <= block.XPos+block.Width-1 {
				return block
			}
		}
	}
	return nil
}

func (m *ShikakuMap) DrawSolution() string {
	output := []string{}

	output = append(output, fmt.Sprintf("+%s", strings.Repeat("---+", m.Width)))
	for y := 0; y < m.Height; y++ {
		// a line: horizontal separators
		lineOutput := "|"
		for x := 0; x < m.Width; x++ {
			if x == m.Width-1 {
				lineOutput += "   |"
			} else {
				if m.GetBlockAtPos(x, y) == m.GetBlockAtPos(x+1, y) {
					lineOutput += "    "
				} else {
					lineOutput += "   |"
				}
			}
		}
		output = append(output, lineOutput)

		// vertical separators
		lineOutput = "+"
		for x := 0; x < m.Width; x++ {
			if y == m.Height-1 {
				lineOutput += "---+"
			} else {
				if m.GetBlockAtPos(x, y) == m.GetBlockAtPos(x, y+1) {
					lineOutput += "   +"
				} else {
					lineOutput += "---+"
				}
			}
		}
		output = append(output, lineOutput)
	}

	return strings.Join(output, "\n")
}
