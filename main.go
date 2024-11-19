package main

import (
        "fmt"
        "strings"
)

var nama string = "Dwi Prasetia"

func writeDwi() {
        fmt.Println(nama)
}

func reverse(s string) string {
        rns := []rune(s)
        for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
                rns[i], rns[j] = rns[j], rns[i]
        }
        return string(rns)
}

func addChar(str string) string {
        var sb strings.Builder
        for _, char := range str {
                sb.WriteString(string(char))
                sb.WriteString("_")
        }
        return sb.String()[:len(sb.String())-1] 
}

func main() {
        writeDwi()
        balik := reverse(nama)
        addedChar := addChar(balik)
        fmt.Println(balik)
        fmt.Println(addedChar)
        fmt.Println("halo coba")
}