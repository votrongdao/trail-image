package geo

type (
	Point [5]float64
	Line  []Point
)

func (p *Point) lon() float64 {
	return p[0]
}

func (p *Point) lat() float64 {
	return p[1]
}

func (p *Point) elevation() float64 {
	return p[2]
}

func (p Point) time() float64 {
	return p[3]
}

func (p Point) speed() float64 {
	return p[4]
}

func (l Line) first() Point {
	return l[0]
}

func (l Line) last() Point {
	return l[len(l)-1]
}

func (l Line) Duration() float64 {
	return l.first().time() - l.last().time()
}