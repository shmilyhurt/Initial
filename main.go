package main

import (
	"Initial/conf"
	"Initial/public"
	"Initial/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf.InitDb()
	//task.InitTask()
	public.InitValidator()
	err := router.InitRouter().Run(":8880")
	if err != nil {
		return
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
