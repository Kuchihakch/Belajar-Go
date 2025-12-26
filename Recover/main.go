package main

import "fmt"

func endTask() {
	fmt.Println("Aplikasi Diberhentikan")
	status := recover()
	fmt.Println("Terjadi Error:", status)
}

func runTask(error bool) {
	defer endTask()
	if error {
		panic("ERROR")
	}
	// Recover dalam defer
	/* if recover() == "ERROR" {
		fmt.Println("Terjadi Error")
	} */
}

func main() {
	runTask(true)
	fmt.Println("Jalan Terus")
}
