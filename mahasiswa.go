package main

import "fmt"

// factory pattern
type Mahasiswa struct {
	Nama  string
	Nim   string
	Email string
	Umur  int
}

func (m Mahasiswa) SayHello() {
	fmt.Printf("Hello %s", m.Nama)
}

// inisiasi object
var mhs = Mahasiswa{
	Nama:  "Sahril Mahendra",
	Nim:   "16012345678",
	Email: "sahrilmahendra@example.com",
	Umur:  18,
}

// implementasi factory pattern
// *simple factory
func NewMahasiswa(nama, nim, email string, umur int) Mahasiswa {
	return Mahasiswa{
		Nama:  nama,
		Nim:   nim,
		Email: email,
		Umur:  umur,
	}
}

func main() {
	mhs.SayHello()
}
