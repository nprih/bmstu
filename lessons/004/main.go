package main

import "fmt"

func main() {
	var r, g, b uint8 = 0, 141, 213
	//var r, g, b uint8 = 0x00, 0x8d, 0xd5
	//fmt.Printf("color: #%02x%02x%02x\n", r16, g16, b16)
	fmt.Printf("color: #%02x%02x%02x\n", r, g, b)
}
