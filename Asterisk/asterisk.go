package main

import (
	"belajar-go/Pointer"
	"fmt"
)

func main() {
	mahasiswa1 := Pointer.Mahasiswa{
		Nama: "Felienz",
		Nim:  "046",
		Jur:  "TI",
	}
	mahasiswa2 := &mahasiswa1
	mahasiswa2.Nama = "Kuchi"
	/* fmt.Println(mahasiswa1)
	fmt.Println(mahasiswa2)
	// sekarang m2 tunjuk ke variabel baru
	mahasiswa2 = &Pointer.Mahasiswa{
		Nama: "Kazu",
		Nim:  "045",
		Jur:  "TI",
	}
	fmt.Println(mahasiswa2)
	fmt.Println(mahasiswa1) //tidak terpengaruh lagi karena pointernya beda */

	//nah kalau tetap ingin mengacu ke mahasiswa1, gunakan asterisk
	*mahasiswa2 = Pointer.Mahasiswa{
		Nama: "Kazuskem",
		Nim:  "045",
		Jur:  "TI",
	}
	// konsepnya variabel apapun yang awalnya menunjuk ke sumber awal (di sini m1), akan pindah ke resource baru karena asterisk
	// sehingga m1 dan m2 akan menunjuk ke m2
	fmt.Println(mahasiswa2)
	fmt.Println(mahasiswa1) // harusnya terpengaruh
}
