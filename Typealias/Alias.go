package main

import "fmt"

func main() {
	type NIM string

	var mahasiswa string = "46"
	var nimMahasiswa NIM = NIM(mahasiswa)
	fmt.Print(mahasiswa)
	fmt.Print(nimMahasiswa)
}
