package main

import "fmt"

type Tagihan struct {
	Nama       string
	Kategori   string
	Nominal    int
	JatuhTempo int
	Lunas      bool
}

var data [100]Tagihan
var jumlah int

func main() {
	var pilih int

	for {
		fmt.Println("\n===== SIMTAB =====")
		fmt.Println("1. Tambah Tagihan")
		fmt.Println("2. Ubah Tagihan")
		fmt.Println("3. Hapus Tagihan")
		fmt.Println("4. Tampilkan Data")
		fmt.Println("5. Sequential Search")
		fmt.Println("6. Binary Search")
		fmt.Println("7. Selection Sort")
		fmt.Println("8. Insertion Sort")
		fmt.Println("9. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahTagihan()
		} else if pilih == 2 {
			ubahTagihan()
		} else if pilih == 3 {
			hapusTagihan()
		} else if pilih == 4 {
			tampilkanData()
		} else if pilih == 5 {
		} else if pilih == 6 {
			sequentialSearch()
		} else if pilih == 7 {
		} else if pilih == 8 {
		} else if pilih == 9 {
		} else if pilih == 0 {
			fmt.Println("Program selesai")
			return
		} else {
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func tambahTagihan() {
	fmt.Print("Nama Tagihan : ")
	fmt.Scan(&data[jumlah].Nama)

	fmt.Print("Kategori     : ")
	fmt.Scan(&data[jumlah].Kategori)

	fmt.Print("Nominal      : ")
	fmt.Scan(&data[jumlah].Nominal)

	fmt.Print("Jatuh Tempo (contoh 15): ")
	fmt.Scan(&data[jumlah].JatuhTempo)

	var status int
	fmt.Print("Lunas? (1=Ya, 0=Tidak): ")
	fmt.Scan(&status)

	if status == 1 {
		data[jumlah].Lunas = true
	} else {
		data[jumlah].Lunas = false
	}

	jumlah++

	fmt.Println("Data berhasil ditambahkan")
}

func tampilkanData() {
	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	fmt.Println("\n===== DATA TAGIHAN =====")

	for i := 0; i < jumlah; i++ {
		status := "Belum"

		if data[i].Lunas {
			status = "Lunas"
		}

		fmt.Println("Data ke-", i+1)
		fmt.Println("Nama       :", data[i].Nama)
		fmt.Println("Kategori   :", data[i].Kategori)
		fmt.Println("Nominal    :", data[i].Nominal)
		fmt.Println("JatuhTempo :", data[i].JatuhTempo)
		fmt.Println("Status     :", status)
		fmt.Println()
	}
}

func ubahTagihan() {
	var index int

	tampilkanData()

	fmt.Print("Pilih data yang diubah: ")
	fmt.Scan(&index)

	index--

	if index < 0 || index >= jumlah {
		fmt.Println("Data tidak ditemukan")
		return
	}

	fmt.Print("Nama Baru : ")
	fmt.Scan(&data[index].Nama)

	fmt.Print("Kategori Baru : ")
	fmt.Scan(&data[index].Kategori)

	fmt.Print("Nominal Baru : ")
	fmt.Scan(&data[index].Nominal)

	fmt.Print("Jatuh Tempo Baru : ")
	fmt.Scan(&data[index].JatuhTempo)

	var status int
	fmt.Print("Lunas? (1=Ya,0=Tidak): ")
	fmt.Scan(&status)

	if status == 1 {
		data[index].Lunas = true
	} else {
		data[index].Lunas = false
	}

	fmt.Println("Data berhasil diubah")
}

func hapusTagihan() {
	var index int

	tampilkanData()

	fmt.Print("Pilih data yang dihapus: ")
	fmt.Scan(&index)

	index--

	if index < 0 || index >= jumlah {
		fmt.Println("Data tidak ditemukan")
		return
	}

	for i := index; i < jumlah-1; i++ {
		data[i] = data[i+1]
	}

	jumlah--

	fmt.Println("Data berhasil dihapus")
}

func sequentialSearch() {
	var cari string
	var ketemu bool

	fmt.Print("Masukkan nama/kategori: ")
	fmt.Scan(&cari)

	for i := 0; i < jumlah; i++ {
		if data[i].Nama == cari || data[i].Kategori == cari {
			fmt.Println("Data ditemukan:")
			fmt.Println(data[i])

			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}
