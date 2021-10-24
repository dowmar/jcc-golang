package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

//Soal 1
func tambahDataHP(merek string, phones *[]string) []string {
	*phones = append(*phones, merek)
	return *phones
}

type lingkaran struct {
	jari2 float64
}

func (x lingkaran) luas() {

	hasil := math.Pi * (math.Pow(x.jari2, 2))
	d := math.Ceil(hasil)
	fmt.Println("Hasil jari-jari lingkaran", d)

}
func (x lingkaran) keliling() {

	hasil := 2 * math.Pi * x.jari2
	d := math.Ceil(hasil)
	fmt.Println("Hasil keliling lingkaran", d)

}

type persegiPanjang struct {
	panjang, lebar int64
}

func (z persegiPanjang) luas2() {
	hasilluas := z.panjang * z.lebar
	fmt.Println("hasil luas : ", hasilluas)
}

func (z persegiPanjang) keliling2() {
	hasilkeliling := 2 * (z.panjang + z.lebar)
	fmt.Println("hasil keliling : ", hasilkeliling)
}

func main() {

	// Soal 1
	fmt.Println("---------Soal 1----------")
	var phones = []string{}

	tambahDataHP("Xiaomi", &phones)
	tambahDataHP("Asus", &phones)
	tambahDataHP("Iphone", &phones)
	tambahDataHP("Samsung", &phones)
	tambahDataHP("Oppo", &phones)
	tambahDataHP("Realme", &phones)
	tambahDataHP("Vivo", &phones)

	//fmt.Println(math.Pi)

	for u, x := range phones {
		fmt.Println(strconv.Itoa(u+1)+".", x)
		//time.Sleep(time.Second * 1)

	}

	// Soal 2
	fmt.Println("---------Soal 2----------")
	luaslingkaran1 := lingkaran{jari2: 7}
	luaslingkaran1.luas()
	luaslingkaran2 := lingkaran{jari2: 10}
	luaslingkaran2.luas()
	luaslingkaran3 := lingkaran{jari2: 15}
	luaslingkaran3.luas()

	kelilingLingkaran1 := lingkaran{jari2: 7}
	kelilingLingkaran1.keliling()
	kelilingLingkaran2 := lingkaran{jari2: 10}
	kelilingLingkaran2.keliling()
	kelilingLingkaran3 := lingkaran{jari2: 15}
	kelilingLingkaran3.keliling()

	// Soal 3
	fmt.Println("---------Soal 3----------")
	var panj = flag.Int64("Panjang", 15, "masukkan panjang")
	var leb = flag.Int64("Lebar", 10, "masukkan lebar")
	flag.Parse()
	luasPersegi := persegiPanjang{*panj, *leb}
	luasPersegi.luas2()
	kelilingPersegi := persegiPanjang{*panj, *leb}
	kelilingPersegi.keliling2()
}
