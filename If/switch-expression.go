package main

import "fmt"

func main () {
	name := "hamid"

	switch name {
	case "hamid" :
		fmt.Println("Hello World")
	
	case "abdul" :
		fmt.Println("Hllo World")
		
	default :
		fmt.Println("salah")
		
	}

	/* 
	Short Statement
	*/

	switch length := le(name); length < 5 {
	case true :
		fmt.Println(Hello World)
	case false :
		fmt.Println(Salah/False)
	}

	/* 
	switch tanpa kondisi

	kondisi di switch expression tidak wajib

	jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut
	di setiap case nya
	*/

	lenght1 := len(name)

	switch {
	case length > 10 :
		fmt.Println("Hello World")

	case length > 5 :
		fmt.Println("Hello World")
	
	default :
		fmt.Println("False")
	}
}