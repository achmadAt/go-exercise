package model

import (
	"errors"
	"time"
)

const (
	CUSTOMER = iota
	MERCHANT
	TOPUP = iota
	PAYMENT
)

type (
	UserAccount struct {
		UserID      string
		Name        string
		Email       string
		Phonenumber string
		UserType    int
	}
	Wallet struct {
		WalletID string
		UserID   string
		Balance  int64
	}
	Transaction struct {
		TransactionID   string
		ReferenceID     string
		CreditWallet    string
		DebitedWallet   string
		Description     string
		TransactionDate time.Time
		Amount          int64
		TransactionType int
	}
)

func (w *Wallet) CreditBalance(amount int64) error {
	w.Balance = w.Balance + amount
	return nil
}

func (w *Wallet) DebitBalance(amount int64) error {
	if w.Balance < amount {
		return errors.New("not enough")
	}
	w.Balance = w.Balance - amount
	return nil
}
