package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Array []float64

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		file_name := args[i]
		fmt.Println(args, ":")
		gaussSimple(readCSVFile(file_name)) //n = system dimensions, A = augmented matrix
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
	fmt.Println(matrix)
	fmt.Println(b)

	return matrix, b
}

func gaussSimple(A []Array, c []float64) {
	println(len(A))
	var n int = len(A)
	
	var slice = make([][]float64, n)
	var b [n]][1] float64 := slice
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// print(A[i][j])
			print(b[i][0])
		}
		println()
	}
}
