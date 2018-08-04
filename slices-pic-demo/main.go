package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)
	for i := range y {
		y[i] = make([]uint8, dx)
	}
	return y
}

func Gradient1(dx, dy int) [][]uint8 {
	p := Pic(dx, dy)
	for i := range p {
		for j := range p[i] {
			p[i][j] = uint8((i + j) / 2)
		}
	}
	return p
}
func Gradient2(dx, dy int) [][]uint8 {
	p := Pic(dx, dy)
	for i := range p {
		for j := range p[i] {
			p[i][j] = uint8(i * j)
		}
	}
	return p
}
func Gradient3(dx, dy int) [][]uint8 {
	p := Pic(dx, dy)
	for i := range p {
		for j := range p[i] {
			p[i][j] = uint8(i ^ j)
		}
	}
	return p
}

func main() {
	pic.Show(Gradient1)
	fmt.Println()
	pic.Show(Gradient2)
	fmt.Println()
	pic.Show(Gradient3)
}
