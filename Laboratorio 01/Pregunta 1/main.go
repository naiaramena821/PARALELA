package main

import (
	"fmt"
	"time"
)

func main() {
	const MAX = 1000 
	var A [MAX][MAX]float64
	var x [MAX]float64
	var y [MAX]float64

	// INICIALIZACIÓN
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			A[i][j] = 1
		}
		x[i] = 2
		y[i] = 0
	}

	// PRIMER LOOP
	start1 := time.Now()
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			y[i] += A[i][j] * x[j]
		}
	}
	elapsed1 := time.Since(start1)
	fmt.Printf("Tiempo Primer Loop: %v\n", elapsed1)

	// y = 0
	for i := 0; i < MAX; i++ {
		y[i] = 0
	}

	// SEGUNDO LOOP
	/*start2 := time.Now()
	for j := 0; j < MAX; j++ {
		for i := 0; i < MAX; i++ {
			y[i] += A[i][j] * x[j]
		}
	}
	elapsed2 := time.Since(start2)
	fmt.Printf("Tiempo Segundo Loop: %v\n", elapsed2)*/
}
