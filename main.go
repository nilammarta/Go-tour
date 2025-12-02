package main

import (
	// package fmt : untuk mengontrol input/output
	"fmt"
	"math/cmplx"

	/**
	PACKAGE
	Setiap program Go terbuat dari paket-paket (package).
	Aturannya, nama paket sama dengan elemen terakhir dari path import. Sebagai contohnya, paket "math/rand"
	terdiri dari berkas-berkas yang dimulai dengan perintah package rand.
	*/
	"math/rand"

	// package time : untuk mendapatkan waktu saat ini
	"time"
)

/**
IMPORT
Selain menggunakan import yang di faktorkan (dalam tanda kurung) seperti diatas,
format lainnya yaitu: menggunakan perintah import
*/
import "math"

/*
FUNCTION
*/
func add(x int, y int) int {
	return x + y
}

// atau inisiasi parameter fungsi juga dapat dijadikan sebagai:
func add1(x, y int) int {
	return x + y
}

// Fungsi dalam GO dapat mengembalikan banyak data
func swap(x, y string) (string, string) {
	return y, x
}

// Fungsi dengan return yg memiliki nama (disebut dengan "naked")
// dimana bagian ini: (x, y int) merupakan return value nya.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/*
VARIABEL
*/
var c, python, java bool

// Inisiasi Variabel u dan j
var u, j int = 1, 2

/*
	TIPE DASAR GO yaitu:

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias untuk uint8
rune // alias untuk int32

	// merepresentasikan sebuah kode Unicode

float32 float64
complex64 complex128

Tipe int, uint, dan uintptr biasanya memiliki panjang 32 bits pada
sistem 32-bit dan 64 bits pada sistem 64-bit.

	**misalnya pada variabel iniL:
*/
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// GO TOUR
	fmt.Println("Hello World!")

	// Go Playground
	fmt.Println("Selamat datang di Playground!")

	fmt.Println("Waktu sekarang adalah", time.Now())

	/**
	PAKET ("package)
	*/
	fmt.Println("Bilangan kesukaan saya adalah", rand.Intn(10))

	// IMPORT
	// Kalau menggunakan "%variabel" maka menggunakan perintah Printf
	fmt.Printf("Sekarang anda memiliki %g masalah.\n", math.Sqrt(7))

	/**
	EKSPOR
	Pada bahasa Go, sebuah nama diekspor jika diawali dengan huruf besar. Contohnya pada ekspor Pi:
	*/
	fmt.Println(math.Pi)

	//	CALL FUNCTION
	fmt.Println(add(32, 38))
	fmt.Println(add1(32, 38))
	// inisiasi parameter
	a, b := swap("Michael", "Bonjour")
	fmt.Println(a, b)
	fmt.Println(split(17))

	// VARIABEL
	var i int
	fmt.Println(i, c, python, java)
	// Inisiasi variabel
	var c, python, java = true, false, "no!"
	fmt.Println(u, j, c, python, java)
	// Deklarasi variabel singkat
	k := 3
	q, r, v := true, false, "no!"
	fmt.Println(u, j, k, q, r, v)

	// Run Tipe Dasar Go
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

}
