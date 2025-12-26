package Pointer

//ini nggak bisa run kalau nggak di main pkg, biasanya jadi helper
import "fmt"

func main() {
	mahasiswa1 := Mahasiswa{"Kuchi", "046", "TI"}
	mahasiswa2 := mahasiswa1
	mahasiswa2.Nama = "Felienz" // Kalo gini pass by Value nggak by ref
	// maka mahasiswa 1 harusnya tetap sama
	fmt.Println(mahasiswa1, "\n", mahasiswa2)

	//pass by ref (pointer) ke source sama
	mahasiswa3 := &mahasiswa1 // ampersand operator menunjuk ke reference
	mahasiswa3.Nama = "Felienz"
	// sekarang harusnya mahasiswa 1 berubah
	fmt.Println(mahasiswa3)
	fmt.Println(mahasiswa1)
}
