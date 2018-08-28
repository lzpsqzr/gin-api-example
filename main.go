package main

import (
	"fmt"
	"github.com/danielgjtan/gin-test/pkg/setting"
	"github.com/danielgjtan/gin-test/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var DB = make(map[string]string)

type error interface {
	Error() string
}

func main() {
	if setting.RunMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else if setting.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Fatalf("Wrong run mode configured in config file ! ")
	}

	r := routers.Initrouters()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
