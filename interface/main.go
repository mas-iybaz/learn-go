package main

import (
	"fmt"
	"math"
)

type bangunDatar interface {
	keliling() float32
	luas() float32
}

type persegi struct {
	sisi float32
}

func (e persegi) luas() float32 {
	return float32(math.Pow(float64(e.sisi), 2))
}

func (e persegi) keliling() float32 {
	return float32(e.sisi * 4)
}

func main() {
	var kertas bangunDatar = persegi{5.0}

	fmt.Printf("Keliling bangun datar adalah %.2f cm \n", kertas.keliling())
	fmt.Printf("Luas bangun datar adalah %.2f cm2 \n", kertas.luas())
}
