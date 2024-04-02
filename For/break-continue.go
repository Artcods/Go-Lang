package main

import "fmt"

/* 
Break & Continue adalah kata kunci yang bisa digunakan dalam perulangan

break digunakan untuk menghentikan seluruh perulangan 

continue adalah digunakan untuk menghentikan perulangan yang berjalan,
dan langsung melanjutkan ke perulangan selanjutnya
*/

func main () {
	// Break 

	for i := 0; i < 10; i++ {
		if i == 5 {
			break // akan stop seluruh perulangan sampai 5
		}

		fmt.Println("Perulangan ke", i) // dieksekusi perulangan ke berapa
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 { // mengecek habis dibagi 2 maka lanjutkan
			continue // berarti jika yang didapatkan bilangan genap maka continue/baris code dibawah tidak dieksekusi
		}

		fmt.Println("Perulangan ke", i)
	}
}