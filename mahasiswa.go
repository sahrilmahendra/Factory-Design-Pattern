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
	fmt.Printf("Hello %s, I'm %d years old\n", m.Nama, m.Umur)
}

func (m Mahasiswa) Next5Years() {
	fmt.Printf("5 years again, I'm %d years old\n\n", m.Umur+5)
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
	mhs.Next5Years()

	NewMahasiswa("ahmad", "1601212034", "ahmad@example.com", 17).SayHello()
	NewMahasiswa("ahmad", "1601212034", "ahmad@example.com", 17).Next5Years()

	NewMahasiswa("mahendra", "1705098247", "mahen@example.com", 16).SayHello()
	NewMahasiswa("mahendra", "1705098247", "mahen@example.com", 16).Next5Years()
}
