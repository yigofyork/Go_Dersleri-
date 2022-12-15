/*

  Data Types (Veri Tipleri)

  Temel Veri Tipleri: Numbers. String, Boolean

  Küme Veri Tipleri: Array ve Structure

  Referans Veri Tipleri: Pointers, slices, maps, functions, channels

  *************************************************
  Numbers: int, float, complex

  int8 = -128   -    127
  int16 = -32768  -    32767
  int32
  int63
  uint8 = 0   -    255
  uint16 = 0    -   65535

*/

package main

import "fmt"

func main() {
	var x int = 12700000000
	fmt.Println(x)

	var y float32 = 32.45
	fmt.Println(y)

	var z complex128 = complex(6, 8)
	fmt.Println(z)
}
