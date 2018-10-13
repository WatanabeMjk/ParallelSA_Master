package main

//#include <stdio.h>
//#include <stdlib.h>
//#include <math.h>
//#include <time.h>
import "C"

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var global int = 123

func test(number int) **int{
  return number
}

func main() {
	log.Print("started.")
	var name int = 1000
	var numberOfCities int = 51
  var pointerTest *int = &name
  var doublePointerTest **int = &pointerTest
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Println("int:", name,numberOfCities)
  fmt.Println("pointer:", pointerTest)
  fmt.Println("doublePointer:", *doublePointerTest)
  fmt.Println("functionPointer:", test(name))
	C.main()
	log.Print("end.")
}
