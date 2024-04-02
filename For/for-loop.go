package main

import "fmt"

/* 
dalam bahasa pemrograman, biasanya ada perulangan
*/

func main () {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke -", counter)
		counter++
	}

	fmt.Println("Finish")

	/* 
	short statement

	dalam for, bisa menambahkan statement, dimana terdapat 2
	statement yang bisa tambahkan di for

	init statement, yaitu statement sebelum for di esksekusi 

	post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan
	*/


	i := 1 // init Statement
	for i <= 10 {
		fmt.Println("Perulangan ke-", i)
		i++ // post statement
	}

	// diubah menjadi

	for i := 1; i <= 10; i++ { // Init statemen U Post Statement
		fmt.Println("Perulangan ke-", i)
	}

	// For Range...


}