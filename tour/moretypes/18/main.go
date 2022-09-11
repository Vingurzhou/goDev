package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

var a int

func Pic(dx, dy int) [][]uint8 {
	myPic := make([][]uint8, dy)
	fmt.Println(dx, dy, a)
	for i := 0; i < dy; i++ {
		myPic[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			myPic[i][j] = uint8(i * j)
		}
	}
	return myPic
}

func main() {
	pic.Show(Pic)
}
