package main

import (
	"fmt"
)

func main() {
	fmt.Println("Logic2Soal01")
	Logic2Soal01(9)

	fmt.Println("Logic2Soal02")
	Logic2Soal02(9)

	fmt.Println("Logic2Soal03")
	Logic2Soal03(9)

	fmt.Println("Logic2Soal04")
	Logic2Soal04(9)

	fmt.Println("Logic2Soal05")
	Logic2Soal05(9)

	fmt.Println("Logic2Soal06")
	Logic2Soal06(9)

	fmt.Println("Logic2Soal07")
	Logic2Soal07(9)

	fmt.Println("Logic2Soal08")
	Logic2Soal08(9)

	fmt.Println("Logic2Soal09")
	Logic2Soal09(9)

	fmt.Println("Logic2Soal10")
	Logic2Soal10(9)
}

func Logic2Soal01(n int) {
	//set variabel a ; value 3
	a := 3
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		} //kolom selesai

		//pindah baris
		fmt.Println("\n")
		//update value variabel a
		a += 3
	}
}

func Logic2Soal02(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		//set variabel a ; value 3
		a := 3
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			//update value variabel a
			a += 3
		} //kolom selesai

		//pindah baris
		fmt.Println("\n")
	}
}

func Logic2Soal03(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		//set variabel a ; value 3 (kali) nilai n
		a := 3 * n
		//looping kolom
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
	//set (new variabel) variabel a ; value 3 (kali) nilai n
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
	//set (new variabel) nt sebagai nilai tengah ; value n (dibagi) 2
	nt := n / 2
	//set variabel a ; value 3
	a := 3
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		} //kolom selesai
		//ke baris selanjutnya
		fmt.Println("\n")
		//update value variabel a
		if i < nt {
			//jika nilai i  kurang dari nilai tengah, maka nilai/value a bertambah 3
			a += 3
		} else {
			//jika nilai i lebih dari nilai tengah, maka nilai/value a berkurang 3
			a -= 3
		}
	}
}

func Logic2Soal06(n int) {
	//set (new variabel) nt sebagai nilai tengah
	nt := n / 2

	//looping baris
	for i := 0; i < n; i++ {
		//set variabel a ; value 3
		a := 3
		//looping kolom
		for j := 0; j < n; j++ {
			if j < nt {
				//jika nilai j kurang dari nilai tengah, maka
				fmt.Print(a, "\t")
				//update value variabel a, bertambah 3
				a += 3
			} else {
				//jika nilai j lebih dari nilai tengah, maka
				fmt.Print(a, "\t")
				//update value variabel a, berkurang 3
				a -= 3
			}
		}
		//pindah baris
		fmt.Println("\n")
	}
}

func Logic2Soal07(n int) {
	//set variabel a ; value 3
	a := 3

	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			if j <= i {
				//jika nilai j kurang dari atau sama dengan nilai i, maka
				fmt.Print(a, "\t")
			}
		}
		//pindah baris
		fmt.Println("\n")
		//update value variabel a, bertambah 3
		a += 3
	}
}

func Logic2Soal08(n int) {
	//set variabel a ; value 3
	a := 3

	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			if i <= j {
				//jika nilai i kurang dari atau sama dengan nilai j, maka
				fmt.Print(a, "\t")
			} else {
				//jika nilai i lebih dari atau sama dengan nilai j, maka
				fmt.Print("\t")
			}
		}
		//pindah baris
		fmt.Println("\n")
		//update value variabel a, bertambah 3
		a += 3
	}
}

func Logic2Soal09(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			//set variabel a ; value 3
			a := 3
			if i+j >= n-1 {
				//jika nilai i ditambah nilai j lebih dari atau sama dengan nilai n dikurang 1, maka
				a *= j + 1 // nilai a dikali nilai (j+1)
				fmt.Print(a, "\t")
			} else {
				//jika nilai i ditambah nilai j kurang dari atau sama dengan nilai n dikurang 1, maka
				fmt.Print("\t")
			}
		}
		//pindah baris
		fmt.Println("\n")

	}
}

func Logic2Soal10(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		//set variabel a ; value 3
		a := 3
		//looping kolom
		for j := 0; j < n; j++ {
			if i+j < n {
				//jika nilai i ditambah nilai j kurang dari nilai n, maka
				fmt.Print(a, "\t")
			}
			//update value variabel a, bertambah 3
			a += 3
		}
		//pindah baris
		fmt.Println("\n")
	}
}
