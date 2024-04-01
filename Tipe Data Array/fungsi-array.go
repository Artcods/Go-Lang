package main

import "fmt"

/* 
len(array) = panjang array
array[index] = mendapatkan data di posisi index
array[index] = value = mengubah data di posisi index
*/

func main () {
	// jika tidak ingin tahu batas kapasitas index-nya
	var values = [...]int {90, 80, 70}
	
	fmt.Println(values)

	fmt.Println(len(values))

	/* 
	tidak ada menghapus data array di go-lang
	*/
}