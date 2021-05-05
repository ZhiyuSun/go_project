package main

import (
	//"context"
	"fmt"
	"github.com/gin-gonic/gin"
	//"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 优雅退出，当我们关闭程序的时候应该做的后续处理
	// 微服务，启动之前或者启动之后会做一件事，将当前的服务的ip和端口号注册到注册中心
	// 我们当前的服务停止了以后并没有告诉注册中心
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	//router.Run(":8083")

	go func() {
		router.Run(":8083")
	}()

	// 如果想要接受到信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	// 处理后续的逻辑
	fmt.Println("关闭server中。。。")
	fmt.Println("注销服务。。。")

	//srv := &http.Server{
	//	Addr:    ":8080",
	//	Handler: router,
	//}
	//
	//go func() {
	//	// service connections
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()
	//
	//// Wait for interrupt signal to gracefully shutdown the server with
	//// a timeout of 5 seconds.
	//quit := make(chan os.Signal)
	//// kill (no param) default send syscanll.SIGTERM
	//// kill -2 is syscall.SIGINT
	//// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//log.Println("Shutdown Server ...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//// catching ctx.Done(). timeout of 5 seconds.
	//select {
	//case <-ctx.Done():
	//	log.Println("timeout of 5 seconds.")
	//}
	//log.Println("Server exiting")
}