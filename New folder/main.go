package main

import "fmt"

func main() {
	var i int
	var j int
	var n int

	n = 10

	for j = 0; j <= n; j++ {
		for i = 0; i <= n; i++ {
			// if i >= 0 && i < 2 || i > 4 && i < 7 || i >= 2 && i <= 4 && j == 1 || i > 6 && i <= 9 && j == 8 || i > 9 && i <= 11 {
			// 	fmt.Printf("**")
			// } else {
			// 	fmt.Printf("  ")
			// }

			if (j == 0 || j == n-1 || j <= n/2) || (j == 0 && i < n-1/2) || (j == n-1 && i > n-1/2) {
				fmt.Printf("*")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}
