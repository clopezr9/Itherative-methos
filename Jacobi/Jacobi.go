package main

import (
	"fmt"
	"math"
)

func main() {

	var A = [][]float64{{5.0, 2.0, -3}, {2.0, 10.0, -8.0}, {3.0, 8.0, 13.0}}
	var b = []float64{1.0, 4.0, 7.0}
	var N int = 500
	var ig = []float64{0.0, 0.0, 0.0}
	fmt.Println("A: ", A)
	fmt.Println("b: ", b)
	fmt.Println("Result:", Jacobi(A, b, N, ig))
}

func dominant(A [][]float64) bool {

	for i := 0; i < len(A); i++ {
		var dom = math.Abs(A[i][i])
		var count = 0.0
		for j := 0; j < len(A); j++ {
			count = count + math.Abs(A[i][j])
		}
		count = count - dom
		if dom < count {
			return false
		}
	}
	return true
}

func dot(A [][]float64, b []float64, l int) []float64 {

	result := make([]float64, l)
	for i := 0; i < l; i++ {
		var aux = 0.0
		for j := 0; j < l; j++ {
			aux = aux + b[j]*A[i][j]
		}
		result[i] = aux
	}
	return result
}

func substraction(a []float64, b []float64, l int) []float64 {

	result := make([]float64, l)
	for i := 0; i < l; i++ {
		result[i] = a[i] - b[i]
	}
	return result
}
func division(a []float64, b []float64, l int) []float64 {

	result := make([]float64, l)
	for i := 0; i < l; i++ {
		result[i] = a[i] / b[i]
	}
	return result
}

func matrix_permutations(arr [][]float64) [][][]float64 {
	var helper func([][]float64, int)
	res := [][][]float64{}

	helper = func(arr [][]float64, n int) {
		if n == 1 {
			tmp := make([][]float64, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func permutations(arr []float64) [][]float64 {

	var helper func([]float64, int)
	res := [][]float64{}

	helper = func(arr []float64, n int) {
		if n == 1 {
			tmp := make([]float64, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1.0]
					arr[n-1.0] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1.0]
					arr[n-1.0] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func make_dominant(A [][]float64, b []float64, l int) [][]float64 {

	C := make([][][]float64, 1)
	D := make([][][][]float64, 1)
	var Dom [][]float64
	D[0] = matrix_permutations(A)

	var bperm = permutations(b)
	C[0] = bperm

	for i := 0; i < len(D[0]); i++ {
		if dominant(D[0][i]) {
			Dom = D[0][i]
			Dom = append(Dom, C[0][i])
		}
	}
	return Dom
}

func Jacobi(A [][]float64, b []float64, N int, ig []float64) []float64 {

	//Creates varible with the length of Matrix A
	var l int = len(A)

	/*Variable that will be used to store te result of make_dominant function in case
	/* that the given matrix is not dominant.
	*/
	var Dom [][]float64
	if dominant(A) != true {
		Dom = make_dominant(A, b, l)
		if len(Dom) <= 0 {
			return Dom[0]
		}
		// Organices the A matrix so it can be dominant.
		for i := 0; i < len(Dom)-1; i++ {
			A[i] = Dom[i]
		}

		fmt.Println(A)
		b = Dom[len(Dom)-1]
		fmt.Println(b)
	}

	//Variable were dominant values are going to be stored.
	D := make([]float64, l)
	for i := 0; i < len(A); i++ {
		//Create a vector of the diagonal elements of A.
		D[i] = A[i][i]
	}

	//Creating Matrix were result from A - D is going to be stored.
	R := make([][]float64, l)
	for i := 0; i < len(A); i++ {
		R[i] = make([]float64, l)
	}

	// Filling R with correct values.
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if i != j {
				R[i][j] = A[i][j]
			}
		}
	}

	// Iterate for N times.
	for i := 0; i < N; i++ {
		ig = division(substraction(b, dot(R, ig, l), l), D, l)
	}

	return ig
}
