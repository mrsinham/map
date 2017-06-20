package main

type samplingUniform struct {
}

func (samplingUniform) GetPoints(m *Map) {
	for i := 0; i < m.Resolution; i++ {
		pts := &Point{}
		pts.X = uint(m.Rand.Int63n(int64(m.Limits.X)))
		pts.Y = uint(m.Rand.Int63n(int64(m.Limits.Y)))
		m.Grid = append(m.Grid, pts)
	}
}
