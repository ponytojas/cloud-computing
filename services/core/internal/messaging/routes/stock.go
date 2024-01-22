package routes

import (
	"bytes"
	"core/internal/logger"
	"core/internal/token"
	"core/shared"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	log = logger.GetLogger()
	storeUrl = os.Getenv("STORE_SERVICE_URL")
}

func StockRoutes(g *gin.RouterGroup) {
	g.GET("/", handleAllStockGet)
	g.GET("/:id", handleStockGet)
	g.POST("/:id", token.CheckTokenMiddleware, handleStockAdd)
}

func handleAllStockGet(c *gin.Context) {
	resp, err := http.Get(storeUrl + "/stock")

	if err != nil {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response shared.ProductStockResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error("Error parsing response JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func handleStockGet(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get(storeUrl + "/stock/" + id)

	if err != nil {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response shared.ProductStockResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error("Error parsing response JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func handleStockAdd(c *gin.Context) {
	id := c.Param("id")
	var stock shared.ProductStock
	err := c.BindJSON(&stock)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestBody, err := json.Marshal(stock)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return
	}

	resp, err := http.Post(storeUrl+"/stock/"+id, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Error("Error on product creation request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Error on product creation request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
