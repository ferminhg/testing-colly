package domain

type Vector struct {
	value []uint64
}

func NewVector(value []uint64) *Vector {
	return &Vector{
		value: value,
	}
}

func NewRandomVector() *Vector {
	return &Vector{
		value: []uint64{1, 2, 3},
	}
}
