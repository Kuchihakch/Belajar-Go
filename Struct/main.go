package main

import "fmt"

func main() {
	var ticket1 Ticket
	fmt.Println("Default Struct Ticket: ", ticket1)
	ticket1.Id = "1"
	ticket1.Name = "Ticket A"
	ticket1.price = 20000
	ticket1.isAvailable = true
	fmt.Println("Ticket 1: ", ticket1)

	// alternatif:
	ticket2 := Ticket{
		Id:          "2",
		Name:        "ticket B",
		price:       50000,
		isAvailable: false,
	}
	fmt.Println(ticket2)

	checkName := ticket1.getNameById("1")
	fmt.Println(checkName)
}
