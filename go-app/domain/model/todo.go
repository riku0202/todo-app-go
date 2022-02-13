package model

import "fmt"

type Todo struct {
	Entity
	userId      UserId
	title       Title
	description Description
	isFinished  bool
}

func NewTodo(userId UserId, title Title, description Description) (*Todo, error) {
	t := &Todo{}

	id, err := NewID()
	if err != nil {
		return nil, fmt.Errorf("IDを作成できませんでした:%v", err)
	}

	e, err := NewEntity(id)
	if err != nil {
		return nil, fmt.Errorf("Entityを作成できません:%v", err)
	}

	t.Entity = e
	t.userId = userId
	t.title = title
	t.description = description
	t.isFinished = false

	return t, nil
}

func (t *Todo) UpdateIsFinished(isFinished bool) {
	t.isFinished = isFinished
}

// // メタ情報を更新します
// func (p *Todo) UpdateMeta(meta meta.Meta) {
// 	p.meta = meta
// }

// // ブロックを更新します
// func (p *Todo) UpdateBlocks(blocks []interface{}) error {
// 	p.blocks = blocks

// 	if err := p.validate(); err != nil {
// 		return errors.NewError(errors.ErrCodeDefault, "検証に失敗しました", err)
// 	}

// 	return nil
// }

// // ブロックのローカルIDを検証します
// func (p *Todo) validateBlockID() error {
// 	m := map[string]string{}

// 	for _, b := range p.blocks {
// 		d, err := json.Marshal(b)
// 		if err != nil {
// 			return errors.NewError(errors.ErrCodeDefault, "構造体をJSONに変換できません", err)
// 		}

// 		var bm map[string]interface{}

// 		if err := json.Unmarshal(d, &bm); err != nil {
// 			return errors.NewError(errors.ErrCodeDefault, "JSONを構造体に変換できません", err)
// 		}

// 		id := seeker.Str(bm, []string{"id", "value"})

// 		if _, ok := m[id]; ok {
// 			return errors.NewError(errors.ErrCodeDefault, "同じIDが存在しています")
// 		}

// 		m[id] = ""
// 	}

// 	return nil
// }

// // ブロック数を検証します
// func (p *Todo) validateBlockLimit() error {
// 	if len(p.blocks) > MaxLimitBlock {
// 		return errors.NewError(errors.ErrCodeDefault, "ブロック数が上限を超えています")
// 	}

// 	return nil
// }

// // JSONから構造体に変換します
// func (p Todo) MarshalJSON() ([]byte, error) {
// 	ej, err := json.Marshal(&p.Entity)
// 	if err != nil {
// 		return nil, errors.NewError(errors.ErrCodeDefault, "構造体をJSONに変換できません", err)
// 	}

// 	var data map[string]interface{}
// 	if err = json.Unmarshal(ej, &data); err != nil {
// 		return nil, errors.NewError(errors.ErrCodeDefault, "変換用のmapに変換できません", err)
// 	}

// 	data["meta"] = p.meta
// 	data["blocks"] = p.blocks
// 	data["asset_id"] = p.assetID

// 	b, err := json.Marshal(&data)
// 	if err != nil {
// 		return nil, errors.NewError(errors.ErrCodeDefault, "構造体をJSONに変換できません", err)
// 	}

// 	return b, nil
// }

// // JSONから構造体に変換します
// func (p *Todo) UnmarshalJSON(b []byte) error {
// 	e := model.Entity{}
// 	if err := json.Unmarshal(b, &e); err != nil {
// 		return errors.NewError(errors.ErrCodeDefault, "JSONを構造体に変換できません", err)
// 	}

// 	j := struct {
// 		Meta    meta.Meta     `json:"meta"`
// 		Blocks  []interface{} `json:"blocks"`
// 		AssetID model.AssetID `json:"asset_id"`
// 	}{}

// 	if err := json.Unmarshal(b, &j); err != nil {
// 		return errors.NewError(errors.ErrCodeDefault, "JSONを構造体に変換できません", err)
// 	}

// 	p.Entity = e
// 	p.meta = j.Meta
// 	p.blocks = j.Blocks
// 	p.assetID = j.AssetID

// 	return nil
// }

// // アセットを作成します
// //
// // アセットコンテキストのAPIをコールしてアセットを作成しアセットIDを返します。
// func (p *Todo) createAsset(ctx context.Context, siteID string) (model.AssetID, error) {
// 	GCSReq := map[string]interface{}{}
// 	gen.Gen(GCSReq, []string{"ctx"}, context.Background())
// 	gen.Gen(GCSReq, []string{"bucket"}, os.Getenv(env.KeyGCSBucketAxisStorage))
// 	gen.Gen(GCSReq, []string{"path"}, GCSPathIndexDist)

// 	GCSRawRes, err := api.Call(api.GCSReadObject, ctx, GCSReq)
// 	if err != nil {
// 		return model.AssetID{}, errors.NewError(errors.ErrCodeDefault, "デフォルトページを取得できません", err)
// 	}

// 	GCSRes, ok := GCSRawRes.(map[string]interface{})
// 	if !ok {
// 		return model.AssetID{}, errors.NewError(errors.ErrCodeDefault, "型アサーションに失敗しました")
// 	}

// 	assetReq := map[string]interface{}{}
// 	gen.Gen(assetReq, []string{"tag"}, siteID)
// 	gen.Gen(assetReq, []string{"path"}, fmt.Sprintf(GCSPathSite, siteID)+"/"+FileName)
// 	gen.Gen(assetReq, []string{"binary"}, GCSRes["binary"])
// 	gen.Gen(assetReq, []string{"binary_limit"}, float64(MaxTodoCap))

// 	assetRawRes, err := api.Call(api.AssetCreateAsset, ctx, assetReq)
// 	if err != nil {
// 		return model.AssetID{}, errors.NewError(errors.ErrCodeDefault, "axis-assetのリソースを作成できません", err)
// 	}

// 	assetRes, ok := assetRawRes.(map[string]interface{})
// 	if !ok {
// 		return model.AssetID{}, errors.NewError(errors.ErrCodeDefault, "型アサーションに失敗しました")
// 	}

// 	i := seeker.Str(assetRes, []string{"id"})
// 	assetID, err := model.NewAssetID(i)
// 	if err != nil {
// 		return model.AssetID{}, errors.NewError(errors.ErrCodeDefault, "アセットIDを作成できません", err)
// 	}

// 	return assetID, nil
// }
