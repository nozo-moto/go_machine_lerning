package main

import (
	"math/rand"

	"github.com/goml/gobrain"
)

type FizzBuzz []float64

func (f FizzBuzz) Type() int {
	for i, f := range f {
		if f > 0.4 {
			return i
		}
	}

	panic("Wrong")
}

func teacher(n int) []float64 {
	switch {
	case n%15 == 0:
		return []float64{1, 0, 0, 0}
	case n%3 == 0:
		return []float64{0, 1, 0, 0}
	case n%5 == 0:
		return []float64{0, 0, 1, 0}
	default:
		return []float64{0, 0, 0, 1}
	}
}

func bin(n int) []float64 {
	f := [8]float64{}
	for i := uint(0); i < 8; i++ {
		f[i] = float64((n >> i) & 1)
	}
	return f[:]
}

func main() {
	rand.Seed(0)

	patterns := [][][]float64{}
	for i := 1; i <= 100; i++ {
		patterns = append(patterns, [][]float64{
			bin(i), teacher(i),
		})
	}
	ff := &gobrain.FeedForward{}

	ff.Init(8, 100, 4)
	ff.Train(patterns, 1000, 0.6, 0.4, false)

	for i := 1; i < 100; i++ {
		switch FizzBuzz(ff.Update(bin(i))).Type() {
		case 0:
			println(i, ":: ", "FizzBuzz")
		case 1:
			println(i, ":: ", "Fizz")
		case 2:
			println(i, ":: ", "Buzz")
		case 3:
			println(i, ":: ", i)
		}
	}
}
