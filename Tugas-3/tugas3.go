package main

import (
	"fmt"
	"strconv"
)

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
}

//Soal 1
func soal1() {
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var kelilingPersegiPanjang int
	var luasSegitiga int

	x, _ := strconv.Atoi(panjangPersegiPanjang)
	y, _ := strconv.Atoi(lebarPersegiPanjang)
	a, _ := strconv.Atoi(alasSegitiga)
	t, _ := strconv.Atoi(tinggiSegitiga)

	kelilingPersegiPanjang = (x*2 + y*2)
	luasSegitiga = (a * t / 2)

	fmt.Println("Keliling Persegi Panjang : ", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga : ", luasSegitiga)

}

// Soal 2
func soal2() {
	var nilaiJohn = 80
	var nilaiDoe = 50
	if nilaiJohn >= 80 {
		fmt.Println("John " + "Index Nilai A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("John " + "Index Nilai B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("John " + "Index Nilai C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("John " + "Index Nilai D")
	} else if nilaiJohn < 50 {
		fmt.Println("John " + "Index Nilai E")
	}
	if nilaiDoe >= 80 {
		fmt.Println("Doe " + "Index Nilai A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Doe " + "Index Nilai B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Doe " + "Index Nilai C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Doe " + "Index Nilai D")
	} else if nilaiDoe < 50 {
		fmt.Println("Doe " + "Index Nilai E")
	}

}

//Soal 3
func soal3() {
	var tanggal = 5
	var bulan = 8
	var tahun = 2000

	switch bulan {
	case 1:
		fmt.Println(tanggal, "Januari", tahun)
	case 2:
		fmt.Println(tanggal, "Februari", tahun)
	case 3:
		fmt.Println(tanggal, "Maret", tahun)
	case 4:
		fmt.Println(tanggal, "April", tahun)
	case 5:
		fmt.Println(tanggal, "Mei", tahun)
	case 6:
		fmt.Println(tanggal, "Juni", tahun)
	case 7:
		fmt.Println(tanggal, "Juli", tahun)
	case 8:
		fmt.Println(tanggal, "Agustus", tahun)
	case 9:
		fmt.Println(tanggal, "September", tahun)
	case 10:
		fmt.Println(tanggal, "Oktober", tahun)
	case 11:
		fmt.Println(tanggal, "November", tahun)
	case 12:
		fmt.Println(tanggal, "Desember", tahun)
	}
}

// Soal 4
func soal4() {
	var kelahiran = 2000

	if kelahiran >= 1944 && kelahiran < 1965 {
		fmt.Println("Anda adalah Baby boomer")
	} else if kelahiran >= 1965 && kelahiran < 1980 {
		fmt.Println("Anda adalah Generasi X")
	} else if kelahiran >= 1980 && kelahiran < 1995 {
		fmt.Println("Anda adalah Generasi Y")
	} else if kelahiran >= 1995 && kelahiran < 2016 {
		fmt.Println("Anda adalah Generasi Z")
	}
}
