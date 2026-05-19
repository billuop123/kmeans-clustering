package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

type Point struct {
	X float64
	Y float64
}

func main() {
	data := make([]Point, 1000000)
	const k int = 3
	seededData, centroid := seedData(data, 1000000, k)
	cluster := initialCluster(seededData, k, centroid)
	i := 0
	for {
		newCluster := (clusterCalculation(cluster))
		// fmt.Println(newCluster)
		i++
		if equalClusters(newCluster, cluster) {
			break
		}
		cluster = newCluster
	}
	fmt.Println(i)
	printCluster(cluster)
}
func seedData(data []Point, totaldata int, k int) ([]Point, []Point) {
	centroid := make([]Point, k)
	for i := 0; i < totaldata; i++ {
		data[i].X = rand.Float64() * 10
		data[i].Y = rand.Float64() * 10
		if i < k {
			centroid[i].X = float64(i)
			centroid[i].Y = float64(i)
		}
	}
	return data, centroid
}
func initialCluster(seededData []Point, k int, centroid []Point) [][]Point {
	n := len(seededData)
	if n < k {
		panic("The number of data should be more than the number of clusters")
	}
	clusters := make([][]Point, k)
	for i := 0; i < len(seededData); i++ {
		minDist := distance(seededData[i], centroid[0])
		minIndex := 0
		for j, v := range centroid {
			d := distance(seededData[i], v)
			if d < minDist {
				minIndex = j
				minDist = d
			}
		}
		clusters[minIndex] = append(clusters[minIndex], seededData[i])
	}
	return clusters
}
func distance(a Point, b Point) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}
func clusterCalculation(cluster [][]Point) [][]Point {
	centroid := make([]Point, len(cluster))
	newCluster := make([][]Point, len(cluster))
	for i, v := range cluster {
		centroid[i].X, centroid[i].Y = centroidCalc(v...)
	}
	for _, v := range cluster {
		for _, subV := range v {
			minDist := distance(subV, centroid[0])
			minIndex := 0
			for i, c := range centroid {
				d := distance(c, subV)
				if d < minDist {
					minDist = d
					minIndex = i
				}
			}
			newCluster[minIndex] = append(newCluster[minIndex], subV)
		}
	}
	return newCluster
}
func centroidCalc(values ...Point) (float64, float64) {
	n := len(values)
	sumX := 0.0
	sumY := 0.0
	for i := 0; i < n; i++ {
		sumX += values[i].X
		sumY += values[i].Y
	}
	return sumX / float64(n), sumY / float64(n)
}
func equalClusters(a, b [][]Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
func printCluster(finalCluster [][]Point) {
	grid := make([][]rune, 10)
	for i := range grid {
		grid[i] = make([]rune, 10)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	sym := [3]rune{'o', '*', '#'}
	for i, outer := range finalCluster {
		for _, inner := range outer {
			x := int(inner.X)
			y := int(inner.Y)
			grid[x][y] = sym[i]
		}
	}
	for _, outer := range grid {
		for _, inner := range outer {
			fmt.Printf("%c  ", inner)
		}
		fmt.Println("")
	}
}
