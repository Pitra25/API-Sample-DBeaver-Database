package messages

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StatusResponse struct {
	Status string `json:"message"`
}

type Message struct {
	Message string `json:"message"`
}

type Level int

const (
	Info Level = iota
	Fatal
	Warning
)

func New(c *gin.Context, statusCode int, message string, typeM Level) {
	switch typeM {
	case Info:
		{
			logrus.Info(message)
			break
		}
	case Fatal:
		{
			logrus.Fatal(message)
			break
		}
	case Warning:
		{
			logrus.Warn(message)
			break
		}
	}
	c.AbortWithStatusJSON(statusCode, Message{message})
}
