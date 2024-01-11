package messaging

import (
	"database/internal/database"
	"database/internal/logger"
	"database/shared"
	"database/sql"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	godotenv.Load()
	log = logger.GetLogger()
}

func SetupHttp(db *sql.DB) {
	gin.ForceConsoleColor()
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/database/health", healthCheckHandler)

	router.POST("/users/create", func(c *gin.Context) { createUserHandler(c, db) })
	router.POST("/users/login", func(c *gin.Context) { loginUserHandler(c, db) })

	router.POST("/products/create", func(c *gin.Context) { createProductHandler(c, db) })
	router.GET("/products", func(c *gin.Context) { getAllProductsHandler(c, db) })
	router.GET("/products/:id", func(c *gin.Context) { getProductHandler(c, db) })
	router.POST("/products/:id/stock", func(c *gin.Context) { setProductStockHandler(c, db) })
	router.GET("/products/stock", func(c *gin.Context) { getAllProductStockHandler(c, db) })
	router.GET("/products/:id/stock/", func(c *gin.Context) { getProductStockHandler(c, db) })
	router.DELETE("/products/:id", func(c *gin.Context) { removeProductHandler(c, db) })

	port := os.Getenv("HTTP_PORT")

	http.ListenAndServe(":"+port, router)
}

func healthCheckHandler(c *gin.Context) {
	c.String(200, "OK")
}

func createUserHandler(c *gin.Context, db *sql.DB) {
	var user shared.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error("Error on json user data:", err)
		c.String(http.StatusBadRequest, "ERROR 1000")
		return
	}

	id, err := database.CreateUser(db, user)
	if err != nil {
		log.Error("Error creating user:", err)
		c.String(http.StatusInternalServerError, "ERROR 1001")
		return
	}

	log.Debug("User created with ID:", id)
	c.String(http.StatusOK, "OK")
}

func loginUserHandler(c *gin.Context, db *sql.DB) {
	var user shared.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error("Error on json user data:", err)
		c.String(http.StatusBadRequest, "ERROR 1000")
		return
	}

	usercheck, err := database.LoginUser(db, user)
	if err != nil {
		log.Error("Error on user login:", err)
		c.String(http.StatusInternalServerError, "ERROR 1003")
		return
	}

	log.Debug("Success login:", usercheck)
	c.JSON(http.StatusOK, usercheck)
}

func createProductHandler(c *gin.Context, db *sql.DB) {
	var product shared.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		log.Error("Error on json product data:", err)
		c.String(http.StatusBadRequest, "ERROR 2000")
		return
	}

	id, err := database.CreateProduct(db, product)
	if err != nil {
		log.Error("Error on product create:", err)
		c.String(http.StatusInternalServerError, "ERROR 2001")
		return
	}

	log.Debug("Success product create:", id)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func getAllProductsHandler(c *gin.Context, db *sql.DB) {
	productData, err := database.GetAllProducts(db)
	if err != nil {
		log.Error("Error on product get:", err)
		c.String(http.StatusInternalServerError, "ERROR 2004")
		return
	}

	log.Debug("Success product get:", productData)
	c.JSON(http.StatusOK, productData)
}

func getProductHandler(c *gin.Context, db *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Error casting'%s' to number: %v", c.Param("id"), err)
		c.String(http.StatusBadRequest, "ERROR 2005")
		return
	}

	productData, err := database.GetProduct(db, id)
	if err != nil {
		log.Error("Error on product get:", err)
		c.String(http.StatusInternalServerError, "ERROR 2006")
		return
	}

	log.Debug("Success product get:", productData)
	c.JSON(http.StatusOK, productData)
}

func getProductStockHandler(c *gin.Context, db *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Error casting'%s' to number: %v", c.Param("id"), err)
		c.String(http.StatusBadRequest, "ERROR 2005")
		return
	}

	productStock, err := database.GetProductStock(db, id)
	if err != nil {
		log.Error("Error on productStock get:", err)
		c.String(http.StatusInternalServerError, "ERROR 2006")
		return
	}

	log.Debug("Success productStock get:", productStock)
	c.JSON(http.StatusOK, productStock)
}

func setProductStockHandler(c *gin.Context, db *sql.DB) {
	var productStock shared.ProductStock
	err := c.ShouldBindJSON(&productStock)
	if err != nil {
		log.Error("Error on json productStock data:", err)
		c.String(http.StatusBadRequest, "ERROR 2007")
		return
	}

	err = database.UpsertProductStock(db, productStock.ProductID, productStock.Quantity)
	if err != nil {
		log.Error("Error on productStock set:", err)
		c.String(http.StatusInternalServerError, "ERROR 2008")
		return
	}

	log.Debug("Success productStock set:", productStock)
	c.String(http.StatusOK, "OK")
}

func removeProductHandler(c *gin.Context, db *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Error casting'%s' to number: %v", c.Param("id"), err)
		c.String(http.StatusBadRequest, "ERROR 2009")
		return
	}

	err = database.RemoveProduct(db, id)
	if err != nil {
		log.Error("Error on product delete:", err)
		c.String(http.StatusInternalServerError, "ERROR 2010")
		return
	}

	log.Debug("Success product delete:", id)
	c.String(http.StatusOK, "OK")
}

func getAllProductStockHandler(c *gin.Context, db *sql.DB) {
	productStocks, err := database.GetAllProductStock(db)
	if err != nil {
		log.Error("Error on productStock get:", err)
		c.String(http.StatusInternalServerError, "ERROR 2006")
		return
	}

	log.Debug("Success productStock get:", productStocks)
	c.JSON(http.StatusOK, productStocks)
}
