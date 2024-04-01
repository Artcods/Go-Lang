package main

import "fmt"

/* 
Operasi untuk tipe data number

+ = pertambahan
- = pengurangan 
/ = pembagian
* = Perkalian
% = sisa bagi/modulo
*/

func main () {
	fmt.Println(10 + 20)

	var i = 10

	i += 10
	fmt.Println(i)
}

/* 
Augmented Assignment
Operasi kedirinya sendiri

a = a + 10 ---> a += 10
a = a - 10
a = a / 10
a = a * 10
a = a % 10
*/

/* 
Unary Operator
ingin menambahkan 1

++ atau a++ ---> a = a + 1
-- atau a-- ---> a = a - 1
- = negatif
+ = Positif
! kebalikan 
*/