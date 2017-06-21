package main

import (
	"fmt"
	"math"
	"math/rand"
)

//
//type samplingPoisson struct {
//	seed int64
//	rand *rand.Rand
//}
//
//func NewSamplingPoisson(seed int64) *samplingPoisson {
//
//	return &samplingPoisson{
//		seed: seed,
//		rand: rand.New(rand.NewSource(seed)),
//	}
//}
//
//func (s *samplingPoisson) GetPoints(m *Map) {
//
//}
//
//func (s *samplingPoisson) randomDouble(min, max uint) uint {
//	return min + uint(uint(s.rand.Int())/(math.MaxInt64/(max-min)))
//}
//
//func (s *samplingPoisson) randomRange(min, max uint) uint {
//	return min + uint(uint(s.rand.Int())%(max-min))
//}
//
//func (s *samplingPoisson) randomPoint(m *Map) *Point {
//	return &Point{
//		X: s.randomDouble(0, m.Limits.X),
//		Y: s.randomDouble(0, m.Limits.Y),
//	}
//}
//
//func (s *samplingPoisson) randomDiscPoint(m *Map) *Point {
//
//}

type samplingPoisson struct {
	seed int64
	rand *rand.Rand
}

func NewSamplingPoisson(seed int64) *samplingPoisson {

	return &samplingPoisson{
		seed: seed,
		rand: rand.New(rand.NewSource(seed)),
	}
}

func (s *samplingPoisson) GetPoints(m *Map) {

}

// Poisson returns a random number of possion distribution
func (s *samplingPoisson) Poisson(lambda float64) int64 {
	if !(lambda > 0.0) {
		panic(fmt.Sprintf("Invalid lambda: %.2f", lambda))
	}
	return s.poisson(lambda)
}

func (s *samplingPoisson) poisson(lambda float64) int64 {
	// algorithm given by Knuth
	L := math.Pow(math.E, -lambda)
	var k int64 = 0
	var p float64 = 1.0

	for p > L {
		k++
		p *= s.rand.Float64()
	}
	return k - 1
}
