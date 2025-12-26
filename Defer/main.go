package main

import "fmt"

func JalankanDulu() {
	//defer di akhir setelah func selesai meski error tetap dijalankan
	defer Defer()
	fmt.Println("Sebelum Defer")
}

func main() {
	JalankanDulu()
}
