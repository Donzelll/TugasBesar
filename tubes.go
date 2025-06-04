package main

import (
	"fmt"
	"time"
)

const NMAX = 100

type tabPolusi [NMAX]struct {
	kota   string
	polusi int
}

func waktu() string {

	t := time.Now()
	fmt.Println(t.Hour())

	if t.Hour() < 12 {
		return "Selamat pagi!"
	} else if t.Hour() < 15 {
		return "Selamat siang!"
	} else if t.Hour() < 18 {
		return "Selamat sore!"
	} else {
		return "Selamat malam!"
	}
}

func tambahData(data *tabPolusi, n *int, kota string, polusi int) {
	fmt.Print("Masukkan nama kota / (\"send\"): ")
	fmt.Scan(&kota)
	for kota != "send" {
		if kota != "send" {
			fmt.Print("Masukkan tingkat polusi: ")
			fmt.Scan(&polusi)
			data[*n].kota = kota
			data[*n].polusi = polusi
			*n++
			fmt.Print("Masukkan nama kota / (\"send\"): ")
			fmt.Scan(&kota)
		}
	}
}

func printData(data tabPolusi, n int) {
	var status string
	var i int

	fmt.Printf("\nBerikut Data Polusi\n%16s %10s %35s\n", "Nama Kota", "Polusi", "Status Udara")

	for i = 0; i < n; i++ {

		status = dataStatus(data, i)
		fmt.Printf("%16s %10d %35s\n", data[i].kota, data[i].polusi, status)
	}
}

func dataStatus(data tabPolusi, i int) string {

	if data[i].polusi < 0 {
		return "Tidak Terdefenisi"
	} else if data[i].polusi <= 50 {
		return "Baik"
	} else if data[i].polusi <= 100 {
		return "Sedang"
	} else if data[i].polusi <= 150 {
		return "Tidak Sehat Untuk Kelompok Sentif"
	} else if data[i].polusi <= 200 {
		return "Tidak Sehat"
	} else if data[i].polusi <= 300 {
		return "Sangat Tidak Sehat"
	} else {
		return "Berbahaya"
	}
}

func deleteData(data *tabPolusi, n *int, index int) {
	if index < 0 || index >= *n {
		fmt.Println("Index tidak valid")
		return
	}

	for i := index; i < *n-1; i++ {
		data[i] = data[i+1]
	}
	*n--
	fmt.Println("Data berhasil dihapus")
}

func modifData(data *tabPolusi, n *int, index int, kota string, polusi int) {
	if index < 0 || index >= *n {
		fmt.Println("Index tidak valid")
		return
	}

	data[index].kota = kota
	data[index].polusi = polusi
	fmt.Println("Data berhasil dimodifikasi")
}
func menu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Modifikasi Data")
	fmt.Println("4. Cari Data Berdasarkan Kota")
	fmt.Println("5. Keluar")
}
func main() {
	var data tabPolusi
	var n int
	var choice int
	var kota string
	var polusi int
	var index int
	var inputKota string

	for {
		menu()
		fmt.Print("\nPilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			tambahData(&data, &n, kota, polusi)
			printData(data, n)
		case 2:
			fmt.Print("Masukkan index data yang ingin dihapus: ")
			fmt.Scan(&index)
			deleteData(&data, &n, index)
			printData(data, n)
		case 3:
			fmt.Print("Masukkan index data yang ingin dimodifikasi: ")
			fmt.Scan(&index)
			fmt.Print("Masukkan nama kota baru: ")
			fmt.Scan(&kota)
			fmt.Print("Masukkan tingkat polusi baru: ")
			fmt.Scan(&polusi)
			modifData(&data, &n, index, kota, polusi)
			printData(data, n)
		case 4:
			fmt.Print("Masukkan nama kota yang ingin dicari: ")
			fmt.Scan(&inputKota)
			cariKota(data, n, inputKota)
		case 5:
			fmt.Println("Terima kasih telah menggunakan program ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func cariKota(data tabPolusi, n int, kota string) {
	for i := 0; i < n; i++ {
		if data[i].kota == kota {
			printData(data, n)
			fmt.Printf("\nBerikut Data Kota yang anda cari\n%16s %10s %35s\n", "Nama Kota", "Polusi", "Status Udara")
			fmt.Printf("%16s %10d %35s\n", data[i].kota, data[i].polusi, dataStatus(data, i))
		}
	}
	return
}

func urutan(data tabPolusi, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if data[i].polusi > data[j].polusi {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	fmt.Println("Data telah diurutkan berdasarkan tingkat polusi.")
	printData(data, n)
}
