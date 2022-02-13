package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	switch req.Method {
	case http.MethodGet:

		params := req.URL.Query().Get("user_id")
		if params == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		req := FindByUserIDRequest{
			UserID: params,
		}

		_, err := FindByUserID(req)
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("検索に失敗しました: %w", err), 400)
		// }
		fmt.Println(err)

		res := Res{
			ID:          "01",
			UserId:      "riku0202",
			Title:       "今日のおかず",
			Description: "コロッケ",
			IsFinished:  false,
		}

		json, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(json)

	// case http.MethodPost:
	// 	if err := post(w, req, db); err != nil {
	// 		log.Fatalf("postTodolist:%v", err)
	// 	}

	// case http.MethodDelete:
	// 	if err := delete(w, req, db); err != nil {
	// 		log.Fatalf("postTodolist:%v", err)
	// 	}

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, fmt.Sprintf("HTTP Method %s Not Allowed", req.Method), http.StatusMethodNotAllowed)
	}
}
