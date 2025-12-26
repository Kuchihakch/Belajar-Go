package main

import (
	"belajar-go/Pointer"
	"fmt"
)

func main() {
	// untuk pointer data kosong bisa pakai new (asalnya &Point{})
	mahasiswa1 := new(Pointer.Mahasiswa)
	mahasiswa2 := mahasiswa1 // ini juga pointer seperti m1, sama2 menunjuk ke Mahasiswa

	mahasiswa2.Nama = "Felienz"
	fmt.Println(mahasiswa1)
	fmt.Println(mahasiswa2) // sekarang keduanya sama2 felienz
}
