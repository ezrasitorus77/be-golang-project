package client

import (
	"be-golang-project/models/handler"

	"gorm.io/gorm"
)

func (parentCtx *Client) Register(context_ *handler.Context) {
	var (
		// client db.Client
		tx *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
	)

	defer func() {
		tx.Rollback()
	}()

	resp.W = context_.Value.Writer

}
