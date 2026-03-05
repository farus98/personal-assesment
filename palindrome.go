package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func canFormPalindrome(s string) bool {
	// Map untuk menyimpan frekuensi setiap karakter (rune)
	freq := make(map[rune]int)

	// Hitung frekuensi tiap huruf
	for _, ch := range s {
		freq[ch]++
	}

	oddCount := 0
	for _, freq := range freq {
		if freq%2 != 0 {
			oddCount++
		}
		if oddCount > 1 {
			return false
		}
	}
	return true
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

	fmt.Printf("%s -> %v\n", input, canFormPalindrome(input))
}
