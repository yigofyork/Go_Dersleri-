package main

import "fmt"

func main() {
	fmt.Println("Hangi hesaba bağlanacaksınız? ")
	var hesap int
	fmt.Scanln(&hesap)

	switch hesap {
	case 1, 5:
		fmt.Println("Hesap 1'e bağlanılıyor ...")
	case 2, 6:
		fmt.Println("Hesap 2'ye bağlanılıyor ...")
	case 3:
		fmt.Println("Hesap 3'e bağlanılıyor ...")
	case 4:
		fmt.Println("Hesap 4'e bağlanılıyor ...")
	default:
		fmt.Println("Geçersiz")
	}
}
