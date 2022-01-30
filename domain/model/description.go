package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Todo: 最大文字数の検討
type Description struct {
	value string `validate:"min=1,max=100"`
}

func NewDescription(v string) (Description, error) {
	d := Description{}
	d.value = v

	if err := d.validateValue(); err != nil {
		return Description{}, fmt.Errorf("値の検証に失敗しました:%v", err)
	}

	return d, nil
}

func (d Description) String() string {
	return d.value
}

func (d Description) Equal(uu Description) bool {
	return d.value == uu.value
}

func (d Description) IsEmpty() bool {
	return d.value == ""
}

func (d Description) validateValue() error {
	if err := validator.New(); err != nil {
		return fmt.Errorf("値の検証に失敗しました:%v", err)
	}

	return nil
}
