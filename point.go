package main

import "math/rand"

type Point struct {
	X, Y float32
}

type Map struct {
	Seed   int64
	Rand   *rand.Rand
	Grid   []*Point
	Limits struct {
		X, Y float32
	}
	Resolution int
}

func NewMap(
	seed int64, sizeX, sizeY float32, resolution int,
) *Map {
	m := &Map{Seed: seed}
	m.Rand = rand.New(rand.NewSource(seed))
	m.Limits.X = sizeX
	m.Limits.Y = sizeY
	m.Resolution = resolution
	return m
}
