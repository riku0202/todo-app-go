package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Todo: 最大文字数の検討
type Title struct {
	value string `validate:"required,min=1,max=50"`
}

func NewTitle(v string) (Title, error) {
	t := Title{}
	t.value = v

	if err := t.validateValue(); err != nil {
		return Title{}, fmt.Errorf("値の検証に失敗しました:%v", err)
	}

	return t, nil
}

func (t Title) String() string {
	return t.value
}

func (t Title) IsEmpty() bool {
	return t.value == ""
}

func (t Title) validateValue() error {
	if err := validator.New(); err != nil {
		return fmt.Errorf("値の検証に失敗しました:%v", err)
	}

	return nil
}
