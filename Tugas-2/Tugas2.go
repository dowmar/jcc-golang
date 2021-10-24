package main

// Soal 1
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
	soal5()
}
func soal1() {
	var txt1, txt2, txt3, txt4, txt5 string
	txt1 = "Jabar"
	txt2 = "Coding"
	txt3 = "Camp"
	txt4 = "Golang"
	txt5 = "2021"

	fmt.Println(txt1, txt2, txt3, txt4, txt5)
}

// Soal 2
func soal2() {

	halo := "Halo Dunia"

	var newhalo = strings.Replace(halo, "Dunia", "Golang", -1)
	fmt.Println(newhalo)
}

// Soal 3
func soal3() {

	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.ToUpper(kataKedua[0:1]) + kataKedua[1:]
	kataKetiga = kataKetiga[0:6] + strings.ToUpper(kataKetiga[6:7])
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)
}

// Soal 4
func soal4() {
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkakeempat = "7"

	var satu, dua, tiga, empat int

	satu, _ = strconv.Atoi(angkaPertama)
	dua, _ = strconv.Atoi(angkaKedua)
	tiga, _ = strconv.Atoi(angkaKetiga)
	empat, _ = strconv.Atoi(angkakeempat)

	fmt.Println(satu + dua + tiga + empat)
}

// Soal 5
func soal5() {

	kalimat := "halo halo bandung"
	angka := 2021

	var kalimat2 = strings.Replace(kalimat, "halo", "Hi", -1)

	fmt.Println(strconv.Quote(kalimat2), "-", strconv.Itoa(angka))

}
