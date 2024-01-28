package routes

import (
	"bytes"
	"core/shared"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(g *gin.RouterGroup) {
	g.POST("/login", handleLogin)
	g.POST("/logout", handleLogout)
	g.POST("/register", handleRegister)
}

func handleLogin(c *gin.Context) {
	var user shared.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return
	}

	resp, err := http.Post(authUrl+"/login", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error("Error sending request:", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Error closing response body:", err)
		}
	}(resp.Body)

	var token shared.TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		log.Error("Error decoding response body:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token.Token, "user": token.User, "status": token.Status})
}

func handleLogout(c *gin.Context) {
	var token shared.TokenLogout
	err := c.ShouldBindJSON(&token)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestBody, err := json.Marshal(token)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return
	}

	resp, err := http.Post(authUrl+"/logout", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error("Error sending request:", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Error closing response body:", err)
		}
	}(resp.Body)

	var response shared.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error("Error decoding response body:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": response.Status})
}

func handleRegister(c *gin.Context) {
	var user shared.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return
	}

	resp, err := http.Post(authUrl+"/register", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error("Error sending request:", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Error closing response body:", err)
		}
	}(resp.Body)

	var response shared.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error("Error decoding response body:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": response.Status})
}
