package main

func main() {

}

func bonus(amount [] int) int {
	const boundForBonus = 10_000
	const percent = 5
	amountBonus := 0
	for _, amount := range amount {
		if amount > boundForBonus {
			amountBonus += (amount - boundForBonus) * percent / 100
		}
	}
	return amountBonus
}
