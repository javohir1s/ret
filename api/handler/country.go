package handler

import (
	"essy_travel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCountry godoc
// @Summary Create Country
// @Description Create Country
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.CreateCountry true "CreateCountryRequestBody"
// @Success 201 {object} Response{data=models.Country} "CountryBody"
// @Failure 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
// @Router /country [post]
func (h *Handler) CreateCountry(c *gin.Context) {
	var country = models.CreateCountry{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		c.JSON(400, "ShouldBindJSON err:"+err.Error())
		return
	}
	resp, err := h.strg.Country().Create(country)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}


// CountryGetById godoc
// @Summary Get Country by ID
// @Description Get Country by ID
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "Country ID"
// @Success 200 {object} Response{data=models.Country} "CountryBody"
// @Failure 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
// @Router /country/{id} [get]
func (h *Handler) CountryGetById(c *gin.Context) {
	var country = models.CountryPrimaryKey{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}
	resp, err := h.strg.Country().GetById(country)
	if err != nil {
		handleResponse(c, 500, "Country does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}


// CountryGetList godoc
// @Summary Get List of Countries
// @Description Get List of Countries
// @Tags Country
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} Response{data=models.GetListCountryResponse} "GetListCountryResponseBody"
// @Router /country [get]
func (h *Handler) CountryGetList(c *gin.Context) {
	var country models.GetListCountryRequest
	err := c.ShouldBindQuery(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while binding data: "+err.Error())
		return
	}
	resp, err := h.strg.Country().GetList(country)
	if err != nil {
		handleResponse(c, 500, "Country does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}


// CountryUpdate godoc
// @Summary Update Country
// @Description Update Country
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.UpdateCountry true "UpdateCountryRequestBody"
// @Success 202 {string} string "Updated"
// @Failure 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
// @Router /country [put]
func (h *Handler) CountryUpdate(c *gin.Context) {
	var country = models.UpdateCountry{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}
	_, err = h.strg.Country().Update(country)
	if err != nil {
		handleResponse(c, 500, "Country does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Updated:")
}


// CountryDelete godoc
// @Summary Delete Country
// @Description Delete Country
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.CountryPrimaryKey true "CountryPrimaryKeyRequestBody"
// @Success 202 {string} string "Deleted"
// @Failure 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
// @Router /country [delete]
func (h *Handler) CountryDelete(c *gin.Context) {
	var country = models.CountryPrimaryKey{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}
	_, err = h.strg.Country().Delete(country)
	if err != nil {
		handleResponse(c, 500, "Country does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted")
}



// CountrySearch godoc
// @Summary Search for Countries
// @Description Search for Countries
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.Country true "CountryRequestBody"
// @Success 202 {object} Response{data=models.GetListCountryResponse} "GetListCountryResponseBody"
// @Failure 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
// @Router /country/search [post]
func (h *Handler) CountrySearch(c *gin.Context) {
	var country = models.Country{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.Country().Search(country)
	if err != nil {
		handleResponse(c, 500, "Country does not find: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}
