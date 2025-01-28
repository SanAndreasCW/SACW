package commons

import (
	"cmp"
	"github.com/kodeyeen/omp"
	"math"
	"slices"
)

type Comparable struct {
	Distance float32
	Point    *omp.Vector3
}

func Distance(startPoint *omp.Vector3, endPoint *omp.Vector3) float32 {
	return float32(math.Sqrt(math.Pow(float64(startPoint.X), 2) + math.Pow(float64(endPoint.Y), 2)))
}

func FindNearestToPoints(startPoint *omp.Vector3, endPoints []*omp.Vector3) *omp.Vector3 {
	var distances []Comparable
	for _, point := range endPoints {
		distances = append(distances, Comparable{
			Distance: Distance(startPoint, point),
			Point:    point,
		})
	}
	return slices.MinFunc(distances, func(a, b Comparable) int {
		return cmp.Compare(a.Distance, b.Distance)
	}).Point
}

func FindFarthestToPoints(startPoint *omp.Vector3, endPoints []*omp.Vector3) *omp.Vector3 {
	var distances []Comparable
	for _, point := range endPoints {
		distances = append(distances, Comparable{
			Distance: Distance(startPoint, point),
			Point:    point,
		})
	}
	return slices.MaxFunc(distances, func(a, b Comparable) int {
		return cmp.Compare(a.Distance, b.Distance)
	}).Point
}
