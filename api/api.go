package api

import (
	"essy_travel/api/handler"
	"essy_travel/config"
	"essy_travel/storage"

	"github.com/gin-gonic/gin"

	_ "essy_travel/api/docs"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {

	handler := handler.NewHandler(cfg, strg)

	// City ...
	r.POST("/city", handler.CreateCity)
	r.GET("/city/:id", handler.CityGetById)
	r.GET("/city", handler.CityGetList)
	r.PUT("/city/:id", handler.CountryUpdate)
	r.DELETE("/city/:id", handler.CityDelete)

	// Country
	r.POST("/country", handler.CreateCountry)
	r.GET("/country/:id", handler.CountryGetById)
	r.GET("/country", handler.CountryGetList)
	r.PUT("/country/:id", handler.CountryUpdate)
	r.DELETE("/country/:id", handler.CountryDelete)

	// Airport
	r.POST("/airport", handler.CreateAirport)
	r.GET("/airport/:id", handler.AirportGetById)
	r.GET("/airport", handler.AirportGetList)
	r.PUT("/airport/:id", handler.AirportUpdate)
	r.DELETE("/airport/:id", handler.AirportDelete)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
