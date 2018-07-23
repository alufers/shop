package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"
)

//AppCtx holds the data needed by all controllers
type AppCtx struct {
	db       *gorm.DB
	r        *gin.Engine
	validate *validator.Validate
	groups   map[string]*gin.RouterGroup
}

func (ctx *AppCtx) RegisterSubrouter(path string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	if ctx.groups == nil {
		ctx.groups = make(map[string]*gin.RouterGroup)
	}
	ctx.groups[path] = ctx.r.Group(path, handlers...)
	return ctx.groups[path]
}

func (ctx *AppCtx) GetSubrouter(path string) *gin.RouterGroup {
	return ctx.groups[path]
}
