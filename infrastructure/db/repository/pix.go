package repository

import (
	"codepix/domain/model"

	"gorm.io/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) RegisterKey(account *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) Addccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) AddTransaction(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}