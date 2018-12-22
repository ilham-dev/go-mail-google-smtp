package main

import (
	"github.com/gin-gonic/gin"
	"go-mail-google-smtp/config"
)

func main(){
	router := gin.Default()

	router.POST("/sendmail", config.Send)

	_ = router.Run()
}
