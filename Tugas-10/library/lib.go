package library

import (
	// . "Tugas-10/main"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (X SegitigaSamaSisi) Luas() int {
	return X.Alas * X.Tinggi / 2
}

func (X SegitigaSamaSisi) Keliling() int {
	return X.Alas * X.Alas * X.Alas
}

func (z PersegiPanjang) Luas() int {
	return z.Panjang * z.Lebar
}

func (z PersegiPanjang) Keliling() int {
	return (z.Panjang + z.Lebar) * 2
}
func (p Tabung) Volume() float64 {
	return math.Pi * math.Pow(p.JariJari, 2) * p.Tinggi
}

func (p Tabung) LuasPermukaan() float64 {
	return (2 * math.Pi * math.Pow(p.JariJari, 2)) + (2 * math.Pi * p.JariJari * p.Tinggi)
}

func (q Balok) Volume() float64 {
	return float64(q.Panjang * q.Lebar * q.Tinggi)
}

func (q Balok) LuasPermukaan() float64 {
	return float64(2 * (q.Panjang*q.Lebar + q.Panjang*q.Tinggi + q.Lebar*q.Tinggi))
}

// Soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type SmartPhone interface {
	Phone() string
}

func (p Phone) Phone() string {
	hasil := "Name : " + p.Name + "\n" + "Brand : " + p.Brand + "\n" + "Year : " + strconv.Itoa(p.Year) + "\n" + "Colors : " + strings.Join(p.Colors, ", ")

	return hasil
}

// Soal 3
func LuasPersegi(nilai int, para bool) interface{} {
	var nama2 interface{}
	if nilai == 0 && para {
		nama2 = "Maaf anda belum menginput sisi dari persegi"
	} else if nilai > 0 && para {
		nama2 = "Luas persegi dengan sisi " + strconv.Itoa(nilai) + " cm adalah " + strconv.Itoa(nilai*2) + " cm"
	} else if nilai > 0 && !para {
		nama2 = nilai
	} else {
		nama2 = nama2
	}
	return nama2
}

// Soal 4

func main() {
	// Soal1()
	// Soal2()
	// Soal3()
	// Soal4()
	//Soal 1
	var BangunDatar HitungBangunDatar
	var bangunRuang HitungBangunRuang

	BangunDatar = SegitigaSamaSisi{4, 2}
	fmt.Println("Segitiga Sama sisi")
	fmt.Println("Luas		:", BangunDatar.Luas())
	fmt.Println("Keliling 	:", BangunDatar.Keliling())

	BangunDatar = PersegiPanjang{6, 8}
	fmt.Println("Persegi Panjang")
	fmt.Println("Luas		:", BangunDatar.Luas())
	fmt.Println("Keliling	:", BangunDatar.Keliling())
	fmt.Println()

	bangunRuang = Tabung{7, 10}
	fmt.Println("Tabung")
	fmt.Println("Volume 		:", bangunRuang.Volume())
	fmt.Println("Luas permukaan	:", bangunRuang.LuasPermukaan())

	bangunRuang = Balok{5, 6, 7}
	fmt.Println("Balok")
	fmt.Println("Volume 		:", bangunRuang.Volume())
	fmt.Println("Luas permukaan	:", bangunRuang.LuasPermukaan())

	//Soal 2
	//var Colors =
	var sentence SmartPhone = Phone{
		Name:   "Samsung GalaXy Note 20",
		Brand:  "Samsung GalaXy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	fmt.Println(sentence.Phone())

	//Soal 3

	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))

	//Soal 4

	var prefiX interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	var abcd string
	kalimat := prefiX.(string)
	angka1 := kumpulanAngkaPertama.([]int)
	angka2 := kumpulanAngkaKedua.([]int)
	indeX := 0
	number1 := 0
	number2 := 0
	result1 := 0
	result2 := 0
	fmt.Printf("%s", kalimat)

	for _, number1 = range angka1 {
		result1 += number1
		fmt.Printf("%d+", number1)
	}
	for indeX, number2 = range angka2 {
		result2 += number2
		if indeX == 0 {
			abcd += "" + strconv.Itoa(number2) + ""
		} else {
			abcd += "+" + strconv.Itoa(number2) + " "
		}
	}
	fmt.Printf("%s", abcd)
	fmt.Println("=", result1+result2)
}

func Soal1() {
	var BangunDatar HitungBangunDatar
	var bangunRuang HitungBangunRuang

	BangunDatar = SegitigaSamaSisi{4, 2}
	fmt.Println("Segitiga Sama sisi")
	fmt.Println("Luas		:", BangunDatar.Luas())
	fmt.Println("Keliling 	:", BangunDatar.Keliling())

	BangunDatar = PersegiPanjang{6, 8}
	fmt.Println("Persegi Panjang")
	fmt.Println("Luas		:", BangunDatar.Luas())
	fmt.Println("Keliling	:", BangunDatar.Keliling())
	fmt.Println()

	bangunRuang = Tabung{7, 10}
	fmt.Println("Tabung")
	fmt.Println("Volume 		:", bangunRuang.Volume())
	fmt.Println("Luas permukaan	:", bangunRuang.LuasPermukaan())

	bangunRuang = Balok{5, 6, 7}
	fmt.Println("Balok")
	fmt.Println("Volume 		:", bangunRuang.Volume())
	fmt.Println("Luas permukaan	:", bangunRuang.LuasPermukaan())

}

func Soal2() {
	//Soal 2
	var Colors = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	var sentence SmartPhone = Phone{
		Name:   "Samsung GalaXy Note 20",
		Brand:  "Samsung GalaXy Note 20",
		Year:   2020,
		Colors: Colors,
	}
	fmt.Println(sentence.Phone())
}
func Soal3() {
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))
}

func Soal4() {

	var prefiX interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	var abcd string
	kalimat := prefiX.(string)
	angka1 := kumpulanAngkaPertama.([]int)
	angka2 := kumpulanAngkaKedua.([]int)
	indeX := 0
	number1 := 0
	number2 := 0
	result1 := 0
	result2 := 0
	fmt.Printf("%s", kalimat)

	for _, number1 = range angka1 {
		result1 += number1
		fmt.Printf("%d+", number1)
	}
	for indeX, number2 = range angka2 {
		result2 += number2
		if indeX == 0 {
			abcd += "" + strconv.Itoa(number2) + ""
		} else {
			abcd += "+" + strconv.Itoa(number2) + " "
		}
	}
	fmt.Printf("%s", abcd)
	fmt.Println("=", result1+result2)
}
