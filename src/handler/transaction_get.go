package handler

import (
	"fmt"
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
	results, err := models.DBH.GetTransactions()
	if err != nil {
		fmt.Println(err)
	}
	context.JSON(http.StatusOK, results)
}
