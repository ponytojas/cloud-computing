package messaging

import (
	"cart/internal/logger"
	sharedcart "cart/internal/sharedCart"
	"cart/internal/store"
	"cart/internal/token"
	"cart/shared"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	v1 := router.Group("/v1")
	{
		v1.GET("/health", handleHealthCheck)
		v1.POST("/add-to-cart", token.CheckTokenMiddleware, store.CheckExistingProduct, handleAddToCart)
		v1.GET("/cart/:userId", token.CheckTokenMiddleware, handleGetCart)
		v1.DELETE("/cart/:userId", token.CheckTokenMiddleware, hanldeClearCart)
	}

	port := os.Getenv("HTTP_PORT")

	log.Infof("Core service started on port %s", os.Getenv("HTTP_PORT"))
	router.Run(":" + port)

}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleAddToCart(c *gin.Context) {
	var toAdd shared.AddToCartRequest
	if err := c.ShouldBindBodyWith(&toAdd, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := sharedcart.AddToUserCart(
		strconv.Itoa(toAdd.UserId),
		strconv.Itoa(toAdd.ProductId),
		int64(toAdd.Quantity),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleGetCart(c *gin.Context) {
	userId := c.Param("userId")
	cart, err := sharedcart.GetUserCart(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func hanldeClearCart(c *gin.Context) {
	userId := c.Param("userId")
	err := sharedcart.ClearCart(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
