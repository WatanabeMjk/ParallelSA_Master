package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

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
	var i int
	var t float64
	var currentTotalDistance float64
	var newTotalDistance float64

	rand.Seed(time.Now().UnixNano())
	currentTotalDistance = totalDistance(route)

	for t = initialT; t > finalT; t *= coolingRate {
		for i = 0; i < n; i++ {
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

func main() {
	route := [][2]int{{37, 52}, {49, 49}, {52, 64}, {20, 26}, {40, 30}, {21, 47}, {17, 63}, {31, 62}, {52, 33}, {51, 21}, {42, 41}, {31, 32}, {5, 25}, {12, 42}, {36, 16}, {52, 41}, {27, 23}, {17, 33}, {13, 13}, {57, 58}, {62, 42}, {42, 57}, {16, 57}, {8, 52}, {7, 38}, {27, 68}, {30, 48}, {43, 67}, {58, 48}, {58, 27}, {37, 69}, {38, 46}, {46, 10}, {61, 33}, {62, 63}, {63, 69}, {32, 22}, {45, 35}, {59, 15}, {5, 6}, {10, 17}, {21, 10}, {5, 64}, {30, 15}, {39, 10}, {32, 39}, {25, 32}, {25, 55}, {48, 28}, {56, 37}, {30, 40}}
	var n int = 1000
	var numberOfCitties int = 51
	var initialT float64 = 100.0
	var finalT float64 = 0.8
	var coolingRate float64 = 0.9

	parallelRouteA := make([][2]int, numberOfCitties)
	parallelRouteB := make([][2]int, numberOfCitties)
	parallelRouteC := make([][2]int, numberOfCitties)

	log.Print("started.")
	start := time.Now()

	finished := make(chan bool)

	funcs := []func(){
		func() {
			fmt.Printf("totalDistance1:%f\n", totalDistance(route))
			sa(route, numberOfCitties, n, initialT, finalT, coolingRate)
			fmt.Printf("totalDistance1:%f\n", totalDistance(route))
			parallelRouteA = route
			fmt.Printf("parallelRouteA:%d\n", parallelRouteA[0])
			finished <- true
		},
		func() {
			fmt.Printf("totalDistance2:%f\n", totalDistance(route))
			sa(route, numberOfCitties, n, initialT, finalT, coolingRate)
			fmt.Printf("totalDistance2:%f\n", totalDistance(route))
			parallelRouteB = route
			fmt.Printf("parallelRouteB:%d\n", parallelRouteB[0])
			finished <- true
		},
		func() {
			fmt.Printf("totalDistance3:%f\n", totalDistance(route))
			sa(route, numberOfCitties, n, initialT, finalT, coolingRate)
			fmt.Printf("totalDistance3:%f\n", totalDistance(route))
			parallelRouteC = route
			fmt.Printf("parallelRouteC:%d\n", parallelRouteC[0])
			finished <- true
		},
	}

	for _, sa := range funcs {
		go sa()
	}

	for i := 0; i < len(funcs); i++ {
		<-finished
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	log.Print("end.")
}