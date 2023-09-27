package controllers

import "github.com/gin-gonic/gin"

func Get(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Get",
	})
}

func Page(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Page",
	})
}

func Read(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Read",
	})
}

func Save(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Save",
	})
}

func Update(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Update",
	})
}

func Delete(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Delete",
	})
}
