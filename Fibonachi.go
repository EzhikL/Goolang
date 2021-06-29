package main

import "fmt"

func rek(n int) int {
	if n == 0 {
		return 0}
	if n ==1 {
		return 1}
	return rek(n-1)+rek(n-2)
}

func sum(n int) int {
	u:=0;
	for i:=0; i<=n; i++ {
		u=rek(i)+u
	}
return u;
}

func main() {

	fmt.Println(rek(5))
	fmt.Println(sum(5))
	}	
