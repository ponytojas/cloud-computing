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

func ProductRoutes(g *gin.RouterGroup) {
	g.GET("/", handleAllProductGet)
	g.POST("/", token.CheckTokenMiddleware, handleProductPost)
	g.GET("/:id", handleProductGet)
	g.DELETE("/:id", token.CheckTokenMiddleware, handleProductDelete)
}

func handleAllProductGet(c *gin.Context) {
	resp, err := http.Get(storeUrl + "/product")

	if err != nil {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response shared.ProductsResponse
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

func handleProductGet(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get(storeUrl + "/product/" + id)

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

	if resp.StatusCode != http.StatusOK {
		log.Error("Error on product get request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "OK", "Product": response.Product})
}

func handleProductPost(c *gin.Context) {
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
	resp, err := http.Post(storeUrl+"/product", "application/json", bytes.NewBuffer(requestBody))

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

func handleProductDelete(c *gin.Context) {
	id := c.Param("id")
	req, err := http.NewRequest(http.MethodDelete, storeUrl+"/product/"+id, nil)
	if err != nil {
		log.Error("Error creating request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Error on product delete request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Error on product delete request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
