/*
	https://leetcode.com/problems/design-an-atm-machine/description/

	There is an ATM machine that stores banknotes of 5 denominations: 20, 50, 100, 200, and 500 dollars. Initially
	the ATM is empty. The user can use the machine to deposit or withdraw any amount of money.

	When withdrawing, the machine prioritizes using banknotes of larger values.

		For example, if you want to withdraw $300 and there are 2 $50 banknotes, 1 $100 banknote, and 1 $200 banknote,
	then the machine will use the $100 and $200 banknotes.
		However, if you try to withdraw $600 and there are 3 $200 banknotes and 1 $500 banknote, then the withdraw
	request will be rejected because the machine will first try to use the $500 banknote and then be unable to use
	banknotes to complete the remaining $100. Note that the machine is not allowed to use the $200 banknotes instead
	of the $500 banknote.

	Implement the ATM class:

		ATM() Initializes the ATM object.
		void deposit(int[] banknotesCount) Deposits new banknotes in the order $20, $50, $100, $200, and $500.
		int[] withdraw(int amount) Returns an array of length 5 of the number of banknotes that will be handed to the
	user in the order $20, $50, $100, $200, and $500, and update the number of banknotes in the ATM after withdrawing.
	Returns [-1] if it is not possible (do not withdraw any banknotes in this case).
*/

package solutions

type ATM struct {
	noms    []int
	amounts []int
}

// NewATM must call Constructor to pass LeetCode tests
func NewATM() ATM {
	return ATM{
		noms:    []int{20, 50, 100, 200, 500},
		amounts: []int{0, 0, 0, 0, 0},
	}
}

func (a *ATM) total() int {
	total := 0
	for i := 0; i < 5; i++ {
		total += a.noms[i] * a.amounts[i]
	}
	return total
}

func (a *ATM) Deposit(banknotesCount []int) {
	for i, amount := range banknotesCount {
		a.amounts[i] += amount
	}
}

func (a *ATM) Withdraw(amount int) []int {
	if amount < 0 {
		return []int{-1}
	}
	if amount > a.total() {
		return []int{-1}
	}

	res := make([]int, 5)
	prev := make([]int, 5)
	copy(prev, a.amounts)
	i := 4
	for amount > 0 && i >= 0 {
		if amount >= a.noms[i] {
			n := min(a.amounts[i], amount/a.noms[i])
			res[i] = min(n, a.amounts[i])
			a.amounts[i] -= n
			amount -= res[i] * a.noms[i]
		}
		i--
	}

	if amount != 0 {
		a.amounts = prev
		return []int{-1}
	}

	return res
}
