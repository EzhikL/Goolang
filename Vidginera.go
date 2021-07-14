package main

import "fmt"

func main() {

	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	
	length := len(cipherText)
	keylength := len(keyword)
	
	for i:=0; i<length; i++ {
		tab := i % keylength
		c := cipherText[i] - 65
		a := keyword[tab] - 65

		c = ((c - a + 26)%26) + 65
		fmt.Printf("%c", c)
		}
	
}
