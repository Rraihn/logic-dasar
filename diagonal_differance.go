package main

import "fmt"

func main() {
	// membuat array 2 dimensi
	array := [][]int32{
		{11, 2, 4},   //baris 0
		{4, 5, 6},    //baris 1
		{10, 8, -12}, //baris 2
	}

	differencial := diagonalDifference(array)
	fmt.Println(differencial)
	array = [][]int32{
		{5, 10, 9, 7, 4},
		{4, 3, 10, 7, 5},
		{2, 1, 3, 9, 8},
		{4, 3, 4, 1, 3},
		{9, 4, 5, 7, 8},
	}
	differencial = diagonalDifference(array)
	fmt.Println(differencial)

	array = [][]int32{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 10},
	}
}

func diagonalDifference(arr [][]int32) int32 {
	var result int32
	var kanan int32 = 0
	var kiri int32 = 0

	// looping baris
	for i := 0; i < len(arr); i++ {
		// looping kolom
		for j := 0; j < len(arr[i]); j++ {
			if i == j {
				kanan = kanan + arr[i][j]
			}
			if i+j == len(arr)-1 {
				kiri = kiri + arr[i][j]
			}
		}
	}
	if kanan > kiri {
		result = kanan - kiri
	} else {
		result = kiri - kanan
	}
	return result
}