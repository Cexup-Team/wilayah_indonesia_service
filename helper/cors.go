package helper

import "net/http"

func Cors(writer *http.ResponseWriter, r *http.Request) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
	(*writer).Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	(*writer).Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		(*writer).Write([]byte("allowed"))
		return
	}
}
