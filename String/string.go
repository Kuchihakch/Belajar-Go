package main

import "fmt"

func main() {
	var teks = "Ini String"
	var length = len((teks))
	fmt.Println(length)        // pjg string
	fmt.Print(string(teks[4])) // 'S' -> aselinya byte konversi string
}
