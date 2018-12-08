package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {
	routeA := [][2]float64{{11003.611100, 42102.500000},{11108.611100, 42373.888900},{11133.333300, 42885.833300},{11155.833300, 42712.500000},{11183.333300, 42933.333300},{11297.500000, 42853.333300},{11310.277800, 42929.444400},{11416.666700, 42983.333300},{11423.888900, 43000.277800},{11438.333300, 42057.222200},{11461.111100, 43252.777800},{11485.555600, 43187.222200},{11503.055600, 42855.277800},{11511.388900, 42106.388900},{11522.222200, 42841.944400},{11569.444400, 43136.666700},{11583.333300, 43150.000000},{11595.000000, 43148.055600},{11600.000000, 43150.000000},{11690.555600, 42686.666700},{11715.833300, 41836.111100},{11751.111100, 42814.444400},{11770.277800, 42651.944400},{11785.277800, 42884.444400},{11822.777800, 42673.611100},{11846.944400, 42660.555600},{11963.055600, 43290.555600},{11973.055600, 43026.111100},{12058.333300, 42195.555600},{12149.444400, 42477.500000},{12286.944400, 43355.555600},{12300.000000, 42433.333300},{12355.833300, 43156.388900},{12363.333300, 43189.166700},{12372.777800, 42711.388900},{12386.666700, 43334.722200},{12421.666700, 42895.555600},{12645.000000, 42973.333300}}
	routeB := [][2]float64{{11003.611100, 42102.500000},{11108.611100, 42373.888900},{11133.333300, 42885.833300},{11155.833300, 42712.500000},{11183.333300, 42933.333300},{11297.500000, 42853.333300},{11310.277800, 42929.444400},{11416.666700, 42983.333300},{11423.888900, 43000.277800},{11438.333300, 42057.222200},{11461.111100, 43252.777800},{11485.555600, 43187.222200},{11503.055600, 42855.277800},{11511.388900, 42106.388900},{11522.222200, 42841.944400},{11569.444400, 43136.666700},{11583.333300, 43150.000000},{11595.000000, 43148.055600},{11600.000000, 43150.000000},{11690.555600, 42686.666700},{11715.833300, 41836.111100},{11751.111100, 42814.444400},{11770.277800, 42651.944400},{11785.277800, 42884.444400},{11822.777800, 42673.611100},{11846.944400, 42660.555600},{11963.055600, 43290.555600},{11973.055600, 43026.111100},{12058.333300, 42195.555600},{12149.444400, 42477.500000},{12286.944400, 43355.555600},{12300.000000, 42433.333300},{12355.833300, 43156.388900},{12363.333300, 43189.166700},{12372.777800, 42711.388900},{12386.666700, 43334.722200},{12421.666700, 42895.555600},{12645.000000, 42973.333300}}

	var n int = 1000
	var numberOfCities int = len(routeA)
	var initialT float64 = 100.0
	var finalT float64 = 0.8
	var coolingRate float64 = 0.9
	var distanceA float64 = 0.0
	var distanceB float64 = 0.0
	var distanceC float64 = 0.0
	var distanceD float64 = 0.0
	fmt.Printf("numberOfCities:%d\n",numberOfCities)

	for i := 0; i < numberOfCities; i++ {
		for j := 0; j < numberOfCities; j++ {
			if i != j {
				if routeA[i] == routeA[j] {
					fmt.Printf("route[%d]:%d, j:%d, かぶりが発生しています。\n", i, routeA[i], j)
					break
				}
			}
		}
	}
	routeC := make([][2]float64, numberOfCities)
	routeD := make([][2]float64, numberOfCities)

	log.Print("started.")
	start := time.Now()
	//SA_1回目
	fmt.Printf("\nSA1回目\n")

	saFinished := make(chan bool)

	saFuncs := []func(){
		func() {
			fmt.Printf("totalDistanceA:%f\n", totalDistance(routeA))
			sa(routeA, numberOfCities, n, initialT, finalT, coolingRate)
			distanceA = totalDistance(routeA)
			fmt.Printf("ResultTotalDistanceA:%f\n", distanceA)
			saFinished <- true
		},
		func() {
			fmt.Printf("totalDistanceB:%f\n", totalDistance(routeB))
			sa(routeB, numberOfCities, n, initialT, finalT, coolingRate)
			distanceB = totalDistance(routeB)
			fmt.Printf("ResultTotalDistanceB:%f\n", distanceB)
			saFinished <- true
		},
	}

	for _, sa := range saFuncs {
		go sa()
	}

	for i := 0; i < len(saFuncs); i++ {
		<-saFinished
	}

	coFinished := make(chan bool)

	coFuncs := []func(){
		func() {
			orderCrossOver(routeA, routeB, routeC, numberOfCities)
			distanceC = totalDistance(routeC)
			fmt.Printf("順序交叉の距離C:%f\n", distanceC)
			coFinished <- true
		},
		func() {
			orderCrossOver(routeB, routeA, routeD, numberOfCities)
			distanceD = totalDistance(routeD)
			fmt.Printf("順序交叉の距離D:%f\n", distanceD)
			coFinished <- true
		},
	}

	for _, orderCrossOver := range coFuncs {
		go orderCrossOver()
	}

	for i := 0; i < len(coFuncs); i++ {
		<-coFinished
	}

	fmt.Printf("前No1.%f,No2.%f\n", totalDistance(routeA), totalDistance(routeB))

	bestOfRouteAB(distanceA, distanceB, distanceC, distanceD, routeA, routeB, routeC, routeD, numberOfCities)

	fmt.Printf("後No1.%f,No2.%f\n", totalDistance(routeA), totalDistance(routeB))

	//SA_2回目
	fmt.Printf("\nSA2回目\n")

	saFinished2 := make(chan bool)

	saFuncs2 := []func(){
		func() {
			fmt.Printf("totalDistanceA:%f\n", totalDistance(routeA))
			sa(routeA, numberOfCities, n, initialT, finalT, coolingRate)
			distanceA = totalDistance(routeA)
			fmt.Printf("ResultTotalDistanceA:%f\n", distanceA)
			saFinished2 <- true
		},
		func() {
			fmt.Printf("totalDistanceB:%f\n", totalDistance(routeB))
			sa(routeB, numberOfCities, n, initialT, finalT, coolingRate)
			distanceB = totalDistance(routeB)
			fmt.Printf("ResultTotalDistanceB:%f\n", distanceB)
			saFinished2 <- true
		},
	}

	for _, sa := range saFuncs2 {
		go sa()
	}

	for i := 0; i < len(saFuncs2); i++ {
		<-saFinished2
	}

	coFinished2 := make(chan bool)

	coFuncs2 := []func(){
		func() {
			orderCrossOver(routeA, routeB, routeC, numberOfCities)
			distanceC = totalDistance(routeC)
			fmt.Printf("順序交叉の距離C:%f\n", distanceC)
			coFinished2 <- true
		},
		func() {
			orderCrossOver(routeB, routeA, routeD, numberOfCities)
			distanceD = totalDistance(routeD)
			fmt.Printf("順序交叉の距離D:%f\n", distanceD)
			coFinished2 <- true
		},
	}

	for _, orderCrossOver := range coFuncs2 {
		go orderCrossOver()
	}

	for i := 0; i < len(coFuncs2); i++ {
		<-coFinished2
	}

	fmt.Printf("前No1.%f,No2.%f\n", totalDistance(routeA), totalDistance(routeB))

	bestOfRouteAB(distanceA, distanceB, distanceC, distanceD, routeA, routeB, routeC, routeD, numberOfCities)

	fmt.Printf("後No1.%f,No2.%f\n", totalDistance(routeA), totalDistance(routeB))

	for i := 0; i < numberOfCities; i++ {
		for j := 0; j < numberOfCities; j++ {
			if i != j {
				if routeA[i] == routeA[j] {
					fmt.Printf("route[%d]:%d, j:%d, かぶりが発生しています。\n", i, routeA[i], j)
					break
				}
			}
		}
	}

	fmt.Printf("routeAの要素数:%d\n",len(routeA))

	end := time.Now()
	fmt.Printf("実行時間：%f秒\n", (end.Sub(start)).Seconds())
	log.Print("end.")
}

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

func orderCrossOver(routeA [][2]float64, routeB [][2]float64, afterRoute [][2]float64, numberOfCities int) {
	var numberSplit int = int(math.Trunc(float64(numberOfCities) / 3.0))
	fmt.Printf("numberSplit:%d\n", numberSplit)
	for i := numberSplit; i < (numberSplit * 2); i++ {
		afterRoute[i] = routeA[i]
		//fmt.Printf("%d個目:Route:%d\n", i, afterRoute[i])
	}

	var f int = 0
	var i int = 0
	for f < numberSplit {
		var h int = 0
		for j := numberSplit; j < (numberSplit * 2); j++ {
			if routeB[i] == afterRoute[j] {
				//fmt.Printf("かぶり発生\n")
				break
			} else {
				h++
				//fmt.Printf("h:%d", h)
				if h == numberSplit {
					//fmt.Printf("f:%d ,", f)
					afterRoute[f] = routeB[i]
					//fmt.Printf("%d個目:Route:%d\n", f, afterRoute[f])
					f++
					//fmt.Printf("i:%d, f:%d. j:%d \n", i, f, j)
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
				//fmt.Printf("かぶり発生\n")
				break
			} else {
				h++
				//fmt.Printf("h:%d", h)
				if h == numberSplit {
					//fmt.Printf("f:%d ,", f)
					afterRoute[f] = routeB[i]
					//fmt.Printf("%d個目:Route:%d\n", f, afterRoute[f])
					f++
					//fmt.Printf("2:i:%d, f:%d. j:%d \n", i, f, j)
				}
			}
		}
		i++
	}
}

func BubbleSort(a []float64) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func bestOfRouteAB(distanceA float64, distanceB float64, distanceC float64, distanceD float64, routeA [][2]float64, routeB [][2]float64, routeC [][2]float64, routeD [][2]float64, numberOfCities int) {
	sortDisc := make([]float64, 4)

	tempRouteA := make([][2]float64, numberOfCities)
	tempRouteB := make([][2]float64, numberOfCities)

	var firstDistance string = ""
	var secondDistance string = ""

	sortDisc[0] = distanceA
	sortDisc[1] = distanceB
	sortDisc[2] = distanceC
	sortDisc[3] = distanceD

	fmt.Printf("前No1:%f,No2:%f,No3:%f,No4:%f\n", sortDisc[0], sortDisc[1], sortDisc[2], sortDisc[3])

	BubbleSort(sortDisc)

	fmt.Printf("後No1:%f,No2:%f,No3:%f,No4:%f\n", sortDisc[0], sortDisc[1], sortDisc[2], sortDisc[3])

	switch sortDisc[0] {
	case distanceA:
		firstDistance = "A"
	case distanceB:
		firstDistance = "B"
	case distanceC:
		firstDistance = "C"
	case distanceD:
		firstDistance = "D"
	}

	switch sortDisc[1] {
	case distanceA:
		secondDistance = "A"
	case distanceB:
		secondDistance = "B"
	case distanceC:
		secondDistance = "C"
	case distanceD:
		secondDistance = "D"
	}

	fmt.Printf("文字列No:1%s,No2:%s\n", firstDistance, secondDistance)

	switch firstDistance {
	case "A":
		for i := 0; i < numberOfCities; i++ {
			tempRouteA[i] = routeA[i]
		}
	case "B":
		for i := 0; i < numberOfCities; i++ {
			tempRouteA[i] = routeB[i]
		}
	case "C":
		for i := 0; i < numberOfCities; i++ {
			tempRouteA[i] = routeC[i]
		}
	case "D":
		for i := 0; i < numberOfCities; i++ {
			tempRouteA[i] = routeD[i]
		}
	}

	switch secondDistance {
	case "A":
		for i := 0; i < numberOfCities; i++ {
			tempRouteB[i] = routeA[i]
		}
	case "B":
		for i := 0; i < numberOfCities; i++ {
			tempRouteB[i] = routeB[i]
		}
	case "C":
		for i := 0; i < numberOfCities; i++ {
			tempRouteB[i] = routeC[i]
		}
	case "D":
		for i := 0; i < numberOfCities; i++ {
			tempRouteB[i] = routeD[i]
		}
	}

	for i := 0; i < numberOfCities; i++ {
		routeA[i] = tempRouteA[i]
	}

	for i := 0; i < numberOfCities; i++ {
		routeB[i] = tempRouteB[i]
	}

}
