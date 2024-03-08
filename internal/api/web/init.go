package web

import (
	"log/slog"

	"parkwise/config"
	"parkwise/internal/domain/definition"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(router *gin.Engine) {
	slog.Warn("Enabling CORS")
	router.Use(cors.New(cors.Config{
		// Only allow your specific domains instead of all
		AllowOrigins: []string{"http://localhost"},

		// Specify the actual methods your frontend uses
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},

		// Specify the headers your frontend sends
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
		},

		AllowCredentials: true,

		// The rest of the properties should be disabled for security reasons unless you specifically need them
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))
}

type repos struct {
	config config.ServerConfig
	ds     definition.DataStore
}

func NewAPIService(cfg config.ServerConfig, ds definition.DataStore) *repos {
	return &repos{
		config: cfg,
		ds:     ds,
	}
}

func (r *repos) InstallRoutes(router *gin.Engine) {
	CORS(router)
	router.GET("", health)
	router.GET("/health", health)

	v1 := router.Group("/v1")

	parkingLots := v1.Group("/parking-lots")
	{
		parkingLots.GET("/:id", r.getStatus)
		parkingLots.POST("", r.addParkingLot)
	}

	slots := v1.Group("/slots")
	{
		slots.POST("/maintenance", r.maintenance)
	}

	reports := v1.Group("/reports")
	{
		reports.GET("/daily", r.getDailyReports)
	}

	park := v1.Group("/park")
	{
		park.POST("", r.park)
	}

	unpark := v1.Group("/unpark")
	{
		unpark.POST("", r.unpark)
	}
}
