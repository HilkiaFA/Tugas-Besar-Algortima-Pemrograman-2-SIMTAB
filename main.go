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
			sequentialSearch()
		} else if pilih == 6 {
			binarySearch()
		} else if pilih == 7 {
			selectionSort()
		} else if pilih == 8 {
			insertionSort()
		} else if pilih == 9 {
			statistik()
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

	data[jumlah].Lunas = status == 1

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

	data[index].Lunas = status == 1

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
	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	var cari string
	var ketemu bool

	fmt.Print("Masukkan nama/kategori: ")
	fmt.Scan(&cari)

	for i := 0; i < jumlah; i++ {
		if data[i].Nama == cari || data[i].Kategori == cari {
			status := "Belum"
			if data[i].Lunas {
				status = "Lunas"
			}

			fmt.Println("\nData ditemukan:")
			fmt.Println("Nama       :", data[i].Nama)
			fmt.Println("Kategori   :", data[i].Kategori)
			fmt.Println("Nominal    :", data[i].Nominal)
			fmt.Println("JatuhTempo :", data[i].JatuhTempo)
			fmt.Println("Status     :", status)
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func binarySearch() {
	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	selectionSort()

	fmt.Println("(Data diurutkan berdasarkan Nama untuk Binary Search)")

	var cari string
	fmt.Print("Masukkan nama tagihan: ")
	fmt.Scan(&cari)

	kr := 0
	kn := jumlah - 1
	ketemu := false

	for kr <= kn && !ketemu {
		med := (kr + kn) / 2

		if data[med].Nama < cari {
			kr = med + 1
		} else if data[med].Nama > cari {
			kn = med - 1
		} else {
			status := "Belum"
			if data[med].Lunas {
				status = "Lunas"
			}

			fmt.Println("\nData ditemukan:")
			fmt.Println("Nama       :", data[med].Nama)
			fmt.Println("Kategori   :", data[med].Kategori)
			fmt.Println("Nominal    :", data[med].Nominal)
			fmt.Println("JatuhTempo :", data[med].JatuhTempo)
			fmt.Println("Status     :", status)

			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func selectionSort() {
	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	var pilih, arahurutan int

	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Kategori")
	fmt.Println("3. Nominal")
	fmt.Println("4. Jatuh Tempo")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	fmt.Println("Urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilih: ")
	fmt.Scan(&arahurutan)

	for i := 1; i <= jumlah-1; i++ {
		idxPilih := i - 1

		for j := i; j < jumlah; j++ {
			lebih := false

			if pilih == 1 {
				if arahurutan == 1 {
					lebih = data[idxPilih].Nama > data[j].Nama
				} else {
					lebih = data[idxPilih].Nama < data[j].Nama
				}
			} else if pilih == 2 {
				if arahurutan == 1 {
					lebih = data[idxPilih].Kategori > data[j].Kategori
				} else {
					lebih = data[idxPilih].Kategori < data[j].Kategori
				}
			} else if pilih == 3 {
				if arahurutan == 1 {
					lebih = data[idxPilih].Nominal > data[j].Nominal
				} else {
					lebih = data[idxPilih].Nominal < data[j].Nominal
				}
			} else if pilih == 4 {
				if arahurutan == 1 {
					lebih = data[idxPilih].JatuhTempo > data[j].JatuhTempo
				} else {
					lebih = data[idxPilih].JatuhTempo < data[j].JatuhTempo
				}
			}

			if lebih {
				idxPilih = j
			}
		}

		t := data[idxPilih]
		data[idxPilih] = data[i-1]
		data[i-1] = t
	}

	fmt.Println("Data sudah terurut")
	tampilkanData()
}

func insertionSort() {
	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	var dipilih, arah int

	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Kategori")
	fmt.Println("3. Nominal")
	fmt.Println("4. Jatuh Tempo")
	fmt.Print("Pilih: ")
	fmt.Scan(&dipilih)

	fmt.Println("Urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilih: ")
	fmt.Scan(&arah)

	for i := 1; i < jumlah; i++ {
		t := data[i]
		j := i - 1

		for j >= 0 {
			lebih := false

			if dipilih == 1 {
				if arah == 1 {
					lebih = data[j].Nama > t.Nama
				} else {
					lebih = data[j].Nama < t.Nama
				}
			} else if dipilih == 2 {
				if arah == 1 {
					lebih = data[j].Kategori > t.Kategori
				} else {
					lebih = data[j].Kategori < t.Kategori
				}
			} else if dipilih == 3 {
				if arah == 1 {
					lebih = data[j].Nominal > t.Nominal
				} else {
					lebih = data[j].Nominal < t.Nominal
				}
			} else if dipilih == 4 {
				if arah == 1 {
					lebih = data[j].JatuhTempo > t.JatuhTempo
				} else {
					lebih = data[j].JatuhTempo < t.JatuhTempo
				}
			}

			if lebih {
				data[j+1] = data[j]
				j--
			} else {
				return
			}
		}

		data[j+1] = t
	}

	fmt.Println("Data sudah terurut")
	tampilkanData()
}

func statistik() {
	var total int
	var totalLunas, totalBelumLunas int
	var jumlahLunas, jumlahBelumLunas int

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < jumlah; i++ {
		total += data[i].Nominal
	}

	fmt.Println("\n===== STATISTIK TAGIHAN =====")
	fmt.Printf("Total seluruh tagihan : Rp.%d\n\n", total)

	fmt.Println("=== Tagihan Sudah Lunas ===")
	for i := 0; i < jumlah; i++ {
		if data[i].Lunas {
			fmt.Printf("Nama      : %s\n", data[i].Nama)
			fmt.Printf("Kategori  : %s\n", data[i].Kategori)
			fmt.Printf("Nominal   : Rp.%d\n\n", data[i].Nominal)

			totalLunas += data[i].Nominal
			jumlahLunas++
		}
	}
	fmt.Printf("Total tagihan lunas : Rp.%d\n\n", totalLunas)

	fmt.Println("=== Tagihan Belum Lunas ===")
	for i := 0; i < jumlah; i++ {
		if !data[i].Lunas {
			fmt.Printf("Nama      : %s\n", data[i].Nama)
			fmt.Printf("Kategori  : %s\n", data[i].Kategori)
			fmt.Printf("Nominal   : Rp.%d\n\n", data[i].Nominal)

			totalBelumLunas += data[i].Nominal
			jumlahBelumLunas++
		}
	}
	fmt.Printf("Total tagihan belum lunas : Rp.%d\n\n", totalBelumLunas)

	fmt.Printf("Persentase tagihan yang sudah lunas : %.2f%%\n",
		float64(jumlahLunas)/float64(jumlah)*100)

	fmt.Printf("Persentase tagihan yang belum lunas : %.2f%%\n",
		float64(jumlahBelumLunas)/float64(jumlah)*100)
}
