package main

import "fmt"

func main() {
	l := list_of_primeNumbers(3, 110)
	fmt.Println(l)
	d := list_of_primeNumbers(110, 200)
	fmt.Println(d)
	s := number_of_primes(3, 100000)
	fmt.Println(s)

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

func number_of_primes(m, n int) int {
	//this func returns number of prime number between m and n
	i := 0
	for v := m; v <= n; v++ {
		if is_prime(v) {
			i += 1
		}
	}
	return i
}
