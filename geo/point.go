package geo

type (
	// Point represents a spacetime location and speed in the format expected
	// by Google Maps Javascript.
	Point [5]float64
	Line  []Point
)

func (p Point) lon() float64       { return p[0] }
func (p Point) lat() float64       { return p[1] }
func (p Point) elevation() float64 { return p[2] }
func (p Point) time() float64      { return p[3] }
func (p Point) speed() float64     { return p[4] }

// SetSpeedFromPoint computes speed based on time and distance change from
// another point.
func (p Point) SetSpeedFromPoint(other Point) {
	p[4] = Speed(other, p)
}

func (l Line) first() Point { return l[0] }
func (l Line) last() Point  { return l[len(l)-1] }

// Duration is the amount of time between the first and last points of the
// line.
func (l Line) Duration() float64 {
	return l.first().time() - l.last().time()
}
