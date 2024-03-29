package main

import "fmt"

/*
1. Tipe Integer (Rentang Angka dari negatif - positif)

int8 = -128 -- 127
int16 = -32768 -- 32767
int32 = -214... -- 2147...
int64 = -92233... -- 92233...

*/

/*
T2. ipe Integer (Angka Positif Saja)

uint8 = 0 -- 255
uint16 = 0 65535
uint32 = 0 -- 42949...
uint64 = 0 -- 184467...

*/

/* 
Tipe Floating Point

float32 = 1.18 x 10^-38 -- 3,4 x 10^38
float64 = 2.23 x 10^-308 -- 1.80 x 10^-308
complex64 = real dan imaginer
complex128 =  real dan imaginer

*/

/* 
Alias

byte = uint8
rune = int 32
int = min int32
uint = min uint32

*/

func main() {
	var a float32
	var b float32
	var c float32

	a = 1.2
	b = 2.1
	c = a / b

	fmt.Println(c)
}
