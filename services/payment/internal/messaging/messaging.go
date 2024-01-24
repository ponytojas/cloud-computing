package messaging

import (
	"encoding/json"
	"net/http"
	"os"
	"payment/internal/cart"
	"payment/internal/logger"
	"payment/internal/token"
	"payment/shared"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var authUrl string
var storeUrl string

func init() {
	godotenv.Load()
	log = logger.GetLogger()
	authUrl = os.Getenv("AUTH_SERVICE_URL")
	storeUrl = os.Getenv("STORE_SERVICE_URL")
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
		v1.POST("/pay/:userId", token.CheckTokenMiddleware, handlePay)
		v1.GET("/total/:userId", token.CheckTokenMiddleware, handleTotal)
	}

	port := os.Getenv("HTTP_PORT")

	log.Infof("Core service started on port %s", os.Getenv("HTTP_PORT"))
	router.Run(":" + port)

}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handlePay(c *gin.Context) {
	userId := c.Param("userId")
	err := cart.ClearUserCart(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleTotal(c *gin.Context) {
	userId := c.Param("userId")

	cartItems, err := cart.GetUserCart(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var totalPrice float64
	totalPrice = 0

	for productId, quantity := range cartItems {
		resp, err := http.Get(storeUrl + "/product/" + productId)
		if err != nil {
			log.Error("Error on product get request:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var response shared.ProductResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Error("Error parsing response JSON:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		parsedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Error("Error parsing quantity:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		totalPrice += response.Product.Pricing * float64(parsedQuantity)
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "totalPrice": totalPrice})
}
