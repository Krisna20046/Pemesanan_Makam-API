package api

import (
	"github.com/Krisna20046/model"
	"github.com/Krisna20046/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PemesananAPI interface {
	AddPemesanan(c *gin.Context)
	UpdatePemesanan(c *gin.Context)
	DeletePemesanan(c *gin.Context)
	GetPemesananByID(c *gin.Context)
	GetPemesananList(c *gin.Context)
}

type pemesananAPI struct {
	pemesananService service.PemesananService
}

func NewPemesananAPI(pemesananService service.PemesananService) *pemesananAPI {
	return &pemesananAPI{pemesananService}
}

func (t *pemesananAPI) AddPemesanan(c *gin.Context) {
	var newPemesanan model.Pemesanan
	if err := c.ShouldBindJSON(&newPemesanan); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.pemesananService.Store(&newPemesanan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add pesan success"})
}

func (t *pemesananAPI) UpdatePemesanan(c *gin.Context) {
	pesanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid pesan ID"})
		return
	}

	var updatePemesanan model.Pemesanan
	if err := c.ShouldBindJSON(&updatePemesanan); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	updatePemesanan.ID = pesanID
	err = t.pemesananService.Update(pesanID, &updatePemesanan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "update pesan success"})
	// TODO: answer here
}

func (t *pemesananAPI) DeletePemesanan(c *gin.Context) {
	pesanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid pesan ID"})
		return
	}
	err = t.pemesananService.Delete(pesanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "delete pesan success"})
	// TODO: answer here
}

func (t *pemesananAPI) GetPemesananByID(c *gin.Context) {
	pesanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid pesan ID"})
		return
	}

	pesan, err := t.pemesananService.GetByID(pesanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, pesan)
}

func (t *pemesananAPI) GetPemesananList(c *gin.Context) {
	pemesanans, err := t.pemesananService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pemesanans)
	// TODO: answer here
}