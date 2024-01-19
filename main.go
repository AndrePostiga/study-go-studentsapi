package main

import (
	"github.com/andrepostiga/api-go-gin/api"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := api.StartWebApi()
	r.Router.Run(":5000")
}
