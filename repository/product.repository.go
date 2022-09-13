package repository

import (
	"net/http"

	"github.com/Sepctrevuln-Sketch/golang-restapi-mux/models"
	"github.com/Sepctrevuln-Sketch/golang-restapi-mux/utils"
)

var ResponseError = utils.ResponseError
var ResponseJson = utils.ResponseJson

func Delete(w http.ResponseWriter, payload interface{}, exact string, successMsg string) map[string]string {
	if models.DB.Delete(payload, exact).RowsAffected == 0 {
		ResponseError(w, http.StatusInternalServerError, "Cannot delete product because notfound")
		return
	}
	responseMessage := map[string]string{
		"message": "test product deleted",
	}

	ResponseJson(w, http.StatusOK, responseMessage)
}
