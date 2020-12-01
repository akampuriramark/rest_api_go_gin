package handler

import (
	"gorestapimysql/src/config"
	"gorestapimysql/src/entities"
	"gorestapimysql/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TransactionGet RETURNS
func TransactionGet() gin.HandlerFunc {
	return func(context *gin.Context) {
		db, errr := config.GetDB()
		if errr != nil {
			respError := entities.Response{
				StatusCode:        100,
				StatusDescription: errr.Error(),
			}
			context.JSON(http.StatusBadRequest, respError)
		} else {
			transDb := models.DbHandler{
				Db: db,
			}
			results, erra := transDb.GetTransactions()
			if erra != nil {
				respError := entities.Response{
					StatusCode:        100,
					StatusDescription: erra.Error(),
				}
				context.JSON(http.StatusBadRequest, respError)
			} else {
				context.JSON(http.StatusOK, results)
			}
		}
	}

}
