package algorithm

import (
	"math"
)

type CalculateDistances struct {
	lat1 float64
	lng1 float64
	lat2 float64
	lng2 float64
	unit string // K or N
}

func (cd *CalculateDistances) Distance() float64 {
	var theta float64 = cd.lng1 - cd.lng2
	var dist float64 = math.Sin(cd.Deg2rad(cd.lat1)) * math.Sin(cd.Deg2rad(cd.lat2)) + math.Cos(cd.Deg2rad(cd.lat1)) * math.Cos(cd.Deg2rad(cd.lat2)) * math.Cos(cd.Deg2rad(theta))
	dist = math.Acos(dist)
	dist = cd.Rad2deg(dist)
	dist = dist * 60 * 1.1515
	if cd.unit == "K" {
		dist = dist * 1.609344
	} else if cd.unit == "N" {
		dist = dist * 0.8684
	}
	return (dist)
}

func (cd *CalculateDistances) Deg2rad(deg float64) float64 {
	return (deg * math.Pi / 180.0)
}

func (cd *CalculateDistances) Rad2deg(rad float64) float64 {
	return (rad * 180 / math.Pi)
}

func (cd *CalculateDistances) DistanceFromTo() float64 {
	var earthRadius float64 = 6371000 //meters
	var dLat float64 = cd.ToRadians(cd.lat2 - cd.lat1)
	var dLng float64 = cd.ToRadians(cd.lng2 - cd.lng1)
	var a float64 = math.Sin(dLat / 2) * math.Sin(dLat / 2) + math.Cos(cd.ToRadians(cd.lat1)) * math.Cos(cd.ToRadians(cd.lat2)) * math.Sin(dLng / 2) * math.Sin(dLng / 2)
	var c float64 = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a))
	var dist float64 = (earthRadius * c)
	return dist
}

//// convert degree data to radians
func (cd *CalculateDistances) ToRadians(angdeg float64) float64 {
	return angdeg * math.Pi / 180
}
