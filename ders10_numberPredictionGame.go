package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	sayi := rand.Intn(max - min)
	hak := 10

	for hak > 0 {
		fmt.Println("1 ile 100 arasında bir sayı tuttum. Tahmin et? ")
		var tahmin int
		fmt.Scanln(&tahmin)

		if tahmin < sayi {
			fmt.Println("Yukarı çık")
			hak--
			if hak == 0 {
				fmt.Println("Haklarınız bitti :(")
			}
		} else if tahmin > sayi {
			fmt.Println("Aşağı in")
			hak--
			if hak == 0 {
				fmt.Println("Haklarınız bitti :(")
			}
		} else {
			fmt.Println("Bildiniz!")
			break // programı sonlandırır
		}
	}

}
