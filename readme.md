# SIMTAB - Sistem Manajemen Tagihan Bulanan

SIMTAB adalah aplikasi sederhana berbasis **Golang** untuk mengelola data tagihan bulanan menggunakan konsep algoritma dan struktur data dasar.

Program ini dibuat untuk melatih pemahaman tentang:

- Array statis
- Struct
- CRUD (Create, Read, Update, Delete)
- Searching (Sequential & Binary Search)
- Sorting (Selection & Insertion Sort)
- Statistik data
- Perulangan & Percabangan

---

# 🚀 Fitur Program

- Tambah data tagihan
- Ubah data tagihan
- Hapus data tagihan
- Menampilkan seluruh data
- Sequential Search (nama / kategori)
- Binary Search (berdasarkan nama)
- Selection Sort
- Insertion Sort
- Statistik tagihan (total, lunas, belum lunas, persentase)

---

# 🧱 Struktur Data

```go
type Tagihan struct {
    Nama       string
    Kategori   string
    Nominal    int
    JatuhTempo int
    Lunas      bool
}

| Atribut    | Tipe Data | Keterangan          |
| ---------- | --------- | ------------------- |
| Nama       | string    | Nama tagihan        |
| Kategori   | string    | Jenis tagihan       |
| Nominal    | int       | Jumlah pembayaran   |
| JatuhTempo | int       | Tanggal jatuh tempo |
| Lunas      | bool      | Status pembayaran   |
```

# 📋 Menu Program

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
10. Keluar

---

# ⚙️ Cara Menjalankan Program

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

# 📌 Cara Penggunaan Program

## ➕ Tambah Tagihan

| Input       | Contoh  |
| ----------- | ------- |
| Nama        | Listrik |
| Kategori    | Rumah   |
| Nominal     | 250000  |
| Jatuh Tempo | 15      |
| Lunas       | 0 / 1   |

---

## 📄 Tampilkan Data

| Menu | Fungsi                 |
| ---- | ---------------------- |
| 4    | Menampilkan semua data |

---

## ✏️ Ubah Data

| Langkah | Keterangan       |
| ------- | ---------------- |
| 1       | Pilih menu 2     |
| 2       | Pilih index data |
| 3       | Input data baru  |

---

## ❌ Hapus Data

| Langkah | Keterangan              |
| ------- | ----------------------- |
| 1       | Pilih menu 3            |
| 2       | Pilih data yang dihapus |

---

## 🔎 Sequential Search

| Fitur     | Keterangan         |
| --------- | ------------------ |
| Menu      | 5                  |
| Pencarian | Nama atau Kategori |
| Metode    | Linear Search      |

---

## ⚡ Binary Search

| Fitur  | Keterangan                     |
| ------ | ------------------------------ |
| Menu   | 6                              |
| Data   | Harus terurut berdasarkan Nama |
| Metode | Binary Search                  |

---

## 📊 Sorting Data

| Fitur  | Keterangan                |
| ------ | ------------------------- |
| Menu   | 7                         |
| Field  | Nama / Kategori / Nominal |
| Urutan | Ascending / Descending    |
| Metode | Selection Sort            |

---

## Insertion Sort

| Fitur  | Keterangan                |
| ------ | ------------------------- |
| Menu   | 8                         |
| Field  | Nama / Kategori / Nominal |
| Urutan | Ascending / Descending    |
| Metode | Insertion Sort            |

---

## 📈 Statistik Tagihan

| Statistik              | Keterangan                       |
| ---------------------- | -------------------------------- |
| Total tagihan          | Jumlah seluruh nominal           |
| Tagihan lunas          | Data & total tagihan lunas       |
| Tagihan belum lunas    | Data & total tagihan belum lunas |
| Persentase lunas       | (lunas / total) × 100            |
| Persentase belum lunas | (belum / total) × 100            |

---

# 🧠 Algoritma yang Digunakan

| Algoritma         | Fungsi                           |
| ----------------- | -------------------------------- |
| Sequential Search | Pencarian data secara linear     |
| Binary Search     | Pencarian data pada data terurut |
| Selection Sort    | Mengurutkan data (selection)     |
| Insertion Sort    | Mengurutkan data (insertion)     |
| Array             | Menyimpan data                   |
| Struct            | Struktur data tagihan            |
| Perulangan        | Menampilkan & memproses data     |
| Percabangan       | Logika menu & kondisi program    |

---

# 📦 Contoh Output

===== DATA TAGIHAN =====

Data ke- 1
Nama : Listrik
Kategori : Rumah
Nominal : 250000
JatuhTempo : 15
Status : Belum
