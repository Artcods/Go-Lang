package main

import "fmt"

/* 
Membuat beberapa data untuk disimpan 
dalam satu variabel sesuai daya tampung array-nya
*/

func main () {
	var names [3]string // Array : daya tampung/kapasitas data 3

	// isi data 
	names[0] = "Hamid"
	names[1] = "Abdul"
	names[3] = "Abdurrahman"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat secara langsung

	var values = [3]int {100, 200, 300,}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	/* 
	Jika data tidak diisi maka default
	*/

}