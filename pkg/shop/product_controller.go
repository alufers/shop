package shop

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type productController struct {
	ctx *AppCtx
}

func (pc *productController) Register(ctx *AppCtx) {
	pc.ctx = ctx
	ctx.db.AutoMigrate(Product{})
	ctx.r.POST("/api/shop/products", pc.postProduct)
	ctx.r.GET("/api/shop/products", pc.getProducts)
	ctx.GetSubrouter("/api/shop/categories/:category_id").GET("/products", pc.getProducts)
}

func (pc *productController) postProduct(c *gin.Context) {
	prod := &Product{}
	c.BindJSON(prod)
	if validationError := pc.ctx.validate.Struct(prod); validationError != nil {
		c.AbortWithError(400, validationError)
		return
	}
	cat := &Category{}
	if err := pc.ctx.db.Save(&prod).Related(cat).Error; err != nil {
		c.AbortWithError(500, err)
		return
	}
	prod.Category = *cat
	c.JSON(200, prod)
}

func (pc *productController) getProducts(c *gin.Context) {
	var products []Product
	db := pc.ctx.db
	if categoryID := getUintParam(c, "category_id"); categoryID != 0 {
		db = db.Where("category_id = ?", categoryID)
	}
	if err := db.Find(&products).Error; err != nil {
		c.AbortWithError(500, err)
		return
	}
	productsMap := make([]map[string]interface{}, 0, len(products))
	for _, prod := range products {
		m := structs.Map(prod)
		delete(m, "Category")
		productsMap = append(productsMap, m)
	}
	c.JSON(200, productsMap)
}
