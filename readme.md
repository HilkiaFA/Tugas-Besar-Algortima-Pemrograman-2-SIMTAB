# SIMTAB - Sistem Manajemen Tagihan Bulanan

SIMTAB adalah aplikasi sederhana berbasis **Golang** untuk mengelola data tagihan bulanan.  
Program ini dibuat menggunakan konsep dasar algoritma seperti:

- Array
- Struct
- CRUD
- Sequential Search
- Perulangan
- Percabangan

---

# Fitur

Tambah data tagihan  
 Ubah data tagihan  
 Hapus data tagihan  
 Tampilkan seluruh data  
 Sequential Search berdasarkan nama atau kategori  
 Menggunakan array statis

---

# Struktur Data

```go
type Tagihan struct {
    Nama       string
    Kategori   string
    Nominal    int
    JatuhTempo int
    Lunas      bool
}
```

| Atribut    | Tipe Data | Keterangan             |
| ---------- | --------- | ---------------------- |
| Nama       | string    | Nama tagihan           |
| Kategori   | string    | Jenis kategori tagihan |
| Nominal    | int       | Jumlah pembayaran      |
| JatuhTempo | int       | Tanggal jatuh tempo    |
| Lunas      | bool      | Status pembayaran      |

---

# Menu Program

Saat program dijalankan, akan muncul menu berikut:

```text
===== SIMTAB =====
1. Tambah Tagihan
2. Ubah Tagihan
3. Hapus Tagihan
4. Tampilkan Data
5. Sequential Search
6. Binary Search
7. Selection Sort
8. Insertion Sort
9. Statistik
0. Keluar
```

---

# Cara Menjalankan Program

## 1. Install Golang

Cek apakah Golang sudah terinstall:

```bash
go version
```

---

## 2. Simpan File

Simpan source code dengan nama:

```text
main.go
```

---

## 3. Jalankan Program

```bash
go run main.go
```

---

# Cara Menggunakan Program

## ➕ Menambah Tagihan

Pilih menu:

```text
1
```

Contoh input:

```text
Nama Tagihan : Listrik
Kategori     : Rumah
Nominal      : 250000
Jatuh Tempo  : 15
Lunas? (1=Ya, 0=Tidak): 0
```

---

## Menampilkan Data

Pilih menu:

```text
4
```

Program akan menampilkan seluruh data tagihan yang telah disimpan.

---

## Mengubah Data

Pilih menu:

```text
2
```

Kemudian pilih nomor data yang ingin diubah.

---

## Menghapus Data

Pilih menu:

```text
3
```

Kemudian pilih nomor data yang ingin dihapus.

---

## Sequential Search

Pilih menu:

```text
5
```

Masukkan nama atau kategori yang ingin dicari.

Contoh:

```text
Masukkan nama/kategori: Rumah
```

---

# Contoh Output

```text
===== DATA TAGIHAN =====

Data ke- 1
Nama       : Listrik
Kategori   : Rumah
Nominal    : 250000
JatuhTempo : 15
Status     : Belum
```

---

# Fitur yang Belum Diimplementasikan

Berikut beberapa fitur yang masih dapat dikembangkan:

- Binary Search
- Selection Sort
- Insertion Sort
- Statistik Tagihan

---

# Algoritma yang Digunakan

| Algoritma         | Fungsi                       |
| ----------------- | ---------------------------- |
| Sequential Search | Mencari data                 |
| Array             | Menyimpan data               |
| Perulangan        | Menampilkan & memproses data |
| Percabangan       | Menu program                 |
