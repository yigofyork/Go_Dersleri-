package main

import "fmt"

func main() {
	var isim_soyisim string // var = değişken tanımlar, değişkenin yanına ise değişkenin cinsi yazılır(string/int/bool).
	var yas int
	var yetiskinmi bool // bool = true/false

	isim_soyisim = "Yiğit GÜRBULAK"
	yas = 17
	yetiskinmi = false

	fmt.Println("İsim Soyisim: ", isim_soyisim)
	fmt.Println("Yaş: ", yas)
	fmt.Println("Yetişkin mi: ", yetiskinmi)

	var username string = "Yiğit"
	fmt.Println(username)

	const password int = 12345 // const : değiştirilemeyen değişken atamaya yarar
	fmt.Println("password: ", password)

	// Inferrring Type
	minute := 60 // := ifadesi tipini kesin bildiğimiz değişkende kullanırız
	hour := 24
	fmt.Println("Bir gün de", minute*hour, "dakika vardır.")
}
