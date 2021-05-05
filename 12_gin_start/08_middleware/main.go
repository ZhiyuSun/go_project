package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123456")
		//让原本改执行的逻辑继续执行
		c.Next()

		end := time.Since(t)
		fmt.Printf("耗时:%v\n", end)
		status := c.Writer.Status()
		fmt.Println("状态", status)
	}
}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			} else {
				fmt.Println(k, v)
			}
		}

		if token != "sunzhiyu" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			//return // return没有起到作用
			//挑战 为什么连return都阻止不了后续逻辑的执行
			c.Abort()
		}

		// auth,ping是gin的队列，会被驱动这一步一步往后走，直到完成，即使return了，后面的也会执行
		c.Next()
	}
}

func main() {
	//router := gin.New()
	router := gin.Default()
	//使用logger和recovery中间件 全局所有
	//router.Use(MyLogger())
	router.Use(TokenRequired())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8083")
}
