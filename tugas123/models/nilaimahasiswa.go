package models

import (
	"time"
)

// type (
//   // Movie
//   Movie struct {
//     ID        int       `json:"id"`
//     Title     string    `json:"title"`
//     Year      int       `json:"year"`
//     CreatedAt time.Time `json:"created_at"`
//     UpdatedAt time.Time `json:"updated_at"`
//   }
// )

type NilaiMahasiswa struct {
	ID          uint      `json:"id"`
	Nmaa        string    `json:"nama"`
	MataKuliah  string    `json:"mata_kuliah"`
	Nilaii      uint      `jdon:"nilai"`
	IndeksNilai string    `json:"indeks_nilai"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
