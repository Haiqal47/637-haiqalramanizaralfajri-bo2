package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	id, nama, alamat, pekerjaan, alasan string
}

func main() {
	var biodata = []Biodata{
		{id: "1", nama: "Haiqal", alamat: "Bogor", pekerjaan: "backend", alasan: "Saya ingin tahu lebih jauh tentang penggunaan Golang untuk menambah pengetahuan"},
		{id: "2", nama: "Irfan", alamat: "Jakrta", pekerjaan: "frontend", alasan: "Nambah ilmu"},
		{id: "3", nama: "Ghifari", alamat: "Jakarta", pekerjaan: "data engineer", alasan: "Menambah Portofolio"},
		{id: "4", nama: "Hasan", alamat: "Jakarta", pekerjaan: "data analyst", alasan: "Menambah ilmu di perkembangan pemrograman"},
		{id: "5", nama: "Mail", alamat: "Jakarta", pekerjaan: "UI/UI Designer", alasan: "Menambah skill"},
	}
	var id = os.Args[1]

	printBiodata(biodata, id)
}

func printBiodata(listBiodata []Biodata, id string) {
	var errorCount = 0
	for _, v := range listBiodata {
		if v.id == id {
			fmt.Printf("Nama\t\t: %s \nPekerjaan\t: %s\nAlamat\t\t: %s\nAlasan\t\t: %s\n", v.nama, v.pekerjaan, v.alamat, v.alasan)
			break
		} else {
			errorCount++
		}
	}

	if errorCount == len(listBiodata) {
		fmt.Println("Data tidak ditemukan")
	}
}
