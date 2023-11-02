package controllers

import (
	"fmt"
	"server/pkg/logger"

	"github.com/gin-gonic/gin"
)

type UserController struct{}
type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
type Query struct {
	Id string `json:"id"`
}

func (o UserController) GetLogin(c *gin.Context) {
	query := make(map[string]interface{})
	c.BindJSON(&query)
	fmt.Println(query["id"])
	list := make([]User, 0)
	user := User{"张三", 20}
	list = append(list, user)
	list = append(list, user)
	list = append(list, user)
	list = append(list, user)
	list = append(list, user)
	list = append(list, user)
	list = append(list, user)

	Success(c, query)
}
func (o UserController) List(c *gin.Context)   {}
func (o UserController) Add(c *gin.Context)    {}
func (o UserController) Delete(c *gin.Context) {}

func (o UserController) Exception(c *gin.Context) {
	/**
	go 的设计思想是让业务处理和异常处理分离 使业务和异常能够解耦
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到异常，如果不捕获，panic会导致程序崩溃")
		}
	}()
	Success(c, nil)
	logger.Write("rizhixinxi", "user")
	panic("执行中断,此时后续代码不会执行")

	Success(c, nil)
}
