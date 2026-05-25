package main

import (
	"fmt"
)

// --- STRUKTUR DATA UTAMA ---
type Meja struct {
	Nomor     int
	Kapasitas int
}

type Pelanggan struct {
	ID   int
	Nama string
}

type Reservasi struct {
	ID        int
	IDMeja    int
	IDPel     int
	Tanggal   string // Format YYYY-MM-DD
	Waktu     string // Format HH:MM
}

// --- DATABASE IN-MEMORY ---
var (
	dataMeja      []Meja
	dataPelanggan []Pelanggan
	dataReservasi []Reservasi
	pelangganSeq  = 1
	reservasiSeq  = 1
)

func main() {
	for {
		fmt.Println("\n=== APLIKASI RESERVA RESTO ===")
		fmt.Println("1. Kelola Data Meja")
		fmt.Println("2. Kelola Data Pelanggan")
		fmt.Println("3. Buat Reservasi")
		fmt.Println("4. Cari Data Meja")
		fmt.Println("5. Urutkan Data Meja")
		fmt.Println("6. Statistik Reservasi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			menuMeja()
		case 2:
			menuPelanggan()
		case 3:
			buatReservasi()
		case 4:
			menuCariMeja()
		case 5:
			menuUrutMeja()
		case 6:
			tampilStatistik()
		case 0:
			fmt.Println("Terima kasih telah menggunakan ReservaResto!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// --- A. KELOLA DATA MEJA & PELANGGAN (CRUD) ---

func menuMeja() {
	fmt.Println("\n--- Kelola Data Meja ---")
	fmt.Println("1. Tambah Meja\n2. Tampil Meja\n3. Hapus Meja")
	fmt.Print("Pilih: ")
	var pil int
	fmt.Scan(&pil)

	if pil == 1 {
		var m Meja
		fmt.Print("Nomor Meja: ")
		fmt.Scan(&m.Nomor)
		fmt.Print("Kapasitas Kursi: ")
		fmt.Scan(&m.Kapasitas)
		dataMeja = append(dataMeja, m)
		fmt.Println("Meja berhasil ditambahkan!")
	} else if pil == 2 {
		for _, m := range dataMeja {
			fmt.Printf("Meja %d | Kapasitas: %d kursi\n", m.Nomor, m.Kapasitas)
		}
	} else if pil == 3 {
		fmt.Print("Masukkan Nomor Meja yang akan dihapus: ")
		var no int
		fmt.Scan(&no)
		for i, m := range dataMeja {
			if m.Nomor == no {
				dataMeja = append(dataMeja[:i], dataMeja[i+1:]...)
				fmt.Println("Meja dihapus!")
				return
			}
		}
		fmt.Println("Meja tidak ditemukan.")
	}
}

func menuPelanggan() {
	fmt.Println("\n--- Kelola Data Pelanggan ---")
	fmt.Println("1. Tambah Pelanggan\n2. Tampil Pelanggan")
	fmt.Print("Pilih: ")
	var pil int
	fmt.Scan(&pil)

	if pil == 1 {
		var p Pelanggan
		p.ID = pelangganSeq
		pelangganSeq++
		fmt.Print("Nama Pelanggan: ")
		fmt.Scan(&p.Nama)
		dataPelanggan = append(dataPelanggan, p)
		fmt.Println("Pelanggan berhasil ditambahkan!")
	} else if pil == 2 {
		for _, p := range dataPelanggan {
			fmt.Printf("ID: %d | Nama: %s\n", p.ID, p.Nama)
		}
	}
}

// --- B. RESERVASI & JADWAL PENGGUNAAN ---

func buatReservasi() {
	var r Reservasi
	r.ID = reservasiSeq
	fmt.Print("ID Pelanggan: ")
	fmt.Scan(&r.IDPel)
	fmt.Print("Nomor Meja: ")
	fmt.Scan(&r.IDMeja)
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&r.Tanggal)
	fmt.Print("Waktu (HH:MM): ")
	fmt.Scan(&r.Waktu)

	// Cek ketersediaan meja pada jadwal tersebut
	tersedia := true
	for _, res := range dataReservasi {
		if res.IDMeja == r.IDMeja && res.Tanggal == r.Tanggal && res.Waktu == r.Waktu {
			tersedia = false
			break
		}
	}

	if tersedia {
		dataReservasi = append(dataReservasi, r)
		reservasiSeq++
		fmt.Println("Reservasi berhasil dicatat!")
	} else {
		fmt.Println("Maaf, meja sudah dipesan pada jadwal tersebut!")
	}
}

// --- C. PENCARIAN (SEQUENTIAL & BINARY) ---

func menuCariMeja() {
	fmt.Println("\n1. Sequential Search (Nomor Meja)")
	fmt.Println("2. Binary Search (Kapasitas Kursi - Meja harus diurutkan dulu)")
	fmt.Print("Pilih: ")
	var pil, target int
	fmt.Scan(&pil)

	if pil == 1 {
		fmt.Print("Masukkan Nomor Meja: ")
		fmt.Scan(&target)
		// Sequential Search
		ditemukan := false
		for _, m := range dataMeja {
			if m.Nomor == target {
				fmt.Printf("Ditemukan: Meja %d, Kapasitas %d\n", m.Nomor, m.Kapasitas)
				ditemukan = true
				break
			}
		}
		if !ditemukan {
			fmt.Println("Meja tidak ditemukan.")
		}
	} else if pil == 2 {
		fmt.Print("Masukkan Kapasitas Kursi: ")
		fmt.Scan(&target)
		// Pastikan data diurutkan berdasarkan kapasitas terlebih dahulu untuk Binary Search
		insertionSortMeja() 
		
		// Binary Search
		kiri, kanan := 0, len(dataMeja)-1
		ditemukan := false
		for kiri <= kanan {
			tengah := (kiri + kanan) / 2
			if dataMeja[tengah].Kapasitas == target {
				fmt.Printf("Ditemukan: Meja %d dengan Kapasitas %d\n", dataMeja[tengah].Nomor, dataMeja[tengah].Kapasitas)
				ditemukan = true
				break // Note: Ini hanya mengambil satu hasil
			} else if dataMeja[tengah].Kapasitas < target {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}
		if !ditemukan {
			fmt.Println("Meja dengan kapasitas tersebut tidak ditemukan.")
		}
	}
}

// --- D. PENGURUTAN (SELECTION & INSERTION SORT) ---

func menuUrutMeja() {
	fmt.Println("\n1. Selection Sort (Kecil ke Besar)")
	fmt.Println("2. Insertion Sort (Kecil ke Besar)")
	fmt.Print("Pilih: ")
	var pil int
	fmt.Scan(&pil)

	if pil == 1 {
		selectionSortMeja() //
		fmt.Println("Berhasil diurutkan dengan Selection Sort.")
	} else if pil == 2 {
		insertionSortMeja() //
		fmt.Println("Berhasil diurutkan dengan Insertion Sort.")
	}
}

func selectionSortMeja() {
	n := len(dataMeja)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if dataMeja[j].Kapasitas < dataMeja[minIdx].Kapasitas {
				minIdx = j
			}
		}
		dataMeja[i], dataMeja[minIdx] = dataMeja[minIdx], dataMeja[i]
	}
}

func insertionSortMeja() {
	n := len(dataMeja)
	for i := 1; i < n; i++ {
		key := dataMeja[i]
		j := i - 1
		for j >= 0 && dataMeja[j].Kapasitas > key.Kapasitas {
			dataMeja[j+1] = dataMeja[j]
			j = j - 1
		}
		dataMeja[j+1] = key
	}
}

// --- E. STATISTIK RESERVASI ---

func tampilStatistik() {
	fmt.Println("\n--- Statistik Reservasi ---")
	
	// 1. Jumlah reservasi per hari
	mapHari := make(map[string]int)
	// 2. Meja yang paling sering dipesan
	mapMeja := make(map[int]int)

	for _, r := range dataReservasi {
		mapHari[r.Tanggal]++
		mapMeja[r.IDMeja]++
	}

	fmt.Println("\nJumlah Reservasi Per Hari:")
	for tgl, jumlah := range mapHari {
		fmt.Printf("- %s: %d reservasi\n", tgl, jumlah)
	}

	maxPesan := 0
	mejaTerlaris := -1
	for idMeja, jumlah := range mapMeja {
		if jumlah > maxPesan {
			maxPesan = jumlah
			mejaTerlaris = idMeja
		}
	}

	if mejaTerlaris != -1 {
		fmt.Printf("\nMeja Paling Sering Dipesan: Nomor Meja %d (Dipesan %d kali)\n", mejaTerlaris, maxPesan)
	} else {
		fmt.Println("\nBelum ada data reservasi untuk menghitung meja terlaris.")
	}
}