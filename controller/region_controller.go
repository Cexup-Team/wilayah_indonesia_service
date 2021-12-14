package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/**
 * Todo: region controller interface
 */
type RegionController interface {
	Find(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
