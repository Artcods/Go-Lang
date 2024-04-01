package main

import "fmt"

/*

Pada Array atau Slice, untuk mengakses data, kita menggunakan index number dimulai dari 0

Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun
kita bisa menentukan jenis tipe data index yang akan kita gunakan

sederhanakannya, Map adalah tipe data kumpulan key-value (kata kunci-nilai), dimana
kata kuncinya bersifat unik, tidak boleh sama.

berbeda dengan array dan slice, jumlah data yang kita masukkan ke dalam map boleh
sebanyak-banyaknya, asalkan kata kuncinya berbeda, jika kita gunakan kata kunci sama,
maka secara otomatis data sebelumnya akan diganti dengan data baru.

*/

func main() {
	orang := map[string]string{
		"name": "Hamid",
		"usia": "18",
	}

	fmt.Println(orang)
	fmt.Println(orang["name"])
	fmt.Println(orang["usia"])

	/* 
	len(map) = untuk mendapatkan jumlah data di map
	map[key] = mengambil data di map dengan key
	map[key] = value, mengubah data di map dengan key
	make(map[typedata]typevalue) = membuat map baru
	delete(map, key) = mengahapus data di map dengan key 
	*/

	book := make(map[string]string)
	book["title"] = "Buku"
	book["author"] = "Hamid"
	book["salah"] = "salah"

	delete(book, "salah")

	fmt.Println(book)
}
