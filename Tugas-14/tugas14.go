package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"matakuliah"`
	IndeksNilai string `json:"indeksnilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))

		return
	})
}

func GetNilaiSiswa(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		dataSiswas, err := json.Marshal(nilaiNilaiMahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataSiswas)
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
}

func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nomor = len(nilaiNilaiMahasiswa)
	var Grade NilaiMahasiswa
	var validation = false
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {

			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Grade); err != nil {
				log.Fatal(err)
			}

		} else {
			nama := r.PostFormValue("nama")
			matkul := r.PostFormValue("matakuliah")
			lnilai := r.PostFormValue("nilai")
			nilay, _ := strconv.Atoi(lnilai)
			var IndexNilai string
			switch {
			case nilay > 100:
				validation = true
				http.Error(w, "Nilai melebihi 100, Tidak diperbolehkan", http.StatusBadRequest)
			case nilay >= 80:
				IndexNilai = "A"
			case nilay >= 70 && nilay < 80:
				IndexNilai = "B"
			case nilay >= 60 && nilay < 70:
				IndexNilai = "C"
			case nilay >= 50 && nilay < 60:
				IndexNilai = "D"
			case nilay < 50:
				IndexNilai = "E"
			}
			Grade = NilaiMahasiswa{
				Nama:        nama,
				MataKuliah:  matkul,
				Nilai:       uint(nilay),
				IndeksNilai: IndexNilai,
				ID:          uint(nomor + 1),
			}
			fmt.Println(Grade)
		}
		if !validation {
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, Grade)
			dataSiswa, _ := json.Marshal(Grade)
			fmt.Println(dataSiswa)
			w.Write(dataSiswa)
		}
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}

func main() {

	http.Handle("/post_nilai", Auth(http.HandlerFunc(PostNilai)))
	http.HandleFunc("/nilaisiswa", GetNilaiSiswa)

	fmt.Println("server running at http://localhost:7000")
	if err := http.ListenAndServe("127.0.0.1:7000", nil); err != nil {
		log.Fatal(err)

	}
}
