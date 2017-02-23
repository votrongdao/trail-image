package gps

import "math"

type (
	Point [5]float64
	Line  []Point
)

const (
	piDeg        = math.Pi / 180.0
	radiusMiles  = 3958.756
	radiusKm     = 6371.0
	feetPerMeter = 3.28084
)

var (
	earthRadius         = radiusMiles
	elevationConversion = feetPerMeter
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

func Speed(p1, p2 Point) float64 {
	t := math.Abs(p2.time() - p1.time())
	d := Distance(p1, p2)
	return 0
}

func SameLocation(p1, p2 Point) bool {
	return p1.lat() == p2.lat() && p1.lon() == p2.lon()
}

func toRadians(degrees float64) float64 {
	return degrees * piDeg
}

// Length gives the total distance between all points.
func Length(points []Point) float64 {
	d := 0.0

	for i := 1; i < len(points); i++ {
		d += Distance(points[i-1], points[i])
	}
	return d
}

// Simplify points using the Douglas-Peucker algorithm with recursion elimination.
func Simplify(points []Point, maxDeviationFeet int) []Point {
	if maxDeviationFeet <= 0 {
		return points
	}
	yard := 3.0
	mile := yard * 1760
	equatorFeet := mile * radiusMiles
	count := len(points)
	// convert tolerance in feet to tolerance in geographic degrees
	// TODO this is a percent, not degrees
	tolerance := float64(maxDeviationFeet) / equatorFeet

	first := 0
	last := count - 1
	maxDistance := 0.0
	distance := 0.0
	index := 0
	keep := make([]bool, count)

	var stack []int
	var simple []Point

	keep[first] = true
	keep[last] = true

	for last >= 0 {
		maxDistance = 0.0

		for i := first + 1; i < last; i++ {
			distance = PointLineDistance(points[i], points[first], points[last])
			if distance > maxDistance {
				index = i
				maxDistance = distance
			}
		}

		if maxDistance > tolerance {
			keep[index] = true // keep the deviant point
			stack = append(stack, first, index, index, last)
		}
		last, stack = pop(stack)
		first, stack = pop(stack)
	}

	for i, p := range points {
		if keep[i] {
			simple = append(simple, p)
		}
	}
	return simple
}

func pop(stack []int) (int, []int) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

// PointLineDistance finds the shortest distance between a point and a line
// defined by two other points.
func PointLineDistance(p, p1, p2 Point) float64 {
	x := p1.lon()
	y := p1.lat()
	Δx := p2.lon() - x
	Δy := p2.lat() - y

	if Δx != 0 || Δy != 0 {
		// non-zero distance
		t := ((p.lon()-x)*Δx + (p.lat()-y)*Δy) / (Δx*Δx + Δy*Δy)

		if t > 1 {
			x = p2.lon()
			y = p2.lat()
		} else if t > 0 {
			x += Δx * t
			y += Δy * t
		}
	}

	Δx = p.lon() - x
	Δy = p.lat() - y

	return Δx*Δx + Δy*Δy
}

// Given φ is latitude radians, λ is longitude radians, R is earth radius:
// a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
// c = 2 ⋅ atan2(√a, √(1−a))
// d = R ⋅ c
func Distance(p1, p2 Point) float64 {
	if SameLocation(p1, p2) {
		return 0
	}
	radLat1 := toRadians(p1.lat())
	radLat2 := toRadians(p2.lat())
	latDistance := toRadians(p2.lat() - p1.lat())
	lonDistance := toRadians(p2.lon() - p1.lon())

	a := math.Pow(math.Sin(latDistance/2), 2) +
		math.Cos(radLat1)*math.Cos(radLat2)*
			math.Pow(math.Sin(lonDistance/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
