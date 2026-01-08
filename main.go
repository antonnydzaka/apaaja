package main

import "fmt"

func main() {
	// Batas bawah dan batas atas
	low := 0
	high := 100

	fmt.Println("=================================================")
	fmt.Println("   Pikirkan sebuah angka antara 0 sampai 100")
	fmt.Println("   Jawab dengan mengetik: true atau false")
	fmt.Println("=================================================")
	fmt.Println("Tekan Enter jika sudah memikirkan angkanya...")
	var dummy string
	fmt.Scanln(&dummy)

	// Loop akan berjalan sampai batas bawah sama dengan batas atas
	for low < high {
		// Komputer mengambil nilai tengah
		mid := (low + high) / 2

		fmt.Printf("Apakah angkanya lebih besar dari %d? (true/false): ", mid)

		var lebihBesar bool
		// Membaca input boolean dari user
		_, err := fmt.Scan(&lebihBesar)

		/* Validasi input sederhana (jika user mengetik selain true/false)
		if err != nil {
			fmt.Println("Error: Harap ketik 'true' atau 'false' saja.")
			// Bersihkan buffer input agar tidak loop infinite (hack sederhana tanpa bufio)
			var buang string
			fmt.Scan(&buang)
			continue
		}*/

		if lebihBesar {
			// Jika ya, berarti angka ada di rentang atas (mid+1 sampai high)
			low = mid + 1
		} else {
			// Jika tidak, berarti angka ada di rentang bawah atau sama dengan mid
			high = mid
		}
	}

	// Ketika loop selesai, low dan high akan bernilai sama
	fmt.Printf("\n------------------------------------\n")
	fmt.Printf("Angka yang Anda pikirkan adalah: %d\n", low)
	fmt.Printf("------------------------------------\n")
}