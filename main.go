package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"minepin-backend/config"
	"minepin-backend/router"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "MinePin config file path.")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(config.GetString(config.MINEPIN_RUNMODE))
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The route has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":"+config.GetString(config.MINEPIN_PORT))
	log.Printf(http.ListenAndServe(":"+config.GetString(config.MINEPIN_PORT), g).Error())
}

func pingServer() error {
	log.Printf("%d", config.GetConfig(config.MINEPIN_MAX_PING_COUNT))
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get("http://127.0.0.1:"+ config.GetString(config.MINEPIN_PORT) + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 seconds.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
