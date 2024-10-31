package main

import "fmt"

func main() {
	l := list_of_primeNumbers(3, 110)
	fmt.Println(l)

}

func is_prime(n int) bool {
	// This function receives an integer and returns true if it is prime and false otherwise.
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func list_of_primeNumbers(m, n int) []int {
	//this func returns list of prime numbers between m and n
	list := []int{}
	for z := m; z <= n; z++ {
		if is_prime(z) {
			list = append(list, z)
		}
	}
	fmt.Printf("list of prime numbers between %d and %d is :", m, n)
	return list
}
