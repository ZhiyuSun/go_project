package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func pong(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"message": "pong",
//	})
//}

//func pong(c *gin.Context) {
//	c.JSON(http.StatusOK, map[string]string{
//		"message": "pong",
//	})
//}

func pong(c *gin.Context) {
	//panic("异常")
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}

func main() {
	//实例化一个gin的server对象
	r := gin.Default()
	r.GET("/ping", pong)
	r.Run(":8083") // listen and serve on 0.0.0.0:8080
}
