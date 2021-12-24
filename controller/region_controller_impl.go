package controller

import (
	"net/http"
	"wilayah_indonesia_service/helper"
	"wilayah_indonesia_service/model/web"
	"wilayah_indonesia_service/services"

	"github.com/julienschmidt/httprouter"
)

type RegionControllerImpl struct {
	service services.RegionService
}

func NewRegionController(services services.RegionService) RegionController {
	return &RegionControllerImpl{service: services}
}

func (c *RegionControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	// helper.Cors(writer)
	regionResponse := c.service.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   regionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *RegionControllerImpl) Find(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	// helper.Cors(writer)
	id := param.ByName("wilayah_id")
	RegionResponse := c.service.Find(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   RegionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
