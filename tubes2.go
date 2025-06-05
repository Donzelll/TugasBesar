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
	var n, choice, polusi, mode int
	var kota string
	var inputKota string

	fmt.Println("\n", waktu())

	for {
		menu()
		fmt.Print("\nPilih menu: ")
		modeSearch := true
		fmt.Scan(&choice)
		switch choice {
		case 1:
			tambahData(&data, &n)
			printData(data, n, modeSearch)
		case 2:
			tambahBanyakData(&data, &n)
			printData(data, n, modeSearch)
		case 3:
			fmt.Print("Masukkan nama kota yang ingin dihapus: ")
			fmt.Scan(&kota)
			deleteData(&data, &n, kota, modeSearch)
			printData(data, n, modeSearch)
		case 4:
			fmt.Print("Masukkan nama kota yang ingin dimodifikasi: ")
			fmt.Scan(&kota)
			fmt.Print("Masukkan tingkat polusi baru: ")
			fmt.Scan(&polusi)
			modifData(&data, &n, kota, polusi, modeSearch)
			printData(data, n, modeSearch)
		case 5:
			fmt.Print("Masukkan nama kota yang ingin dicari: ")
			fmt.Scan(&inputKota)
			cariKota(data, n, inputKota, modeSearch)
		case 6:
			insertionSortNaik(&data, n)
			printData(data, n, modeSearch)
		case 7:
			insertionSortTurun(&data, n)
			printData(data, n, modeSearch)
		case 8:
			fmt.Println("Mode pencarian: (1) Sequential\n(2) Binary")
			fmt.Scan(&mode)
			if mode == 1 {
				modeSearch = true
			} else if mode == 2 {
				modeSearch = false
			} else {
				fmt.Println("Pilihan tidak valid. Kembali ke menu utama.")
			}
		case 9:
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
	fmt.Println("2. Tambah Banyak Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Modifikasi Data")
	fmt.Println("5. Cari Data Berdasarkan Kota")
	fmt.Println("6. Urutkan Data Naik Berdasarkan Polusi")
	fmt.Println("7. Urutkan Data Turun Berdasarkan Polusi")
	fmt.Println("8. Mode Pencarian")
	fmt.Println("9. Keluar")
}

func waktu() string {

	t := time.Now()

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

// Sequential Search
func sequentialSearch(data tabPolusi, n int, kota string) int {
	var i, index int
	index = -1
	for index == -1 && i < n {
		if data[i].kota == kota {
			index = i
		}
		i++
	}
	return index
}

func searchData(data tabPolusi, n int, searchMode bool, kota string) int {
	if searchMode {
		return sequentialSearch(data, n, kota)
	} else {
		insertionSortNaik(&data, n)
		return binarySearch(data, n, kota)
	}
}

// Binary Search
func binarySearch(data tabPolusi, n int, kota string) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].kota == kota {
			return mid
		} else if data[mid].kota < kota {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func printData(data tabPolusi, n int, modeSearch bool) {
	var status string
	var i int

	if modeSearch {
		fmt.Println("\nMode Pencarian: Sequential")
	} else {
		fmt.Println("\nMode Pencarian: Binary")
	}

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

func deleteData(data *tabPolusi, n *int, kota string, modeSearch bool) {
	index := searchData(*data, *n, modeSearch, kota)
	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	for index < *n-1 {
		(*data)[index] = (*data)[index+1]
		index++
	}
	*n--
	fmt.Println("Data berhasil dihapus")
}

func modifData(data *tabPolusi, n *int, kota string, polusi int, modeSearch bool) {
	index := searchData(*data, *n, modeSearch, kota)
	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return ``
	}

	data[index].kota = kota
	data[index].polusi = polusi
	fmt.Println("Data berhasil dimodifikasi")
}

func cariKota(data tabPolusi, n int, kota string, modeSearch bool) {
	index := searchData(data, n, modeSearch, kota)
	printData(data, n, modeSearch)
	if index == -1 {
		fmt.Println("Kota tidak ditemukan.")
	} else {
		status := dataStatus(data, index)
		fmt.Printf("Berikut adalah data Kota yang anda cari: \nKota: %s, Polusi: %d, Status Udara: %s\n", data[index].kota, data[index].polusi, status)
	}
}

func insertionSortNaik(data *tabPolusi, n int) {
	var pass, i int
	var temp tabPolusi
	pass = 1
	for pass < n {
		i = pass
		temp[0] = (*data)[i]
		for i > 0 && (*data)[i-1].polusi > temp[0].polusi {
			(*data)[i] = (*data)[i-1]
			i--
		}
		(*data)[i] = temp[0]
		pass++
	}
}

func insertionSortTurun(data *tabPolusi, n int) {
	var pass, i int
	var temp tabPolusi
	pass = 1
	for pass < n {
		i = pass
		temp[0] = (*data)[i]
		for i > 0 && (*data)[i-1].polusi < temp[0].polusi {
			(*data)[i] = (*data)[i-1]
			i--
		}
		(*data)[i] = temp[0]
		pass++
	}
}

func tambahBanyakData(data *tabPolusi, n *int) {
	var tempKota string
	fmt.Print("Masukkan nama kota dan polusi / (\"send\"): ")
	fmt.Scan(&tempKota)
	for tempKota != "send" {
		if tempKota != "send" {
			data[*n].kota = tempKota
			fmt.Scan(&data[*n].polusi)
			*n++
			fmt.Scan(&tempKota)
		}
	}
	fmt.Println("Data berhasil ditambahkan")
}
