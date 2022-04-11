package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

//estrutura basica do banco
type Bank struct {
	Base `valid:"required"`
	Code 		string 		`json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name 		string 		`json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:BankID" valid:"-"`
}

//método por attach no go
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	} 
		return nil
}


//função - *Bank - ponteiro pra referenciar, q nem no C
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	//setando novo ID
	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	//validação de erro
	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	//retornar por referencia
	return &bank,nil
}