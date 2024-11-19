package main

import ("fmt")

type Biodata struct {
	Nama string
	Umur int
	Tinggi int
	JenisKelamin string
}

func (b *Biodata) Detail() {
	fmt.Printf("Nama dia adalah %s dan berjenis kelamin %s/n", b.Nama, b.JenisKelamin)
}

func main() {
	Biodata1 := &Biodata{
		Nama: "Dwi Prasetia",
		Umur: 23,
		Tinggi: 170,
		JenisKelamin: "Laki-Laki",
	}

	Biodata1.Detail()
}