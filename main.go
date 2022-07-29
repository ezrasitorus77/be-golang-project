package main

import (
	"be-golang-project/delivery/middleware"
	"be-golang-project/infrastructure"
	"be-golang-project/models/handler"
	"be-golang-project/models/interface_"
	"be-golang-project/usecase/client"
	"be-golang-project/usecase/vendor"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var (
		midd      *middleware.Middleware = &middleware.Middleware{}
		parentCtx *handler.ParentContext
		newVendor interface_.Vendor = vendor.New()
		newClient interface_.Client = client.New()
		err       error
	)

	if err := infrastructure.ConfigInit(); err != nil {
		fmt.Println(err)
	}

	parentCtx, err = handler.ParentContext{}.Init()
	if err != nil {
		fmt.Println(err)
	}

	midd.RouteCollections = make([]middleware.Route, 0)
	midd.Mux = mux.NewRouter()
	midd.Ctx = parentCtx

	midd.AddRoute("/vendor-register", []string{"POST"}, newVendor.Register)
	midd.AddRoute("/login", []string{"POST"}, newVendor.Login)
	midd.AddRoute("/index", []string{"GET"}, newVendor.Index)
	midd.AddRoute("/vendor-profile", []string{"GET", "POST", "DELETE"}, newVendor.Profile)

	midd.AddRoute("/client-register", []string{"POST"}, newClient.Register)
	// midd.AddRoute("/client-profile", []string{"GET", "POST", "DELETE"}, newClient.Profile)

	midd.Serve()
	err = http.ListenAndServe(":8000", midd.Mux)
	if err != nil {
		fmt.Println(err)
	}
}
