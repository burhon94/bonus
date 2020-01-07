package main

import "fmt"

func main() {
	sales := [] int{12_000, 8_000, 15_000, 8_000}
	amountBonus := 0
	for _, sales := range sales {
		amountBonus += bonus(sales)
	}
	fmt.Println(amountBonus)
}

func bonus(amount int) int {
	const boundPurchanse = 10_000
	const percent = 5
	amountBonus := 0
	if amount > boundPurchanse {
		amountBonus = (amount - boundPurchanse) * percent / 100
	}
	return amountBonus
}
