package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func distance(point1 [2]float64, point2 [2]float64) float64 {
	var dx float64 = point1[0] - point2[0]
	var dy float64 = point1[1] - point2[1]

	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}

func totalDistance(points [][2]float64) float64 {
	var length int = len(points)
	var total float64 = 0.0

	for i := 1; i < length; i++ {
		total += distance(points[i], points[i-1])
	}

	total += distance(points[0], points[length-1])

	return total
}

func swap(array [][2]float64, index1 int, index2 int) {
	var tmp [2]float64 = array[index1]
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

func sa(route [][2]float64, numberOfCities int, n int, initialT float64, finalT float64, coolingRate float64) {
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

func orderCrossOver(routeA [][2]float64, routeB [][2]float64, afterRoute [][2]float64, numberOfCities int) {
	var numberSplit int = int(math.Trunc(float64(numberOfCities) / 3.0))

	for i := numberSplit; i < (numberSplit * 2); i++ {
		afterRoute[i] = routeA[i]
	}

	for i := 0; i < numberSplit; i++ {
		afterRoute[i] = routeB[i]
		for j := numberSplit; j < (numberSplit * 2); j++ {
			if afterRoute[i] == routeB[j] {
				i--
				break
			}
		}
	}

	for i := (numberSplit * 2); i < numberOfCities; i++ {
		afterRoute[i] = routeB[i]
		for j := numberSplit; j < (numberSplit * 2); j++ {
			if afterRoute[i] == routeB[j] {
				i--
				break
			}
		}
	}
}

func main() {
	routeA := [][2]float64{{11003.611100, 42102.500000},{11108.611100, 42373.888900},{11133.333300, 42885.833300},{11155.833300, 42712.500000},{11183.333300, 42933.333300},{11297.500000, 42853.333300},{11310.277800, 42929.444400},{11416.666700, 42983.333300},{11423.888900, 43000.277800},{11438.333300, 42057.222200},{11461.111100, 43252.777800},{11485.555600, 43187.222200},{11503.055600, 42855.277800},{11511.388900, 42106.388900},{11522.222200, 42841.944400},{11569.444400, 43136.666700},{11583.333300, 43150.000000},{11595.000000, 43148.055600},{11600.000000, 43150.000000},{11690.555600, 42686.666700},{11715.833300, 41836.111100},{11751.111100, 42814.444400},{11770.277800, 42651.944400},{11785.277800, 42884.444400},{11822.777800, 42673.611100},{11846.944400, 42660.555600},{11963.055600, 43290.555600},{11973.055600, 43026.111100},{12058.333300, 42195.555600},{12149.444400, 42477.500000},{12286.944400, 43355.555600},{12300.000000, 42433.333300},{12355.833300, 43156.388900},{12363.333300, 43189.166700},{12372.777800, 42711.388900},{12386.666700, 43334.722200},{12421.666700, 42895.555600},{12645.000000, 42973.333300}}
	var n int = 1000
	var numberOfCitties int = len(routeA)
	var initialT float64 = 100.0
	var finalT float64 = 0.8
	var coolingRate float64 = 0.9

	log.Print("started.")
	start := time.Now()

	fmt.Printf("totalDistanceA:%f\n", totalDistance(routeA))
	sa(routeA, numberOfCitties, n, initialT, finalT, coolingRate)
	fmt.Printf("ResultTotalDistanceA:%f\n", totalDistance(routeA))

	end := time.Now()
	fmt.Printf("実行時間%f秒\n", (end.Sub(start)).Seconds())

	/*for i := 0; i < numberOfCitties; i++ {
		fmt.Printf("%d個目:RouteC:%d\n", i, routeC[i])
	}

	//重複チェック
	var k int = 0
	for i := 0; i < numberOfCitties; i++ {
		for j := 0; j < numberOfCitties; j++ {
			if routeC[i] == routeA[j] {
				k++
			}
		}
	}
	fmt.Printf("一致した数:%d\n", k)*/

	log.Print("end.")
}
