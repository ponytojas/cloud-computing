package store

import (
	"cart/internal/logger"
	"cart/shared"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var storeUrl string

func init() {
	godotenv.Load()
	log = logger.GetLogger()
	storeUrl = os.Getenv("STORE_SERVICE_URL")
}

func CheckExistingProduct(c *gin.Context) {
	var toAdd shared.AddToCartRequest
	var productId int

	if err := c.ShouldBindBodyWith(&toAdd, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	productId = toAdd.ProductId

	var id = strconv.Itoa(productId)
	res, err := http.Get(storeUrl + "/stock/" + id)
	if err != nil {
		log.Error("Error sending request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if res.StatusCode == http.StatusNotFound || res.StatusCode == http.StatusBadRequest || res.StatusCode == http.StatusInternalServerError {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		c.Abort()
		return
	}

	var response shared.ProductStockSingleGetResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Error("Error parsing response JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	if response.Stock.Quantity <= 0 || response.Stock.Quantity-int(toAdd.Quantity) < 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not enough in stock"})
		c.Abort()
		return
	}

	c.Next()
}
