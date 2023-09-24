package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/repository"
)

func (c Controllers) GetProduct() {
	c.Group.GET("/product/:id", getProduct)
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetProduct(c.Param("id")))
}

func (c Controllers) PurchaseProduct() {
	c.Group.POST("/product/:id/purchase", purchaseProduct)
}

func purchaseProduct(c *gin.Context) {
	s, err := repository.PurchaseProduct()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, pkg.Error(err))
		return
	}

	c.Redirect(http.StatusSeeOther, s.URL)
}
