package service

import (
	"errors"
	"minipay/dto"
	"minipay/model"
	"minipay/repository"
	"time"

	"github.com/google/uuid"
)

type (
	WalletService interface {
		RegisterCustomer(reg *dto.RegisterCustomer) (string, error)
		RegisterMerchant(reg *dto.RegisterMerchant) (string, error)
		Topup(up *dto.TopUp) (string, error)
		Pay(pay *dto.Payment) (string, error)
	}
	WalletServiceImpl struct {
		userRepo        repository.UserRepo
		walletRepo      repository.WalletRepo
		transactionRepo repository.TransactionRepo
	}
)

func (w *WalletServiceImpl) Pay(pay *dto.Payment) (string, error) {
	wallet, err := w.findMerchantWallet(pay.Merchant)
	if err != nil {
		return "", errors.New("error")
	}
	cWallet, err := w.walletRepo.FindByID(pay.WalletID)
	if err != nil {
		return "", errors.New("error")
	}
	tcnID := uuid.NewString()
	tcn := model.Transaction{
		TransactionID:   tcnID,
		ReferenceID:     pay.ReferenceID,
		CreditWallet:    wallet.WalletID,
		DebitedWallet:   pay.WalletID,
		Description:     pay.Description,
		Amount:          pay.Amount,
		TransactionDate: time.Now().UTC(),
		TransactionType: model.PAYMENT,
	}
	if err := w.transactionRepo.Save(&tcn); err != nil {
		return "", err
	}
	if err := wallet.CreditBalance(pay.Amount); err != nil {
		return "", err
	}
	if err := w.walletRepo.Save(wallet); err != nil {
		return "", err
	}
	if err := cWallet.DebitBalance(pay.Amount); err != nil {
		return "", err
	}
	if err := w.walletRepo.Save(cWallet); err != nil {
		return "", err
	}
	return tcnID, nil
}

func (w *WalletServiceImpl) Topup(up *dto.TopUp) (string, error) {
	wallet, err := w.walletRepo.FindByID(up.WalletID)
	if err != nil {
		return "", errors.New("error")
	}
	tcdId := uuid.NewString()
	transaction := model.Transaction{
		TransactionID:   tcdId,
		ReferenceID:     up.ReferenceID,
		CreditWallet:    up.WalletID,
		Description:     up.Description,
		Amount:          up.Amount,
		TransactionDate: time.Now().UTC(),
	}
	if err := w.transactionRepo.Save(&transaction); err != nil {
		return "", errors.New("error")
	}
	wallet.CreditBalance(up.Amount)
	if err := w.walletRepo.Save(wallet); err != nil {
		return "", errors.New("error")
	}
	return tcdId, nil
}

func (w *WalletServiceImpl) RegisterCustomer(reg *dto.RegisterCustomer) (string, error) {
	account := model.UserAccount{
		UserID:      uuid.NewString(),
		Name:        reg.Name,
		Email:       reg.Email,
		Phonenumber: reg.Phonenumber,
		UserType:    model.CUSTOMER,
	}
	return w.createWallet(&account)
}

func (w *WalletServiceImpl) RegisterMerchant(reg *dto.RegisterCustomer) (string, error) {
	account := model.UserAccount{
		UserID:      uuid.NewString(),
		Name:        reg.Name,
		Email:       reg.Email,
		Phonenumber: reg.Phonenumber,
		UserType:    model.CUSTOMER,
	}
	return w.createWallet(&account)
}

func (w *WalletServiceImpl) createWallet(account *model.UserAccount) (string, error) {
	if err := w.userRepo.Save(account); err != nil {
		return "", errors.New("error")
	}
	walletId := uuid.NewString()
	wallet := model.Wallet{
		WalletID: walletId,
		UserID:   account.UserID,
	}
	if err := w.walletRepo.Save(&wallet); err != nil {
		return "", errors.New("errpr")
	}
	return walletId, nil
}

func (w *WalletServiceImpl) findMerchantWallet(name string) (*model.Wallet, error) {
	merchant, err := w.userRepo.FindMerchantByName(name)
	if err != nil {
		return nil, err
	}
	return w.walletRepo.FindByID(merchant.UserID)
}
