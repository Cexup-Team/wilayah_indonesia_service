package helper

import (
	"encoding/json"
	"net/http"
	"wilayah_indonesia_service/model/web"
)

func WriteToRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, webResponse web.WebResponse) {
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}
