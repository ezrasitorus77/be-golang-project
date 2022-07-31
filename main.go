package main

import (
	"be-golang-project/delivery/middleware"
	"be-golang-project/infrastructure"
	"be-golang-project/models/handler"
	"be-golang-project/models/interface_"
	"be-golang-project/usecase/user"
	"be-golang-project/usecase/vendor"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var (
		midd      *middleware.Middleware = &middleware.Middleware{}
		parentCtx *handler.ParentContext
		newUser   interface_.User   = user.New()
		newVendor interface_.Vendor = vendor.New()
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

	midd.AddRoute("/vendor/register", []string{"POST"}, newVendor.Register)

	midd.AddRoute("/user/register", []string{"POST"}, newUser.Register)
	midd.AddRoute("/login", []string{"POST"}, newUser.Login)
	midd.AddRoute("/index", []string{"GET"}, newUser.Index)
	midd.AddRoute("/profile", []string{"GET", "POST", "DELETE"}, newUser.Profile)

	midd.Serve()
	err = http.ListenAndServe(":8000", midd.Mux)
	if err != nil {
		fmt.Println(err)
	}
}
