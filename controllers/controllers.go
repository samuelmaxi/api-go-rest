package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuelmaxi/api-go-rest/database"
	"github.com/samuelmaxi/api-go-rest/models"
)

func Home(c *gin.Context) {
	fmt.Fprint(c.Writer, "Home page")
}

func AllPersonalities(c *gin.Context) {
	var p []models.Personality
	database.DB.Find(&p)
	c.JSON(http.StatusOK, p)
}

func ReturnPersonalitie(c *gin.Context) {
	vars := c.Param("id")
	var p models.Personality

	id, err := strconv.Atoi(vars)
	if err != nil {
		fmt.Println("error: ", err)
	}
	database.DB.First(&p, id)
	c.JSON(http.StatusOK, p)
}

func CreateNewPersonalitie(c *gin.Context) {
	var newPersonalitie models.Personality
	c.BindJSON(&newPersonalitie)
	database.DB.Create(&newPersonalitie)
	c.JSON(http.StatusOK, newPersonalitie)
}

func DeletePersonalitie(c *gin.Context) {
	vars := c.Param("id")
	var p models.Personality
	id, err := strconv.Atoi(vars)
	if err != nil {
		fmt.Println("error: ", err)
	}
	database.DB.Delete(&p, id)
	c.JSON(http.StatusOK, p)
}

func EditPersonalitie(c *gin.Context) {
	vars := c.Param("id")
	var p models.Personality

	id, err := strconv.Atoi(vars)
	if err != nil {
		fmt.Println("error: ", err)
	}
	database.DB.First(&p, id)
	c.BindJSON(&p)
	database.DB.Save(&p)
	c.JSON(http.StatusOK, &p)
}
