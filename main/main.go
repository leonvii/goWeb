package main

import (
	"context"
	"goWeb/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)


func main() {
	//mux := http.NewServeMux()
	//mux.Handle("/", &helloHandler{})
	muxRouter := mux.NewRouter()//mux.Router将会定义一个路由列表，其中每一个路由都会定义对应的请求url，及其处理方法。

	router.RegisterRoutes(muxRouter)

	server := &http.Server{
		Addr:    ":8080",
		Handler: muxRouter,
	}

	// 创建系统信号接收器
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}