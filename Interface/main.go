package main

import "fmt"

func SayHello(value HasName) {
	fmt.Println("Hello " + value.GetName())
}

// struct harus mengimplementasikan kontrak di interface
func (user User) GetName() string {
	// return user.age -> tidak sesuai kontrak: return int
	return user.Name
}

func main() {
	userA := User{
		Name:   "Kuchi",
		Gender: "Male",
		age:    21,
	}
	SayHello(userA)
	// interface Kosong -> berarti kontraknya tidak ada
	fmt.Println(empty())
	fmt.Println("Apala")
}
