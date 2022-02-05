package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	switch req.Method {
	case http.MethodGet:
		defer req.Body.Close()

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		var r FindByUserIDRequest

		r.UserID = string(body)

		res, err := FindByUserID(r)
		if err != nil {
			panic(err)
		}

		fmt.Println(res)
		fmt.Println(string(body))

		// if err := getTodolist(w, req, db); err != nil {
		// 	log.Fatalf("getTodoList:%v", err)
		// }

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
