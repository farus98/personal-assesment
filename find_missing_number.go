package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findMissingNummber(input string) int {
	arr := strings.Split(input, " ")
	sort.Strings(arr)

	first, _ := strconv.Atoi(arr[0])
	num := first

	for i, p := range arr {
		if i == 0 {
			continue
		}
		cur, _ := strconv.Atoi(p)
		num++
		if cur != num {
			return num
		} else {
			num = cur
		}
	}

	return num
}

func main() {
	fmt.Print("Masukkan string: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Gagal membaca input:", err)
		return
	}

	// Menghapus karakter newline (\n) dari input
	input = strings.TrimSpace(input)

	fmt.Println(findMissingNummber(input))
}
