package main

import "fmt"

func main() {
	type Nilai int8
	var angka = [...]Nilai{1, 2, 3, 4, 5}

	// ini slice ambil dari low ke high atau tidak peduli langsung dari awal(0) hingga akhir(n)

	var sliceAngka = angka[:4] // artinya ambil dari idx 0 sampai sebelum 4
	fmt.Println("cek slice 1: ", sliceAngka)

	// ambil semua
	var sliceAngka2 = angka[:]
	fmt.Println("cek slice 2: ", sliceAngka2)

	// ambil slice dari idx 1 sampai sebelum 3
	var sliceAngka3 = angka[1:3] // harusnya 2, 3 alias setelah idx 1 sampai sebelum 3
	//length 2 capacity 4  ptr 1
	fmt.Println("cek slice 3: ", sliceAngka3)

	// ambil dari 2 sampai akhir
	var sliceAngka4 []Nilai = angka[2:] // wujud aselinya seperti ini, tidak seperti array yang ditentukan lengthnya
	fmt.Println("cek slice 4: ", sliceAngka4)

	// util Slice :
	// len (pjg Slice),
	// cap (capacity Slice -> merujuk len dari ptr sampai akhir arr),
	// append(slice, data) -> (tambah data baru di akhir slice, kalau > cap bikin arr baru)
	// make([] dataType, len, cap) -> bikin slice baru
	// copy(dest, src) -> salin slice dari src tertentu ke tujuan dest

	// misal:
	fmt.Println(len(sliceAngka3)) // harusnya 2
	fmt.Println(cap(sliceAngka3)) // harusnya 4

	// aksi slice: (slice3 , cap: 4)
	var sliceBaru = append(sliceAngka3, 10, 20) // lennya jadi 4 = cap, tapi lanjut bisa
	fmt.Println(sliceBaru)
	fmt.Println(len(sliceBaru))

	// memutasi slice berarti memutasi array
}
