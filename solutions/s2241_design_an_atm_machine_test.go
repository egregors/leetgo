package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewATM(t *testing.T) {
	/**
	 * Your ATM object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Deposit(banknotesCount);
	 * param_2 := obj.Withdraw(amount);
	 */
	atm := NewATM()
	atm.Deposit([]int{0, 0, 1, 2, 1})
	assert.Equal(t, []int{0, 0, 1, 0, 1}, atm.Withdraw(600))
	atm.Deposit([]int{0, 1, 0, 1, 1})
	assert.Equal(t, []int{-1}, atm.Withdraw(600))
	assert.Equal(t, []int{0, 1, 0, 0, 1}, atm.Withdraw(550))
}
