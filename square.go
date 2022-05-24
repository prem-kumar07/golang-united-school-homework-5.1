package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	endPoint:=Point{
		x:s.start.x+int(s.a),
		y:s.start.y-int(s.a),
	}
	return endPoint
}

func (s Square) Area() uint {
	area:=s.a*s.a
	return area
}

func (s Square) Perimeter() uint {
	perimeter:=2*s.a
	return perimeter
}
