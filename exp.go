package main

import (
 "fmt"
)
func main() {

	var ip int
	var rem int
	var rev int
	var act int=0
	fmt.Println("Enter Number:")
	fmt.Scanf("%d\n", &ip)

	act=ip
	fmt.Println("This is the number",ip)

	for ip>0 {
	rem=ip%10
	rev=(rev*10)+rem
	ip=ip/10
	}
	if act==rev {
		fmt.Println("Polindrome")
	}else{
		fmt.Println("Not Poloindrome")
	}
	}
   