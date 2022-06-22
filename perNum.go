package main

import "fmt"

func main() {
	var userNum int
	fmt.Println("Enter your wanted number and: ")
	fmt.Scan(&userNum)
	d := 5
	for userNum != 0 {
		d++
		if Perf(d) {
			userNum--
			if userNum == 0 {
				fmt.Println(d)
			}
		}
	}
}

func Perf(d int) bool {
	sum := 1
	for i := 2; i <= d/2; i++ {
		if d%i == 0 {
			sum += i
		}
	}
	if sum == d {
		return true
	}
	return false
}
