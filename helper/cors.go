package helper

import "net/http"

func Cors(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
}
