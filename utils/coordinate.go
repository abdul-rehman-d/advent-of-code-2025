package utils

type Coord struct {
	R int
	C int
}

func (c Coord) Add(a Coord) Coord {
	return Coord{
		R: c.R + a.R,
		C: c.C + a.C,
	}
}
