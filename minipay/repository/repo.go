package repository

import "minipay/model"

type (
	UserRepo interface {
		FindById(userID string) (*model.UserAccount, error)
		FindMerchantByName(name string) (*model.UserAccount, error)
		Save(account *model.UserAccount) error
	}
	WalletRepo interface {
		FindByID(walletID string) (*model.Wallet, error)
		FindByUserID(userid string) (*model.Wallet, error)
		Save(*model.Wallet) error
	}
	TransactionRepo interface {
		FindByID(id string) (*model.Transaction, error)
		Save(trc *model.Transaction) error
	}
)
