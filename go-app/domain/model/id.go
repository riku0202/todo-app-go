package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ID struct {
	value string
}

func NewID() (ID, error) {
	i := ID{}
	i.value = uuid.New().String()

	if err := i.validateValue(); err != nil {
		return ID{}, fmt.Errorf("値の検証に失敗しました:%v", err)
	}

	return i, nil
}

func RestoreID(value string) (ID, error) {
	i := ID{}
	i.value = value

	if err := i.validateValue(); err != nil {
		return ID{}, fmt.Errorf("値の検証に失敗しました:%v", err)
	}

	return i, nil
}

func (i ID) String() string {
	return i.value
}

func (i ID) Equal(ii ID) bool {
	return i.value == ii.value
}

func (i ID) IsEmpty() bool {
	return i.value == ""
}

func (i ID) validateValue() error {
	if err := validator.New().Var(i.value, "uuid"); err != nil {
		return fmt.Errorf("値の検証に失敗しました:%v", err)
	}

	return nil
}

// 構造体をJSONに変換します
// func (i ID) MarshalJSON() ([]byte, error) {
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
// func (i *ID) UnmarshalJSON(b []byte) error {
// 	j := struct {
// 		Value string `json:"value"`
// 	}{}

// 	if err := json.Unmarshal(b, &j); err != nil {
// 		return errors.NewError(errors.ErrCodeDefault, "JSONを構造体に変換できません", err)
// 	}

// 	i.value = j.Value

// 	return nil
// }
