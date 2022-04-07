package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

//init - toda vez que carrega, faz algo (tipo um watch em js)
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

//Criar keys base pra serem herdadas pelo bank
type Base struct {
	ID string 					`json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}