package main

import (
	"fmt"
	"time"
)

func main() {
	MAX := 100
	A := make([][]float64, MAX)
	B := make([][]float64, MAX)
	resultado := make([][]float64, MAX)

	for i := 0; i < MAX; i++ {
		A[i] = make([]float64, MAX)
		B[i] = make([]float64, MAX)
		resultado[i] = make([]float64, MAX)
	}

	// INICIALIZACIÓN
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			A[i][j] = 2
			B[i][j] = 4
			resultado[i][j] = 0
		}
	}

	start := time.Now()

	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			for k := 0; k < MAX; k++ {
				resultado[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Multiplicación de matrices: %v ms\n", elapsed.Milliseconds())
}
