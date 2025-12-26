package main

import "fmt"

func main() {
	// map[keyType]valueType
	mahasiswa := map[string]string{
		"nama": "Kuchi",
		"NIM":  "046",
		"skem": "skem",
	}
	fmt.Println(mahasiswa["NIM"])
	fmt.Println(mahasiswa["nama"])
	fmt.Println(mahasiswa)
	//kalau ternyata key tidak ada direturn default value, string = ''

	//util di map
	// map[key] -> akses data
	// map[key] = value -> reassign value
	// make(map[keyType]dataType, Value) -> bikin map baru dengan cara lain
	// delete(map, key) -> menghapus total data dengan key tertentu di map
	delete(mahasiswa, "skem")
	fmt.Println(mahasiswa)
}
