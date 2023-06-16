package api

import (
	"github.com/Krisna20046/model"
	"github.com/Krisna20046/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MakamAPI interface {
	AddMakam(c *gin.Context)
	UpdateMakam(c *gin.Context)
	DeleteMakam(c *gin.Context)
	GetMakamByID(c *gin.Context)
	GetMakamList(c *gin.Context)
}

type makamAPI struct {
	makamService service.MakamService
}

func NewMakamAPI(makamRepo service.MakamService) *makamAPI {
	return &makamAPI{makamRepo}
}

func (t *makamAPI) AddMakam(c *gin.Context) {
	var newMakam model.DataMakam
	if err := c.ShouldBindJSON(&newMakam); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.makamService.Store(&newMakam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add makam success"})
}

func (t *makamAPI) UpdateMakam(c *gin.Context) {
	makamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid makam ID"})
		return
	}

	var updateMakam model.DataMakam
	if err := c.ShouldBindJSON(&updateMakam); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	updateMakam.ID = makamID
	err = t.makamService.Update(makamID, &updateMakam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "update makam success"})
	// TODO: answer here
}

func (t *makamAPI) DeleteMakam(c *gin.Context) {
	makamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid makam ID"})
		return
	}
	err = t.makamService.Delete(makamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "delete makam success"})
	// TODO: answer here
}

func (t *makamAPI) GetMakamByID(c *gin.Context) {
	makamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid makam ID"})
		return
	}

	makam, err := t.makamService.GetByID(makamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, makam)
}

func (t *makamAPI) GetMakamList(c *gin.Context) {
	makams, err := t.makamService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, makams)
	// TODO: answer here
}