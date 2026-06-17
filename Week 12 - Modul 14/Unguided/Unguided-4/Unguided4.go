package main

import "fmt"

const nMax = 7919

type Buku struct {
    id        string
    judul     string
    penulis   string
    penerbit  string
    eksemplar int
    tahun     int
    rating    int
}

type DaftarBuku []Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
    *pustaka = make(DaftarBuku, n)
    for i := 0; i < n; i++ {
        fmt.Printf("Masukkan data buku ke-%d\n", i+1)
        fmt.Print("ID: ")
        fmt.Scanln(&(*pustaka)[i].id)
        fmt.Print("Judul: ")
        fmt.Scanln(&(*pustaka)[i].judul)
        fmt.Print("Penulis: ")
        fmt.Scanln(&(*pustaka)[i].penulis)
        fmt.Print("Penerbit: ")
        fmt.Scanln(&(*pustaka)[i].penerbit)
        fmt.Print("Tahun: ")
        fmt.Scanln(&(*pustaka)[i].tahun)
        fmt.Print("Eksemplar: ")
        fmt.Scanln(&(*pustaka)[i].eksemplar)
        fmt.Print("Rating: ")
        fmt.Scanln(&(*pustaka)[i].rating)
        fmt.Println()
    }
}

func CetakTerfavorit(pustaka DaftarBuku) {
    if len(pustaka) == 0 {
        fmt.Println("Tidak ada buku dalam pustaka.")
        return
    }
    maxRating := pustaka[0].rating
    idx := 0
    for i := 1; i < len(pustaka); i++ {
        if pustaka[i].rating > maxRating {
            maxRating = pustaka[i].rating
            idx = i
        }
    }
    buku := pustaka[idx]
    fmt.Println("Buku terfavorit:")
    fmt.Printf("Judul: %s\n", buku.judul)
    fmt.Printf("Penulis: %s\n", buku.penulis)
    fmt.Printf("Penerbit: %s\n", buku.penerbit)
    fmt.Printf("Tahun: %d\n", buku.tahun)
}

func UrutBuku(pustaka DaftarBuku) {
    for i := 1; i < len(pustaka); i++ {
        key := pustaka[i]
        j := i - 1
        for j >= 0 && pustaka[j].rating < key.rating {
            pustaka[j+1] = pustaka[j]
            j--
        }
        pustaka[j+1] = key
    }
}

func Cetak5Terbaru(pustaka DaftarBuku) {
    jumlah := 5
    if len(pustaka) < 5 {
        jumlah = len(pustaka)
    }
    fmt.Printf("%d buku dengan rating tertinggi:\n", jumlah)
    for i := 0; i < jumlah; i++ {
        fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
    }
}

func CariBuku(pustaka DaftarBuku, r int) {
    low := 0
    high := len(pustaka) - 1
    for low <= high {
        mid := (low + high) / 2
        if pustaka[mid].rating == r {
            buku := pustaka[mid]
            fmt.Println("Buku ditemukan:")
            fmt.Printf("Judul: %s\n", buku.judul)
            fmt.Printf("Penulis: %s\n", buku.penulis)
            fmt.Printf("Penerbit: %s\n", buku.penerbit)
            fmt.Printf("Tahun: %d\n", buku.tahun)
            fmt.Printf("Eksemplar: %d\n", buku.eksemplar)
            fmt.Printf("Rating: %d\n", buku.rating)
            return
        }
        if pustaka[mid].rating > r {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
    var n int
    fmt.Print("Masukkan jumlah buku: ")
    fmt.Scanln(&n)
    if n <= 0 || n > nMax {
        fmt.Println("Jumlah buku tidak valid.")
        return
    }
    var pustaka DaftarBuku
    DaftarkanBuku(&pustaka, n)
    CetakTerfavorit(pustaka)
    UrutBuku(pustaka)
    Cetak5Terbaru(pustaka)
    var ratingCari int
    fmt.Print("Masukkan rating buku yang dicari: ")
    fmt.Scanln(&ratingCari)
    CariBuku(pustaka, ratingCari)
}
