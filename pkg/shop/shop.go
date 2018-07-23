package shop

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

func setupCORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowMethods("DELETE", "PATCH")
	return cors.New(config)
}

//InitAppCtx prepares AppCtx with passed config
func InitAppCtx(cfg *AppConfig) *AppCtx {
	db, err := gorm.Open(cfg.DBDialect, cfg.DBArgs...)
	if err != nil {
		log.Panicf("Failed to init databse %v", err)
	}
	log.Info("DB ok")
	r := gin.Default()
	//r.RedirectTrailingSlash = true
	r.Use(errorPrinterMiddleware())
	r.Use(setupCORS())
	validate := validator.New()
	ctx := &AppCtx{
		validate: validate,
		r:        r,
		db:       db,
	}
	(&categoryController{}).Register(ctx)
	(&productController{}).Register(ctx)
	return ctx
}

// Run forest run
func Run() {
	ctx := InitAppCtx(&AppConfig{
		DBDialect: "sqlite3",
		DBArgs:    []interface{}{"development.db"},
	})
	ctx.r.Run()
}
