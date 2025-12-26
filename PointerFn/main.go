package main

import (
	"fmt"
)

// kalau seperti ini nggak berubah karena bukan pointer, melainkan value maka dicopy nggak memutasi
/* func ChangeMhsJur(mhs Mahasiswa) {
	mhs.Jur = "SI"
} */

// Harusnya seperti ini sudah merujuk ke pointer, sehingga bisa berubah
func ChangeMhsJur(mhs *Mahasiswa) {
	mhs.Jur = "SI"
}

func (mhs *Mahasiswa) Graduation() {
	mhs.Nama += " S.Kom"
}

func main() {
	mahasiswa1 := Mahasiswa{
		Nama: "Kuchi",
		Nim:  "046",
		Jur:  "TI",
	}
	ChangeMhsJur(&mahasiswa1)
	fmt.Println(mahasiswa1)

	mahasiswa1.Graduation()
	fmt.Println(mahasiswa1)
}
