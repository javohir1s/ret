package handler

import (
	"essy_travel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


// CreateAirport godoc
// @Summary Create Airport
// @Description Create Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param object body models.CreateAirport true "CreateAirportRequestBody"
// @Success 201 {object} Response{data=models.Airport} "AirportBody"
// @Router /airport [post]
func (h *Handler) CreateAirport(c *gin.Context) {
	var airport models.CreateAirport
	err := c.ShouldBind(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while binding data: "+err.Error())
		return
	}
	_, err = h.strg.Airport().Create(airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, "Airport created")
}


// AirportGetById godoc
// @Summary Get Airport by ID
// @Description Get Airport by ID
// @Tags Airport
// @Accept json
// @Produce json
// @Param id path string true "Airport ID"
// @Success 200 {object} Response{data=models.Airport} "AirportBody"
// @Router /airport/{id} [get]
func (h *Handler) AirportGetById(c *gin.Context) {
	var airport models.AirportPrimaryKey
	err := c.ShouldBind(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while binding data: "+err.Error())
		return
	}
	resp, err := h.strg.Airport().GetById(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// AirportGetList godoc
// @Summary Get List of Airports
// @Description Get List of Airports
// @Tags Airport
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} Response{data=models.GetListAirportResponse} "GetListAirportResponseBody"
// @Router /airport [get]
func (h *Handler) AirportGetList(c *gin.Context) {
	var airport models.GetListAirportRequest
	err := c.ShouldBindQuery(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while binding data: "+err.Error())
		return
	}
	resp, err := h.strg.Airport().GetList(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}


// AirportUpdate godoc
// @Summary Update Airport
// @Description Update Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param object body models.UpdateAirport true "UpdateAirportRequestBody"
// @Success 202 {string} string "Updated"
// @Router /airport [put]
func (h *Handler) AirportUpdate(c *gin.Context) {
	var airport = models.UpdateAirport{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	_, err = h.strg.Airport().Update(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Updated")
}


// AirportDelete godoc
// @Summary Delete Airport
// @Description Delete Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param object body models.AirportPrimaryKey true "AirportPrimaryKeyRequestBody"
// @Success 202 {string} string "Deleted"
// @Router /airport [delete]
func (h *Handler) AirportDelete(c *gin.Context) {
	var airport = models.AirportPrimaryKey{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	_, err = h.strg.Airport().Delete(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted:")
}



// AirportSearch godoc
// @Summary Search for Airports
// @Description Search for Airports
// @Tags Airport
// @Accept json
// @Produce json
// @Param object body models.Airport true "AirportRequestBody"
// @Success 202 {object} Response{data=models.GetListAirportResponse} "GetListAirportResponseBody"
// @Router /airport/search [post]
func (h *Handler) AirportSearch(c *gin.Context) {
	var airport = models.Airport{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.Airport().Search(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not find: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}
