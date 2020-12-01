package handler

import (
	"gorestapimysql/src/config"
	"gorestapimysql/src/entities"
	"gorestapimysql/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//TransactionPost posts transactions to the db
func TransactionPost() gin.HandlerFunc {
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
			transaction := entities.Transaction{
				CustomerRef:       "testref2",
				AgentCode:         "UMEME",
				AgentID:           "UMEME2341",
				Amount:            101.5,
				PaymentDate:       time.Now(),
				CompletedAtVendor: true,
				SentToUtility:     false,
				UtilityReference:  "",
				RecordDate:        time.Now(),
			}
			insertedID, err := transDb.SaveTransaction(transaction)
			if err != nil {
				respError := entities.Response{
					StatusCode:        100,
					StatusDescription: err.Error(),
				}
				context.JSON(http.StatusBadRequest, respError)
			} else {
				response := entities.PostResponse{
					Response: entities.Response{
						StatusCode:        0,
						StatusDescription: "SUCCESS",
					},
					InsertedID: insertedID,
				}
				context.JSON(http.StatusOK, response)
			}
		}
	}
}
