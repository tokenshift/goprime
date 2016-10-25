package main

import (
	"fmt"
	"math/big"
	"os"
)

var bigZero = big.NewInt(0)
var bigOne  = big.NewInt(1)

func isPrime(num *big.Int) bool {
	candidate := big.NewInt(2)
	for {
		// If num/candidate < candidate, then we've passed (or reached) sqrt(num),
		// and can terminate.
		var div big.Int
		div.Div(num, candidate)
		if div.Cmp(candidate) < 0 {
			break
		}

		var mod, quot big.Int
		quot.DivMod(num, candidate, &mod)
		if bigZero.Cmp(&mod) == 0 {
			return false
		}

		candidate.Add(candidate, bigOne)
	}

	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a number to check for primality.")
		os.Exit(1)
	}

	input := os.Args[1]

	var num big.Int
	_, err := fmt.Sscan(input, &num)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Checking %s for primality...\n", num.String())
	if isPrime(&num) {
		fmt.Println(num.String(), "is prime.")
		os.Exit(0)
	} else {
		fmt.Println(num.String(), "is not prime.")
		os.Exit(1)
	}
}
