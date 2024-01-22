package messaging

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"store/internal/logger"
	"store/shared"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var dbUrl string

func init() {
	godotenv.Load()
	log = logger.GetLogger()
	dbUrl = os.Getenv("DB_SERVICE_URL")
}
func SetupHTTPServer() {
	gin.ForceConsoleColor()
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/product", handleProductAllGet)
	router.POST("/product", handleProductCreate)
	router.GET("/product/:id", handleProductGet)
	router.DELETE("/product/:id", handleProductRemove)
	router.POST("/stock/:id", handleStockSet)
	router.GET("/stock/:id", handleStockGet)
	router.GET("/stock", handleStockAllGet)
	router.GET("/health", handleHealthCheck)

	port := os.Getenv("HTTP_PORT")

	http.ListenAndServe(":"+port, router)
}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleProductCreate(c *gin.Context) {
	var product shared.Product
	err := c.BindJSON(&product)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestBody, err := json.Marshal(product)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return
	}
	resp, err := http.Post(dbUrl+"/products/create", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Error("Error on product creation request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response shared.ProductCreationResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error("Error parsing response JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Error on product creation request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "product_id": response.Id})
}

func handleProductGet(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get(dbUrl + "/products/" + id)

	if err != nil {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response shared.Product
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

	c.JSON(http.StatusOK, gin.H{"status": "OK", "product": response})
}

func handleProductAllGet(c *gin.Context) {
	resp, err := http.Get(dbUrl + "/products")

	if err != nil {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response []shared.Product
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

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": response})
}

func handleStockSet(c *gin.Context) {
	id := c.Param("id")
	var stock shared.ProductStockRequest
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

	resp, err := http.Post(dbUrl+"/products/"+id+"/stock", "application/json", bytes.NewBuffer(requestBody))

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

func handleProductRemove(c *gin.Context) {
	id := c.Param("id")
	req, err := http.NewRequest(http.MethodDelete, dbUrl+"/products/"+id, nil)
	if err != nil {
		log.Error("Error on product remove request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Error on product remove request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Error on product remove request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleStockGet(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get(dbUrl + "/products/" + id + "/stock/")

	if err != nil {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response shared.ProductStock
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

	c.JSON(http.StatusOK, gin.H{"status": "OK", "stock": response})
}

func handleStockAllGet(c *gin.Context) {
	resp, err := http.Get(dbUrl + "/products/stock")

	if err != nil {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response []shared.ProductStock
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

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": response})
}
