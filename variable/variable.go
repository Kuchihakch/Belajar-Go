package main

import "fmt"

// buat var -> wajib tipe adata
// inisialisasi data pada var -> tidak wajib
func main() {
	var teks string
	teks = "Hello Worl"
	fmt.Println(teks)

	teks = "0" // number tidak bisa diassign, jadi string saja
	fmt.Println(teks)
}
