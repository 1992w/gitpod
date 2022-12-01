package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET(`/gitpod-core-dev/build/supervisor:commit-b4fc228990e4325ebf7b3a8079c41ac1437b1d2c`, func(context *gin.Context) {
		context.File("supervisor.html")
	})
	r.Run(":80")
}
