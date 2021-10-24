package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"quiz3/buku"
	"quiz3/models"
	"quiz3/utils"
)

func main() {
	// http.HandleFunc("/bangun-datar/segitiga-sama-sisi?hitung=luas&alas=7&tinggi=10",GetSegitigaSamaSisi)
	http.HandleFunc("/books/", GetBooks)
	http.HandleFunc("/books/create", PostBooks)
	err := http.ListenAndServe("127.0.0.1:7000", nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func GetSegitigaSamaSisi() {}

//func

func GetBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		mahasiswas, err := buku.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, mahasiswas, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func PostBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var kmk models.Books

		//if kmk.Image_Url

		if kmk.Release_Year <= 1980 && kmk.Release_Year >= 2021 {
			//kmk.NIM = 90
			fmt.Println(kmk)
			http.Error(w, "tahun minimal 1980 dan tahun maksimal 2021", http.StatusBadRequest)
			//	return
		}
		// else {
		// 	kmk.NIM = indeksu.GetIndeksNilai(int(kmk.Semester))
		// 	//Mahasiswas = append(Mahasiswas, kmk)
		// 	//kmk = append(kmk, konak)
		// 	fmt.Println(kmk)
		// 	fmt.Println(Mahasiswas)
		// }
		// return
		//}
		//}
		// if kmk.Name == "utep" {

		// 	http.Error(w, "nilai minimal 0 dan maksimal 100", http.StatusBadRequest)
		// 	return
		// }

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var bk models.Books

		if err := json.NewDecoder(r.Body).Decode(&bk); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := buku.Insert(ctx, bk); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
