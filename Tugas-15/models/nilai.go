package models

import (
	"time"
)

// type (
// 	// Movie
// 	Movie struct {
// 		ID        int       `json:"id"`
// 		Title     string    `json:"title"`
// 		Year      int       `json:"year"`
// 		CreatedAt time.Time `json:"created_at"`
// 		UpdatedAt time.Time `json:"updated_at"`
// 	}
// )

type NilaiMahasiswa struct {
	ID          uint      `json:"id"`
	Nama        string    `json:"nama"`
	MataKuliah  string    `json:"matakuliah"`
	IndeksNilai string    `json:"indeksnilai"`
	Nilai       uint      `json:"nilai"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
