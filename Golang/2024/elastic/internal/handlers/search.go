package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	searchStr := c.Param("search")

	products, err := h.service.SearchProduct(context.TODO(), searchStr)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	
	c.JSON(http.StatusOK, products)
}