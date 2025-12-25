package main

import "fmt"

func main() {
	type Nilai int

	listNilai := [3]Nilai{90, 80} // pasti jumlah 3, kalau kurang default number 0
	fmt.Print(listNilai)

	// untuk bikin kapasitas tidak tentu: (harus langsung declare nilai)
	newNilai := [...]Nilai{
		1, 2, 3, 4, 5,
	}
	fmt.Print(newNilai)

	// tidak bisa
	/* var nilai2 [...]Nilai
	nilai2[0] = 1
	nilai2[1] = 2
	fmt.Print(nilai2) */
}
