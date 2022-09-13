package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sepctrevuln-Sketch/golang-restapi-mux/models"
	"github.com/Sepctrevuln-Sketch/golang-restapi-mux/repository"
	"github.com/Sepctrevuln-Sketch/golang-restapi-mux/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = utils.ResponseJson
var ResponseError = utils.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var product []models.Product
	if err := models.DB.Find(&product).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, product)
}
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var product models.Product
	if err := models.DB.Where("id =?", id).First(&product).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "id not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, product)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	if err := models.DB.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, product)
}
func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	if models.DB.Where("id=?", id).Updates(&product).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Cannot update product")
		return
	}

	product.Id = id

	ResponseJson(w, http.StatusOK, product)

}
func Delete(w http.ResponseWriter, r *http.Request) {
	id := map[string]string{
		"id": "",
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&id); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	var product models.Product

	repository.Delete(w, &product, id["id"], "Berhasil Delete Product")
	// if models.DB.Delete(&product, id["id"]).RowsAffected == 0 {
	// 	ResponseError(w, http.StatusInternalServerError, "Cannot delete product because notfound")
	// 	return
	// }

	// responseMessage := map[string]string{
	// 	"message": fmt.Sprintf("%s deleted", product.Name),
	// }
	// ResponseJson(w, http.StatusOK, responseMessage)

}
