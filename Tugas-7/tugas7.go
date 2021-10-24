package main

import (
	"fmt"
)

func main() {
	//Soal 1
	soal1()

	fmt.Println()
	//Soal 2
	luassegitiga := Segitiga{alas: 5, tinggi: 4}
	luassegitiga.luas()
	luaspersegi := persegi{sisi: 7}
	luaspersegi.luas2()
	luaspersegipanjang := persegiPanjang{panjang: 10, lebar: 8}
	luaspersegipanjang.luas3()
	fmt.Println()
	//Soal3
	var colors = []string{"putih", "biru"}
	var hp = phone{
		name:   "galaxy",
		brand:  "samsung",
		year:   12,
		colors: colors[0:1],
	}
	fmt.Println(hp)
	//Soal4
	soal4()

}

//Soal 1
type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

func soal1() {
	var fruit1 = buah{
		nama:       "Nanas",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      9000,
	}
	var fruit2 = buah{
		nama:       "Jeruk",
		warna:      "Oranye",
		adaBijinya: true,
		harga:      8000,
	}
	var fruit3 = buah{
		nama:       "Semangka",
		warna:      "Hijau & Merah",
		adaBijinya: true,
		harga:      10000,
	}
	var fruit4 = buah{
		nama:       "Pisang",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      5000,
	}

	fmt.Println(fruit1)
	fmt.Println(fruit2)
	fmt.Println(fruit3)
	fmt.Println(fruit4)
}

//Soal 2
type Segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (x Segitiga) luas() {
	hasil := x.alas * x.tinggi / 2
	fmt.Println("luas segitigaa:", hasil)
}

func (y persegi) luas2() {
	hasil2 := y.sisi * y.sisi
	fmt.Println("luas persegi: ", hasil2)
}
func (z persegiPanjang) luas3() {
	hasil3 := z.panjang * z.lebar
	fmt.Println("luas persegipanjang: ", hasil3)
}

// Soal3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type movie struct {
	title, genre   string
	duration, year int
}

//Soal 4
func tambahDataFilm(judul string, waktu int, genre string, tahun int, dataFilm *[]movie) {
	item := movie{
		title:    judul,
		genre:    genre,
		duration: waktu,
		year:     tahun,
	}
	*dataFilm = append(*dataFilm, item)
	for x, _ := range *dataFilm {
		fmt.Printf("%d. ", x+1)
		fmt.Printf("title: %s\n genre: %s\n duration: %d\n year: %d\n", item.title, item.genre, item.duration, item.year)
		fmt.Println()
	}

}

func soal4() {
	var dataFilm = []movie{}
	tambahDataFilm("LOTR", 2, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 2, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 2, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 2, "horror", 2004, &dataFilm)

}
