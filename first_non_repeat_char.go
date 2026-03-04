package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func firstNonRepeatingChar(s string) rune {
	// Map untuk menyimpan frekuensi setiap karakter (rune)
	freq := make(map[rune]int)

	// Iterasi pertama: hitung frekuensi kemunculan tiap karakter
	for _, ch := range s {
		freq[ch]++
	}

	// Iterasi kedua: cari karakter pertama dengan frekuensi 1
	for _, ch := range s {
		if freq[ch] == 1 {
			return ch
		}
	}

	// default tidak ada karakter unik
	return 0

}

func main() {
	// Membaca input dari pengguna
	fmt.Print("Masukkan string: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Gagal membaca input:", err)
		return
	}

	// Menghapus karakter newline (\n) dari input
	input = strings.TrimSpace(input)

	// Panggil fungsi
	result := firstNonRepeatingChar(input)

	// Tampilkan hasil
	if result == 0 {
		fmt.Println("None")
	} else {
		// fmt.Println("Karakter pertama yang tidak berulang: ", string(result))
		fmt.Println(string(result))
	}

}
