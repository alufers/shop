package shop

import (
	"errors"
	"fmt"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type categoryController struct {
	ctx *AppCtx
}

func (cc *categoryController) Register(ctx *AppCtx) {
	cc.ctx = ctx
	ctx.db.AutoMigrate(Category{})
	ctx.r.POST("/api/shop/categories", cc.postCategory)
	ctx.r.GET("/api/shop/categories", cc.getCategories)

	categorySubrouter := ctx.RegisterSubrouter("/api/shop/categories/:category_id")
	categorySubrouter.POST("/children", cc.postCategory)
	categorySubrouter.GET("/children", cc.getCategories)
	categorySubrouter.GET("", cc.getCategory)
	categorySubrouter.DELETE("", cc.deleteCategory)
	//ctx.r.POST("/api/shop/categories/:parent_i
}

func (cc *categoryController) postCategory(c *gin.Context) {
	cat := &Category{}
	c.BindJSON(cat)
	cat.ParentID = getUintParam(c, "category_id")

	if validationError := cc.ctx.validate.Struct(cat); validationError != nil {
		c.AbortWithError(400, validationError)
		return
	}
	if err := cc.ctx.db.Save(&cat).Error; err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, cat)
}

func (cc *categoryController) getCategories(c *gin.Context) {
	var categories []Category
	parentID := getUintParam(c, "category_id")
	if err := cc.ctx.db.Find(&categories, "parent_id = ?", parentID).Error; err != nil {
		c.AbortWithError(500, err)
		return
	}

	categoriesMap := make([]map[string]interface{}, 0, len(categories))
	for _, cat := range categories {
		m := structs.Map(cat)
		delete(m, "Parent")
		delete(m, "ParentID")

		var childrenCount uint

		if err := cc.ctx.db.Model(&Category{}).Where("parent_id = ?", cat.ID).Count(&childrenCount).Error; err != nil {
			c.AbortWithError(500, err)
			return
		}
		m["ChildrenCount"] = childrenCount
		categoriesMap = append(categoriesMap, m)
	}

	c.JSON(200, categoriesMap)
}

func (cc *categoryController) getCategory(c *gin.Context) {
	categoryID := getUintParam(c, "category_id")
	cat := &Category{}
	if err := cc.ctx.db.First(cat, categoryID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithError(404, errors.New("Category not found"))
			return
		}
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, cc.populateCategoryMap(c, structs.Map(cat)))
}

func (cc *categoryController) deleteCategory(c *gin.Context) {
	categoryID := getUintParam(c, "category_id")
	cat := &Category{}
	if err := cc.ctx.db.First(cat, categoryID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithError(404, errors.New("Category not found"))
			return
		}
		c.AbortWithError(500, err)
		return
	}
	if err := cc.ctx.db.Delete(cat).Error; err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, map[string]interface{}{})
}

func (cc *categoryController) populateCategoryMap(c *gin.Context, categoryMap map[string]interface{}) map[string]interface{} {
	if categoryMap["ParentID"].(uint) == 0 {
		return categoryMap
	}
	par := &Category{}
	if err := cc.ctx.db.First(par, categoryMap["ParentID"]).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithError(404, fmt.Errorf("broken parent relation, ID = %v", categoryMap["ParentID"]))
			return nil
		}
		c.AbortWithError(500, err)
		return nil
	}
	parentMap := cc.populateCategoryMap(c, structs.Map(par))
	categoryMap["Parent"] = parentMap
	return categoryMap

}
