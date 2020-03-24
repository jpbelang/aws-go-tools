package tools

import (
	"fmt"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func LaunchUsingEnvironment(addr string, engine *gin.Engine) {

	mode := os.Getenv("GIN_MODE")

	if mode == "debug" {
		fmt.Printf("Launching in HTTP mode (%s)\n", mode)
		log.Fatal(http.ListenAndServe(addr, engine))
	} else {
		fmt.Printf("Launching in lambda mode (%s)\n", mode)
		log.Fatal(gateway.ListenAndServe(addr, engine))
	}
}
