package mahasiswa

import (
	"api-test/config"
	"api-test/models"
	"context"
	"database/sql"

	//"errors"
	"fmt"
	"log"

	//"strconv"
	"time"
)

const (
	//table  = "mahasiswa"
	table1 = "tb_admin"
	table2 = "tb_makan"
	table3 = "tb_minum"
	table4 = "tb_keranjang"
	table5 = "tb_customer"
	table6 = "tb_order"
	table7 = "tb_shipper"
	table8 = "tb_supplier"
	//table9 = "tb_users"

	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Admin, error) {

	var admins []models.Admin

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table1)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var admin models.Admin
		var loginAt, updatedAt string

		if err = rowQuery.Scan(
			&admin.ID,
			//&mahasiswa.NIM,
			&admin.Name,
			&admin.Password,
			&loginAt,
			&updatedAt); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		admin.LoginAt, err = time.Parse(layoutDateTime, loginAt)

		if err != nil {
			log.Fatal(err)
		}

		admin.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		admins = append(admins, admin)
	}

	return admins, nil
}

func GetAllmakan(ctx context.Context) ([]models.Makanan, error) {

	var makanans []models.Makanan

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id_makan DESC", table2)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var makanan models.Makanan

		if err = rowQuery.Scan(
			&makanan.Id_Makanan,
			&makanan.Nama_Makanan,
			&makanan.Stock_Makanan,
			&makanan.Harga_Makanan,
			&makanan.Id_Supplier); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		makanans = append(makanans, makanan)
	}

	return makanans, nil
}

func GetAllminum(ctx context.Context) ([]models.Minuman, error) {

	var minumans []models.Minuman

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id_minum DESC", table3)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var minuman models.Minuman

		if err = rowQuery.Scan(
			&minuman.Id_Minuman,
			&minuman.Nama_Minuman,
			&minuman.Stock_Minuman,
			&minuman.Harga_Minuman,
			&minuman.Id_Supplier); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		minumans = append(minumans, minuman)
	}

	return minumans, nil
}

func GetAllKeranjang(ctx context.Context) ([]models.Keranjang, error) {

	var keranjangs []models.Keranjang

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id_keranjang DESC", table4)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		// var minuman models.Minuman
		// var makanan models.Makanan
		var keranjang models.Keranjang

		if err = rowQuery.Scan(
			&keranjang.Id_Keranjang,
			&keranjang.Id_Makanan,
			&keranjang.Id_Minuman,
			&keranjang.Total_Harga); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		keranjangs = append(keranjangs, keranjang)
	}

	return keranjangs, nil
}

func GetAllCustomer(ctx context.Context) ([]models.Customer, error) {

	var customers []models.Customer

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id_cust DESC", table5)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		// var minuman models.Minuman
		// var makanan models.Makanan
		var customer models.Customer

		if err = rowQuery.Scan(
			&customer.Id_Customer,
			&customer.Nama_Cust); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func GetAllOrder(ctx context.Context) ([]models.Order, error) {

	var orders []models.Order

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id_order DESC", table6)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		// var minuman models.Minuman
		// var makanan models.Makanan
		var order models.Order

		if err = rowQuery.Scan(
			&order.Id_Order,
			&order.Id_Keranjang,
			&order.Id_Customer,
			&order.Id_Shipper); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
func GetAllShipper(ctx context.Context) ([]models.Shippers, error) {

	var shippers []models.Shippers

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id_shipper DESC", table7)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		// var minuman models.Minuman
		// var makanan models.Makanan
		var shipper models.Shippers

		if err = rowQuery.Scan(
			&shipper.Id_Shipper,
			&shipper.Nama_Shipper); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		shippers = append(shippers, shipper)
	}

	return shippers, nil
}

func GetAllSupplier(ctx context.Context) ([]models.Supplier, error) {

	var suppliers []models.Supplier

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id_supplier DESC", table8)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		// var minuman models.Minuman
		// var makanan models.Makanan
		var supplier models.Supplier

		if err = rowQuery.Scan(
			&supplier.Id_Supplier,
			&supplier.Nama_Supplier); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

// Insert
func Insert(ctx context.Context, mhs models.Admin) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name, password, login_at, updated_at) values('%v','%v','%v','%v')",
		table1,
		mhs.Name,
		mhs.Password,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func InsertMakan(ctx context.Context, mkn models.Makanan) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id_makan, nama_makan, stock_makan, harga_makan, id_supplier) values(%v,'%v',%v,%v, %v)",
		table2,
		mkn.Id_Makanan,
		mkn.Nama_Makanan,
		mkn.Stock_Makanan,
		mkn.Harga_Makanan,
		mkn.Id_Supplier)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func InsertMinum(ctx context.Context, min models.Minuman) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id_minum, nama_minum, stock_minum, harga_minum, id_supplier) values(%v,'%v',%v,%v,%v)",
		table3,
		min.Id_Minuman,
		min.Nama_Minuman,
		min.Stock_Minuman,
		min.Harga_Minuman,
		min.Id_Supplier)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func InsertKeranjang(ctx context.Context, ker models.Keranjang) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id_keranjang, id_makan, id_minum,total_harga) values(%v,%v,%v,%v)",
		table4,
		ker.Id_Keranjang,
		ker.Id_Makanan,
		ker.Id_Minuman,
		ker.Total_Harga)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func InsertCustomer(ctx context.Context, cus models.Customer) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id_cust, nama_cust) values(%v,'%v')",
		table5,
		cus.Id_Customer,
		cus.Nama_Cust)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func InsertOrder(ctx context.Context, ord models.Order) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id_order, id_keranjang,id_cust,id_shipper) values(%v,%v,%v,%v)",
		table6,
		ord.Id_Order,
		ord.Id_Keranjang,
		ord.Id_Customer,
		ord.Id_Shipper)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}
func InsertShipper(ctx context.Context, sip models.Shippers) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id_shipper, nama_shipper) values(%v,'%v')",
		table7,
		sip.Id_Shipper,
		sip.Nama_Shipper)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}
func InsertSupplier(ctx context.Context, sup models.Supplier) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id_supplier, nama_supplier) values(%v,'%v')",
		table8,
		sup.Id_Supplier,
		sup.Nama_Supplier)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Update
// func Update(ctx context.Context, mhs models.Mahasiswa) error {

// 	db, err := config.MySQL()

// 	if err != nil {
// 		log.Fatal("Can't connect to MySQL", err)
// 	}

// 	queryText := fmt.Sprintf("UPDATE %v set nim = %d, name ='%s', semester = %d, updated_at = '%v' where id = '%d'",
// 		table,
// 		mhs.NIM,
// 		mhs.Name,
// 		mhs.Semester,
// 		time.Now().Format(layoutDateTime),
// 		mhs.ID,
// 	)
// 	fmt.Println(queryText)

// 	_, err = db.ExecContext(ctx, queryText)

// 	if err != nil && err != sql.ErrNoRows {
// 		return err
// 	}

// 	return nil
// }
// func Update(ctx context.Context, ker models.Keranjang) error {

// 	db, err := config.MySQL()

// 	if err != nil {
// 		log.Fatal("Can't connect to MySQL", err)
// 	}

// 	queryText := fmt.Sprintf("UPDATE %v set id_makan = %d, id_minum = %d, total_harga %d where id_keranjang = '%d'",
// 		table4,
// 		ker.Id_Makanan,
// 		ker.Id_Minuman,
// 		ker.Total_Harga,
// 		//time.Now().Format(layoutDateTime),
// 		ker.Id_Keranjang,
// 	)
// 	fmt.Println(queryText)

// 	_, err = db.ExecContext(ctx, queryText)

// 	if err != nil && err != sql.ErrNoRows {
// 		return err
// 	}

// 	return nil
// }

func UpdateKeranjang(ctx context.Context, ker models.Keranjang, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set id_makan = %d, id_minum = %d, total_harga = %d where id_keranjang = %s",
		table4,
		ker.Id_Makanan,
		ker.Id_Minuman,
		ker.Total_Harga,
		id,
	)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// Delete
// func Delete(ctx context.Context, mhs models.Mahasiswa) error {

// 	db, err := config.MySQL()

// 	if err != nil {
// 		log.Fatal("Can't connect to MySQL", err)
// 	}

// 	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, mhs.ID)

// 	s, err := db.ExecContext(ctx, queryText)

// 	if err != nil && err != sql.ErrNoRows {
// 		return err
// 	}

// 	check, err := s.RowsAffected()

// 	if check == 0 {
// 		return errors.New("id tidak ada ")
// 	}

// 	return nil
// }
