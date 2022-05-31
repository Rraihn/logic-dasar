package main

import (
	"fmt"
	"math"
	"testing"
)

func TestSoal01(t *testing.T) {
	n := 10
	a := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

func TestSoal02(t *testing.T) {
	n := 10
	a := 1

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

func TestLogic03(t *testing.T) {
	n := 10
	x := 99
	a := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print(x, "\t")
		}
	}
}

func TestSoal04(t *testing.T) {
	n := 10
	x := 777
	a := 1

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print(x, "\t")
		}
	}
}

func TestSoal05(t *testing.T) {
	n := 15

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizz buzz", "\t")
		} else if i%3 == 0 {
			fmt.Print("fizz", "\t")
		} else if i%5 == 0 {
			fmt.Print("buzz", "\t")
		} else {
			fmt.Print(i, "\t")
		}
	}
}

func TestSoal06(t *testing.T) {
	n := 15
	num := 1
	numb := 24
	numbr := 12
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			num = i * i
			fmt.Print(num, "\t")
		} else if i <= 7 {
			num = i * 3
			fmt.Print(num, "\t")
		} else if i <= 11 {
			numb -= 3
			fmt.Print(numb, "\t")
		} else {
			numbr -= 3
			fmt.Print(numbr, "\t")
		}
	}
}

func TestSoal07(t *testing.T) {
	n := 10
	m := 4
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 3
			fmt.Print(num, "\t")
		}
	}
}

func TestSoal08(t *testing.T) {
	n := 10
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 2
			fmt.Print(a, "\t")
		} else {
			b *= 5
			fmt.Print(b, "\t")
		}
	}
}

func TestSoal09(t *testing.T) {
	n := 12
	a := 1
	b := 2
	c := 3
	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 2
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func TestSoal10(t *testing.T) {
	n := 12
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 1
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func TestSoalNo6Simple(t *testing.T) {
	n := 15
	nt := n / 2
	x := 3

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			z := math.Pow(float64(i), 2)
			fmt.Print(z, "\t")
		} else {
			fmt.Print(x, "\t")
		}
		if i <= nt {
			x += 3
		} else {
			x -= 3
		}
	}
}
