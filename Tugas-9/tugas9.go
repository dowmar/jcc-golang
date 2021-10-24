package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Soal 1
func endApp(kalimat string, tahun int) {
	fmt.Println(kalimat, tahun)
}

func runApp() {
	defer endApp("Candradimuka Jabar Coding Camp", 2021)
	fmt.Println("test")
}

// Soal 2
func exceptionCase() {
	if e := recover(); e != nil {
		fmt.Println(e)
	}
}
func kelilingSegitigaSamaSisi(nilai int, para bool) (string, error) {
	switch {
	case nilai == 0 && para == true:
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		return "", err
	case nilai == 0 && para == false:
		panic("Maaf and belum menginput sisi dari segitiga sama sisi")
	case para == false:
		return strconv.Itoa(int(nilai * 3)), nil
	default:
		return "keliling segitiga dengan sisi " + strconv.Itoa(int(nilai)) + "cm adalah " + strconv.Itoa(int(nilai*3)) + " cm", nil
	}
}
func testo() {
	defer exceptionCase()
	defer kelilingSegitigaSamaSisi(4, true)
	defer kelilingSegitigaSamaSisi(8, false)
	defer kelilingSegitigaSamaSisi(0, true)
	defer kelilingSegitigaSamaSisi(0, false)

	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}

// Soal 3
func tambahAngka(x int, angka *int) int {
	result := 0
	numbe1 := 0
	numbe2 := 0
	numbe3 := 0
	numbe4 := 0
	result += x
	if x == 7 {
		numbe1 = x
	} else if x == 6 {
		numbe2 = x
	} else if x == -1 {
		numbe3 = x
	} else if x == 9 {
		numbe4 = x
	}
	fmt.Println(numbe1, numbe2, numbe3, numbe4)
	return result
}

type movie struct {
	nilais int
}

func totalangka() {
	angka := 1

	numba1 := tambahAngka(7, &angka)
	numba2 := tambahAngka(6, &angka)
	numba3 := tambahAngka(-1, &angka)
	numba4 := tambahAngka(9, &angka)
	totalnumba := angka + numba1 + numba2 + numba3 + numba4
	fmt.Printf("Hasil dari %d + %d + %d + %d + %d adalah %d", angka, numba1, numba2, numba3, numba4, totalnumba)
}

// Soal 2

func main() {
	// Soal 3
	defer totalangka()

	//Soal 1
	runApp()

	// Soal 2
	testo()

}
