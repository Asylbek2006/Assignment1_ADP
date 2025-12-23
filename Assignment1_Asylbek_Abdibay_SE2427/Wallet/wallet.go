package Wallet

import "fmt"

type Wallet struct {
	Owner        string
	Balance      float64
	Transactions []string
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	w.Transactions = append(w.Transactions, fmt.Sprintf("Deposited: +$%.2f", amount))
}

func (w *Wallet) SpendMoney(amount float64) bool {
	if amount > w.Balance {
		fmt.Println("Insufficient funds!")
		return false
	}
	w.Balance -= amount
	w.Transactions = append(w.Transactions, fmt.Sprintf("Spent: -$%.2f", amount))
	return true
}

func (w *Wallet) GetBalance() {
	fmt.Printf("Current Balance: $%.2f\n", w.Balance)
}

func (w *Wallet) PrintHistory() {
	fmt.Println("Transaction History:")
	for _, t := range w.Transactions {
		fmt.Println(t)
	}
}
