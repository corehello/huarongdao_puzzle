package main

import (
  "strings"
  "strconv"
  "bufio"
  "io/ioutil"
  "os"
  "fmt"
)


// return the
func waitInput() string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}


func readFile(fname string) (nums [][]int, err error) {
    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }
    lines := strings.Split(string(b), "\n")
    for _, l := range lines {
        if len(l) == 0 { continue }
        arr := strings.Split(l, " ")
        row := []int{}
        for j:=0; j< len(arr); j++ {
          n, err := strconv.Atoi(arr[j])
          if err != nil { return nil, err }
          row = append(row, n)
        }
        nums = append(nums, row)
    }

    return nums, nil
}
