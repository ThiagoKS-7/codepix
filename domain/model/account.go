package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base 				  `valid:"required"`
	Owner string  `json:"owner_name" valid:"notnull"`
	Bank 	*Bank	  `valid:"-"`
	Number string `json:"number" valid:"notnull"`
	PixKeys []*PixKey `valid:"-"`
}

//método por attach no go
func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	} 
		return nil
}


// função - *Account - ponteiro pra referenciar, q nem no C
func NewAccount(bank *Bank, number string, owner string) (*Account, error) {
	account := Account {
		Owner:  owner,
		Bank:	bank,
		Number: number,

	}

	//setando novo ID
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	//validação de erro
	err := account.isValid()
	if err != nil {
		return nil, err
	}

	//retornar por referencia
	return &account,nil
}