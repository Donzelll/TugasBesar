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

func main() {
	var data tabPolusi
	var n, choice, polusi int
	var kota, inputKota string

	for {
		menu()
		fmt.Print("\nPilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			tambahData(&data, &n)
			printData(data, n)
		case 2:
			fmt.Print("Masukkan nama kota yang ingin dihapus: ")
			fmt.Scan(&kota)
			deleteData(&data, &n, kota)
			printData(data, n)
		case 3:
			fmt.Print("Masukkan nama kota yang ingin dimodifikasi: ")
			fmt.Scan(&kota)
			fmt.Print("Masukkan tingkat polusi baru: ")
			fmt.Scan(&polusi)
			modifData(&data, &n, kota, polusi)
			printData(data, n)
		case 4:
			fmt.Print("Masukkan nama kota yang ingin dicari: ")
			fmt.Scan(&inputKota)
			cariKota(data, n, inputKota)
		case 5:
			urutanNaik(&data, n)
			printData(data, n)
		case 6:
			urutanTurun(&data, n)
			printData(data, n)
		case 7:
			fmt.Println("Terima kasih telah menggunakan program ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func menu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Modifikasi Data")
	fmt.Println("4. Cari Data Berdasarkan Kota")
	fmt.Println("5. Urutkan Data Naik Berdasarkan Polusi")
	fmt.Println("6. Urutkan Data Turun Berdasarkan Polusi")
	fmt.Println("7. Keluar")
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

func tambahData(data *tabPolusi, n *int) {
	var kota string
	var polusi int
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
	fmt.Println("Data berhasil ditambahkan")
}

func searchData(data tabPolusi, n int, kota string) int {
	for i := 0; i < n; i++ {
		if data[i].kota == kota {
			return i
		}
	}
	return -1
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

func deleteData(data *tabPolusi, n *int, kota string) {
	index := searchData(*data, *n, kota)
	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	for i := index; i < *n-1; i++ {
		data[i] = data[i+1]
	}
	*n--
	fmt.Println("Data berhasil dihapus")
}

func modifData(data *tabPolusi, n *int, kota string, polusi int) {
	index := searchData(*data, *n, kota)
	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	data[index].kota = kota
	data[index].polusi = polusi
	fmt.Println("Data berhasil dimodifikasi")
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

func urutanNaik(data *tabPolusi, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if data[i].polusi > data[j].polusi {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

func urutanTurun(data *tabPolusi, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if data[i].polusi < data[j].polusi {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}
