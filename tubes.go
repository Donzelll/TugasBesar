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
	var n int

	fmt.Println(waktu(), " Selamat datang di program kami")

	fmt.Println()

	fmt.Println("Silahkan masukan data")
	bacaData(&data, &n)

	printData(data, n)
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

func bacaData(data *tabPolusi, n *int) {
	var tempKota string

	fmt.Scan(&tempKota)
	for tempKota != "send" {
		if tempKota != "send" {
			data[*n].kota = tempKota
			fmt.Scan(&data[*n].polusi)
			*n++
			fmt.Scan(&tempKota)
		}
	}
}

func printData(data tabPolusi, n int) {
	var status string
	var i int

	fmt.Printf("Berikut Data Polusi\n%16s %10s %35s\n", "Nama Kota", "Polusi", "Status Udara")

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
