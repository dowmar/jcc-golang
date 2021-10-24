package main

import (
	"Tugas-15/models"
	"Tugas-15/nilai"
	"Tugas-15/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/nilai", GetNilai)
	router.POST("/nilai/create", PostNilai)
	router.PUT("/nilai/id/update", UpdateNilai)
	// http.HandleFunc("/mahasiswa", GetNilai)
	// http.HandleFunc("/mahasiswa/create", PostNilai)
	// http.HandleFunc("/mahasiswa/update", UpdateNilai)
	// http.HandleFunc("/nilai/delete", DeleteNilai)
	router.DELETE("/nilai/delete", DeleteNilai)

	fmt.Println("Server Running at Port 7000")
	log.Fatal(http.ListenAndServe("127.0.0.1:7000", router))

}

func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nilay, err := nilai.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilay, http.StatusOK)
}

func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var Scr models.NilaiMahasiswa
	// if r.Header.Get("Content-Type") == "application/json" {
	// 	decodeJSON := json.NewDecoder(r.Body)
	// 	if err := decodeJSON.Decode(&Scr); err != nil {
	// 		log.Fatal(err)
	// 	}
	if Scr.Nilai > 100 || Scr.Nilai == 0 {
		//validation = true
		http.Error(w, "nilai minimal 0 dan maksimal 100", http.StatusBadRequest)
	}

	//	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mov models.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := nilai.Insert(ctx, mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mov models.NilaiMahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idNilai = ps.ByName("id")

	if err := nilai.Update(ctx, mov, idNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idNilai = ps.ByName("id")
	if err := nilai.Delete(ctx, idNilai); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

// func GetNilai(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		ctx, cancel := context.WithCancel(context.Background())

// 		defer cancel()

// 		mahasiswas, err := nilai.GetAll(ctx)

// 		if err != nil {
// 			kesalahan := map[string]string{
// 				"error": fmt.Sprintf("%v", err),
// 			}

// 			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
// 			return
// 		}

// 		utils.ResponseJSON(w, mahasiswas, http.StatusOK)
// 		return
// 	}

// 	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
// 	return
// }

// // PostMahasiswa
// func PostNilai(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {

// 		if r.Header.Get("Content-Type") != "application/json" {
// 			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
// 			return
// 		}

// 		ctx, cancel := context.WithCancel(context.Background())
// 		defer cancel()

// 		var mhs models.NilaiMahasiswa

// 		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
// 			utils.ResponseJSON(w, err, http.StatusBadRequest)
// 			return
// 		}

// 		if err := nilai.Insert(ctx, mhs); err != nil {
// 			utils.ResponseJSON(w, err, http.StatusInternalServerError)
// 			return
// 		}

// 		nilai := r.URL.Query().Get("nilai")

// 		if nilai == "100" {
// 			utils.ResponseJSON(w, "nilai  tidak boleh 100", http.StatusBadRequest)
// 			return
// 		}

// 		res := map[string]string{
// 			"status": "Succesfully",
// 		}

// 		utils.ResponseJSON(w, res, http.StatusCreated)
// 		return
// 	}

// 	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
// 	return
// }

// // UpdateMahasiswa
// func UpdateNilai(w http.ResponseWriter, r *http.Request, ) {
// 	if r.Method == "PUT" {

// 		if r.Header.Get("Content-Type") != "application/json" {
// 			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
// 			return
// 		}

// 		ctx, cancel := context.WithCancel(context.Background())
// 		defer cancel()

// 		var mhs models.NilaiMahasiswa

// 		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
// 			utils.ResponseJSON(w, err, http.StatusBadRequest)
// 			return
// 		}
// 		var idNilai = ps.ByName("id")

// 		if err := nilai.Update(ctx, mhs, idNilai); err != nil {
// 			kesalahan := map[string]string{
// 				"error": fmt.Sprintf("%v", err),
// 			}

// 			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
// 			return
// 		}

// 		res := map[string]string{
// 			"status": "Succesfully",
// 		}

// 		utils.ResponseJSON(w, res, http.StatusCreated)
// 		return
// 	}

// 	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
// 	return
// }
// func DeleteNilai(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == "DELETE" {

// 		ctx, cancel := context.WithCancel(context.Background())
// 		defer cancel()

// 		var mhs models.NilaiMahasiswa

// 		id := r.URL.Query().Get("id")

// 		if id == "" {
// 			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
// 			return
// 		}
// 		mhs.ID, _ = strconv.Atoi(id)

// 		if err := nilai.Delete(ctx, mhs); err != nil {

// 			kesalahan := map[string]string{
// 				"error": fmt.Sprintf("%v", err),
// 			}

// 			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
// 			return
// 		}

// 		res := map[string]string{
// 			"status": "Succesfully",
// 		}

// 		utils.ResponseJSON(w, res, http.StatusOK)
// 		return
// 	}

// 	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
// 	return
// }
