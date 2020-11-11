package statemap

/*
	This package contains the map where State names are the key and lat/lon Points are the values
*/

import "github.com/AnonymousRecess/flight_cost/point"

var StateLocations = map[string]point.Point{
	"Alabama":    point.NewPoint(32.806671, 86.791130),
	"Alaska":     point.NewPoint(64.2008, 149.4937),
	"Arizona":    point.NewPoint(33.729759, 111.431221),
	"Arkansas":   point.NewPoint(34.969704, 92.373123),
	"California": point.NewPoint(36.7783, 119.4179),
	"Colorado":   point.NewPoint(39.059811, 105.311104),

	"Connecticut": point.NewPoint(41.597782, 72.755371),
	"Delaware":    point.NewPoint(39.318523, 75.507141),
	"Florida":     point.NewPoint(27.766279, 81.686783),

	"Georgia":  point.NewPoint(33.040619, 83.643074),
	"Hawaii":   point.NewPoint(21.094318, 157.498337),
	"Idaho":    point.NewPoint(44.240459, 114.478828),
	"Illinois": point.NewPoint(40.349457, 88.986137),
	"Indiana":  point.NewPoint(39.849426, 86.258278),

	"Iowa":      point.NewPoint(42.011539, 93.210526),
	"Kansas":    point.NewPoint(38.526600, 96.726486),
	"Kentucky":  point.NewPoint(37.668140, 84.670067),
	"Louisiana": point.NewPoint(31.169546, 91.867805),

	"Maine":         point.NewPoint(44.693947, 69.381927),
	"Maryland":      point.NewPoint(39.063946, 76.802101),
	"Massachusetts": point.NewPoint(42.230171, 71.530106),
	"Michigan":      point.NewPoint(43.326618, 84.536095),

	"Minnesota":   point.NewPoint(45.694454, 93.900192),
	"Mississippi": point.NewPoint(32.741646, 89.678696),
	"Missouri":    point.NewPoint(38.456085, 92.288368),

	"Montana":       point.NewPoint(46.921925, 110.454353),
	"Nebraska":      point.NewPoint(41.125370, 98.268082),
	"Nevada":        point.NewPoint(38.313515, 117.055374),
	"New Hampshire": point.NewPoint(43.452492, 71.563896),

	"New Jersey": point.NewPoint(40.298904, 74.521011),
	"New Mexico": point.NewPoint(34.840515, 106.248482),
	"New York":   point.NewPoint(42.165726, 74.948051),

	"North Carolina": point.NewPoint(35.630066, 79.806419),
	"North Dakota":   point.NewPoint(47.528912, 99.784012),
	"Ohio":           point.NewPoint(40.388783, 82.764915),
	"Oklahoma":       point.NewPoint(35.565342, 96.928917),

	"Oregon":         point.NewPoint(44.572021, 122.070938),
	"Pennsylvania":   point.NewPoint(40.590752, 77.209755),
	"Rhode Island":   point.NewPoint(41.680893, 71.511780),
	"South Carolina": point.NewPoint(33.856892, 80.945007),

	"South Dakota":  point.NewPoint(44.299782, 99.438828),
	"Tennessee":     point.NewPoint(35.747845, 86.692345),
	"Texas":         point.NewPoint(31.054487, 97.563461),
	"Utah":          point.NewPoint(40.150032, 111.862434),
	"Vermont":       point.NewPoint(44.045876, 72.710686),
	"Virginia":      point.NewPoint(37.769337, 78.169968),
	"Washington":    point.NewPoint(47.400902, 121.490494),
	"West Virginia": point.NewPoint(38.491226, 80.954453),
	"Wisconsin":     point.NewPoint(44.268543, 89.616508),
	"Wyoming":       point.NewPoint(42.755966, 107.302490),
}
