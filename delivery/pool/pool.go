package pool

import (
	"sync"

	"gorm.io/gorm"
)

type (
	BasePool struct {
		Err         error
		DB          *gorm.DB
		HttpMessage string
		HttpStatus  int
	}
)

var (
	SyncPool sync.Pool = sync.Pool{
		New: func() interface{} {
			return &BasePool{
				Err:         nil,
				DB:          nil,
				HttpMessage: "",
				HttpStatus:  0,
			}
		},
	}
)
