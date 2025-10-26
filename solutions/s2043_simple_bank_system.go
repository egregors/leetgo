/*
	https://leetcode.com/problems/simple-bank-system/

	You have been tasked with writing a program for a popular bank that will
		automate all its incoming transactions (transfer, deposit, and withdraw). The
		bank has n accounts numbered from 1 to n. The initial balance of each account
		is stored in a 0-indexed integer array balance, with the (i + 1)th account
		having an initial balance of balance[i].

	Execute all the valid transactions. A transaction is valid if:

		The given account number(s) are between 1 and n, and
		The amount of money withdrawn or transferred from is less than or equal to the
			balance of the account.

	Implement the Bank class:

		Bank(long[] balance) Initializes the object with the 0-indexed integer array
			balance.
		boolean transfer(int account1, int account2, long money) Transfers money
			dollars from the account numbered account1 to the account numbered account2.
			Return true if the transaction was successful, false otherwise.
		boolean deposit(int account, long money) Deposit money dollars into the
			account numbered account. Return true if the transaction was successful,
			false otherwise.
		boolean withdraw(int account, long money) Withdraw money dollars from the
			account numbered account. Return true if the transaction was successful,
			false otherwise.
*/

package solutions

type Bank struct {
	accs []int64
}

// NewBank must call Constructor to pass LeetCode tests
func NewBank(balance []int64) Bank {
	return Bank{accs: append([]int64{-1}, balance...)}
}

func (b *Bank) apply(fn func() bool, ids ...int) bool {
	for _, id := range ids {
		if id < 0 || id > len(b.accs) {
			return false
		}
	}

	return fn()
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	return b.apply(func() bool {
		if b.accs[account1]-money < 0 {
			return false
		}
		b.accs[account1] -= money
		b.accs[account2] += money

		return true
	}, account1, account2)
}

func (b *Bank) Deposit(account int, money int64) bool {
	return b.apply(func() bool {
		if money < 0 {
			return false
		}
		b.accs[account] += money

		return true
	}, account)
}

func (b *Bank) Withdraw(account int, money int64) bool {
	return b.apply(func() bool {
		if money < 0 {
			return false
		}
		if b.accs[account]-money < 0 {
			return false
		}
		b.accs[account] -= money
		return true
	}, account)
}
