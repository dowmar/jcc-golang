package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (x segitigaSamaSisi) luas() int {
	return x.alas * x.tinggi / 2
}

func (x segitigaSamaSisi) keliling() int {
	return x.alas * x.alas * x.alas
}

func (z persegiPanjang) luas() int {
	return z.panjang * z.lebar
}

func (z persegiPanjang) keliling() int {
	return (z.panjang + z.lebar) * 2
}
func (p tabung) volume() float64 {
	return math.Pi * math.Pow(p.jariJari, 2) * p.tinggi
}

func (p tabung) luasPermukaan() float64 {
	return (2 * math.Pi * math.Pow(p.jariJari, 2)) + (2 * math.Pi * p.jariJari * p.tinggi)
}

func (q balok) volume() float64 {
	return float64(q.panjang * q.lebar * q.tinggi)
}

func (q balok) luasPermukaan() float64 {
	return float64(2 * (q.panjang*q.lebar + q.panjang*q.tinggi + q.lebar*q.tinggi))
}

// Soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type smarthphone interface {
	phone() string
}

func (p phone) phone() string {
	hasil := "name : " + p.name + "\n" + "brand : " + p.brand + "\n" + "year : " + strconv.Itoa(p.year) + "\n" + "colors : " + strings.Join(p.colors, ", ")

	return hasil
}

// Soal 3
func luasPersegi(nilai int, para bool) interface{} {
	var nama2 interface{}
	if nilai == 0 && para == true {
		nama2 = "Maaf anda belum menginput sisi dari persegi"
	} else if nilai > 0 && para == true {
		nama2 = "luas persegi dengan sisi " + strconv.Itoa(nilai) + " cm adalah " + strconv.Itoa(nilai*2) + " cm"
	} else if nilai > 0 && para == false {
		nama2 = nilai
	} else {
		nama2 = nama2
	}
	return nama2
}

// Soal 4

func main() {
	//Soal 1
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	bangunDatar = segitigaSamaSisi{4, 2}
	fmt.Println("Segitiga Sama sisi")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling 	:", bangunDatar.keliling())

	bangunDatar = persegiPanjang{6, 8}
	fmt.Println("Persegi Panjang")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling	:", bangunDatar.keliling())
	fmt.Println()

	bangunRuang = tabung{7, 10}
	fmt.Println("Tabung")
	fmt.Println("volume 		:", bangunRuang.volume())
	fmt.Println("luas permukaan	:", bangunRuang.luasPermukaan())

	bangunRuang = balok{5, 6, 7}
	fmt.Println("Balok")
	fmt.Println("volume 		:", bangunRuang.volume())
	fmt.Println("luas permukaan	:", bangunRuang.luasPermukaan())

	//Soal 2
	var colors = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	var sentence smarthphone = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: colors,
	}
	fmt.Println(sentence.phone())

	//Soal 3

	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//Soal 4

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
	fmt.Printf(abcd)
	fmt.Println("=", result1+result2)
}
