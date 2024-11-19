package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Angka struct {
	Num1 int
	Num2 int
	Num3 int
}

func (a *Angka) kali() {
	hasilkali := a.Num1 * a.Num2 * a.Num3
	fmt.Printf("Hasil dari %d x %d x %d adalah= %d \n", a.Num1, a.Num2, a.Num3, hasilkali)
}

func getInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error membaca input. Silakan coba lagi!")
			continue
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Masukkan angka saja!")
			continue
		}
		return num
	}
}

func main() {
	num1 := getInput("Pilih angka ke-1 = ")
	num2 := getInput("Pilih angka ke-2 = ")
	num3 := getInput("Pilih angka ke-3 = ")

	BarisAngka := &Angka{
		Num1: num1,
		Num2: num2,
		Num3: num3,
	}

	BarisAngka.kali()
	return 
}
