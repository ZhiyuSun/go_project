package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", welcome) // http://localhost:8083/welcome?firstname=geliang&lastname=zhu
	router.POST("/form_post", formPost)
	router.POST("/post", getPost) // http://127.0.0.1:8083/post?id=1&page=2 加body

	router.Run(":8083")
}

func getPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.DefaultPostForm("message", "信息")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}

func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "zhiyu") // 取get参数
	lastName := c.DefaultQuery("lastname", "sun")
	c.JSON(http.StatusOK, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})
}
