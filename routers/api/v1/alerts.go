package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lzpsqzr/gin-api-example/pkg/e"
	"time"
)

type AlertContent struct {
	Status   string `json:"status" binding:"required"`
	Alerting Alerts `json:"alerts" binding:"required"`
}

type KV map[string]string

type Alert struct {
	Status       string    `json:"status"`
	Labels       KV        `json:"labels"`
	Annotations  KV        `json:"annotations"`
	StartsAt     time.Time `json:"startsAt"`
	EndsAt       time.Time `json:"endsAt"`
	GeneratorURL string    `json:"generatorURL"`
}

type Alerts []Alert

func GetTags(c *gin.Context) {
	c.String(e.HTTP_200_OK, "get")
}

func PostTag(c *gin.Context) {
	// buf := make([]byte, 1024)
	// num, _ := c.Request.Body.Read(buf)
	// reqBody := string(buf[0:num])
	// fmt.Println(buf[0:num])
	// fmt.Println(reflect.TypeOf(buf[0:num]))
	// c.JSON(200, reqBody)
	var alerts AlertContent
	c.BindJSON(&alerts)
	fmt.Println(alerts)
	// fmt.Println(reflect.TypeOf(buf[0:num]))
	c.JSON(e.HTTP_201_CREATED, gin.H{
		"alerts": alerts.Alerting,
		"Status": alerts.Status,
	})
}

func PutTag(c *gin.Context) {
	c.String(200, "put")
}

func DeleteTag(c *gin.Context) {
	c.String(200, "delete")
}
