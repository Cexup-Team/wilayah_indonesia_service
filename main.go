package main

import (
	"net/http"
	"wilayah_indonesia_service/app"
	"wilayah_indonesia_service/controller"
	"wilayah_indonesia_service/helper"
	"wilayah_indonesia_service/repository"
	"wilayah_indonesia_service/services"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func Hello(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	println("hello")
}

func main() {
	// inisiate data
	DB := app.NewDB()
	validate := validator.New()
	regionRepository := repository.NewRegionRepository()
	regionService := services.NewRegionService(regionRepository, DB, validate)
	regionController := controller.NewRegionController(regionService)
	router := httprouter.New()

	// router list
	router.GET("/api/wilayah", regionController.FindAll)
	router.GET("/api/wilayah/:wilayah_id", regionController.Find)

	handler := cors.Default().Handler(router)
	// running server
	server := http.Server{
		Addr:    ":3000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
