package handler

import (
	"gorestapimysql/src/config"
	"gorestapimysql/src/entities"
	"gorestapimysql/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TransactionGet RETURNS
func TransactionGet(context *gin.Context) {
	// results := []entities.Transaction{
	// 	{
	// 		Idnumber:   12,
	// 		Amount:     100.45,
	// 		Utility:    "UMEME",
	// 		Vendorcode: "MTN",
	// 		// Affected: map[string]string{
	// 		// 	"Balance": "ACCL",
	// 		// 	"Ledgers": "GL",
	// 		// },
	// 	},
	// 	{
	// 		Idnumber:   13,
	// 		Amount:     101.45,
	// 		Utility:    "UMEME",
	// 		Vendorcode: "UTL",
	// 		// Affected: map[string]string{
	// 		// 	"Balance": "ACCL",
	// 		// 	"Ledgers": "GL",
	// 		// },
	// 	},
	// }
	db, errr := config.GetDB()
	if errr != nil {
		respError := entities.Error{
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
			respError := entities.Error{
				StatusCode:        100,
				StatusDescription: erra.Error(),
			}
			context.JSON(http.StatusBadRequest, respError)
		} else {
			context.JSON(http.StatusOK, results)
		}
	}
}
