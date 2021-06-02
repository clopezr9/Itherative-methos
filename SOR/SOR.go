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
	for i := 1; i < len(args); i++ {
		file_name := args[i]
		fmt.Println(args[i], ":")
		sor_solver(readCSVFile(file_name)) //n = system dimensions, A = augmented matrix
	}
}
func readCSVFile(file_name string) ([]Array, []float64) {
	csvFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Successfully Opened CSV file")
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

func euclidian_distance(A []float64, b []float64) float64 {
	var distance = 0.0
	for fila := 0; fila < len(A); fila++ {
		distance = distance + math.Pow(A[fila]-b[fila], 2)
	}

	distance = math.Sqrt(distance)
	return distance
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

func sor_solver(A []Array, b []float64) {

	//Create where the solution will be store
	var long int = len(b)
	var X0 = make([]float64, long)

	var tolerancia float64 = 0.000001 //1.0071743059050822
	var iteracionMax int = 100

	var n int = len(A)
	var m int = len(A)
	var omega float64 = 1.5
	var X = X0
	var tol float64 = euclidian_distance(dot(A, X, n), b)

	var iteracion int = 0
	start := time.Now()

	for tol > tolerancia && iteracion < iteracionMax {

		for fila := 0; fila < n; fila++ {
			var suma float64 = 0
			for columna := 0; columna < m; columna++ {
				if fila != columna {
					suma = (A[fila][columna]) * (X[columna])
				}
			}
			var nuevo float64 = (1.0-omega)*X[fila] + (omega/A[fila][fila])*(b[fila]-suma)

			X[fila] = nuevo

		}

		tol = euclidian_distance(dot(A, X, n), b)
		iteracion = iteracion + 1
	}
	// fmt.Println("Respuesta X: ", X)
	fmt.Println("")
	elapsed := time.Since(start)
	fmt.Println(elapsed.Nanoseconds())
	// fmt.Println("EXECUTION TIME: ", elapsed.Nanoseconds(), " nanoseconds")
}
