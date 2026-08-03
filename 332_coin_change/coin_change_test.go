package coinChange

import "testing"

func Test_coinChange(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "Case 1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "Case 2",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "Case 3",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "Fail Case",
			coins:  []int{186, 419, 83, 408},
			amount: 6249,
			want:   20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("coins:%v, amount:%v, coinChange() = %v, want:%v, ",
					tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}
