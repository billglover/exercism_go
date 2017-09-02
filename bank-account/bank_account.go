package account

import "sync"

const testVersion = 1

// Account represents a bank account
type Account struct {
	balance int64
	open    bool
	mutex   *sync.RWMutex
}

// Open takes an initial deposit and returns an Account. It returns nil if the
// initial deposit is less than zero.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	a := Account{
		open:    true,
		mutex:   &sync.RWMutex{},
		balance: initialDeposit,
	}

	return &a
}

// Close marks an account as closed and returns the final payout. If the
// account is already closed it returns ok = false.
func (a *Account) Close() (payout int64, ok bool) {
	defer a.mutex.Unlock()
	a.mutex.Lock()

	if a.open == false {
		return 0, false
	}
	payout = a.balance
	a.balance = 0
	a.open = false

	return payout, true
}

// Balance returns the balance on an account If the account is already closed
// it returns ok = false.
func (a *Account) Balance() (balance int64, ok bool) {
	defer a.mutex.RUnlock()
	a.mutex.RLock()

	return a.balance, a.open
}

// Deposit makes a deposit against an account and returns the new balance. If
// the account is already closed it returns ok = false. Deposits can be
// positive or negative but must not result in a balance that falls below zero.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	defer a.mutex.Unlock()
	a.mutex.Lock()

	if a.open == false {
		return a.balance, false
	}
	newBalance = a.balance + amount

	if newBalance < 0 {
		return a.balance, false
	}
	a.balance = newBalance

	return a.balance, true
}
