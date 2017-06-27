package main

import (
	"math"
	"math/rand"
)

//
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

func (s *samplingPoisson) GetPoints(m *Map) []*Point {

}

func (s *samplingPoisson) randomDouble(min, max float32) float32 {
	return min + s.rand.Float32()/(math.MaxFloat32/(max-min))
}

func (s *samplingPoisson) randomRange(min, max uint) uint {
	return min + uint(uint(s.rand.Int())%(max-min))
}

func (s *samplingPoisson) randomPoint(m *Map) *Point {
	return &Point{
		X: s.randomDouble(0, m.Limits.X),
		Y: s.randomDouble(0, m.Limits.Y),
	}
}

func (s *samplingPoisson) randomDiscPoint(center *Point, r float32) *Point {
	angle := s.randomDouble(0, math.Pi)
	nx := math.Sin(float64(angle))
	ny := math.Cos(float64(angle))
	rl := s.randomDouble(r, 2*r)
	return &Point{center.X + float32(nx)*rl, center.Y + float32(ny)*rl}
}

func (s *samplingPoisson) findDiscPoint(center *Point, r float32, k int, points []*Point) {

}

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
//// Poisson returns a random number of possion distribution
//func (s *samplingPoisson) Poisson(lambda float64) int64 {
//	if !(lambda > 0.0) {
//		panic(fmt.Sprintf("Invalid lambda: %.2f", lambda))
//	}
//	return s.poisson(lambda)
//}
//
//func (s *samplingPoisson) poisson(lambda float64) int64 {
//	// algorithm given by Knuth
//	L := math.Pow(math.E, -lambda)
//	var k int64 = 0
//	var p float64 = 1.0
//
//	for p > L {
//		k++
//		p *= s.rand.Float64()
//	}
//	return k - 1
//}
