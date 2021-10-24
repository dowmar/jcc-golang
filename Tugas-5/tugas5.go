package main

import (
	"fmt"
)

func main() {
	//untuk soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	fmt.Println(luas)

	keliling := kelilingPersegiPanjang(panjang, lebar)
	fmt.Println(keliling)

	volume := volumeBalok(panjang, lebar, tinggi)
	fmt.Println(volume)

	//-------------------------pembatas-----------------------
	//untuk soal 2

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	//-------------------------pembatas-----------------------

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)

	//-------------------------pembatas-----------------------

	soal4()
	//saya memanggil func soal4 yg berisi closure ke main, hasilnya sama
	//dan lebih rapih

}

//Soal 1
func luasPersegiPanjang(panjang int, lebar int) int {

	return (panjang * lebar)

}

func kelilingPersegiPanjang(panjang int, lebar int) int {

	return (panjang + lebar) * 2

}

func volumeBalok(panjang int, lebar int, tinggi int) int {

	return (panjang * lebar * tinggi)

}

//Soal 2
func introduce(nama, gender, occupation, age string) string {
	var nama2 string

	if gender == "laki-laki" {
		nama2 = "Pak " + nama
	} else {
		nama2 = "Bu " + nama
	}

	hasil := (nama2 + " adalah seorang " + occupation + " yang berusia " + age + " tahun")
	return hasil
}

//Soal 3
func buahFavorit(name1 string, buah ...string) (fruits string) {

	fmt.Printf("Halo nama saya %s dan buah favorit saya adalah", name1)
	for _, fruits := range buah {
		fmt.Printf(" %+q, ", fruits)
	}
	return
}

//Soal 4
func soal4() {
	var dataFilm = map[string]string{}

	var tambahDataFilm = func(judul, waktu, genre, tahun string) {
		dataFilm["genre"] = genre
		dataFilm["jam"] = waktu
		dataFilm["tahun"] = tahun
		dataFilm["title"] = judul
		fmt.Println(dataFilm)

	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")
}
