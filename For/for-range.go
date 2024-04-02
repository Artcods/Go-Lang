package main

import "fmt" 

/* ]
for bisa digunakan untuk melakukan iterasi terhadap semua data collection 
data collection contohnya Array, Slice dan Map
*/

func main () {

	// manual

	name := []string {"Hamid", "Abdul", "Abdurrahman"}
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	// for range

	names := []string {"Hamid", "Abdul", "Abdurrahman"}
	for index, names := range names { // jika map maka index menjadi key dan name menjadi value
		fmt.Println("Index", index, "=", names)
	}

	// jika tidak butuh index nya, hanya names saja
	names1 := []string {"Hamid", "Abdul", "Abdurrahman"}
	for _, names := range names1 { // jika map maka index menjadi key dan name menjadi value
		fmt.Println(names)
	}
}