package main

import "fmt"

const NMAX int = 100

type Film struct {
	ID       int
	Judul    string
	Genre    string
	Durasi   int
	Penonton int
}

type Jadwal struct {
	ID            int
	FilmID        int
	WaktuTayang   string
	TotalKursi    int
	KursiTerpesan int
}

type arrFilm [NMAX]Film
type arrJadwal [NMAX]Jadwal

var tabFilm arrFilm
var nFilm int

var tabJadwal arrJadwal
var nJadwal int

func main() {
	dummyData()
	var pilihan int
	var selesai bool = false

	for !selesai {
		fmt.Println("\n🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️🎞️")
		fmt.Println("APLIKASI PEMESANAN TIKET BIOSKOP 🎦")
		fmt.Println("====================================\n")
		fmt.Println("1. Kelola Data Film 📊")
		fmt.Println("2. Kelola Jadwal Tayang 📋")
		fmt.Println("3. Pesan Tiket 🎟️")
		fmt.Println("4. Cari Film (Judul/Genre) 🔍")
		fmt.Println("5. Lihat Daftar Film Terurut 📋")
		fmt.Println("0. Keluar 🏃")
		fmt.Print("\nPilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuKelolaFilm()
		} else if pilihan == 2 {
			menuKelolaJadwal()
		} else if pilihan == 3 {
			pesanTiket()
		} else if pilihan == 4 {
			cariFilm(tabFilm, nFilm)
		} else if pilihan == 5 {
			listFilm(&tabFilm, nFilm)
		} else if pilihan == 0 {
			selesai = true
			fmt.Println("\nTerima Kasih! 🙏")
		} else {
			fmt.Println("\nPilihan tidak valid.")
		}
	}
}

// BAGIAN A & B: KELOLA FILM & JADWAL
func menuKelolaFilm() {
	var pilihan int
	fmt.Println("\n🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️")
	fmt.Println("KELOLA DATA FILM 🎦")
	fmt.Println("====================================\n")
	fmt.Println("1. Tambah Film ✏️")
	fmt.Println("2. Lihat Semua Film 🔍")
	fmt.Println("3. Hapus Film 🧼")
	fmt.Print("\nPilih aksi: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		if nFilm < NMAX {
			var f Film
			f.ID = nFilm + 1
			if nFilm-1 == tabFilm[nFilm-1].ID {
				f.ID = nFilm + 2
			}
			fmt.Print("Judul (tanpa spasi): ")
			fmt.Scan(&f.Judul)
			fmt.Print("Genre: ")
			fmt.Scan(&f.Genre)
			fmt.Print("Durasi (menit): ")
			fmt.Scan(&f.Durasi)
			f.Penonton = 0

			tabFilm[nFilm] = f
			nFilm++
			fmt.Println("Film berhasil ditambahkan! ✅")
		} else {
			fmt.Println("Kapasitas film penuh. 😢")
		}
	} else if pilihan == 2 {
		lihatSemuaFilm()
	} else if pilihan == 3 {
		var id int
		fmt.Print("Masukkan ID Film yang mau dihapus: ")
		fmt.Scan(&id)
		hapusFilm(id)
	}
}

func hapusFilm(id int) {
	var denganJadwal bool = false
	var idx int = sequentialSearchDenganID(id, denganJadwal)
	var i int
	if idx != -1 {
		for i = idx; i < nFilm-1; i++ {
			tabFilm[i] = tabFilm[i+1]
			tabJadwal[i] = tabJadwal[i+1]
		}
		nFilm--
		nJadwal--
		fmt.Println("Film berhasil dihapus. 🫧")
	} else {
		fmt.Println("Film tidak ditemukan. 😢")
	}
}

func lihatSemuaFilm() {
	var i int
	if nFilm == 0 {
		fmt.Println("\nBelum ada data film. 😢")
	} else {
		fmt.Println("\n🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️")
		fmt.Println("Daftar Film 📋:")
		fmt.Println("====================================\n")
		for i = 0; i < nFilm; i++ {
			fmt.Printf("ID: %d | %s | %s | %d menit | Penonton: %d\n",
				tabFilm[i].ID, tabFilm[i].Judul, tabFilm[i].Genre, tabFilm[i].Durasi, tabFilm[i].Penonton)
		}
	}
}

func menuKelolaJadwal() {
	var pilihan int
	var denganJadwal bool = false

	fmt.Println("\n🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️🎟️")
	fmt.Println("KELOLA JADWAL TAYANG 🎦")
	fmt.Println("====================================\n")
	fmt.Println("1. Tambah Jadwal ✏️")
	fmt.Println("2. Lihat Semua Jadwal 🔍")
	fmt.Print("\nPilih aksi: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		if nJadwal < NMAX {
			lihatSemuaFilm()
			var jdl Jadwal
			jdl.ID = nJadwal + 1

			fmt.Print("\nMasukkan ID Film: ")
			fmt.Scan(&jdl.FilmID)

			if sequentialSearchDenganID(jdl.FilmID, denganJadwal) != -1 {
				fmt.Print("Waktu Tayang (00:00 - 23:59): ")
				fmt.Scan(&jdl.WaktuTayang)
				fmt.Print("Kapasitas Kursi: ")
				fmt.Scan(&jdl.TotalKursi)
				jdl.KursiTerpesan = 0

				tabJadwal[nJadwal] = jdl
				nJadwal++
				fmt.Println("Jadwal berhasil ditambahkan.")
			} else {
				fmt.Println("ID Film tidak valid. 😢")
			}
		} else {
			fmt.Println("Kapasitas jadwal penuh. 😢")
		}
	} else if pilihan == 2 {
		lihatSemuaJadwal()
	}
}

func lihatSemuaJadwal() {
	var i, idxFilm, sisa int
	var judul string
	var denganJadwal bool = false

	if nJadwal == 0 {
		fmt.Println("Belum ada jadwal tayang. 😢")
	} else {
		fmt.Println("\nDaftar Jadwal 📋:")
		fmt.Println("====================================\n")
		for i = 0; i < nJadwal; i++ {
			idxFilm = sequentialSearchDenganID(tabJadwal[i].FilmID, denganJadwal)
			judul = "Unknown"
			if idxFilm != -1 {
				judul = tabFilm[idxFilm].Judul
			}
			sisa = tabJadwal[i].TotalKursi - tabJadwal[i].KursiTerpesan
			fmt.Printf("ID: %d | %s | Jam: %s | Sisa Kursi: %d\n",
				tabJadwal[i].ID, judul, tabJadwal[i].WaktuTayang, sisa)
		}
	}
}

// BAGIAN C: PESAN TIKET
func pesanTiket() {
	var idxJadwal, sisaKursi, idxFilm int
	var denganJadwal bool = false

	lihatSemuaJadwal()
	if nJadwal == 0 {
		return
	}

	var jadwalID, jumlah int
	fmt.Print("\nMasukkan ID Jadwal yang ingin dipesan: ")
	fmt.Scan(&jadwalID)

	denganJadwal = true
	idxJadwal = sequentialSearchDenganID(jadwalID, denganJadwal)

	if idxJadwal != -1 {
		sisaKursi = tabJadwal[idxJadwal].TotalKursi - tabJadwal[idxJadwal].KursiTerpesan
		fmt.Printf("Sisa kursi: %d\n", sisaKursi)
		fmt.Print("Jumlah tiket: ")
		fmt.Scan(&jumlah)

		if jumlah > 0 && jumlah <= sisaKursi {
			tabJadwal[idxJadwal].KursiTerpesan += jumlah

			denganJadwal = true
			idxFilm = sequentialSearchDenganID(tabJadwal[idxJadwal].FilmID, denganJadwal)

			if idxFilm != -1 {
				tabFilm[idxFilm].Penonton += jumlah
			}

			fmt.Println("Tiket berhasil dipesan. ✅")
		} else {
			fmt.Println("Gagal: kursi tidak cukup atau jumlah salah. 😢")
		}
	} else {
		fmt.Println("Jadwal tidak ditemukan. 😢")
	}
}

// BAGIAN D: CARI FILM
func cariFilm(info arrFilm, n int) {
	if n == 0 {
		fmt.Println("Data film masih kosong.")
		return
	}

	var i, count int
	var judul, genre string

	selectionSort(&info, n)

	fmt.Println("\nPENCARIAN FILM 🎦")
	fmt.Println("====================================\n")
	fmt.Print("Search judul (ketik - jika kosong): ")
	fmt.Scan(&judul)

	fmt.Print("Search genre (ketik - jika kosong): ")
	fmt.Scan(&genre)

	for judul == "-" && genre == "-" {
		fmt.Println("Judul dan Genre tidak ada masukan, mohon untuk diketik kembali.")
		fmt.Print("Search judul (ketik - jika kosong): ")
		fmt.Scan(&judul)

		fmt.Print("Search genre (ketik - jika kosong): ")
		fmt.Scan(&genre)
	}

	fmt.Println("\nHasil Pencarian:")
	fmt.Println("====================================\n")

	count = binarySearch(info, n, judul)

	if genre != "-" {
		for i = 0; i < n; i++ {
			if info[i].Genre == genre && info[i].Judul != judul {
				count++
				fmt.Printf("ID: %d | %s | %s | %d menit\n", info[i].ID, info[i].Judul, info[i].Genre, info[i].Durasi)
			}
		}
	}

	if count > 0 {
		fmt.Printf("✅ Ditemukan %d film.\n", count)
	} else {
		fmt.Println("Film tidak ditemukan. 😢")
	}
}

// BAGIAN E & F: LIST & TERLARIS
func listFilm(info *arrFilm, n int) {
	var j int
	if n == 0 {
		fmt.Println("Data film masih kosong. 😢")
		return
	}
	fmt.Println("\nUrutan daftar film berdasarkan jumlah penonton 📋:")
	fmt.Println("====================================\n")
	insertionSort(info, n)

	for j = 0; j < n; j++ {
		fmt.Printf("%d. %s : %d Penonton\n", j+1, info[j].Judul, info[j].Penonton)
	}

	fmt.Println()
	filmPalingLaris(*info)
}

func filmPalingLaris(info arrFilm) {
	fmt.Printf("Film paling laris dengan %d penonton! ✅:\n", info[0].Penonton)
	fmt.Printf("%s | %s | %d menit\n", info[0].Judul, info[0].Genre, info[0].Durasi)
}

// sequential & binary search ############################################################
// sequential search untuk berbagai hal ##################################################

func sequentialSearchDenganID(id int, denganJadwal bool) int {
	var foundIdx int = -1
	var i int = 0
	if denganJadwal == true {
		for i < nJadwal && foundIdx == -1 {
			if tabJadwal[i].ID == id {
				foundIdx = i
			}
			i++
		}
	} else {
		for i < nFilm && foundIdx == -1 {
			if tabFilm[i].ID == id {
				foundIdx = i
			}
			i++
		}
	}
	return foundIdx
}

// binary search untuk mencari film sesuai judul #########################################
// sudah diurutkan menggunakan binary search #############################################

func binarySearch(info arrFilm, n int, judul string) int {
	var left, right, mid, count int
	if judul != "-" {
		left = 0
		right = n - 1

		for left <= right {
			mid = (left + right) / 2
			if info[mid].Judul == judul {
				count++
				fmt.Printf("%d | %s | %s | %d menit\n", info[mid].ID, info[mid].Judul, info[mid].Genre, info[mid].Durasi)
				left = right + 1
			} else if judul < info[mid].Judul {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return count
}

// selection & insertion sort ############################################################
// selection sort untuk binary search ####################################################

func selectionSort(info *arrFilm, n int) {
	var pass, idx, i int
	var temp Film
	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if info[i].Judul < info[idx].Judul {
				idx = i
			}
			i = i + 1
		}
		temp = info[pass-1]
		info[pass-1] = info[idx]
		info[idx] = temp
		pass++
	}
}

// insertion sort untuk urut film berdasarkan jumlah penonton ############################

func insertionSort(info *arrFilm, n int) {
	var pass, i int
	var temp Film
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = info[pass]
		for i > 0 && (temp.Penonton > info[i-1].Penonton) {
			info[i] = info[i-1]
			i = i - 1
		}
		info[i] = temp
		pass = pass + 1
	}
}

// dummy data isi 4 film #################################################################

func dummyData() {
	tabFilm[0] = Film{
		ID:       1,
		Judul:    "Avatar",
		Genre:    "Action",
		Durasi:   162,
		Penonton: 90,
	}
	nFilm++

	tabFilm[1] = Film{
		ID:       2,
		Judul:    "Titanic",
		Genre:    "Drama",
		Durasi:   195,
		Penonton: 70,
	}
	nFilm++

	tabFilm[2] = Film{
		ID:       3,
		Judul:    "Zootopia",
		Genre:    "Animation",
		Durasi:   108,
		Penonton: 100,
	}
	nFilm++

	tabFilm[3] = Film{
		ID:       4,
		Judul:    "Star_Wars:_The_Force_Awakens",
		Genre:    "Sci-fi",
		Durasi:   138,
		Penonton: 80,
	}
	nFilm++

	// jadwal
	tabJadwal[0] = Jadwal{
		ID:            1,
		FilmID:        1,
		WaktuTayang:   "10:00",
		TotalKursi:    150,
		KursiTerpesan: 90,
	}
	nJadwal++

	tabJadwal[1] = Jadwal{
		ID:            2,
		FilmID:        2,
		WaktuTayang:   "13:00",
		TotalKursi:    150,
		KursiTerpesan: 70,
	}
	nJadwal++

	tabJadwal[2] = Jadwal{
		ID:            3,
		FilmID:        3,
		WaktuTayang:   "12:00",
		TotalKursi:    150,
		KursiTerpesan: 100,
	}
	nJadwal++

	tabJadwal[3] = Jadwal{
		ID:            4,
		FilmID:        4,
		WaktuTayang:   "14:20",
		TotalKursi:    150,
		KursiTerpesan: 80,
	}
	nJadwal++
}
