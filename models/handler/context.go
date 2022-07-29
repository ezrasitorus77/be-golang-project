package handler

import (
	"be-golang-project/consts"
	"be-golang-project/models/validation_"
	"be-golang-project/repository/orm"
	"context"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

var (
	DB *orm.Database
)

type (
	Context struct {
		ChildCtx context.Context
		Value    ContextValue
	}

	ContextValue struct {
		Request    *http.Request
		Writer     http.ResponseWriter
		Payload    interface{}
		HandleFunc contextFunc
		Error      error
		Key        string
		JWTKey     string
	}

	contextFunc func(*Context)

	ParentContext struct {
		C context.Context
		// Pool *pool.BasePool
		Salt []byte
		DB   *gorm.DB
	}
)

func (parentCtx *ParentContext) Set(ctx *Context, w http.ResponseWriter, r *http.Request, hFunc func(*Context)) *Context {
	ctx.Value.Writer = w
	ctx.Value.Request = r
	ctx.Value.HandleFunc = hFunc
	ctx.ChildCtx = context.WithValue(parentCtx.C, "DB", parentCtx.DB)

	return ctx
}

func (ctx *Context) ParseRequest(model interface{}) {
	var decoder *json.Decoder = json.NewDecoder(ctx.Value.Request.Body)

	err := decoder.Decode(&model)
	if err != nil {
		ctx.Value.Error = err
	}

	return
}

func (parentCtx ParentContext) GetDBSession() *gorm.DB {
	return orm.CreateDBSession()
}

func (parentCtx ParentContext) Init() (*ParentContext, error) {
	parentCtx.C = context.Background()
	// parentCtx.Pool = pool.SyncPool.Get().(*pool.BasePool)
	parentCtx.DB = parentCtx.GetDBSession()

	salt, err := validation_.GenerateRandomSalt(consts.SaltSize)
	if err != nil {
		return nil, err
	}
	parentCtx.Salt = salt

	return &parentCtx, nil
}

// func (c *Context) WithValue(val interface{}) context.Context {
// 	var newCtx context.Context

// 	newCtx = context.WithValue(c.C, consts.ContextKey, val)

// 	return newCtx
// }
