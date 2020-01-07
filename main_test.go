package main

import "testing"

func Test_bonus(t *testing.T) {
	tests := []struct {
		name string
		amount int
		want int
	}{
		{"Punches bonus", 12_000, 100 },
		{"No bonus", 5_000, 0},
	}
	for _, test := range tests {
		got := bonus(test.amount)
		if got != test.want {
			t.Error("For bonus test:", test.name, "got:", got, "want:", test.want)
		}
	}
}