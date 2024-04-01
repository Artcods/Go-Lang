package main

import "fmt"

/* 
Function slice

len(slice) = untuk mendapatkan panjang

cap(slice) = untuk mendapatka kapasitas

append(slide, data) - membuat slice baru dengan
menambahkan data ke posisi terakhir slice, jika kapasitas sudah penuh,
maka akan membuat array baru.

make([]typedata, length, capacity) = membuat slice baru

copy(destination, source) = menyalis slice dari source ke destination
*/

func main () {
	// Append case

	days := [...]string {"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// Buat slice
	daySlice := days[5:]

	// Ganti Data Didalam Slice
	daySlice[0] = "Ganti"
	daySlice[1] = "Ganti"

	// Membuat Array baru karena sudah penuh yaitu menambahkan Hari Libur
	daySlice1 := append(daySlice, "Hari Libur") // Menambahkan data ke slice

	// Ganti Data Didalam Slice 2
	daySlice1[0] = "Ganti"
	fmt.Println(daySlice1)

	// Data Array Lama
	fmt.Println(days)

	// Data Array Baru dibikin
	fmt.Println(daySlice1)  // Tidak Berubah
	

	// Make case

	// Membuat slice/deklarasi slice terlebih dahulu
	newSlice := make([]string, 2, 5) // Panjang 2, dan capasitas 5
	newSlice[0] = "Hamid"
	newSlice[1] = "Abdul"
	fmt.Println(newSlice)

	// mengecek len dan cap (slice 1)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))


	// menambah data menggunakan append karena panjangnya 2
	newSlice1 := append(newSlice, "Abdurrahman")
	fmt.Println(newSlice1)

	// mengecek len dan cap (slice 2)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))


	/* 
	Kapasitasnya sama dengan slice 1 karena kapasitas 5
	tidak melebihi batasnya. kita append hanya 1 data yaitu 3 array saja
	*/


	// Pembuktian 
	newSlice1[0] = "Ganti"
	// Slice 2 
	fmt.Println(newSlice1)
	// Slice 2
	fmt.Println(newSlice) // index 0 akan tetap ada 
	/* 
	kesimpulan jika data yang ditambahkan tidak melebihi
	kapasitas maka array tidak dibuat baru.
	*/


	// Copy Slice Case

	fromSlice := days[:] // dari slice days
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // menuju slice baru/copy

	copy(toSlice, fromSlice) // DiCOPY!

	fmt.Println(fromSlice) // Slice yang belum di copy
	fmt.Println(toSlice) // Sudah di Copy


	// Perbedaan Array dan Slice saat membuat

	Array := [...]string {"a", "b",}
	// vs
	Slice := []int {1, 2, 3} // Slice lebih bagus

	fmt.Println(Array)
	fmt.Println(Slice)
}