package main

import "fmt"

func main() {
	const usd2Eur = 0.88
	const usd2Rub = 83.5
	var rubAmmount float64 = 100
	eur2rub := (rubAmmount / usd2Rub) * usd2Eur

	fmt.Print(eur2rub)
}
