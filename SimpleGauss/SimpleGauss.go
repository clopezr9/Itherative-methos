package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Array []float64

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		file_name := args[i]
		fmt.Println(file_name, ":")
		gaussSimple(readCSVFile(file_name)) //n = system dimensions, A = augmented matrix
	}
}

func readCSVFile(file_name string) ([]Array, int) {
	csvFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	str := strings.Split(file_name, "/")[1]
	var a_len = strings.Split(str, "-")[0]
	line_length, err := strconv.Atoi(a_len)
	augmented_matrix := []Array{}
	for _, line := range csvLines {
		line_length := len(line)
		array_aux := []float64{}
		for i := 0; i < line_length; i++ {
			line_i, err := strconv.ParseFloat(line[i], 64)
			if err != nil {
				fmt.Println(err)
			}
			array_aux = append(array_aux, line_i)
		}
		augmented_matrix = append(augmented_matrix, array_aux)
	}

	return augmented_matrix, line_length
}

func gaussSimple(augmented_matrix []Array, n int) {
	x := make([]float64, n)
	start := time.Now()
	for i := 0; i < n; i++ {
		if augmented_matrix[i][i] == 0.0 {
			return
		}

		for j := i + 1; j < n; j++ {
			var ratio = augmented_matrix[j][i] / augmented_matrix[i][i]

			for k := 0; k < n+1; k++ {
				augmented_matrix[j][k] = augmented_matrix[j][k] - ratio*augmented_matrix[i][k]
			}
		}
	}

	x[n-1] = augmented_matrix[n-1][n] / augmented_matrix[n-1][n-1]

	for i := n - 2; i > -1; i-- {
		x[i] = augmented_matrix[i][n]

		for j := i + 1; j < n; j++ {
			x[i] = x[i] - augmented_matrix[i][j]*x[j]
		}
		x[i] = x[i] / augmented_matrix[i][i]
	}
	duration := time.Since(start)
	fmt.Println("EXECUTION TIME: ", duration.Seconds())
}
