package main

import "fmt"

type Mobil struct {
    merk    string
    tahun   int
    warna   string
    bensin  int
}

// Method untuk struct Mobil
func (m *Mobil) Nyalakan() {
    fmt.Printf("Mobil %s dinyalakan\n", m.merk)
}

func (m *Mobil) Isi(liter int) {
    m.bensin += liter
    fmt.Printf("Bensin ditambah %d liter. Total bensin sekarang: %d liter\n", liter, m.bensin)
}

func (m *Mobil) Info() {
    fmt.Printf("Informasi Mobil:\n")
    fmt.Printf("Merk: %s\n", m.merk)
    fmt.Printf("Tahun: %d\n", m.tahun)
    fmt.Printf("Warna: %s\n", m.warna)
    fmt.Printf("Bensin: %d liter\n", m.bensin)
}

func main() {
    // Membuat instance mobil
    mobilSaya := &Mobil{
        merk:    "Toyota",
        tahun:   2020,
        warna:   "Merah",
        bensin:  40,
    }
    mobilDia := &Mobil{
        merk:    "Honda",
        tahun:   2020,
        warna:   "Biru",
        bensin:  40,
    }

    // Memanggil method-method
    mobilSaya.Info()
	
    mobilDia.Info()
    fmt.Println()
    
    mobilSaya.Nyalakan()
    mobilSaya.Isi(20)
    fmt.Println()
    
    mobilSaya.Info()
}
