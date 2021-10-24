package main

import (
	"api-test/mahasiswa"
	"api-test/models"
	"api-test/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	//"strconv"

	"github.com/julienschmidt/httprouter"
)

const USERNAME = "admin"
const PASSWORD = "admin"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//handlefunc 404 need fix
	router := httprouter.New()
	//router.Handle("/",Auth(router.GET(GetAdmin)))
	router.GET("/admin", GetAdmin)
	router.GET("/makanan", GetMakan)
	router.GET("/minuman", GetMinum)
	router.GET("/keranjang", GetKeranjang)
	router.GET("/login", GetCustomer)
	router.GET("/order", GetOrder)
	router.GET("/shipper", GetShipper)
	router.GET("/supplier", GetSupplier)
	router.POST("/makanan/create", PostMakan)
	router.POST("/minuman/create", PostMinum)
	router.POST("/keranjang/create", PostKeranjang)
	router.POST("/login/create", PostCustomer)
	router.POST("/order/create", PostOrder)
	router.POST("/shipper/create", PostShipper)
	router.POST("/supplier/create", PostSupplier)
	router.POST("/admin/create", PostAdmin)
	router.PUT("/keranjang/:id/update", UpdateKeranjang)
	router.NotFound = http.FileServer(http.Dir("public"))
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":"+port, router))
	
}
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
		//	return
	})
}



func GetAdmin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	uname, pwd, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("Username atau Password tidak boleh kosong"))
		return
	}

	if uname == "admin" && pwd == "admin" {
		//http.Handle.ServeHTTP(w, r).
		barang, err := mahasiswa.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	
	utils.ResponseJSON(w, barang, http.StatusOK)
		return
	}
	

	w.Write([]byte("Username atau Password tidak sesuai"))
	return

}

func GetMakan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	makanan, err := mahasiswa.GetAllmakan(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, makanan, http.StatusOK)

}

func GetMinum(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	minuman, err := mahasiswa.GetAllminum(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, minuman, http.StatusOK)

}

func GetKeranjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	keranjang, err := mahasiswa.GetAllKeranjang(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, keranjang, http.StatusOK)

}
func GetCustomer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	keranjang, err := mahasiswa.GetAllCustomer(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, keranjang, http.StatusOK)

}
func GetOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	keranjang, err := mahasiswa.GetAllOrder(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, keranjang, http.StatusOK)

}

func GetShipper(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	shipper, err := mahasiswa.GetAllShipper(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, shipper, http.StatusOK)

}

func GetSupplier(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	supplier, err := mahasiswa.GetAllSupplier(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, supplier, http.StatusOK)

}



// PostMahasiswa
func PostAdmin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mhs models.Admin

		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := mahasiswa.Insert(ctx, mhs); err != nil {
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
	//return
}

func PostMakan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mkn models.Makanan

		if err := json.NewDecoder(r.Body).Decode(&mkn); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := mahasiswa.InsertMakan(ctx, mkn); err != nil {
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
	//return
}
func PostMinum(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var min models.Minuman

		if err := json.NewDecoder(r.Body).Decode(&min); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := mahasiswa.InsertMinum(ctx, min); err != nil {
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
	//return
}

func PostKeranjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var ker models.Keranjang

		if err := json.NewDecoder(r.Body).Decode(&ker); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := mahasiswa.InsertKeranjang(ctx, ker); err != nil {
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
	//return
}

func PostCustomer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var ker models.Customer

		if err := json.NewDecoder(r.Body).Decode(&ker); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := mahasiswa.InsertCustomer(ctx, ker); err != nil {
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
	//return
}

func PostOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var ord models.Order

		if err := json.NewDecoder(r.Body).Decode(&ord); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := mahasiswa.InsertOrder(ctx, ord); err != nil {
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
	//return
}

func PostShipper(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var sip models.Shippers

		if err := json.NewDecoder(r.Body).Decode(&sip); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := mahasiswa.InsertShipper(ctx, sip); err != nil {
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
	//return
}

func PostSupplier(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var sup models.Supplier

		if err := json.NewDecoder(r.Body).Decode(&sup); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := mahasiswa.InsertSupplier(ctx, sup); err != nil {
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
	//return
}

// UpdateMahasiswa
// func UpdateKeranjang(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "PUT" {

// 		if r.Header.Get("Content-Type") != "application/json" {
// 			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
// 			return
// 		}

// 		ctx, cancel := context.WithCancel(context.Background())
// 		defer cancel()

// 		var mhs models.Mahasiswa

// 		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
// 			utils.ResponseJSON(w, err, http.StatusBadRequest)
// 			return
// 		}

// 		if err := mahasiswa.Update(ctx, mhs); err != nil {
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
// 	//return
// }
// func UpdateKeranjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	if r.Method == "PUT" {

// 		if r.Header.Get("Content-Type") != "application/json" {
// 			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
// 			return
// 		}

// 		ctx, cancel := context.WithCancel(context.Background())
// 		defer cancel()

// 		var mhs models.Keranjang

// 		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
// 			utils.ResponseJSON(w, err, http.StatusBadRequest)
// 			return
// 		}

// 		if err := mahasiswa.Update(ctx, mhs); err != nil {
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
// 	//return
// }

// UpdateMovie
func UpdateKeranjang(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mov models.Keranjang

	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMovie = ps.ByName("id")

	if err := mahasiswa.UpdateKeranjang(ctx, mov, idMovie); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// DeleteMahasisw
// func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == "DELETE" {

// 		ctx, cancel := context.WithCancel(context.Background())
// 		defer cancel()

// 		var mhs models.Mahasiswa

// 		id := r.URL.Query().Get("id")

// 		if id == "" {
// 			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
// 			return
// 		}
// 		mhs.ID, _ = strconv.Atoi(id)

// 		if err := mahasiswa.Delete(ctx, mhs); err != nil {

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
// 	//return
// }
