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

func (p *Point) Longitude() float64 {
	return p[0]
}

func (p *Point) Latitude() float64 {
	return p[1]
}

func (p *Point) Elevation() float64 {
	return p[2]
}

func (p Point) Time() float64 {
	return p[3]
}

func (p Point) Speed() float64 {
	return p[4]
}

func (l Line) First() Point {
	return l[0]
}

func (l Line) Last() Point {
	return l[len(l)-1]
}

func (l Line) Duration() float64 {
	return l.First().Time() - l.Last().Time()
}

func Speed(p1, p2 Point) float64 {
	t := math.Abs(p2.Time() - p1.Time())
	d := Distance(p1, p2)
	return 0
}

func SameLocation(p1, p2 Point) bool {
	return p1.Latitude() == p2.Latitude() && p1.Longitude() == p2.Longitude()
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

// Given φ is latitude radians, λ is longitude radians, R is earth radius:
// a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
// c = 2 ⋅ atan2(√a, √(1−a))
// d = R ⋅ c
func Distance(p1, p2 Point) float64 {
	if SameLocation(p1, p2) {
		return 0
	}
	radLat1 := toRadians(p1.Latitude())
	radLat2 := toRadians(p2.Latitude())
	latDistance := toRadians(p2.Latitude() - p1.Latitude())
	lonDistance := toRadians(p2.Longitude() - p1.Longitude())

	a := math.Pow(math.Sin(latDistance/2), 2) +
		math.Cos(radLat1)*math.Cos(radLat2)*
			math.Pow(math.Sin(lonDistance/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
