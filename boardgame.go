package main

import (
	"fmt"
	"os"
	"bufio"
)

type boardGame struct {
	height int
	width int
	chessmans []chessMan
	bitmap [][]int
}


// initialite the boardGame structure and set the bitmap as zero
func newBoardGame(height, width int) boardGame {
	status := [][]int{}
	for i:=0; i< height; i++ {
		row := []int{}
		for i:=0; i< width; i++ {
			row = append(row, 0)
		}

		status = append(status, row)
	}
	return boardGame{height: height, width: width, bitmap: status}
}


// check if player win, hardcode now.
func (b boardGame)checkWin() bool {
	for i:=0; i<len(DST_LOCATION);i++ {
		if b.chessmans[DST_CHESSMAN].location[i] != DST_LOCATION[i] {
			return false
		}
	}
	return true
}


// check if move block will touch the bound or obstacle
func (b boardGame)checkCollision(c *chessMan, direction int) bool {
	switch direction{
	case 1:
		// check if the chessman touch the bound
		if c.location[0] -1 < 0 {
			return true
		} else {
			// check if all the block can move without obstacle
			r := 0
			for i:=c.location[1]; i<c.width+c.location[1]; i++ {
				if b.bitmap[c.location[0]-1][i] == 0 {
					r++
				}
			}
			if r ==  c.width {
				return false
			} else {
				return true
			}
		}
	case 2:
		if c.location[0] + c.height  > b.height {
			return true
		} else {
			r := 0
			for i:=c.location[1]; i< c.width+c.location[1]; i++{
				if b.bitmap[c.location[0]+c.height ][i] == 0 {
					r++
				}
			}
			if r == c.width {
				return false
			} else {
				return true
			}
		}
	case 3:
		if c.location[1] - 1 < 0 {
			return true
		} else {
			r := 0
			for i:= c.location[0]; i<c.height+c.location[0]; i++ {
				if b.bitmap[i][c.location[1]-1] == 0 {
					r++
				}
			}
			if r == c.height {
				return false
			} else {
				return true
			}
		}
	case 4:
		if c.location[1]  + c.width > b.width {
			return true
		} else {
			r := 0
			for i:=c.location[0]; i<c.height + c.location[0]; i++ {
				if b.bitmap[i][c.location[1]+c.width] == 0 {
					r++
				}
			}
			if r == c.height {
				return false
			} else {
				return true
			}
		}
	default:
		fmt.Println("wrong direction")
	}
	return true
}

// move do two things
// 1, update the bitmap of the boardGame as the current status
// 2, update the current location of the block
func (b boardGame)move(direction int, c *chessMan) bool {
	if col:=b.checkCollision(c, direction); col {
		fmt.Println("Can not move!!!")
		return false
	} else {
		switch direction{
		case 1:
			for i:=c.location[1]; i< c.location[1] + c.width; i++ {
				b.bitmap[c.location[0]-1][i] = c.name
				b.bitmap[c.location[0]+c.height-1][i] = 0
			}
			c.location[0]--
			return false
		case 2:
			for i:= c.location[1]; i< c.location[1]  + c.width; i++ {
				b.bitmap[c.location[0]+c.height][i] = c.name
				b.bitmap[c.location[0]][i] = 0
			}
			c.location[0]++
			return false
		case 3:
			for i:= c.location[0]; i< c.location[0] + c.height; i++ {
				b.bitmap[i][c.location[1]-1] = c.name
				b.bitmap[i][c.location[1]+c.width -1] = 0
			}
			c.location[1]--
			return false
		case 4:
			for i:= c.location[0]; i< c.location[0]+ c.height; i++ {
				b.bitmap[i][c.location[1]+c.width] = c.name
				b.bitmap[i][c.location[1]] = 0
			}
			c.location[1]++
			return false
		}
	}
	return false
}

// render will draw current status, and most simple here, no mapping
func (b boardGame)render() bool {
	for i:=0; i< len(b.bitmap); i++{
		fmt.Println(b.bitmap[i])
	}
	return true
}

// save current status into a file to load next time
func (b boardGame)save(path string) bool {
	f, err := os.Create(path)
	if err != nil {
		return false
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "%d %d\n", b.height, b.width)
	for i:= 0; i<len(b.chessmans); i++ {
		c := b.chessmans[i]
		fmt.Fprintf(w, "%d %d %d %d %d\n", c.name,c.height,c.width, c.location[0], c.location[1])
	}
	w.Flush()
	f.Sync()
	fmt.Println("Save succennfully in: ", path)
	return true
}

func (b boardGame)usage() {
	fmt.Println("Usage:")
	fmt.Println("\tcommands: move <chessman number> | save filepath")
	fmt.Println("\texamples:")
	fmt.Println("\t\tmove 9 4 | move 10 3 | move 7 2 | move 8 2")
	fmt.Println("\t\tsave /tmp/hrdsave")
}

//func(b boardGame)autoSolve() {
	//solution := []string{}
//	for b.checkWin() {
		//b.searchStatus()
		//b.moveNext()
//	}
//}
