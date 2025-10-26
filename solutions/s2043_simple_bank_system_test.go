package solutions

import (
	"testing"
)

func TestNewBank(t *testing.T) {
	//Input
	//["Bank", "withdraw", "transfer", "deposit", "transfer", "withdraw"]
	//[[[10, 100, 20, 50, 30]], [3, 10], [5, 1, 20], [5, 20], [3, 4, 15], [10, 50]]
	//Output
	//[null, true, true, true, false, false]
	bank := NewBank([]int64{10, 100, 20, 50, 30})
	if !bank.Withdraw(3, 10) {
		t.Errorf("Withdraw(3, 10) = false; want true")
	}
	if !bank.Transfer(5, 1, 20) {
		t.Errorf("Transfer(5, 1, 20) = false; want true")
	}
	if !bank.Deposit(5, 20) {
		t.Errorf("Deposit(5, 20) = false; want true")
	}
	if bank.Transfer(3, 4, 15) {
		t.Errorf("Transfer(3, 4, 15) = true; want false")
	}
	if bank.Withdraw(10, 50) {
		t.Errorf("Withdraw(10, 50) = true; want false")
	}
}
