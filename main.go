package main

import (
	"be-golang-project/delivery/middleware"
	"be-golang-project/infrastructure"
	"be-golang-project/models/handler"
	"be-golang-project/models/interface_"
	"be-golang-project/usecase/user"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var (
		midd      *middleware.Middleware = &middleware.Middleware{}
		parentCtx *handler.ParentContext
		newUser   interface_.User = user.New()
		existUser interface_.User = user.Login()
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

	midd.AddRoute("/register", "POST", newUser.Register)
	midd.AddRoute("/login", "POST", existUser.Login)
	midd.AddRoute("/index", "GET", existUser.Index)

	midd.Serve()
	err = http.ListenAndServe(":8000", midd.Mux)
	if err != nil {
		fmt.Println(err)
	}
}
