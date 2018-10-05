package main

import (
    "fmt"
    "log"
    "time"
    "math/rand"
)

func main() {
    log.Print("started.")

    rand.Seed(time.Now().UnixNano())
    fmt.Println("My favorite number is", rand.Intn(10))
    log.Print("end.")
}
