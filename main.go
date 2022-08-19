package main

import "github.com/gin-gonic/gin"

func GetAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Rodrigo Santos",
	})
}

func main() {
	r := gin.Default()
	r.GET("/student", GetAllStudents)
	r.Run()
}
