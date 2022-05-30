package main

import "fmt"

func main() {
	/*fmt.Println("Logic2Soal01")
	Logic2Soal01(9)*/

	/*fmt.Println("Logic2Soal02")
	Logic2Soal02(9)*/

	/*fmt.Println("Logic2Soal03")
	Logic2Soal03(9)*/

	/*fmt.Println("Logic2Soal04")
	Logic2Soal04(9)*/

	fmt.Println("Logic2Soal05")
	Logic2Soal05(9)
}

func Logic2Soal01(n int) {
	a := 3

	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		}
		//kolom selesai

		//ke baris selanjutnya
		fmt.Println("\n")
		//update value variabel a
		a += 3
	}
}

func Logic2Soal02(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		a := 3
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			//update value variabel a
			a += 3
		}
		//pindah baris
		fmt.Println("\n")
	}
}

func Logic2Soal03(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		a := 3 * n
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			//update value variabel a
			a -= 3
		}
		//pindah baris
		fmt.Println("\n")
	}
}

func Logic2Soal04(n int) {
	a := 3 * n
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		}
		//update value variabel a
		a -= 3
		//pindah baris
		fmt.Println("\n")
	}
}

func Logic2Soal05(n int) {
	x := n / 2
	y := 3

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

		}
	}
}
