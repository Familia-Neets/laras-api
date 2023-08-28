package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	listen(800)
}

func listen(p int) {
	port := fmt.Sprintf(":%d", p)

	r := gin.Default()
	users(r)
	content(r)
	reviews(r)

	r.Run(port)
}
