package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	//router inits here, so newer call
	//to cloud function would use this global value
	//and wont create router again
	router = gin.Default()
	router.GET("/", MyHandler)
	router.GET("/stuff", MyHandler2)
}

func main() {

}

func MyHandler(c *gin.Context) {
	//user := c.Params.ByName("name")

	c.JSON(http.StatusOK, gin.H{"user": "val", "value": "test"})
}

func MyHandler2(c *gin.Context) {
	//user := c.Params.ByName("name")

	c.JSON(http.StatusOK, gin.H{"user": "stuff", "value": "its working"})
}

func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(*r)
	router.ServeHTTP(w, r) // ServeHTTP conforms to the http.Handler interface (https://godoc.org/github.com/gin-gonic/gin#Engine.ServeHTTP)
}
