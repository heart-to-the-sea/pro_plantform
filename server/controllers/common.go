package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type JsonResult struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg:`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

func Success(c *gin.Context, data any) {
	json := JsonResult{0, "success", data, 0}
	fmt.Println("--------->", json, data)
	c.JSON(200, json)
}
func Error(c *gin.Context, data interface{}) {
	json := JsonResult{500, "error by server!", data, 0}
	c.JSON(200, json)
}
