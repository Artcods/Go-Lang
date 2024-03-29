package main

import "fmt"

/*
di Golang kadang kita butuh melakukan konversi
tipe data dari satu tipe ke tipe lain

misal ingin mengkonversi tipe data int32 ke int64
*/

func main() {
	var a int32 = 323232
	var b int64 = int64(a)

	var c int16 = int16(a)

	fmt.Println(a) 
	fmt.Println(b)
	fmt.Println(c)
}
