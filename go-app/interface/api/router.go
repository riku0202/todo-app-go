package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/riku0202/todo-app-go/go-app/infrastructure/di"
)

// func todoHandler(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

// 	switch req.Method {
// 	case http.MethodGet:
// 		FindByUserID()

// 		if err := getTodolist(w, req, db); err != nil {
// 			log.Fatalf("getTodoList:%v", err)
// 		}

// 	case http.MethodPost:
// 		if err := post(w, req, db); err != nil {
// 			log.Fatalf("postTodolist:%v", err)
// 		}

// 	case http.MethodDelete:
// 		if err := delete(w, req, db); err != nil {
// 			log.Fatalf("postTodolist:%v", err)
// 		}

// 	case http.MethodOptions:
// 		w.WriteHeader(http.StatusOK)

// 	default:
// 		http.Error(w, fmt.Sprintf("HTTP Method %s Not Allowed", req.Method), http.StatusMethodNotAllowed)
// 	}
// }

type CreateTodoRequest struct {
	UserID      string `validate:"required,uuid"`
	title       string `validate:"required,min=1,max=50"`
	description string `validate:"max=100"`
}

func CreateTodo(req CreateTodoRequest) (Res, error) {
	var res Res

	if err := validator.New().Struct(req); err != nil {
		return res, fmt.Errorf("リクエストが不正です:%v", err)
	}

	app, err := di.InitApp()
	if err != nil {
		return res, fmt.Errorf("アプリケーションを初期化できません:%v", err)
	}

	err = app.CreateTodo(req.UserID, req.title, req.description)
	if err != nil {
		return res, fmt.Errorf("Todoを作成できません:%v", err)
	}

	return res, nil
}

type FindByUserIDRequest struct {
	UserID string `json:"user_id"`
	// validate:"required"
}

func FindByUserID(req FindByUserIDRequest) (Res, error) {
	var res Res

	if err := validator.New().Struct(req); err != nil {
		return res, fmt.Errorf("リクエストが不正です:%v", err)
	}

	query, err := di.InitQuery()
	if err != nil {
		return res, fmt.Errorf("クエリーを初期化できません:%v", err)
	}

	m, err := query.FindByUserID(res.UserId)
	if err != nil {
		return res, fmt.Errorf("Todoを取得できませんでした:%v", err)
	}

	// res, err = shared.NewCardRes(m)
	// if err != nil {
	// 	return res, errors.NewError("レスポンスを作成できませんでした", err)
	// }

	fmt.Println(m)

	return res, nil
}

type Res struct {
	ID          string
	UserId      string
	Title       string
	Description string
	IsFinished  bool
}

func NewRes(d map[string]interface{}) (Res, error) {
	response := Res{
		ID: d["id"].(string),
	}

	// response := CardRes{
	// 	Res: shared.Res{
	// 		ID:            parser.Str(d, []string{"id", "value"}),
	// 		PaymentMethod: parser.Str(d, []string{"method", "value"}),
	// 		PaymentDate:   parser.Time(d, []string{"date"}),
	// 		Billing: struct {
	// 			ID   string
	// 			Date time.Time
	// 		}{
	// 			ID:   parser.Str(d, []string{"billing", "id", "value"}),
	// 			Date: parser.Time(d, []string{"billing", "date"}),
	// 		},
	// 	},
	// 	ContractID: parser.Str(d, []string{"contract_id", "value"}),
	// 	PaidStatus: struct {
	// 		IsPaid     bool
	// 		IsExecuted bool
	// 	}{
	// 		IsPaid:     parser.Bool(d, []string{"paid_status", "is_paid"}),
	// 		IsExecuted: parser.Bool(d, []string{"paid_status", "is_executed"}),
	// 	},
	// 	Amount: uint(parser.Float64(d, []string{"amount", "value"})),
	// 	StripeID: struct {
	// 		PaymentIntentID string
	// 		PaymentMethodID string
	// 	}{
	// 		PaymentIntentID: parser.Str(d, []string{"stripe_id", "payment_intent_id", "value"}),
	// 		PaymentMethodID: parser.Str(d, []string{"stripe_id", "method_id", "value"}),
	// 	},
	// }

	return response, nil
}
