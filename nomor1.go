package main

import (
	"fmt"
	"strings"
	"sync"
)

var tampung = sync.WaitGroup{}

func frekuensi(j string, s int, y bool) {
	N := 97
	if y {
		N = 65
	}
	tes := string(rune(s + N))
	hitung := strings.Count(j, tes)
	if hitung != 0 {
		fmt.Printf("%s : %d \n", tes, hitung)
	}
	tampung.Done()
}

func main() {
	j := "aku lapar sekali sehingga pingsan tadi pagi"

	for i := 0; i < 26; i++ {
		tampung.Add(2)
		go frekuensi(j, i, false)
		go frekuensi(j, i, true)  
	}
	tampung.Wait()
}