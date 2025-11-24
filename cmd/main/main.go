package main

import "C"
import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		panic(err)
	}

	if err = r.Run(":8080"); err != nil {
		panic(err)
	}

}
