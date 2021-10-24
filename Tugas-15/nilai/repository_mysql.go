package nilai

import (
	"Tugas-15/config"
	"Tugas-15/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"


	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	
)

const (
	table          = "nilai"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.NilaiMahasiswa, error) {

	var nilais []models.NilaiMahasiswa

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
		var nilai models.NilaiMahasiswa
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Nama,
			&nilai.MataKuliah,
			&nilai.Nilai,
			&nilai.IndeksNilai,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		nilai.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		nilais = append(nilais, nilai)
	}
	return nilais, nil
}

// Insert Movie
func Insert(ctx context.Context, nilai models.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	//queryText := fmt.Sprintf("INSERT INTO %v (nama, matakuliah, created_at, updated_at) values('%v',%v, NOW(), NOW())", table,
	queryText := fmt.Sprintf("INSERT INTO %v (nama, matakuliah, indeksnilai, nilai) values('%v','%v', '%v', %v)", table,
		nilai.Nama,
		nilai.MataKuliah,
		GetIndeksNilai(int(nilai.Nilai)),
		nilai.Nilai)
	//switch {
	// case nilai.Nilai > 100:
	// 	//validation = true
	// 	http.Error(w, "Nilai melebihi 100, Tidak diperbolehkan", http.StatusBadRequest)
	// case nilai.Nilai >= 80:
	// 	nilai.IndeksNilai = "A"
	// case nilai.Nilai >= 70 && nilai.Nilai < 80:
	// 	nilai.IndeksNilai = "B"
	// case nilai.Nilai >= 60 && nilai.Nilai < 70:
	// 	nilai.IndeksNilai = "C"
	// case nilai.Nilai >= 50 && nilai.Nilai < 60:
	// 	nilai.IndeksNilai = "D"
	// case nilai.Nilai < 50:
	// 	nilai.IndeksNilai = "E"
	// }
	_, err = db.ExecContext(ctx, queryText)

	// switch {
	// case nilai.Nilai > 100:
	// 	//validation = true
	// 	//http.Error(w, "Nilai melebihi 100, Tidak diperbolehkan", http.StatusBadRequest)
	// case nilai.Nilai >= 80:
	// 	nilai.IndeksNilai = "A"
	// case nilai.Nilai >= 70 && nilai.Nilai < 80:
	// 	nilai.IndeksNilai = "B"
	// case nilai.Nilai >= 60 && nilai.Nilai < 70:
	// 	nilai.IndeksNilai = "C"
	// case nilai.Nilai >= 50 && nilai.Nilai < 60:
	// 	nilai.IndeksNilai = "D"
	// case nilai.Nilai < 50:
	// 	nilai.IndeksNilai = "E"
	// }

	if err != nil {
		return err
	}
	return nil
}

func GetIndeksNilai(angka int) (nilai string) {
	if angka >= 80 {
		nilai = "A"
	} else if angka >= 70 && angka < 80 {
		nilai = "B"
	} else if angka >= 60 && angka < 70 {
		nilai = "C"
	} else if angka >= 50 && angka < 60 {
		nilai = "D"
	} else {
		nilai = "E"
	}
	return

}

// Update Movie
func Update(ctx context.Context, nilai models.NilaiMahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set indeksnilai ='%s',nilai = %d where id = %s",
		table,
		nilai.IndeksNilai,
		nilai.Nilai,
		id,
	)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// Delete Movie
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
