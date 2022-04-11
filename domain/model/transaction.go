package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending string = "pending"
	TransactionCompleted string = "completed"
	TransactionError string = "error"
	TransactionCancel string = "canceled"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}


type Transaction struct {
	Base 													`valid:"required"`
	AccountFrom 			*Account		`valid:"-"`
	AccountFromID 		string 			`gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Amount 						float64			`json:"ammount" gorm:"type:float" valid:"notnull"`
	PixKeyTo 					*PixKey			`valid:"-"`
	PixKeyIdTo				string 			`gorm:"column:pix_key_id_to; type:uuid; not null" valid:"notnull"`
	Status 						string			`json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description 			string			`json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	CancelDescription string 			`json:"cancel_description" gorm:"type:varchar(255)" valid:"notnull"`
}


//método por attach no go
func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("The amount must be greater than zero")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("Invalid status for the transaction")
	}

	if t.PixKeyTo.Account.ID == t.AccountFrom.ID {
		return errors.New("The source account cant trasfer to itself")
	}

	if err != nil {
		return err
	} 
		return nil
}


//função - *Bank - ponteiro pra referenciar, q nem no C
func NewTransaction(AccountFrom *Account, Amount float64, PixKeyTo *PixKey, description string, ) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: AccountFrom,
		Amount: Amount,
		PixKeyTo: PixKeyTo,
		Status: TransactionPending,
		Description: description,
	}


	//setando novo ID
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	//validação de erro
	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	//retornar por referencia
	return &transaction,nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionCancel
	t.UpdatedAt = time.Now()
	t.Description = description
	err := t.isValid()
	return err
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}