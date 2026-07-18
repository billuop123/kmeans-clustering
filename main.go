package main

import (
	"fmt"
	"math"
	"math/rand"
)

type coord struct {
	X       float64
	Y       float64
	cluster int
}
type centroid struct {
	X float64
	Y float64
}

func main() {
	maxLimit := 1000
	minLimit := 0
	_ = minLimit
	var c []coord
	coordNum := 50
	k := 3
	scaleFactor := float64(coordNum) / float64(maxLimit)
	for range coordNum {
		xRand := rand.Float64() * float64(maxLimit)
		yRand := rand.Float64() * float64(maxLimit)
		c = append(c, coord{
			X: xRand * float64(scaleFactor),
			Y: yRand * float64(scaleFactor),
		})
	}
	centroids := []centroid{}
	for i := range k {
		centroids = append(centroids, centroid{c[i].X, c[i].Y})
	}
	loops := 0
	_ = loops
	assignment := make([]int, coordNum)
	changed := true
	for changed {
		changed = false
		loops++
		for i, val := range c {
			best := 0
			minD := dist(centroid{X: val.X, Y: val.Y}, centroids[best])
			for j := range k {
				d := dist(centroid{X: val.X, Y: val.Y}, centroids[j])
				if d < minD {
					minD = d
					best = j
				}
			}
			if assignment[i] != best {
				changed = true
			}
			assignment[i] = best
			c[i].cluster = best
		}
		average := make([]centroid, k)
		count := make([]int, k)
		for i := range c {
			average[c[i].cluster].X += c[i].X
			average[c[i].cluster].Y += c[i].Y
			count[c[i].cluster] += 1
		}
		for j := range average {
			average[j].X = average[j].X / float64(count[j])
			average[j].Y = average[j].Y / float64(count[j])
		}
		centroids = average
	}
	for _, val := range c {
		fmt.Printf("%f, %f      ======>cluster %d\n", val.X, val.Y, val.cluster)
	}
	fmt.Printf("Total number of loops %d", loops)
}

func dist(x, y centroid) float64 {
	return math.Sqrt((x.Y-y.Y)*(x.Y-y.Y) - (y.X-y.X)*(y.X-y.X))
}
