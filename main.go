package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"minepin-backend/config"
	"minepin-backend/model"
	"minepin-backend/pkg/logger"
	v "minepin-backend/pkg/version"
	"minepin-backend/router"
	"net/http"
	"os"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "MinePin config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()

	if *version {
		vv := v.Get()
		marshalled, err := json.MarshalIndent(&vv, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Printf(string(marshalled))
		return
	}

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	logger.InitLogger()

	model.DB.Init()

	gin.SetMode(config.GetMinePinRunMode())
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			logger.ErrorF("The router has no response, or it might took too long to start up.", err)
			os.Exit(1)
		}
		logger.Info("The route has been deployed successfully.")
	}()

	logger.InfoF("Start to listening the incoming requests on http address: '%s'", ":"+config.GetMinePinPort())
	logger.Info(http.ListenAndServe(":"+config.GetMinePinPort(), g).Error())
}

func pingServer() error {
	for i := 0; i < config.GetMinePinMaxPingCount(); i++ {
		resp, err := http.Get("http://127.0.0.1:" + config.GetMinePinPort() + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		logger.Info("Waiting for the router, retry in 1 seconds.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
