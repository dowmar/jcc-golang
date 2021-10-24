package main

import (
	"fmt"
	"strings"
)

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
	soal5()
}

//Soal 1
func soal1() {
	x := "JCC"
	y := "Candradimuka"
	z := "I Love Coding"

	for i := 1; i <= 20; i++ {
		if i%2 == 1 && i%3 == 1 {
			fmt.Println(i, x)
		} else if i%2 == 0 {
			fmt.Println(i, y)
		} else if i%3 == 0 && i%2 == 1 {
			fmt.Println(i, z)
		}

	}
}

//Soal 2
func soal2() {
	a := 7
	for i := 1; i <= a; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}

//Soal 3
func soal3() {
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var newkalimat = kalimat[2:7]
	fmt.Println(strings.Join(newkalimat, " "))
}

//Soal 4
func soal4() {
	start := 1
	var sayuran = [...]string{
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	}

	for i, j := range sayuran {
		fmt.Printf("%d.%s\n", i+start, j)
	}

}

//Soal 5
func soal5() {
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	var volume = (satuan["panjang"] * satuan["lebar"] * satuan["tinggi"])

	for i, j := range satuan {
		fmt.Printf("%s = %d\n", i, j)

	}
	fmt.Printf("Volume Balok = %d", volume)

}
