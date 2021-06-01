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
		fmt.Println(file_name, ":")
		sor_solver(readCSVFile(file_name)) //n = system dimensions, A = augmented matrix
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

func sor_solver(A []Array, b []float64) {

	//Create where the solution will be store
	var long int = len(b)
	var X0 = make([]float64, long)

	var tolerancia float64 = 0.0000001
	var iteracionMax int = 100

	var n int = len(A)
	var m int = len(A)
	var omega float64 = 0.5
	var X = X0
	var diferencia = make([]float64, long)

	var errado = tolerancia * 2
	var iteracion int = 0
	start := time.Now()

	for !(errado <= tolerancia || iteracion > iteracionMax) {
		for fila := 0; fila < n; fila++ {
			var suma float64 = 0
			for columna := 0; columna < m; columna++ {
				if fila != columna {
					suma = (A[fila][columna]) * (X[columna])
				}
			}
			var nuevo float64 = (1.0-omega)*X[fila] + (omega/A[fila][fila])*(b[fila]-suma)
			diferencia[fila] = math.Abs(nuevo - X[fila])

			X[fila] = nuevo

		}
		var max float64
		for i := 0; i < len(diferencia)-1; i++ {

			if max < diferencia[i] {

				max = diferencia[i]
			}
		}
		errado = max
		iteracion = iteracion + 1
	}
	// fmt.Println("Respuesta X: ", X)
	// fmt.Println("")
	elapsed := time.Since(start)
	fmt.Println("EXECUTION TIME: ", elapsed.Nanoseconds(), " nanoseconds")
}
