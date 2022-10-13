package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/config"
	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/router"
	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/utils"
)

func main() {
	// Hello world, the web server
	config.InitializeConfig()
	utils.InitializeDatabase()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic("[ERROR] Failed to listen and serve")
	}
}
