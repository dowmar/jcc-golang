package buku

import (
	"context"
	"fmt"
	"log"
	"quiz3/config"
	"quiz3/indeksu"
	"time"
	"quiz3/models"
	
)

const (
	table          = "books"
	layoutDateTime = "2006-01-02 15:04:05"
)

//GetIndex

// Insert

func Insert(ctx context.Context, bk models.Books) error {

	//var mahasiswas models.Mahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, kategori_ketebalan,created_at,updated_at) values('%v','%v','%v',%v,'Rp. %v','%v','%v','%v','%v',)", table,
		bk.Title,
		bk.Description,
		bk.Image_Url,
		bk.Release_Year,
		bk.Price,
		bk.Total_Page,
		indeksu.GetKetebalan(string(bk.Total_Page)),
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

func GetAll(ctx context.Context) ([]models.Books, error) {

	var bukus []models.Books

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var buku models.Books
		var createdAt, updatedAt string

		//var tot = indeksu.GetIndeksNilai(int(mahasiswa.Semester))
		//var Price = 
		
		if err = rowQuery.Scan(&buku.ID,
			&buku.Title,
			&buku.Description,
			&buku.Image_Url,
			&buku.Release_Year,
			&buku.Price,
			&buku.Total_Page,
			&buku.Kategori_Ketebalan,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		buku.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		buku.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		// if buku.Price != 0 {
		// 	buku.Price = string(buku.Price)
		// }


		// if bukus.Semester == 90 {
		// 	mahasiswa.NIM = 90
		// }

		bukus = append(bukus, buku)
	}
	return bukus, nil
}
