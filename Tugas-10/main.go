package main

import (
	. "Tugas-10/library"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("-------------Soal 1------------")

	var BangunDatar HitungBangunDatar
	var bangunRuang HitungBangunRuang

	BangunDatar = SegitigaSamaSisi{4, 2}
	fmt.Println("Segitiga Sama sisi")
	fmt.Println("luas		:", BangunDatar.Luas())
	fmt.Println("keliling 	:", BangunDatar.Keliling())

	BangunDatar = PersegiPanjang{6, 8}
	fmt.Println("Persegi Panjang")
	fmt.Println("luas		:", BangunDatar.Luas())
	fmt.Println("keliling	:", BangunDatar.Keliling())
	fmt.Println()

	bangunRuang = Tabung{7, 10}
	fmt.Println("Tabung")
	fmt.Println("volume 		:", bangunRuang.Volume())
	fmt.Println("luas permukaan	:", bangunRuang.LuasPermukaan())

	bangunRuang = Balok{5, 6, 7}
	fmt.Println("Balok")
	fmt.Println("volume 		:", bangunRuang.Volume())
	fmt.Println("luas permukaan	:", bangunRuang.LuasPermukaan())

	//Soal 2
	fmt.Println("-------------Soal 2------------")
	var Colors = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	var sentence SmartPhone = Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: Colors,
	}
	fmt.Println(sentence.Phone())

	//Soal 3
	fmt.Println("-------------Soal 3------------")
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))

	//Soal 4
	fmt.Println("-------------Soal 4------------")
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	var abcd string
	kalimat := prefix.(string)
	angka1 := kumpulanAngkaPertama.([]int)
	angka2 := kumpulanAngkaKedua.([]int)
	index := 0
	number1 := 0
	number2 := 0
	result1 := 0
	result2 := 0
	fmt.Printf("%s", kalimat)

	for _, number1 = range angka1 {
		result1 += number1
		fmt.Printf("%d+", number1)
	}
	for index, number2 = range angka2 {
		result2 += number2
		if index == 0 {
			abcd += "" + strconv.Itoa(number2) + ""
		} else {
			abcd += "+" + strconv.Itoa(number2) + " "
		}
	}
	fmt.Printf("%s", abcd)
	fmt.Println("=", result1+result2)
	fmt.Println()
}
