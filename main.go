package main

import (
	"Initial/router"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	err := router.InitRouter().Run(":8000")
	if err != nil{
		return
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}