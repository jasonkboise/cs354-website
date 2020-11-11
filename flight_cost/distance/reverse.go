package distance

/*
	This package contains methods utilized for computing the distance between two points on a sphere.
	Despite the fact that earth is not a perfect sphere, this formula has a pruported error of up to 0.3%
	https://gis.stackexchange.com/questions/25494/how-accurate-is-approximating-the-earth-as-a-sphere#25580
*/

import "math"

func Length(lat1, lon1, lat2, lon2 float64) float64 {
	/*
		It is important to note that, during these calculations, directly assigning the type became important.
		Float 32 was the typed infered by Go for some values while float 64 was assigned for others. Go being strongly typed
		can take issue when performing operations on differing types
	*/

	var dlat float64 = toRadians(lat2 - lat1)
	var dlon float64 = toRadians(lon2 - lon1)
	lat1 = toRadians(lat1)
	lat2 = toRadians(lat2)
	var distance float64 = haversine(dlat, dlon, lat1, lon1, lat2, lon2)
	return distance
}

// Convert degrees to radians
func toRadians(lat1 float64) float64 { // function name is not capitalized, ergo it is private to this package (helper function)
	var radianUnit float64 = lat1 * (math.Pi / 180)
	return radianUnit
}

// haversine formula
func haversine(dlat, dlon, lat1, lon1, lat2, lon2 float64) float64 {
	var a float64 = math.Pow(math.Sin(dlat/2.0), 2.0)
	a += math.Pow(math.Sin(dlon/2.0), 2.0)
	a = a * math.Cos(lat1) * math.Cos(lat2)
	var rad float64 = 6371.0
	var c float64 = 2.0 * math.Asin(math.Sqrt(a))
	return rad * c
}
