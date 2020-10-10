package main

import (
	"math"
)

func router(returnToOrigin bool, origin Point, points []Point) ([]Point, error) {
	var result []Point
	result = append(result, origin)

	current := origin
	rest := points

	for len(rest) > 0 {
		current, rest = findNextPoint(current, rest)
		result = append(result, current)
	}

	if returnToOrigin {
		result = append(result, origin)
	}

	return result, nil
}

func distance(pointA Point, pointB Point) float64 {
	R := 6373.0
	rad := math.Pi / 180
	lat1 := pointA.Latitude * rad
	lat2 := pointB.Latitude * rad
	dlon := (pointB.Longitude - pointA.Longitude) * rad
	dlat := (pointB.Latitude - pointA.Latitude) * rad
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func findNextPoint(current Point, points []Point) (Point, []Point) {
	var next Point
	var rest []Point

	for _, point := range points {
		if next.Latitude == 0 && next.Longitude == 0 {
			next = point
			continue
		}

		if distance(current, next) > distance(current, point) {
			rest = append(rest, next)
			next = point
		} else {
			rest = append(rest, point)
		}
	}

	return next, rest
}
