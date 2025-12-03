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

/*
NILAI KOSONG

0 untuk tipe numerik,
false untuk tipe boolean, dan
"" (string kosong) untuk string.
*/

/*
KONVERSI TIPE
Ekspresi T(v): artinyta mengkonversi nilai v ke tipe T
misalnya:
var i int = 42 || i := 42
var f float64 = float64(i) ||
var u uint = uint(f)

contoh di main()
*/

/*
INFERENSI TIPE
misalnya:
var i int
j := i // j adalah sebuah int

Namun jika,sisi kakan berupa konstanta numerik, misalnya (3.142) mak tipe data akan berubah menjadi float
*/

/*
KONSTANTA
Konstanta dideklarasikan seperti variabel, tapi dengan kata const.
Konstanta tidak bisa dideklarasikan dengan sintaks (:=).
*/
const Pi = 3.14

/*
KONSTANTA NUMERIK
Konstanta numerik adalah nilai dengan presisi-tinggi.
Konstanta yang tidak bertipe menggunakan tipe yang dibutuhkan sesuai dengan konteksnya.
*/
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

/*
PERULANGAN (FOR)
dalam bahasa Go hanya memiliki 1 perulangan yaitu FOR.
Syntaxnya yaitu:
for i:= 0; i < 10; i++{}

tanpa tanda kurung untuk komponen perintah.

JIKA tanpa perintah setelah for, disebut infinity loop,
seingga untuk keluar dari loop hanya dapat menggunakan break atau return.

contoh penerapan ada di main.
*/

/*
KONDISI IF
Perintah IF tidak harus menggunakan ditutupi dengan tanda kurung

contohnya yaitu:
*/
func sqrt(x float64) string {
	if x < 0 {
		// sqrt (square root) digunakan untuk menghitung akar kuadrat
		// (square root) dari sebuah bilangan bertipe float64
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*
*
KONDISI IF SINGKAT
contohnya:
*/
func pow(x, n, lim float64) float64 {
	// function math.pow digunakan untuk menghitung pangkat
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

/**
KONDISI IF-ELSE
*/

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// v tidak dapat digunakan disini
	return lim
}

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
	// Kalau menggunakan "%g" maka menggunakan perintah Printf
	// %g : digunakan untuk format generik untuk angka floating point
	// %T : digunakan untuk menampilkan tipe data
	// %v : digunakan untuk menampilkan nilai default
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

	// NILAI KOSONG
	var h int
	var f float64
	var t bool
	var s string
	fmt.Printf("%v %v %v %q\n", h, f, t, s)

	// KONVERSI TIPE
	var x, y int = 3, 4
	var l float64 = math.Sqrt(float64(x*x + y*y))
	var w uint = uint(l)
	fmt.Println(x, y, w)

	// INFERENSI TIPE
	o := 3.2890342 // ubahlah nilai v!
	fmt.Printf("variabel o bertipe %T\n", o)

	// contoh penerapan KONSTANTA
	const World = "Hola"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// KONSTANTA NUMERIK
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// PERULANGAN (LOOPING)
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// looping juga dapat dibuat seperti while yaitu sebagai berikut:
	total := 1
	for total < 1000 {
		total += total
	}
	fmt.Println(total)

	// KONDISI ("if")
	fmt.Println(sqrt(2), sqrt(-4))
	//if singkat
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	// kondisi if else
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
