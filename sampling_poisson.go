package main

type samplingPoisson struct{
	seed int64
}

func NewSamplingPoisson(seed int64) *samplingPoisson {
	return &samplingPoisson{
		seed:seed,
		rand:math.New(),
	}
}

func (s *samplingPoisson) GetPoints(m *Map) {

}


func (s *samplingPoisson) randomDouble(min, max float32) float32 {
	return min +
}