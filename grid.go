package main

import (
	"io"

	"github.com/ajstarks/svgo"
)

func renderGrid(w io.Writer, g *Map) {

	s := svg.New(w)
	s.Start(int(g.Limits.X), int(g.Limits.Y))
	for i := range g.Grid {
		s.Circle(int(g.Grid[i].X), int(g.Grid[i].Y), 1, "fill:black;stroke:black")
	}
	s.End()
}
