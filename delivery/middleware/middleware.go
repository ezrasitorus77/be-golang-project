package middleware

import (
	"be-golang-project/consts"
	"be-golang-project/models/handler"
	"be-golang-project/models/interface_"
	"be-golang-project/models/response"
	"be-golang-project/models/token"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	resp *response.Response = &response.Response{}
)

type (
	Middleware struct {
		Mux              *mux.Router
		RouteCollections []Route
		Ctx              *handler.ParentContext
	}

	Route struct {
		path   string
		method string
		hFunc  func(*handler.Context)
	}
)

func (m *Middleware) AddRoute(path, method string, hFunc func(*handler.Context)) error {
	if err := m.routeAvailabilityCheck(path, method, hFunc); err != nil {
		return err
	}

	return nil
}

func (m *Middleware) routeAvailabilityCheck(path, method string, hFunc func(*handler.Context)) error {
	for _, route := range m.RouteCollections {
		if route.path == path {
			if route.method == method {
				err := errors.New("Path with specific method exists")
				return err
			}
		}
	}

	m.RouteCollections = append(m.RouteCollections, Route{
		path:   path,
		method: method,
		hFunc:  hFunc,
	})

	return nil
}

func (m *Middleware) methodCheck(w http.ResponseWriter, r *http.Request) {
	var (
		determinant  bool = false
		hFunc        func(*handler.Context)
		incomingPath string           = r.URL.Path
		ctx          *handler.Context = &handler.Context{}
	)

	for _, route := range m.RouteCollections {
		if route.path == incomingPath {
			if route.method == r.Method {
				determinant = true
				hFunc = route.hFunc

				break
			}
		}
	}

	if !determinant {
		resp.W = w
		err := errors.New(fmt.Sprintf("%s doesn't accept %s method", incomingPath, r.Method))
		resp.SendResponse(http.StatusMethodNotAllowed, consts.MethodNotAllowedRC, consts.MethodNotAllowedMessage, nil, err)

		return
	}

	ctx = m.Ctx.Set(ctx, w, r, hFunc)
	if incomingPath == "/login" || incomingPath == "/register" {
		ctx.Value.HandleFunc(ctx)
	} else {
		m.tokenCheck(ctx, m.Ctx.DB, r)
	}
}

func (m *Middleware) tokenCheck(ctx *handler.Context, parentCtxDB *gorm.DB, r *http.Request) {
	var (
		userToken  string = ctx.Value.Request.Header.Get("Token")
		err        error
		tokenMaker interface_.Maker
	)

	resp.W = ctx.Value.Writer

	tokenMaker, err = token.NewJWTMaker(consts.JWTSecretKey)
	if err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)
	}

	ctx.Value.Payload, err = tokenMaker.VerifyToken(userToken)
	if err != nil {
		if strings.Contains(err.Error(), consts.ErrInvalidToken) {
			resp.SendResponse(http.StatusForbidden, consts.InvalidTokenRC, consts.GeneralForbiddenMessage, err.Error(), err)

			return
		} else {
			resp.SendResponse(http.StatusForbidden, consts.ExpiredTokenRC, consts.GeneralForbiddenMessage, err.Error(), err)

			return
		}
	}

	ctx.Value.HandleFunc(ctx)
}

func (m *Middleware) Serve() {
	for _, route := range m.RouteCollections {
		m.Mux.HandleFunc(route.path, m.methodCheck)
	}
}
