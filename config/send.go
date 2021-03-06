package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send(c *gin.Context)  {
	sender := NewSender("<Your Email>", "<Your Password>")

	email := c.PostForm("email")
	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{email}

	Subject := "Testing HTLML Email from golang"
	message := `
	<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
	</head>
	<body>Hellow World<br>
	<div class="moz-signature"><i><br>
	<br>
	Regards<br>
	Ilham-Dev<br>
	<i></div>
	</body>
	</html>
	`
	bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, message)

	sender.SendMail(Receiver, Subject, bodyMessage)


	c.JSON(http.StatusOK,gin.H{"status": 200, "message": "success send to "+email})
}