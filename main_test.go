package main

import "testing"

func Test_bonus(t *testing.T) {
	tests := []struct {
		name   string
		amount [] int
		want   int
	}{
		{"Purchanse bonus", [] int{12_000, 8_000, 15_000, 8_000}, 350},
		{"No bonus", [] int{5_000}, 0},
	}
	for _, test := range tests {
		got := bonus(test.amount)
		if got != test.want {
			t.Error("For bonus test:", test.name, "got:", got, "want:", test.want)
		}
	}
}
