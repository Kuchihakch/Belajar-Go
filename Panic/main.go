package main

import "fmt"

func endTask() {
	fmt.Println("Ini Defer")
}

func runTask(error bool) {
	//defer duluan, baru log error & panic
	defer endTask()
	if error {
		panic("Terjadi Error")
	}
}

func main() {
	runTask(true)
}
