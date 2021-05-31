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
		fmt.Println(args, ":")
		sor_solver(readCSVFile(file_name)) //n = system dimensions, A = augmented matrix
	}
}
func readCSVFile(file_name string) ([]Array, []float64) {
	csvFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
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
	var long int = len(A)
	var X0 = make([]float64, long)

	var tolerancia float64 = 0.0000001
	var iteracionMax int = 100

	var n int = len(A)
	var m int = len(A)
	var omega float64 = 0.5
	var X = X0

	diferencia := [][]float64{
		{1, 1, 1},
	}
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
			var nuevo = (1.0-omega)*X[fila] + (omega/A[fila][fila])*(b[fila]-suma)

			diferencia[0][fila] = math.Abs(nuevo - X[fila])

			X[fila] = nuevo

		}
		var max float64
		for _, j := range diferencia {
			var n float64
			for _, v := range j {
				if v > n {
					n = v
					max = n
				}
			}
		}
		errado = max
		iteracion = iteracion + 1
	}
	// fmt.Println("Respuesta X: ", X)
	// fmt.Println("")
	elapsed := time.Since(start)
	fmt.Println("Tiempo :", elapsed.Nanoseconds())
}
