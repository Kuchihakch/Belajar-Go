package main

import "fmt"

func main() {
	var num32 int32 = 32768
	var num64 int64 = int64(num32)
	var num16 int16 = int16(num32) //overflow kapasitas (balik ke batas belakang)

	fmt.Println(num32)
	fmt.Println(num64)
	fmt.Println(num16)
}
