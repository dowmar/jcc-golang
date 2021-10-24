package models

import (
	"time"
)

type (
	// Mahasiswa
	Books struct {
		ID                 int       `json:"id"`
		Title              string    `json:"title"`
		Description        string    `json:"description"`
		Image_Url          string    `json:"image_url"`
		Release_Year       int       `json:"release_year"`
		Price              string       `json:"price"`
		Total_Page         string    `json:"total_page"`
		Kategori_Ketebalan string    `json:"Kategori_Ketebalan"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
	}
)

// title(tipe data string)
// description (tipe data string)
// image_url(tipe data string)
// release_year(tipe data int)
// price (tipe data int)
// total_page(tipe data string)
// created_at (tipe data time)
// updated_at (tipe data time)
