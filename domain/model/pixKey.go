package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKey struct {
	Base 				  `valid:"required"`
	Kind string 	`json:"kind" valid:"notnull"`
	Key string 		`json:"key" valid:"notnull"`
	AuthID string `json:"auth_id" valid:"notnull"`
	Account *Account `valid="-"`
	Status string `json:"status" valid:"notnull"`
}

//método por attach no go
func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}
	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid type of key")
	}


	if err != nil {
		return err
	} 
		return nil
}


// função - *PixKey- ponteiro pra referenciar, q nem no C
func NewPixKey(kind string, account *Account,  key string) (*PixKey, error) {
	pixKey := PixKey {
		Kind:  kind,
		Key:	key,
		Account:	account,
		Status: "active",

	}

	//setando novo ID
	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	//validação de erro
	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	//retornar por referencia
	return &pixKey,nil
}