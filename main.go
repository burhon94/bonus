package main

func main() {

}

func bonus(amount [] int) int {
	const boundPurchanse = 10_000
	const percent = 5
	amountBonus := 0
	for _, amount := range amount {
		if amount > boundPurchanse {
			amountBonus += (amount - boundPurchanse) * percent / 100
		}
	}
	return amountBonus
}
