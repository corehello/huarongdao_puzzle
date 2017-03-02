package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


const (
	DEFAULT_SAVE_PATH = "/tmp/hrdsave"
	DST_CHESSMAN = 2
)

var 	DST_LOCATION [2]int = [2]int {3,1}

// read boardGame status from a file

func initGameWithFile(path string) boardGame{
	max,_ := readFile(path)
	bg := newBoardGame(max[0][0], max[0][1])
	for i:= 1; i< len(max); i++ {
		nm := newChessMan(max[i][0], max[i][1], max[i][2], [2]int{max[i][3], max[i][4]})
		bg.chessmans = append(bg.chessmans, nm)
	}
	for i:=0; i< len(bg.chessmans); i++ {
		fill(bg, bg.chessmans[i])
	}
	return bg
}

func fill(b boardGame, man chessMan){
	for i:= man.location[0]; i< man.location[0]+ man.height; i++ {
		for j:= man.location[1]; j< man.location[1]+man.width; j++ {
			b.bitmap[i][j] = man.name
		}
	}
}

func processCommand(b boardGame, s string) bool{
	command := strings.Split(s, " ")
	if len(command) >= 1 {
		switch command[0]{
		case "move":
			if len(command) == 3{
				name,_ := strconv.Atoi(command[1])
				direction,_ := strconv.Atoi(command[2]) ///// can not prase string as int
				if name > len(b.chessmans) {
					fmt.Println("chessman is out if index")
					return false
				}
				b.move(direction, &b.chessmans[name-1])
			}
		case "save":
			if len(command) == 1 {
				b.save(DEFAULT_SAVE_PATH)
			} else {
				b.save(command[1])
			}
			return true
		default:
			fmt.Println("Not support this command")
		}
	} else {
		fmt.Println("Please input the correct command")
		return false
	}
	return false
}

func main() {
	hrd := initGameWithFile(os.Args[1])
	fmt.Println("welcome to 华容道")
	hrd.usage()
	hrd.render()
	for {
		if processCommand(hrd, waitInput()) {
			return
		}
		if hrd.checkWin() {
			fmt.Println("you win")
		}
		hrd.render()
	}
}
