package main

import "fmt"

func main() {
	type Hari string
	ListHari := [...]Hari{"senin,", "selasa,", "rabu,", "kamis,", "jumat,", "sabtu,", "minggu"}
	var slice1 = ListHari[5:]
	fmt.Println("Slice 1 Awal: ", slice1)
	//mutasi slice => array
	slice1[0] = "sabtu Baru,"
	slice1[1] = "minggu Baru,"
	fmt.Println("Slice 1 akhir: ", slice1)
	fmt.Println("Array Asli akhir: ", ListHari) // ada perubahan

	var slice2 = append(slice1, "Hari baru")
	fmt.Println("Slice 2 awal: ", slice2)
	slice2[0] = "sabtu ubah 2, "
	fmt.Println("Slice 2 akhir: ", slice2)
	fmt.Println("Array Asli setelah append: ", ListHari)
}
