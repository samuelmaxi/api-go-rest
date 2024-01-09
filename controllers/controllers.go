package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuelmaxi/api-go-rest/models"
)

func Home(c *gin.Context) {
	fmt.Fprint(c.Writer, "Home page")
}

func AllPersonalities(c *gin.Context) {
	c.JSON(http.StatusOK, models.Personalities)
}

func FirstPersanalitie(c *gin.Context) {
	c.JSON(http.StatusOK, models.Personalities)
}

func SecondPersonalitie(c *gin.Context) {
	c.JSON(http.StatusOK, models.Personalities[1])
}

func ReturnPersonalitie(c *gin.Context) {
	vars := c.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		fmt.Println("error: ", err)
	}

	for _, personalitie := range models.Personalities {
		if personalitie.Id == id {
			c.JSON(http.StatusOK, personalitie)
		}
	}
	fmt.Printf("Valor do id: %d", id)
}
