package main

import (
	"fmt"
	"math/rand"
	"time"
)

var route **int

func buildRoute(mapData [][2]int, numberOfCities int) {
	for i := 0; i < numberOfCities; i++ {
		route[i] = mapData[i]
	}
}

func distance(point1 *int, point2 *int) float64 {
	var dx int = point1[0] - point2[0]
	var dy int = point1[1] - point2[1]

	return sqrt(dx*dx + dy*dy)
}

func totalDistance(points **int, length int) float64 {
	var total float64 = 0.0
	for i := 0; i < length; i++ {
		total += distance(points[i], points[i-1])
	}
	return total
}

func main() {
	var n int = 1000
	var numberOfCities int = 51
	mapData := [][2]int{{37, 52}, {49, 49}, {52, 64}, {20, 26}, {40, 30}, {21, 47}, {17, 63}, {31, 62}, {52, 33}, {51, 21}, {42, 41}, {31, 32}, {5, 25}, {12, 42}, {36, 16}, {52, 41}, {27, 23}, {17, 33}, {13, 13}, {57, 58}, {62, 42}, {42, 57}, {16, 57}, {8, 52}, {7, 38}, {27, 68}, {30, 48}, {43, 67}, {58, 48}, {58, 27}, {37, 69}, {38, 46}, {46, 10}, {61, 33}, {62, 63}, {63, 69}, {32, 22}, {45, 35}, {59, 15}, {5, 6}, {10, 17}, {21, 10}, {5, 64}, {30, 15}, {39, 10}, {32, 39}, {25, 32}, {25, 55}, {48, 28}, {56, 37}, {30, 40}}
	//var initialT float64 = 100.0
	//var finalT float64 = 0.8
	//var coolingRate float64 = 0.9
	buildRoute(mapData, numberOfCities)
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("mapData:%d\n", mapData[0])
	fmt.Printf("totalDistance:%f\n", totalDistance(points, numberOfCities))
}
