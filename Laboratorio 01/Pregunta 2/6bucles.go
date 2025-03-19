package main

import (
	"fmt"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	const MAX = 1000
	const BLOCK = 30

	// Creación de matrices A, B y resultado
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

	for i := 0; i < MAX; i += BLOCK {
		for j := 0; j < MAX; j += BLOCK {
			for k := 0; k < MAX; k += BLOCK {
				for l := i; l < min(i+BLOCK, MAX); l++ {
					for m := j; m < min(j+BLOCK, MAX); m++ {
						for n := k; n < min(k+BLOCK, MAX); n++ {
							resultado[l][m] += A[l][n] * B[n][m]
						}
					}
				}
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Multiplicación de matrices: %.2f ms\n", elapsed.Seconds()*1000)
}
