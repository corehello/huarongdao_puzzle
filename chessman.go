package main


type chessMan struct {
	name int
	height int
	width int
	location [2]int
}


func newChessMan(name, height, width int, location [2]int) chessMan {
	return chessMan{name: name, height: height, width: width, location: location}
}
