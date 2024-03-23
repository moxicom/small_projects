package main

import (
	_ "api_first/docs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListAccounts godoc
//
//	@Summary		List accounts
//	@Description	get accounts
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	false	"name search by q"	Format(email)
//	@Success		200	{array}		models.User
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/users [get]
func (h *Handler) GetUsers(ctx *gin.Context) {
	fmt.Println("asd")
	ctx.Status(http.StatusOK)
}
