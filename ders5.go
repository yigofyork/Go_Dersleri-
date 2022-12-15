package main

import "fmt"

func main() {
	var takim1, takim2 string = "A takımı", "B takımı"
	fmt.Println(takim1, "ile", takim2, "bu aksam saat 19 da karşılaşacaklar.")

	fmt.Printf("%s ile %s bu akşam saat 19 da karşılaşacaklar\n", takim1, takim2) // \n : çıktının alt satıra geçmesine yarar, printf çıktısı kendisi al satıra atmaz yana yazar
	var team1Goal, team2Goal int = 3, 2
	fmt.Printf("Maç Sonu Skoru %d-%d\n", team1Goal, team2Goal)

	var siddet float32 = 3.4
	fmt.Printf("Ege Açıklarında %g şiddetinde sarsıntı oldu", siddet)
}

// string ifadeler için %s , integer ifadeler için ise %d,  floatta ise %g
