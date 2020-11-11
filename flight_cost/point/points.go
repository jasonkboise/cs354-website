package point

/*
	X,Y ordered pairs used to represent lat/lon
*/

type Point struct { // The capital letter for the variable type indicates that this type is exported (not private to this class)
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}
