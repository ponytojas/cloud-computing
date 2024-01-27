package messaging

import (
	"core/internal/logger"
	"core/internal/messaging/routes"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var authUrl string

func init() {
	godotenv.Load()
	log = logger.GetLogger()
	authUrl = os.Getenv("AUTH_SERVICE_URL")
}

func SetupHTTPServer() {
	gin.ForceConsoleColor()
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(cors.Default())
	v1 := router.Group("/v1")
	{
		v1.GET("/health", handleHealthCheck)
		routes.UserRoutes(v1.Group("/user"))
		routes.ProductRoutes(v1.Group("/product"))
		routes.StockRoutes(v1.Group("/stock"))
	}

	port := os.Getenv("HTTP_PORT")

	log.Infof("Core service started on port %s", os.Getenv("HTTP_PORT"))
	router.Run(":" + port)

}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
