package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	data := []float64{}
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
}
func seedData(data []float64, totaldata int, k int) ([]float64, []int) {
	centroid := make([]int, k)
	for i := 0; i < totaldata; i++ {
		data = append(data, rand.Float64()*10)
		if i < k {
			centroid[i] = i
		}
	}
	return data, centroid
}
func initialCluster(seededData []float64, k int, centroid []int) [][]float64 {
	n := len(seededData)
	if n < k {
		panic("The number of data should be more than the number of clusters")
	}
	clusters := make([][]float64, k)
	for i := 0; i < len(seededData); i++ {
		minDist := distance(seededData[i], seededData[centroid[0]])
		minIndex := 0
		for _, v := range centroid {
			d := distance(seededData[i], seededData[v])
			if d < minDist {
				minIndex = v
				minDist = d
			}
		}
		clusters[minIndex] = append(clusters[minIndex], seededData[i])
	}
	return clusters
}
func distance(x, y float64) float64 {
	return math.Abs(x - y)
}
func clusterCalculation(cluster [][]float64) [][]float64 {
	centroid := make([]float64, len(cluster))
	newCluster := make([][]float64, len(cluster))
	for i, v := range cluster {
		centroid[i] = centroidCalc(v...)
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
func centroidCalc(values ...float64) float64 {
	n := len(values)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += values[i]
	}
	return sum / float64(n)
}
func equalClusters(a, b [][]float64) bool {
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
