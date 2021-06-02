package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Array []float64

func main() {
	args := os.Args
	var N int = 100
	for i := 1; i < len(args); i++ {
		file_name := args[i]
		fmt.Println(file_name, ":")
		var a, b = readCSVFile(file_name)
		jacobi(a, b, N)
	}
}

func readCSVFile(file_name string) ([]Array, []float64) {
	csvFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	matrix := []Array{}
	b := []float64{}
	for _, line := range csvLines {
		line_length := len(line)
		array_aux := []float64{}
		for i := 0; i < line_length; i++ {
			line_i, err := strconv.ParseFloat(line[i], 64)
			if err != nil {
				fmt.Println(err)
			}
			if i == line_length-1 {
				b = append(b, line_i)
				break
			}
			array_aux = append(array_aux, line_i)
		}
		matrix = append(matrix, array_aux)
	}

	return matrix, b
}

func dominant(A []Array) bool {

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

func dot(A []Array, b []float64, l int) []float64 {

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

func euclidian_distance(A []float64, b []float64) float64 {

	var distance = 0.0
	for fila := 0; fila < len(A); fila++ {
		distance = distance + math.Pow(A[fila]-b[fila], 2)
	}
	distance = math.Sqrt(distance)
	return distance
}

func jacobi(A []Array, b []float64, N int) {

	//Creates varible with the length of Matrix A
	var l int = len(A)
	var ig = make([]float64, l)
	var tolerancia float64 = 0.000001
	var tol float64 = euclidian_distance(dot(A, ig, l), b)
	//Variable were dominant values are going to be stored.
	D := make([]float64, l)
	for i := 0; i < len(A); i++ {
		//Create a vector of the diagonal elements of A.
		D[i] = A[i][i]
	}

	//Creating Matrix were result from A - D is going to be stored.
	R := make([]Array, l)
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

	var i int = 0
	start := time.Now()
	// Iterate for N times.
	for tol > tolerancia && i < N {
		ig = division(substraction(b, dot(R, ig, l), l), D, l)
		tol = euclidian_distance(dot(A, ig, l), b)
		i = i + 1
	}
	duration := time.Since(start)
	fmt.Println("EXECUTION TIME: ", duration.Nanoseconds())
	//fmt.Println("Ans:", ig)
}
