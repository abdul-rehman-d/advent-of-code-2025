package utils

type Grid struct {
	data [][]byte
	rows int
	cols int
}

func (g *Grid) Get(r, c int) (byte, bool) {
	if r < 0 || r >= g.rows {
		return 0, false
	}
	if c < 0 || c >= g.cols {
		return 0, false
	}
	return g.data[r][c], true
}

func (g *Grid) Find(x byte) (Coord, bool) {
	for r := range g.rows {
		for c := range g.cols {
			if ch, ok := g.Get(r, c); ok && ch == x {
				return Coord{r, c}, true
			}
		}
	}
	return Coord{-1, -1}, false
}

func NewGrid(data string) *Grid {
	lines := GetLines(data)
	if len(lines) == 0 {
		return nil
	}
	d := make([][]byte, len(lines))
	for i, line := range lines {
		a := make([]byte, len(line))
		for i, ch := range line {
			a[i] = byte(ch)
		}
		d[i] = a
	}
	return &Grid{data: d, rows: len(d), cols: len(d[0])}
}
