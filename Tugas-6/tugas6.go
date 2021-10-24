package main

import (
	"fmt"
)

func main() {
	soal1()
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	soal3()
	soal4()
}

// Soal 1
func soal1() {
	var luasLingkaran float64
	var kelilingLingkaran float64
	var r float64 = 4
	const phi float64 = 3.14

	luasLingkaran = phi * r * r
	kelilingLingkaran = 2 * phi * r

	fmt.Println("Before")
	fmt.Println("luas (value)       :", luasLingkaran)
	fmt.Println("luas (address)     :", &luasLingkaran)

	fmt.Println("keliling (value)   :", kelilingLingkaran)
	fmt.Println("keliling (address) :", &kelilingLingkaran)

	r = 7
	fmt.Println("")
	fmt.Println("After")
	luasLingkaran = phi * r * r
	kelilingLingkaran = 2 * phi * r
	fmt.Println("luas (value)       :", luasLingkaran)
	fmt.Println("luas (address)     :", &luasLingkaran)

	fmt.Println("keliling (value)   :", kelilingLingkaran)
	fmt.Println("keliling (address) :", &kelilingLingkaran)
}

//Soal 2
func introduce(sentence *string, nama, gender, occupation, age string) {
	var nama2 string

	if gender == "laki-laki" {
		nama2 = "Pak " + nama
	} else {
		nama2 = "Bu " + nama
	}

	*sentence = (nama2 + " adalah seorang " + occupation + " yang berusia " + age + " tahun")

}

//Soal 3
func soal3() {
	var buah = []string{}
	start := 1
	p := &buah
	*p = append(*p, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i, buah2 := range *p {

		fmt.Printf("%d. %s\n", i+start, buah2)
	}

}

//Soal 4
func soal4() {

	var dataFilm = []map[string]string{}

	var tambahDataFilm = func(judul, waktu, genre, tahun string, dataFilm *[]map[string]string) {
		item := map[string]string{
			"genre":    genre,
			"duration": waktu,
			"year":     tahun,
			"title":    judul,
		}
		*dataFilm = append(*dataFilm, item)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for z, x := range dataFilm {
		fmt.Printf("%d.", z+1)
		for i, j := range x {
			fmt.Println(i, ":", j)

		}
		fmt.Println("")
	}

}
