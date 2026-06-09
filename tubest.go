package main

import "fmt"

const NMAX int = 100

type film struct {
	judul, genre                   string
	durasi, jadwalTayang, penonton int
}

type arrFilm [NMAX]film

func cariFilm(info arrFilm, n int) {
	var i, count int
	var left, right, mid int
	var judul, genre string
	selectionSortUntukBinarySearch(&info, n)

	// input judul dan genre yang ingin disearch
	fmt.Println("Search judul: ")
	fmt.Scan(&judul)

	fmt.Println("Search genre: ")
	fmt.Scan(&genre)

	for judul == "" && genre == "" {
		fmt.Println("Judul dan Genre tidak ada masukan, mohon untuk diketik kembali.")
		fmt.Println("Search judul: ")
		fmt.Scan(&judul)

		fmt.Println("Search genre: ")
		fmt.Scan(&genre)
	}

	// binary search mencari judul
	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2
		if info[mid].judul == judul {
			count++
			fmt.Println(info[mid].judul, info[mid].genre, info[mid].durasi, info[mid].jadwalTayang)
			left = right + 1
		} else if judul < info[mid].judul {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// sequential search mencari genre
	for i = 0; i < n; i++ {
		if info[i].genre == genre {
			count++
			fmt.Println(info[i].judul, info[i].genre, info[i].durasi, info[i].jadwalTayang)
		}
	}

	if count > 0 {
		fmt.Printf("Ditemukan %d film:\n", count)
	} else {
		fmt.Println("Film tidak ditemukan.")
	}
}

func selectionSortUntukBinarySearch(info *arrFilm, n int) {
	var pass, idx, i int
	var temp film
	pass = 1

	// selection sort descending untuk func cariFilm
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if info[idx].judul > info[i].judul {
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

func listFilm(info *arrFilm, n int) {
	var pass, i int
	var temp film
	pass = 1
	fmt.Println("Urutan daftar film berdasarkan jumlah penonton: ")

	// insertion sort descending untuk penonton
	for pass <= n-1 {
		i = pass
		temp = info[pass]
		for i > 0 && (temp.penonton < info[i-1].penonton) {
			info[i] = info[i-1]
			i = i - 1
		}
		info[i] = temp
		pass = pass + 1
	}
	filmPalingLaris(*info)
}

// sudah disort jadi langsung print :)
func filmPalingLaris(info arrFilm) {
	fmt.Printf("Film paling laris dengan %d penonton:\n", info[0].penonton)
	fmt.Println(info[0].judul, info[0].genre, info[0].durasi, info[0].jadwalTayang)
}
