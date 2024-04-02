package main

import "fmt"

/*
if adalah salah satu kata kunci yang digunakan untuk percabangan

percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi

hampir di semua bahasa pemrograman mendukung if expression
*/

func main() {
	name := "hamid"

	if name == "hamid" {
		fmt.Println("Hello World")
	}

	/*
		Else Expression

		blok if akan dieksekusi kondisi if bernilai true
		kadang kita ingin melakukan eksekusi program tertentu jika bernilai false
		hali ini bisa dilakukan menggunakan else expression
	*/

	if name == "Hamid" {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Salah/False")
	}

	/* 
	Else If Expression

	kadang dalam if, kita butuh membuat beberapa
	kasus seperti ini, kita bisa menggunakan else if expression
	*/

	if name == "hamid" {
		fmt.Println("Hello World") 

	} else if name == "abdul" {
		fmt.Println("Hello World")

	} else if name == "abdurrahman" {
		fmt.Println("False")

	} else {
		fmt.Println("False") // Program False-nya
	}

	/* 
	if dengan short statement

	if mendukung short statement sebelum kondisi

	hal ini sangat cocok untuk membuat statement sederhana sebelum melakukan 
	pengecekan terhadap kondisi
	*/

	if length := len(name); length > 5 { // if (short statement; boolean)
		fmt.Println("Hello World")
	} else {
		fmt.Println("False")
	}

	/* 
	digunakan jika variabelnya digunakan hanya program if itu saja
	*/
}


