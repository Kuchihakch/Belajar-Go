package main

import "fmt"

func main() {
	type Nilai uint8
	type Kelulusan bool
	UAS := Nilai(90)
	Absen := Nilai(75)

	var isLulus Kelulusan = UAS > 85 && Absen > 75

	if isLulus {
		fmt.Print("Ya")
	} else {
		fmt.Print("Tidak")
	}
}
