package api

import (
	"github.com/Krisna20046/model"
	"github.com/Krisna20046/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JenazahAPI interface {
	AddJenazah(c *gin.Context)
	UpdateJenazah(c *gin.Context)
	DeleteJenazah(c *gin.Context)
	GetJenazahByID(c *gin.Context)
	GetJenazahList(c *gin.Context)
}

type jenazahAPI struct {
	jenazahService service.JenazahService
}

func NewJenazahAPI(jenazahRepo service.JenazahService) *jenazahAPI {
	return &jenazahAPI{jenazahRepo}
}

func (t *jenazahAPI) AddJenazah(c *gin.Context) {
	var newJenazah model.DataJenazah
	if err := c.ShouldBindJSON(&newJenazah); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.jenazahService.Store(&newJenazah)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add jenazah success"})
}

func (t *jenazahAPI) UpdateJenazah(c *gin.Context) {
	jenazahID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid jenazah ID"})
		return
	}

	var updateJenazah model.DataJenazah
	if err := c.ShouldBindJSON(&updateJenazah); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	updateJenazah.ID = jenazahID
	err = t.jenazahService.Update(jenazahID, &updateJenazah)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "update jenazah success"})
	// TODO: answer here
}

func (t *jenazahAPI) DeleteJenazah(c *gin.Context) {
	jenazahID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid jenazah ID"})
		return
	}
	err = t.jenazahService.Delete(jenazahID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "delete jenazah success"})
	// TODO: answer here
}

func (t *jenazahAPI) GetJenazahByID(c *gin.Context) {
	jenazahID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid jenazah ID"})
		return
	}

	jenazah, err := t.jenazahService.GetByID(jenazahID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, jenazah)
}

func (t *jenazahAPI) GetJenazahList(c *gin.Context) {
	jenazahs, err := t.jenazahService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jenazahs)
	// TODO: answer here
}