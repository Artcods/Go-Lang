package main

import "fmt"
/* 

Potongan dari data array

Bisa bertambah secara dinamis

slice dan array selalu tekoneksi = dimana slice mengatur seluruh data array

*/

/* 
Tipe data slice memiliki 3 tipe:
1. pointer = penunjuk data pertama dalam array
2. lengtg = panjang dari  slice
3. capacity = length tidak boleh melebihi kapasitasnya
*/

/* 
Membuat Slice dari Array

Array[low:high] = membuat slice dari array dimulai index low sampai index sebelum high
array[low:] = membuat slice dari array dimulai index low sampai index akhir di array
array[:high] = membuat slice dari array dimulai index 0 sampai index sebelum high
array[:] = membuat slice dair array dimulai index 0 sampai index di array
*/

func main () {
	values := [...]int {20, 30, 40, 50, 60, 70, 80}

	slice := values[3:7]
	fmt.Println(slice)
	
	slice1 := values[2:]
	fmt.Println(slice1)
	
	slice2 := values[:3]
	fmt.Println(slice2)

}