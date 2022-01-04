package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"minepin-backend/router"
	"net/http"
	"time"
)

func main() {
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...
		)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The route has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":6010")
	log.Printf(http.ListenAndServe(":6010", g).Error())
}

func pingServer() error {
	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:6010" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 seconds.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}