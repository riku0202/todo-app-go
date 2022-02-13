package model

type Entity struct {
	id ID
}

func NewEntity(id ID) (Entity, error) {
	e := Entity{}
	e.id = id

	return e, nil
}

func (e *Entity) ID() ID {
	return e.id
}

// // 構造体をJSONに変換します
// func (e *Entity) MarshalJSON() ([]byte, error) {
// 	j := struct {
// 		ID ID `json:"id"`
// 	}{}

// 	j.ID = e.id

// 	b, err := json.Marshal(j)
// 	if err != nil {
// 		return nil, errors.NewError(errors.ErrCodeDefault, "構造体をJSONに変換できません", err)
// 	}

// 	return b, nil
// }

// // JSONを構造体に変換します
// func (e *Entity) UnmarshalJSON(b []byte) error {
// 	j := struct {
// 		ID ID `json:"id"`
// 	}{}

// 	if err := json.Unmarshal(b, &j); err != nil {
// 		return errors.NewError(errors.ErrCodeDefault, "JSONを構造体に変換できません", err)
// 	}

// 	e.id = j.ID

// 	return nil
// }
