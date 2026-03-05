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

	// Membaca input dari user
	fmt.Print("Masukkan string: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Gagal membaca input:", err)
		return
	}

	input = strings.TrimSpace(input)

	result := firstNonRepeatingChar(input)

	if result == 0 {
		fmt.Println("None")
	} else {
		fmt.Println(string(result))
	}

}
