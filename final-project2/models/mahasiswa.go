package models

import (
	"time"
)

type (
	// Mahasiswa
	// Mahasiswa struct {
	// 	ID        int       `json:"id"`
	// 	NIM       int       `json:"nim"`
	// 	Name      string    `name:"name"`
	// 	Semester  int       `json:"semester"`
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// }

	Admin struct {
		ID int `json:"id"`
		//NIM       int       `json:"nim"`
		Name      string    `name:"name"`
		Password  string    `json:"password"`
		LoginAt   time.Time `json:"login_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Makanan struct {
		Id_Makanan    uint   `json:"id_makan"`
		Nama_Makanan  string `json:"nama_makan"`
		Stock_Makanan int    `json:"stock_makan"`
		Harga_Makanan int    `json:"harga_makan"`
		Id_Supplier   uint   `json:"id_supplier"`
	}

	Minuman struct {
		Id_Minuman    uint   `json:"id_minum"`
		Nama_Minuman  string `json:"nama_minum"`
		Stock_Minuman int    `json:"stock_minum"`
		Harga_Minuman int    `json:"harga_minum"`
		Id_Supplier   uint   `json:"id_supplier"`
	}

	Customer struct {
		Id_Customer uint   `json:"id_cust"`
		Nama_Cust   string `json:"nama_cust"`
	}

	Keranjang struct {
		Id_Keranjang uint `json:"id_keranjang"`
		Id_Makanan   uint `json:"id_makan"`
		Id_Minuman   uint `json:"id_minum"`
		Total_Harga  int  `jdon:"total_harga"`
	}

	Order struct {
		Id_Order     uint   `json:"id_order"`
		Id_Keranjang uint   `json:"id_keranjang"`
		Id_Customer  uint   `json:"id_cust"`
		Id_Shipper   string `json:"id_shipper"`
	}

	Shippers struct {
		Id_Shipper   uint   `json:"id_shipper"`
		Nama_Shipper string `json:"nama_shipper"`
	}

	Supplier struct {
		Id_Supplier   uint   `json:"id_supplier"`
		Nama_Supplier string `json:"nama_supplier"`
	}

	Users struct {
		Id_Users   uint   `json:"id_users"`
		Nama_Users string `json:"nama_users"`
		Tipe_Users string `json:"tipe_users"`
	}
)
