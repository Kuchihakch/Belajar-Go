package main

// berarti any bebas, bahkan bisa ganti any untuk interface kosong
func empty(val ...interface{}) interface{} {
	return val
}
