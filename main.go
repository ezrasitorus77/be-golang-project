package main

import (
	// "be-golang-project/delivery/middleware"
	_ "be-golang-project/config"
	"log"
	"net/http"

	// "be-golang-project/models/handler"
	// "be-golang-project/models/interface_"
	"be-golang-project/controller"
	// "be-golang-project/repository"
	// "be-golang-project/service"

	handlerController "github.com/ezrasitorus77/http-handler/controller"
	handlerDelivery "github.com/ezrasitorus77/http-handler/domain/delivery"
	handlerService "github.com/ezrasitorus77/http-handler/service"
	// vendor "be-golang-project/usecase/vendor_"
	// "fmt"
	// "net/http"
	// "github.com/gorilla/mux"
)

func main() {
	// Since this project utilizez init(), it is IMPORTANT to declare the variables in SEQUENCE
	// 1. repository
	// 2. service
	// 3. controller

	var (
		e             error
		router        handlerDelivery.Router            = handlerService.RouterService
		midService    handlerDelivery.MiddlewareService = handlerService.MiddlewareService
		midController handlerDelivery.Handler
		server        http.Server
	)

	router.GET("/->d:id", controller.ClientController.Get)

	midController = handlerController.NewMiddleware(router, midService)

	server = http.Server{
		Addr:    "localhost:8080",
		Handler: midController,
	}

	e = server.ListenAndServe()
	if e != nil {
		log.Fatal(e)
	}
}
