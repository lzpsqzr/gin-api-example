package v1

import (
	"fmt"
	"github.com/danielgjtan/gin-test/pkg/e"
	"github.com/gin-gonic/gin"
)

type AlertContent struct {
	Alerts string `json:"alerts" binding:"required"`
}

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
	c.JSON(e.HTTP_201_CREATED, gin.H{"alert": alerts.Alerts})
}

func PutTag(c *gin.Context) {
	c.String(200, "put")
}

func DeleteTag(c *gin.Context) {
	c.String(200, "delete")
}
