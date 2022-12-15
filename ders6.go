package main

import "fmt"

func main() {

	/*toplam := 0
	for i:=0;i<10;i++ {
		toplam += i
		fmt.Println("Toplam: ", toplam)
	}*/

	toplam := 1

	for toplam < 1000 {
		toplam += toplam
		fmt.Println("Toplam:", toplam)
	}
}
