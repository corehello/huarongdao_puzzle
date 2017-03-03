package main

import (
  "fmt"
  "time"
  "math/rand"
)


var possiblesteps map[int][][]int
var resultsteps [][]int
var currentlevel int
var passedmap map[string]bool

func makeslicetostring(s [][]int)string{
  return fmt.Sprint(s)
}

func setpassedmap(s string){
  passedmap[s] = true
}

func cp(a *boardGame, b *boardGame){
  a.height = b.height
  a.width = b.width
  a.chessmans = b.chessmans
  a.bitmap = b.bitmap
}

func possibleSteps(b boardGame) [][]int{
  ps := [][]int{}
  for i:=1; i<5; i++ {
    for j:=0; j<len(b.chessmans); j++ {
      if b.checkCollision(&b.chessmans[j], i){
        continue
      } else {
        n := boardGame{}
        cp(&n, &b)
        n.move(i, &n.chessmans[j])
        if !n.passed() {
          ps = append(ps, []int{j+1, i})
        }
        dir := 0
        if i>2 {
          dir = (i-3+1)%2+3
        } else {
          dir = (i-1+1)%2+1
        }
        n.move(dir, &n.chessmans[j])
      }
    }
  }
  return ps
}


func pickNext(level int) []int{
  step := []int{}
  l := len(possiblesteps[level])
  rand.Seed(time.Now().UnixNano())
  rd := rand.Intn(l)
  step = possiblesteps[level][rd]
  switch rd {
  case l-1:
    possiblesteps[level] = possiblesteps[level][:l-1]
  case 0:
    possiblesteps[level] = possiblesteps[level][1:]
  default:
    a := possiblesteps[level][0:rd]
    for i:=rd+1; i<l; i++{
      a = append(a, possiblesteps[level][i])
    }
    possiblesteps[level] = a
  }
  return step
}

func (b boardGame)moveNext(){
  //fmt.Println("Moving next ......")
  step := pickNext(currentlevel)
  //fmt.Println("Next step is: move", step)
  resultsteps = append(resultsteps, step)
  b.move(step[1], &b.chessmans[step[0]-1])
  currentlevel++
  possiblesteps[currentlevel] = possibleSteps(b)
}


func (b boardGame)moveBack(){
  //fmt.Println("Moving back ......")
  currentlevel--
  l := len(resultsteps)
  step := resultsteps[l-1]
  dir := 0
  if step[1]>2 {
    dir = (step[1]-3+1)%2+3
  } else {
    dir = (step[1]-1+1)%2+1
  }
  b.move(dir, &b.chessmans[step[0]-1])
  resultsteps=resultsteps[:l-1]
}


func checkDead() bool {
  //fmt.Println("Checking dead ......")
  if len(possiblesteps[currentlevel]) == 0{
    return true
  } else {
    return false
  }
}

func (b boardGame)passed()bool{
  _,ok := passedmap[makeslicetostring(b.bitmap)]
  return ok
}

func autoSolve(b boardGame) bool {
  start := time.Now()
  passedmap = make(map[string]bool)
  currentlevel=1
  possiblesteps = make(map[int][][]int)
  possiblesteps[currentlevel] = possibleSteps(b)
	for !b.checkWin() {
    b.render()
		if checkDead() {
      if currentlevel == 1 {
        fmt.Println("No solution found.")
        return true
      } else {
		    b.moveBack()
      }
    } else {
      b.moveNext()
      if !b.passed() {
        setpassedmap(makeslicetostring(b.bitmap))
      }
    }
	}
  end := time.Now()
  fmt.Println()
  fmt.Println("Totally spend ", end.Sub(start), " time")
  fmt.Println("Have tried ", len(passedmap), " step")
  return false
}
