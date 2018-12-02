package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	routeA := [][2]int{{37, 52}, {49, 49}, {52, 64}, {20, 26}, {40, 30}, {21, 47}, {17, 63}, {31, 62}, {52, 33}, {51, 21}, {42, 41}, {31, 32}, {5, 25}, {12, 42}, {36, 16}, {52, 41}, {27, 23}, {17, 33}, {13, 13}, {57, 58}, {62, 42}, {42, 57}, {16, 57}, {8, 52}, {7, 38}, {27, 68}, {30, 48}, {43, 67}, {58, 48}, {58, 27}, {37, 69}, {38, 46}, {46, 10}, {61, 33}, {62, 63}, {63, 69}, {32, 22}, {45, 35}, {59, 15}, {5, 6}, {10, 17}, {21, 10}, {5, 64}, {30, 15}, {39, 10}, {32, 39}, {25, 32}, {25, 55}, {48, 28}, {56, 37}}
	routeB := [][2]int{{37, 52}, {49, 49}, {52, 64}, {20, 26}, {40, 30}, {21, 47}, {17, 63}, {31, 62}, {52, 33}, {51, 21}, {42, 41}, {31, 32}, {5, 25}, {12, 42}, {36, 16}, {52, 41}, {27, 23}, {17, 33}, {13, 13}, {57, 58}, {62, 42}, {42, 57}, {16, 57}, {8, 52}, {7, 38}, {27, 68}, {30, 48}, {43, 67}, {58, 48}, {58, 27}, {37, 69}, {38, 46}, {46, 10}, {61, 33}, {62, 63}, {63, 69}, {32, 22}, {45, 35}, {59, 15}, {5, 6}, {10, 17}, {21, 10}, {5, 64}, {30, 15}, {39, 10}, {32, 39}, {25, 32}, {25, 55}, {48, 28}, {56, 37}}

	var n int = 1000
	var numberOfCities int = len(routeA)
	fmt.Printf("numberOfCities:%d\n", numberOfCities)
	var initialT float64 = 100.0
	var finalT float64 = 0.8
	var coolingRate float64 = 0.9

	routeC := make([][2]int, numberOfCities)

	sa(routeA, numberOfCities, n, initialT, finalT, coolingRate)
	orderCrossOver(routeA, routeB, routeC, numberOfCities)

	for i := 0; i < numberOfCities; i++ {
		fmt.Printf("resultRoute[%d]:%d\n", i, routeC[i])
	}

	for i := 0; i < numberOfCities; i++ {
		for j := 0; j < numberOfCities; j++ {
			if i != j {
				if routeC[i] == routeC[j] {
					fmt.Printf("route[%d]:%d, j:%d, かぶりが発生しています。\n", i, routeC[i], j)
					break
				}
			}
		}
	}
}

func orderCrossOver(routeA [][2]int, routeB [][2]int, afterRoute [][2]int, numberOfCities int) {
	var numberSplit int = int(math.Trunc(float64(numberOfCities) / 3.0))
	fmt.Printf("numberSplit:%d\n", numberSplit)
	for i := numberSplit; i < (numberSplit * 2); i++ {
		afterRoute[i] = routeA[i]
		fmt.Printf("%d個目:Route:%d\n", i, afterRoute[i])
	}
	//かぶり数
	var k int = 0

	var f int = 0
	var i int = 0
	for f < numberSplit {
		var h int = 0
		for j := numberSplit; j < (numberSplit * 2); j++ {
			if routeB[i] == afterRoute[j] {
				k++
				fmt.Printf("かぶり発生\n")
				break
			} else {
				h++
				//fmt.Printf("h:%d", h)
				if h == numberSplit {
					fmt.Printf("f:%d ,", f)
					afterRoute[f] = routeB[i]
					fmt.Printf("%d個目:Route:%d\n", f, afterRoute[f])
					f++
				}
			}
		}
		i++
	}

	f = (numberSplit * 2)
	for f < numberOfCities {
		var h int = 0
		for j := numberSplit; j < (numberSplit * 2); j++ {
			if routeB[i] == afterRoute[j] {
				k++
				fmt.Printf("かぶり発生\n")
				break
			} else {
				h++
				//fmt.Printf("h:%d", h)
				if h == numberSplit {
					fmt.Printf("f:%d ,", f)
					afterRoute[f] = routeB[i]
					fmt.Printf("%d個目:Route:%d\n", f, afterRoute[f])
					f++
				}
			}
		}
		i++
	}
	fmt.Printf("i:%d, f:%d, k:%d\n", i, f, k)
}

func distance(point1 [2]int, point2 [2]int) float64 {
	var dx int = point1[0] - point2[0]
	var dy int = point1[1] - point2[1]

	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}

func totalDistance(points [][2]int) float64 {
	var length int = len(points)
	var total float64 = 0.0

	for i := 1; i < length; i++ {
		total += distance(points[i], points[i-1])
	}

	total += distance(points[0], points[length-1])

	return total
}

func swap(array [][2]int, index1 int, index2 int) {
	var tmp [2]int = array[index1]
	array[index1] = array[index2]
	array[index2] = tmp
}

func shouldChange(delta float64, t float64) bool {
	if delta <= 0 {
		return true
	}
	var frand float64 = rand.Float64() / 32767.0
	if frand < math.Exp(-delta/t) {
		return true
	}
	return false
}

func sa(route [][2]int, numberOfCities int, n int, initialT float64, finalT float64, coolingRate float64) {
	var randomIndex1 int
	var randomIndex2 int
	var currentTotalDistance float64
	var newTotalDistance float64

	rand.Seed(time.Now().UnixNano())
	currentTotalDistance = totalDistance(route)

	for t := initialT; t > finalT; t *= coolingRate {
		for i := 0; i < n; i++ {
			randomIndex1 = rand.Int() % numberOfCities
			randomIndex2 = rand.Int() % numberOfCities

			swap(route, randomIndex1, randomIndex2)

			newTotalDistance = totalDistance(route)

			if shouldChange(newTotalDistance-currentTotalDistance, t) {
				currentTotalDistance = newTotalDistance
			} else {
				swap(route, randomIndex1, randomIndex2)
			}
		}
	}
}
