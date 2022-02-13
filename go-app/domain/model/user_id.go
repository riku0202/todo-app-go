package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserId struct {
	value string `validate:"required,uuid"`
}

func NewUserId(v string) (UserId, error) {
	u := UserId{}
	u.value = v

	if err := u.validateValue(); err != nil {
		return UserId{}, fmt.Errorf("値の検証に失敗しました:%v", err)
	}

	return u, nil
}

func (u UserId) String() string {
	return u.value
}

func (u UserId) Equal(uu UserId) bool {
	return u.value == uu.value
}

func (t UserId) validateValue() error {
	if err := validator.New(); err != nil {
		return fmt.Errorf("値の検証に失敗しました:%v", err)
	}

	return nil
}

// // 構造体をJSONに変換します
// func (i UserId) MarshalJSON() ([]byte, error) {
// 	j := struct {
// 		Value string `json:"value"`
// 	}{
// 		Value: i.value,
// 	}

// 	b, err := json.Marshal(j)
// 	if err != nil {
// 		return nil, errors.NewError(errors.ErrCodeDefault, "構造体をJSONに変換できません", err)
// 	}

// 	return b, nil
// }

// // JSONを構造体に変換します
// func (i *UserId) UnmarshalJSON(b []byte) error {
// 	j := struct {
// 		Value string `json:"value"`
// 	}{}

// 	if err := json.Unmarshal(b, &j); err != nil {
// 		return errors.NewError(errors.ErrCodeDefault, "JSONを構造体に変換できません", err)
// 	}

// 	i.value = j.Value

// 	return nil
// }
