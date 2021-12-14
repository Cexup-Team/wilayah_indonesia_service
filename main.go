package main

import (
	"api-test/app"
	"api-test/controller"
	"api-test/helper"
	"api-test/repository"
	"api-test/services"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
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
	router.GET("/api/provinsi", regionController.FindAll)
	router.GET("/api/wilayah/:wilayah_id", regionController.Find)

	// running server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
