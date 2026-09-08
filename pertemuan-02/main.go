package main

import "fmt"

// TODO(Level 1): lihat SOAL.md untuk kontrak lengkap tiap fungsi di bawah.
// Ganti setiap "panic" dengan implementasi yang benar.

func HitungSubtotal(qty int, hargaSatuan float64) float64 {
	return float64(qty) * hargaSatuan
}

func HitungTotalPesanan(qty []int, hargaSatuan []float64) float64 {
	if len(qty) != len(hargaSatuan) {
		return 0
	}
	var total float64
	for i := 0; i < len(qty); i++ {
		total += HitungSubtotal(qty[i], hargaSatuan[i])
	}
	return total
}

func TerapkanPajak(total float64, tarifPajak float64) float64 {
	pajak := total * tarifPajak
	return total + pajak
}

func HitungDiskon(total float64) float64 {
	switch {
	case total >= 1000000:
		return total * 10 / 100
	case total >= 500000:
		return total * 5 / 100
	default:
		return 0
	}
}

func TotalSetelahDiskon(qty []int, hargaSatuan []float64, tarifPajak float64) float64 {
	total := HitungTotalPesanan(qty, hargaSatuan)
	totalSetelahPajak := TerapkanPajak(total, tarifPajak)
	diskon := HitungDiskon(totalSetelahPajak)
	return totalSetelahPajak - diskon
}

func ValidasiPesanan(qty []int, hargaSatuan []float64) (bool, string) {
	panic("belum diimplementasikan")
}

func TentukanStatus(total float64) string {
	panic("belum diimplementasikan")
}

func RingkasanPesanan(qty []int, hargaSatuan []float64, tarifPajak float64) string {
	panic("belum diimplementasikan")
}

// TODO(Level 9): signature ini SUDAH benar (cari tahu sendiri kenapa
// bentuknya begini - lihat SOAL.md) - tinggal implementasikan isinya.
func Total(harga ...float64) float64 {
	panic("belum diimplementasikan")
}

// TODO(Level 10, bonus): signature ini SUDAH benar (cari tahu sendiri
// kenapa ada dua nilai balik - lihat SOAL.md) - tinggal implementasikan isinya.
func HitungOngkosKirim(beratKg float64, jarakKm float64) (float64, error) {
	panic("belum diimplementasikan")
}

func main() {
	fmt.Println("Sales Order Processor - pertemuan 2")
}
